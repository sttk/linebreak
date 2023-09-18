// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package linebreak

import (
	"strings"
	"text/scanner"
	"unicode"

	"golang.org/x/text/width"
)

// Line break opprtunity type
type lboType int

const (
	lbo_never lboType = iota
	lbo_before
	lbo_after
	lbo_both
	lbo_break
	lbo_space
)

type lboState struct {
	lboType  lboType
	lboPrev  lboType
	openApos int8 // 0:not, 1:opened, 2:opened inside "..."
	openQuot int8 // 0:not, 1:opened, 2:opened inside '...'
}

// LineIter is the struct that output the given string line by line.
// This struct can control the overall line witdh and the indentation from any
// desired line.
type LineIter struct {
	scanner  *scanner.Scanner
	buffer   runeBuffer
	width    [2]int /* 0: width before lbo, 1: width after lbo */
	lboPos   int
	limit    int
	indent   string
	openQuot int8
	openApos int8
}

// New is the function that creates a LineIter instance which outputs the given
// string line by line.
// The second arguument is the width of the output lines.
func New(text string, lineWidth int) LineIter {
	sc := new(scanner.Scanner)
	sc.Init(strings.NewReader(text))

	iter := LineIter{}
	iter.scanner = sc
	iter.buffer = newRuneBuffer(lineWidth)
	iter.limit = lineWidth
	return iter
}

// SetIndent is the method to set an indentation for the subsequent lines.
func (iter *LineIter) SetIndent(indent string) {
	iter.indent = indent
}

// Init is the method to re-initialize with an argument string and reuse this
// instance.
func (iter *LineIter) Init(text string) {
	iter.scanner.Init(strings.NewReader(text))
	iter.buffer.length = 0
	iter.width[0] = 0
	iter.width[1] = 0
	iter.lboPos = 0
	iter.openQuot = 0
	iter.openApos = 0
}

// Next is the method that returns a string of a next line and a bool which
// indicates whether there are more next lines or not.
func (iter *LineIter) Next() (string, bool) {
	limit := iter.limit - len(iter.indent)

	var line string

	var state lboState
	state.openQuot = iter.openQuot
	state.openApos = iter.openApos

	for r := iter.scanner.Next(); r != scanner.EOF; r = iter.scanner.Next() {
		lineBreakOppotunity(r, &state)

		if state.lboType == lbo_break {
			line = string(trimRight(iter.buffer.full()))
			iter.buffer.length = 0
			iter.width[0] = 0
			iter.width[1] = 0
			iter.openQuot = 0
			iter.openApos = 0
			iter.lboPos = 0
			if len(line) > 0 {
				line = iter.indent + line
			}
			return line, true
		}

		if iter.buffer.length == 0 && state.lboType == lbo_space {
			continue
		}

		runeW := runeWidth(r)
		lboPos := iter.lboPos

		if (iter.width[0] + iter.width[1] + runeW) > limit {
			if state.lboPrev == lbo_before {
				line := string(trimRight(iter.buffer.runes[0:lboPos]))
				iter.buffer.cr(lboPos)

				iter.buffer.add(r)
				iter.width[0] = iter.width[1] + runeW
				iter.width[1] = 0
				iter.lboPos = iter.buffer.length

				iter.openQuot = state.openQuot
				iter.openApos = state.openApos

				if len(line) > 0 {
					line = iter.indent + line
				}
				return line, true
			}

			switch state.lboType {
			case lbo_before, lbo_both, lbo_space:
				lboPos = iter.buffer.length
			}
			// break forcely when no lbo in the current line.
			if lboPos == 0 {
				iter.width[0] += iter.width[1]
				iter.width[1] = 0
				lboPos = iter.buffer.length
			}

			line := string(trimRight(iter.buffer.runes[0:lboPos]))
			iter.buffer.cr(lboPos)

			switch state.lboType {
			case lbo_space:
				iter.width[0] = 0
				iter.width[1] = 0
				iter.lboPos = 0
			case lbo_before, lbo_both:
				iter.buffer.add(r)
				iter.width[0] = runeW
				iter.width[1] = 0
				iter.lboPos = 0
			case lbo_after:
				iter.buffer.add(r)
				iter.width[0] = iter.width[1] + runeW
				iter.width[1] = 0
				iter.lboPos = iter.buffer.length
			default:
				iter.buffer.add(r)
				iter.width[0] = iter.width[1] + runeW
				iter.width[1] = 0
				iter.lboPos = 0
			}

			iter.openQuot = state.openQuot
			iter.openApos = state.openApos

			if len(line) > 0 {
				line = iter.indent + line
			}
			return line, true
		}

		if runeW > 0 {
			iter.buffer.add(r)
		}
		switch state.lboType {
		case lbo_before:
			if state.lboPrev != lbo_before {
				iter.lboPos = iter.buffer.length - 1
			}
			iter.width[0] += iter.width[1]
			iter.width[1] = runeW
		case lbo_both:
			iter.lboPos = iter.buffer.length - 1
			iter.width[0] += iter.width[1]
			iter.width[1] = runeW
		case lbo_after, lbo_space:
			iter.lboPos = iter.buffer.length
			iter.width[0] += iter.width[1] + runeW
			iter.width[1] = 0
		default:
			iter.width[1] += runeW
		}
	}

	line = string(trimRight(iter.buffer.full()))
	iter.buffer.length = 0

	if len(line) > 0 {
		line = iter.indent + line
	}
	return line, false
}

func lineBreakOppotunity(r rune, state *lboState) {
	state.lboPrev = state.lboType

	switch r {
	case 0x22: // "
		if state.openQuot == 0 { // open
			state.openQuot = state.openApos + 1
			state.lboType = lbo_before
		} else { // close
			state.openQuot = 0
			if state.openQuot < state.openApos {
				state.openApos = 0
			}
			state.lboType = lbo_after
		}
		return
	case 0x27: // '
		if state.openApos == 0 { // open
			state.openApos = state.openQuot + 1
			state.lboType = lbo_before
		} else { // close
			state.openApos = 0
			if state.openApos < state.openQuot {
				state.openQuot = 0
			}
			state.lboType = lbo_after
		}
		return
	}

	if contains(lboBreaks, r) {
		state.lboType = lbo_break
		return
	}

	if contains(lboBefores, r) {
		state.lboType = lbo_before
		return
	}

	if contains(lboAfters, r) {
		state.lboType = lbo_after
		return
	}

	if unicode.IsSpace(r) {
		state.lboType = lbo_space
		return
	}

	switch width.LookupRune(r).Kind() {
	case width.EastAsianWide, width.EastAsianFullwidth:
		state.lboType = lbo_both
		return
	}

	state.lboType = lbo_never
}

func contains(candidates []rune, r rune) bool {
	for _, e := range candidates {
		if e == r {
			return true
		}
	}
	return false
}

func runeWidth(r rune) int {
	if !unicode.IsPrint(r) {
		return 0
	}

	switch width.LookupRune(r).Kind() {
	case width.EastAsianNarrow, width.EastAsianHalfwidth, width.Neutral:
		return 1
	case width.EastAsianWide, width.EastAsianFullwidth:
		return 2
	default: // width.EastAsianAmbiguous
		return 2
	}
}

func trimRight(runes []rune) []rune {
	i := len(runes) - 1
	for ; i >= 0; i-- {
		if !unicode.IsSpace(runes[i]) {
			return runes[0 : i+1]
		}
	}
	return runes
}

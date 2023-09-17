// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package linebreak

import (
	"strings"
	"text/scanner"
	"unicode"
)

// Line break opprtunity type
type lboType int

const (
	lbo_before lboType = iota
	lbo_after
	lbo_both
	lbo_never
	lbo_break
	lbo_space
)

// LineIter is the struct that output the given string line by line.
// This struct can control the overall line witdh and the indentation from any
// desired line.
type LineIter struct {
	scanner *scanner.Scanner
	buffer  runeBuffer
	width   [2]int /* 0: width before lbo, 1: width after lbo */
	lboPos  int
	limit   int
	indent  string
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
}

// Next is the method that returns a string of a next line and a bool which
// indicates whether there are more next lines or not.
func (iter *LineIter) Next() (string, bool) {
	limit := iter.limit - len(iter.indent)

	var line string

	for r := iter.scanner.Next(); r != scanner.EOF; r = iter.scanner.Next() {
		lboTyp := lineBreakOppotunity(r)

		if lboTyp == lbo_break {
			line = string(trimRight(iter.buffer.full()))
			iter.buffer.length = 0
			iter.width[0] = 0
			iter.width[1] = 0
			iter.lboPos = 0
			if len(line) > 0 {
				line = iter.indent + line
			}
			return line, true
		}

		if iter.buffer.length == 0 && lboTyp == lbo_space {
			continue
		}

		runeW := runeWidth(r)
		lboPos := iter.lboPos

		if (iter.width[0] + iter.width[1] + runeW) > limit {
			switch lboTyp {
			case lbo_before, lbo_both, lbo_space:
				lboPos = iter.buffer.length
			}
			if lboPos == 0 {
				//iter.width[0] += iter.width[1]
				iter.width[1] = 0
				lboPos = iter.buffer.length
			}

			line := string(trimRight(iter.buffer.runes[0:lboPos]))
			iter.buffer.cr(lboPos)

			switch lboTyp {
			case lbo_space:
				iter.width[0] = 0
				iter.width[1] = 0
				iter.lboPos = 0
			//case lbo_before:
			//	iter.buffer.add(r)
			//	iter.width[0] = runeW
			//	iter.width[1] = 0
			//	iter.lboPos = 0
			case lbo_after, lbo_both:
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

			if len(line) > 0 {
				line = iter.indent + line
			}
			return line, true
		}

		if runeW > 0 {
			iter.buffer.add(r)
		}
		switch lboTyp {
		//case lbo_before:
		//	iter.lboPos = iter.buffer.length - 1
		//	iter.width[0] += iter.width[1]
		//	iter.width[1] = runeW
		case lbo_after, lbo_both, lbo_space:
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

func lineBreakOppotunity(r rune) lboType {
	if r == 0x0a || r == 0x0d {
		return lbo_break
	}
	if unicode.IsSpace(r) {
		return lbo_space
	}
	if unicode.IsPunct(r) {
		return lbo_after
	}
	return lbo_never
}

func runeWidth(r rune) int {
	if !unicode.IsPrint(r) {
		return 0
	}
	return 1
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

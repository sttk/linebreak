// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package linebreak

import (
	"os"
	"unicode"

	"golang.org/x/term"
	"golang.org/x/text/width"
)

// TermCols is the function that returns the column count of the current
// terminal.
// This count is the number of ASCII printable characters.
// If it failed to get the count, this function returns the fixed number:
// 80.
func TermCols() int {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		w = 80
	}
	return w
}

// TermSize is the function that returns the column count and row count of the
// current terminal.
// These counts are the numbers of ASCII printable characters.
// If it failed to get these counts, this function returns the fixed numbers:
// 80 columns and 24 rows..
func TermSize() (cols, rows int) {
	var err error
	cols, rows, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		cols = 80
		rows = 24
	}
	return
}

// RuneWidth is the function that returns the display width of the specified
// rune.
// A display width is determined by the Unicode Standard Annex #11 (UAX11)
// East-Asian-Width.
func RuneWidth(r rune) int {
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

// TextWidth is the function that returns the display width of the specified
// text.
// This function calculates the width of the text taking into account the
// letter width determined by the Unicode Standard Annex #11 (UAX11)
// East-Asian-Width.
func TextWidth(text string) int {
	w := 0
	for _, r := range text {
		w += RuneWidth(r)
	}
	return w
}

// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package linebreak

import (
	"os"
	"strings"

	"golang.org/x/term"
)

// TermWidth is the function that returns terminal width.
// The width is the count of ASCII printable character.
// If it failed to get terminal width, this function returns the fixed number:
// 80.
func TermWidth() int {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		w = 80
	}
	return w
}

// TextWidth is the function that calculates a text width.
// This function calculates the width of the specified text taking into
// account the letter width determined by the Unicode Standard Annex #11
// (UAX11) East-Asian-Width.
func TextWidth(text string) int {
	w := 0
	for _, r := range text {
		w += runeWidth(r)
	}
	return w
}

// Spaces is the function that generates a consecutive ASCII spaces of which
// count is specified by the argument.
// If the count is negative, this function returns an empty strings.
func Spaces(count int) string {
	if count < 0 {
		count = 0
	}
	return strings.Repeat(" ", count)
}

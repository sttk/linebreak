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

// Spaces is the function that generates a consecutive ASCII spaces of which
// count is specified by the argument.
// If the count is negative, this function returns an empty strings.
func Spaces(count int) string {
	if count < 0 {
		count = 0
	}
	return strings.Repeat(" ", count)
}

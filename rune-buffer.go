// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package linebreak

type runeBuffer struct {
	runes  []rune
	length int
}

func newRuneBuffer(capacity int) runeBuffer {
	return runeBuffer{runes: make([]rune, capacity)}
}

func (rb *runeBuffer) add(runes ...rune) bool {
	n := len(runes)
	if rb.length+n > len(rb.runes) {
		return false
	}
	for i, r := range runes {
		rb.runes[rb.length+i] = r
	}
	rb.length += n
	return true
}

func (rb *runeBuffer) cr(start int) {
	if start < 0 {
		return
	}
	if start >= rb.length {
		rb.length = 0
		return
	}

	n := rb.length - start
	for i := 0; i < n; i++ {
		rb.runes[i] = rb.runes[i+start]
	}
	rb.length = n
}

func (rb runeBuffer) slice() []rune {
	return rb.runes[0:rb.length]
}

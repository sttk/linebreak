package linebreak

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRuneWidth(t *testing.T) {
	assert.Equal(t, runeWidth('a'), 1)
	assert.Equal(t, runeWidth('あ'), 2)
	assert.Equal(t, runeWidth('ｱ'), 1)
}

func TestTrimRight(t *testing.T) {
	assert.Equal(t, trimRight([]rune{0x31, 0x20, 0x20}), []rune{0x31})
	assert.Equal(t, trimRight([]rune{0x31, 0x32}), []rune{0x31, 0x32})
	assert.Equal(t, trimRight([]rune{0x20, 0x20, 0x20}), []rune{})
}

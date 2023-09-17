package linebreak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRuneBuffer_empty(t *testing.T) {
	rb := newRuneBuffer(0)
	assert.Equal(t, rb.runes, []rune{})
	assert.Equal(t, rb.length, 0)
	assert.Equal(t, len(rb.runes), 0)
	assert.Equal(t, cap(rb.runes), 0)
	assert.Equal(t, rb.full(), []rune{})
}

func TestRuneBuffer_add(t *testing.T) {
	rb := newRuneBuffer(5)
	assert.Equal(t, rb.runes, []rune{0, 0, 0, 0, 0})
	assert.Equal(t, rb.length, 0)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{})

	assert.True(t, rb.add('1'))
	assert.Equal(t, rb.runes, []rune{'1', 0, 0, 0, 0})
	assert.Equal(t, rb.length, 1)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{'1'})

	assert.True(t, rb.add('2', '3'))
	assert.Equal(t, rb.runes, []rune{'1', '2', '3', 0, 0})
	assert.Equal(t, rb.length, 3)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{'1', '2', '3'})

	assert.False(t, rb.add('x', 'y', 'z'))
	assert.Equal(t, rb.runes, []rune{'1', '2', '3', 0, 0})
	assert.Equal(t, rb.length, 3)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{'1', '2', '3'})

	assert.True(t, rb.add('4', '5'))
	assert.Equal(t, rb.runes, []rune{'1', '2', '3', '4', '5'})
	assert.Equal(t, rb.length, 5)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{'1', '2', '3', '4', '5'})

	assert.False(t, rb.add('6'))
	assert.Equal(t, rb.runes, []rune{'1', '2', '3', '4', '5'})
	assert.Equal(t, rb.length, 5)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{'1', '2', '3', '4', '5'})
}

func TestRuneBuffer_cr(t *testing.T) {
	rb := newRuneBuffer(5)
	assert.Equal(t, rb.runes, []rune{0, 0, 0, 0, 0})
	assert.Equal(t, rb.length, 0)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{})

	assert.True(t, rb.add('1', '2', '3', '4', '5'))
	assert.Equal(t, rb.runes, []rune{'1', '2', '3', '4', '5'})
	assert.Equal(t, rb.length, 5)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{'1', '2', '3', '4', '5'})

	rb.cr(3)
	assert.Equal(t, rb.runes, []rune{'4', '5', '3', '4', '5'})
	assert.Equal(t, rb.length, 2)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{'4', '5'})

	assert.True(t, rb.add('6'))
	assert.Equal(t, rb.runes, []rune{'4', '5', '6', '4', '5'})
	assert.Equal(t, rb.length, 3)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{'4', '5', '6'})

	rb.cr(3)
	assert.Equal(t, rb.runes, []rune{'4', '5', '6', '4', '5'})
	assert.Equal(t, rb.length, 0)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{})

	assert.True(t, rb.add('1', '2', '3', '4', '5'))
	assert.Equal(t, rb.runes, []rune{'1', '2', '3', '4', '5'})
	assert.Equal(t, rb.length, 5)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{'1', '2', '3', '4', '5'})

	rb.cr(0)
	assert.False(t, rb.add('1', '2', '3', '4', '5'))
	assert.Equal(t, rb.runes, []rune{'1', '2', '3', '4', '5'})
	assert.Equal(t, rb.length, 5)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{'1', '2', '3', '4', '5'})

	rb.cr(-1)
	assert.False(t, rb.add('1', '2', '3', '4', '5'))
	assert.Equal(t, rb.runes, []rune{'1', '2', '3', '4', '5'})
	assert.Equal(t, rb.length, 5)
	assert.Equal(t, len(rb.runes), 5)
	assert.Equal(t, cap(rb.runes), 5)
	assert.Equal(t, rb.full(), []rune{'1', '2', '3', '4', '5'})
}

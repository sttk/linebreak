package linebreak_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/linebreak"
)

func TestTermCols(t *testing.T) {
	assert.Equal(t, linebreak.TermCols(), 80)
}

func TestTermSize(t *testing.T) {
	cols, rows := linebreak.TermSize()
	assert.Equal(t, cols, 80)
	assert.Equal(t, rows, 24)
}

func TestTextWidth(t *testing.T) {
	assert.Equal(t, linebreak.TextWidth("abc"), 3)
	assert.Equal(t, linebreak.TextWidth("あいう"), 6)
}

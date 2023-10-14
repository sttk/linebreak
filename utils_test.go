package linebreak_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/linebreak"
)

func TestTermWidth(t *testing.T) {
	assert.Equal(t, linebreak.TermWidth(), 80)
}

func TestTextWidth(t *testing.T) {
	assert.Equal(t, linebreak.TextWidth("abc"), 3)
	assert.Equal(t, linebreak.TextWidth("あいう"), 6)
}

func TestSpaces(t *testing.T) {
	assert.Equal(t, linebreak.Spaces(3), "   ")
	assert.Equal(t, linebreak.Spaces(-1), "")
}

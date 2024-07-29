package linebreak_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/linebreak"
)

func TestLineIter_Next_emptyText(t *testing.T) {
	text := ""
	iter := linebreak.New(text, 20)

	line, more := iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, text)

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_oneCharText(t *testing.T) {
	text := "a"
	iter := linebreak.New(text, 20)

	line, more := iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, text)

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_lessThanLineWidth(t *testing.T) {
	text := "1234567890123456789"
	iter := linebreak.New(text, 20)

	line, more := iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, text)

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_equalToLineWidth(t *testing.T) {
	text := "12345678901234567890"
	iter := linebreak.New(text, 20)

	line, more := iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, text)

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_breakAtLineBreakOpportunity(t *testing.T) {
	text := "1234567890 abcdefghij"
	iter := linebreak.New(text, 20)

	line, more := iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, text[0:10])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, text[11:21])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_removeHeadingSpaceOfEachLine(t *testing.T) {
	text := "12345678901234567890   abcdefghij"
	iter := linebreak.New(text, 20)

	line, more := iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, text[0:20])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, text[23:])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_removeTailingSpaceOfEachLine(t *testing.T) {
	text := "12345678901234567      abcdefghij"
	iter := linebreak.New(text, 20)

	line, more := iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, text[0:17])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, text[23:])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_removeSpacesOfAllSpaceLine(t *testing.T) {
	text := "       "
	iter := linebreak.New(text, 10)

	line, more := iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_thereIsNoLineBreakOpportunity(t *testing.T) {
	text := "12345678901234567890abcdefghij"
	iter := linebreak.New(text, 20)

	line, more := iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, text[0:20])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, text[20:])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "")
}

func TestLineIter_SetIndent(t *testing.T) {
	text := "12345678901234567890abcdefghij"
	iter := linebreak.New(text, 10)

	line, more := iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, text[0:10])

	iter.SetIndent("   ")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "   "+text[10:17])

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "   "+text[17:24])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "   "+text[24:])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "")
}

func TestLineIter_breakPositionAfterIndentWidthIsIncreased(t *testing.T) {
	lineWidth := 30
	indent := strings.Repeat(" ", 7)
	text := "aaaaa " + strings.Repeat("b", lineWidth-7) + strings.Repeat("c", lineWidth-7) + "ddd"

	iter := linebreak.New(text, lineWidth)

	line, more := iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "aaaaa")
	assert.Equal(t, linebreak.TextWidth(line), 5)

	iter.SetIndent(indent)

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, strings.Repeat(" ", 7)+strings.Repeat("b", lineWidth-7))
	assert.Equal(t, linebreak.TextWidth(line), lineWidth)

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, strings.Repeat(" ", 7)+strings.Repeat("c", lineWidth-7))
	assert.Equal(t, linebreak.TextWidth(line), lineWidth)

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "       ddd")
	assert.Equal(t, linebreak.TextWidth(line), 10)
}

func TestLineIter_breakPositionIfIndentContainsFullWidthChars(t *testing.T) {
	lineWidth := 30
	indent := "__ああ__" // width is 8.
	text := "aaaaa " + strings.Repeat("b", lineWidth-8) + strings.Repeat("c", lineWidth-8) + "ddd"

	iter := linebreak.New(text, lineWidth)

	line, more := iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "aaaaa")
	assert.Equal(t, linebreak.TextWidth(line), 5)

	iter.SetIndent(indent)

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, indent+strings.Repeat("b", lineWidth-8))
	assert.Equal(t, linebreak.TextWidth(line), lineWidth)

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, indent+strings.Repeat("c", lineWidth-8))
	assert.Equal(t, linebreak.TextWidth(line), lineWidth)

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, indent+"ddd")
	assert.Equal(t, linebreak.TextWidth(line), 11)
}

func TestLineIter_Init(t *testing.T) {
	text := "12345678901234567890"
	iter := linebreak.New(text, 12)

	line, more := iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, text[0:12])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, text[12:])

	text = "abcdefghijklmnopqrstuvwxyz"
	iter.Init(text)

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, text[0:12])

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, text[12:24])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, text[24:])

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "")
}

// This text is quoted from https://go.dev/doc/
const longText string = `The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.    `

func TestLineIter_Next_tryLongText(t *testing.T) {
	iter := linebreak.New(longText, 20)

	line, more := iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "The Go programming")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "language is an open")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "source project to")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "make programmers")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "more productive.")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "Go is expressive,")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "concise, clean, and")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "efficient. Its")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "concurrency")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "mechanisms make it")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "easy to write")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "programs that get")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "the most out of")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "multicore and")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "networked machines,")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "while its novel type")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "system enables")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "flexible and modular")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "program")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "construction. Go")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "compiles quickly to")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "machine code yet has")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "the convenience of")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "garbage collection")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "and the power of")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "run-time reflection.")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "It's a fast,")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "statically typed,")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "compiled language")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "that feels like a")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "dynamically typed,")

	line, more = iter.Next()
	assert.Equal(t, more, true)
	assert.Equal(t, line, "interpreted")

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "language.")

	line, more = iter.Next()
	assert.Equal(t, more, false)
	assert.Equal(t, line, "")
}

func TestLineIter_Next_printLongText(t *testing.T) {
	iter := linebreak.New(longText, 20)

	for {
		line, more := iter.Next()
		fmt.Println(line)
		if more == false {
			break
		}
	}
}

func TestLineIter_SetIndentToLongText(t *testing.T) {
	iter := linebreak.New(longText, 40)

	line, more := iter.Next()
	fmt.Println(line)

	iter.SetIndent(strings.Repeat(" ", 8))

	for {
		if more == false {
			break
		}
		line, more = iter.Next()
		fmt.Println(line)
	}
}

func TestLineIter_textContainsNonPrintChar(t *testing.T) {
	text := "abcdefg\u0002hijklmn"
	iter := linebreak.New(text, 10)

	line, more := iter.Next()
	assert.Equal(t, line, "abcdefghij")
	assert.Equal(t, more, true)

	line, more = iter.Next()
	assert.Equal(t, line, "klmn")
	assert.Equal(t, more, false)
}

func TestLineIter_letterWidthOfEastAsianWideLetter(t *testing.T) {
	text := "東アジアの全角文字は２文字分の幅をとります。"
	iter := linebreak.New(text, 20)

	line, more := iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "東アジアの全角文字は")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "２文字分の幅をとりま")

	line, more = iter.Next()
	assert.False(t, more)
	assert.Equal(t, line, "す。")
}

func TestLineIter_lineBreaksOfEastAsianWideLetter(t *testing.T) {
	text := "東アジアの全角文字は基本的に、文字の前後どちらに行の終わりが来て" +
		"も改行が行われます。"
	iter := linebreak.New(text, 28)

	line, more := iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "東アジアの全角文字は基本的")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "に、文字の前後どちらに行の終")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "わりが来ても改行が行われま")

	line, more = iter.Next()
	assert.False(t, more)
	assert.Equal(t, line, "す。")
}

func TestLineIter_prohibitionsOfLineBreakOfJapanese_start(t *testing.T) {
	text := "句読点は、行頭に置くことは禁止である。"
	iter := linebreak.New(text, 8)

	line, more := iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "句読点")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "は、行頭")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "に置くこ")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "とは禁止")

	line, more = iter.Next()
	assert.False(t, more)
	assert.Equal(t, line, "である。")
}

func TestLineIter_prohibitionsOfLineBreakOfJapanese_end(t *testing.T) {
	text := "開き括弧は「行末に置く」ことは禁止である。"
	iter := linebreak.New(text, 12)

	line, more := iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "開き括弧は")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "「行末に置")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "く」ことは禁")

	line, more = iter.Next()
	assert.False(t, more)
	assert.Equal(t, line, "止である。")
}

func TestLineIter_prohibitionsOfLineBreakOfEnglish(t *testing.T) {
	text := "abc def ghi(jkl mn opq rst uvw xyz)"
	iter := linebreak.New(text, 11)

	line, more := iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "abc def ghi")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "(jkl mn opq")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "rst uvw")

	line, more = iter.Next()
	assert.False(t, more)
	assert.Equal(t, line, "xyz)")
}

func TestLineIter_prohibitionsOfLineBreakOfEnglish_quots(t *testing.T) {
	text := `abc def " ghi j " kl mno pq" rst uvw" xyz`
	iter := linebreak.New(text, 9)

	line, more := iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "abc def")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "\" ghi j \"")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "kl mno pq")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "\" rst")

	line, more = iter.Next()
	assert.False(t, more)
	assert.Equal(t, line, "uvw\" xyz")
}

func TestLineIter_prohibitionsOfLineBreakOfEnglish_mixedQuots(t *testing.T) {
	text := `abc def " ghi j ' kl mno pq' rst uvw" xyz`
	iter := linebreak.New(text, 9)

	line, more := iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "abc def")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "\" ghi j")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "' kl mno")

	line, more = iter.Next()
	assert.True(t, more)
	assert.Equal(t, line, "pq' rst")

	line, more = iter.Next()
	assert.False(t, more)
	assert.Equal(t, line, "uvw\" xyz")

	iter = linebreak.New(text, 9)

	for {
		line, more := iter.Next()
		fmt.Println(line)
		if !more {
			break
		}
	}
}

func TestLineIter_japanese(t *testing.T) {
	text := "私はその人を常に先生と呼んでいた。だからここでもただ先生と書くだ" +
		"けで本名は打ち明けない。これは世間を憚かる遠慮というよりも、その方が私" +
		"にとって自然だからである。私はその人の記憶を呼び起すごとに、すぐ「先生" +
		"」といいたくなる。筆を執っても心持は同じ事である。よそよそしい頭文字な" +
		"どはとても使う気にならない。\n（夏目漱石「こころ」から引用）"

	iter := linebreak.New(text, 50)

	for {
		line, more := iter.Next()
		fmt.Println(line)
		if !more {
			break
		}
	}
}

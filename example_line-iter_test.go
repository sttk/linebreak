package linebreak_test

import (
	"fmt"
	"strings"

	"github.com/sttk/linebreak"
)

func ExampleLineIter() {
	text := "Go is a new language. Although it borrows ideas from existing " +
		"languages, it has unusual properties that make effective Go programs " +
		"different in character from programs written in its relatives. " +
		"\n\n(Quoted from 'Effective Go')"

	fmt.Println("....:....1....:....2....:....3....:....4....:....5")
	iter := linebreak.New(text, 50)
	for {
		line, more := iter.Next()
		fmt.Println(line)
		if !more {
			break
		}
	}

	// Output:
	// ....:....1....:....2....:....3....:....4....:....5
	// Go is a new language. Although it borrows ideas
	// from existing languages, it has unusual properties
	// that make effective Go programs different in
	// character from programs written in its relatives.
	//
	// (Quoted from 'Effective Go')
}

func ExampleLineIter_SetIndent() {
	text := "Go is a new language. Although it borrows ideas from existing " +
		"languages, it has unusual properties that make effective Go programs " +
		"different in character from programs written in its relatives. " +
		"\n\n(Quoted from 'Effective Go')"

	fmt.Println("....:....1....:....2....:....3....:....4....:....5")

	iter := linebreak.New(text, 50)
	line, more := iter.Next()
	fmt.Println(line)

	if more {
		for i := 1; ; i++ {
			iter.SetIndent(strings.Repeat(" ", i*2))
			line, more := iter.Next()
			fmt.Println(line)
			if !more {
				break
			}
		}
	}

	// Output:
	// ....:....1....:....2....:....3....:....4....:....5
	// Go is a new language. Although it borrows ideas
	//   from existing languages, it has unusual
	//     properties that make effective Go programs
	//       different in character from programs written
	//         in its relatives.
	//
	//             (Quoted from 'Effective Go')
}

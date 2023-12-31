# [linebreak][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]

A library for breaking a given text into
lines within a specified width.
This library also supports per-line indentation.

## Import declaration

To use this package in your code, the following import declaration is necessary.

```
import "github.com/sttk/linebreak"
```

## Usage

The following code breaks the argument text into lines within the terminal width, and outputs them to stdout.

```
iter := linebreak.New(text, linebreak.TermCols())
for {
	line, more := iter.Next()
	fmt.Println(line)
	if !more {
		break
	}
}
```

## Supporting Go versions

This library supports Go 1.18 or later.

### Actual test results for each Go version:

```
% gvm-fav
Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk/linebreak	0.152s	coverage: 98.9% of statements

Now using version go1.19.13
go version go1.19.13 darwin/amd64
ok  	github.com/sttk/linebreak	0.166s	coverage: 98.9% of statements

Now using version go1.20.8
go version go1.20.8 darwin/amd64
ok  	github.com/sttk/linebreak	0.162s	coverage: 98.9% of statements

Now using version go1.21.1
go version go1.21.1 darwin/amd64
ok  	github.com/sttk/linebreak	0.174s	coverage: 98.9% of statements

Back to go1.21.1
Now using version go1.21.1
```

## License

Copyright (C) 2023 Takayuki Sato

This program is free software under MIT License.<br>
See the file LICENSE in this distribution for more details.


[repo-url]: https://github.com/sttk/linebreak
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk/linebreak.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk/linebreak
[ci-img]: https://github.com/sttk/linebreak/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk/linebreak/actions
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT

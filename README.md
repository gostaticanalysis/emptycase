# emptycase

[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/gostaticanalysis/emptycase)
[![Test](https://github.com/gostaticanalysis/emptycase/actions/workflows/go.yml/badge.svg)](https://github.com/gostaticanalysis/emptycase/actions/workflows/go.yml)


Analyzer: emptycase finds case statements with no body

## Intall

```bash
$ go get github.com/gostaticanalysis/emptycase/cmd/emptycase
```


## Usage

```bash
$ go vet -vettool=`which emptycase` pkgname
```

## Motivation

It is common knowledge for gopher that there is no fall through in the switch-case statement of the Go language.
However, people who are used to C or Java may inadvertently write code like the following.

```go
func f() {
	v := 1
	switch v {
	case 1:
	case 2:
		fmt.Println("case 1 or 2")
	}
}
```

Of course, Println() donesn't be executed.
However, it takes a while to realize this when you write it.

This Analyzer detects empty cases, making it faster to notice problems.

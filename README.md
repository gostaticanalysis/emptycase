# emptycase

Analyzer: emptycase finds case statements with no body

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

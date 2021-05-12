package a

import "fmt"

func f() {
	v := 1
	switch v {
	case 1: // want "empty case"
	case 2: // Intentional empty case. Not treated as an error.
	case 3:
		// Intentional empty case. Not treated as an error.

	case 4:
		fmt.Println("execute")
	}
}

func g() {
	ch := make(chan int)

	select {
	case <-ch: // want "empty case"
	default:
		fmt.Println("execute")
	}
}

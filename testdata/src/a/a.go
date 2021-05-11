package a

import "fmt"

func f() {
	v := 1
	switch v {
	case 1: // want "empty case"
	case 2:
		fmt.Println("case 2")
	}
}

func g() {
	ch := make(chan int)

	select {
	case <-ch: // want "empty case"
	default:
		fmt.Println("no value")
	}
}

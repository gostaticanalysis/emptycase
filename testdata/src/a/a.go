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

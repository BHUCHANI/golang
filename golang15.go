package main

import "fmt"

func returnFunc(x string) func() {
	sum := 0
	return func() {
		fmt.Println(sum)
	}
}

func main() {
	returnFunc("hola")()
	x := returnFunc("its me bhuchani")
	x()
}

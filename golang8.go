package main

import "fmt"

func main() {
	x := 8
	val := (true || false) && !false
	val2 := val || false
	fmt.Printf("%t", val)
}

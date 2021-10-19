package main

import "fmt"

func add(x, y int) (z1 int, z2 int) {
	defer fmt.Println("jo na nacha uskiii....")
	z1 = x + y
	z2 = x - y
	fmt.Println("andi mandi sandiiiiiiii")
	return
}

func main() {
	ans1, ans2 := add(14, 7)
	fmt.Println(ans1, ans2)
}

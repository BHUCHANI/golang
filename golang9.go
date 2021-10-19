package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("u can ride alone")
	} else if age >= 14 {
		fmt.Println("u can ride with parents")
	} else {
		fmt.Println("nikal pakode")
	}

}

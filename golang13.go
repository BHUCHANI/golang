package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple": 5,
		"guava": 6,
		"aam":   9,
	}
	val, ok := mp["apple"]
	fmt.Println(val, ok)
	//fmt.Println(len(mp)) counting keys in a mp
	fmt.Println(mp)

}

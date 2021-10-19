package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years old at end of 2021", 2021-input)
}
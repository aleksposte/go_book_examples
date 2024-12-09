package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	res, sep := "", ""
	for _, arg := range os.Args[1:] {
		res += sep + arg
		sep = " "
	}
	fmt.Println(res)

	// Modify the echo program to also print os.Args[0], the name of the command that invoked it.
	fmt.Println(strings.Join(os.Args[:], " "))
	fmt.Println("test")
	fmt.Println(strings.Join(os.Args[1:], " "))

	// Modify the echo program to print the index and value of each of its arguments, one per line.
	for i, value := range os.Args[1:] {
		fmt.Printf("index: %v value: %v\n", i+1, value)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Print(`Input string, for "exit" print - "exit" `)
	for input.Scan() {
		fmt.Print("Input string: ")
		counts[input.Text()]++

		if input.Text() == "exit" {
			result(counts)
			os.Exit(1)
		}
	}
}

func result(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("String: %s - was wrout: %d times\n", line, n)
		}
	}
}

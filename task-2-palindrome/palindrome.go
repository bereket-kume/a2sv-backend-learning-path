package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func palindromeCheck(a, b string) bool {
	return a == b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	reversed := reverse(input)
	pali := palindromeCheck(input, reversed)

	if pali {
		fmt.Printf("%s is palindrome\n", input)
	} else {
		fmt.Printf("%s is not palindrome\n", input)
	}

}

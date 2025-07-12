package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordCount(sentence string) map[string]int {
	result := make(map[string]int)
	sentence = strings.ToLower(sentence)
	words := strings.Fields(sentence)
	for _, word := range words {
		result[word]++
	}
	return result
}

func main() {
	fmt.Println("enter your sentence:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sentence := scanner.Text()
	answer := wordCount(sentence)
	fmt.Println(answer)
}

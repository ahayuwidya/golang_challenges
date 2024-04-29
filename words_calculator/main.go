package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter words to calculate: ")
	reader := bufio.NewReader(os.Stdin)
	inputSentence, _ := reader.ReadString('\n')

	for _, char := range inputSentence {
		fmt.Printf("%c\n", char)
	}

	charCount := make(map[string]int)
	for _, char := range inputSentence {
		charCount[string(char)]++
	}
	fmt.Println(charCount)
}

package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(flipOverWords("snow dog sun"))
}

func flipOverWords(str string) string {
	words := strings.Split(str, " ")
	slices.Reverse(words)
	return strings.Join(words, " ")
}

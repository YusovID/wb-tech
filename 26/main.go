package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "abcd"
	str2 := "abca"

	fmt.Printf("str1 has unique symbols: %t\n", hasUniqueSymbols(str1))
	fmt.Printf("str2 has unique symbols: %t\n", hasUniqueSymbols(str2))
}

func hasUniqueSymbols(str string) bool {
	uniqueSymbols := make(map[rune]struct{})
	// приводим к нижнему регистру, потому что а и А - это одна и та же буква
	str = strings.ToLower(str)

	for _, symbol := range str {
		if _, ok := uniqueSymbols[symbol]; ok {
			return false
		}
		uniqueSymbols[symbol] = struct{}{}
	}

	return true
}

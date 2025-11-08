package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(flipOverString(`Разработать программу, которая переворачивает подаваемую на вход строку.

Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.), то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).`))
}

func flipOverString(str string) string {
	runes := []rune(str)
	slices.Reverse(runes)
	return string(runes)
}

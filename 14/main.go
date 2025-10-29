package main

import (
	"fmt"
)

func main() {
	var1 := 1
	var2 := "string"
	var3 := true
	var4 := make(chan interface{})
	var5 := []int{}

	fmt.Printf("type of var1: %s\n", typeOf(var1))
	fmt.Printf("type of var2: %s\n", typeOf(var2))
	fmt.Printf("type of var3: %s\n", typeOf(var3))
	fmt.Printf("type of var4: %s\n", typeOf(var4))
	fmt.Printf("type of var5: %s\n", typeOf(var5))
}

func typeOf(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "chan"
	default:
		return "unknown type"
	}
}

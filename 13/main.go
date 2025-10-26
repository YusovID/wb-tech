package main

import "fmt"

func main() {
	num1, num2 := 0, 1
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)

	switchByAddAndSub(&num1, &num2)
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)

	switchByXOR(&num1, &num2)
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
}

func switchByAddAndSub(num1, num2 *int) {
	*num1 = *num1 - *num2
	*num2 = *num2 + *num1
	*num1 = *num2 - *num1
}

func switchByXOR(num1, num2 *int) {
	*num1 ^= *num2
	*num2 ^= *num1
	*num1 ^= *num2
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Создаем ридер для чтения пользовательского ввода.
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("To leave, press Ctrl+C or type 'exit'")

	// Бесконечный цикл для повторного выполнения программы.
	for {
	numberInput:
		// Получаем число от пользователя.
		fmt.Print("Input number: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Используем TrimSpace для чистоты.

		if input == "exit" {
			return
		}

		// Конвертируем строку в число.
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Invalid number: %v\nTry again\n", err)
			goto numberInput // В случае ошибки, возвращаемся к вводу числа.
		}

		// Выводим число в десятичном и двоичном виде.
		fmt.Printf("inputed number = %d (%b)\n", inputInt, inputInt)

	bitNumberInput:
		// Получаем номер бита для изменения.
		fmt.Print("Input bit number: ")
		bitNumberStr, _ := reader.ReadString('\n')
		bitNumberStr = strings.TrimSpace(bitNumberStr)

		bitNumber, err := strconv.Atoi(bitNumberStr)
		if err != nil {
			fmt.Printf("Invalid bit number: %v\nTry again\n", err)
			goto bitNumberInput
		}

		// Номер бита должен быть положительным (т.к. считаем с 1).
		if bitNumber <= 0 {
			fmt.Println("Invalid bit number: bit number can't be less than or equal to zero!")
			goto bitNumberInput
		}

	toOneInput:
		var setToOne bool

		// Определяем, установить бит в 1 или 0.
		fmt.Print("Set to one? (true/no OR 0/1): ")
		setToOneStr, _ := reader.ReadString('\n')
		setToOneStr = strings.TrimSpace(setToOneStr)

		switch setToOneStr {
		case "true", "1":
			setToOne = true
		case "no", "0":
			setToOne = false
		default:
			fmt.Println("Invalid input\nTry again")
			goto toOneInput
		}

		// Создаем переменную нашего типа и вызываем метод изменения бита.
		num := Num(inputInt)
		num.setBit(bitNumber, setToOne)

		// Проверяем, изменилось ли число после операции.
		if num == Num(inputInt) {
			if setToOne {
				fmt.Println("Nothing changed.\nTry another bit number or try to set it to zero")
			} else {
				fmt.Println("Nothing changed.\nTry another bit number or try to set it to one")
			}
			goto bitNumberInput // Предлагаем ввести другой бит.
		}

		// Выводим результат.
		fmt.Printf("New number is %d (%b)\n", num, num)
		fmt.Println()
	}
}

// Num - пользовательский тип на основе int64 для добавления методов.
type Num int64

// setBit устанавливает или сбрасывает i-й бит числа.
func (n *Num) setBit(bitNumber int, toOne bool) {
	// Конвертируем номер бита (с 1) в позицию (с 0) для сдвига.
	pos := bitNumber - 1

	if toOne {
		// Установка бита: n = n | (1 << pos).
		// Операция ИЛИ с маской (где 1 только в нужной позиции) устанавливает бит.
		*n |= (1 << pos)
	} else {
		// Сброс бита: n = n &^ (1 << pos).
		// Операция "И-НЕ" с маской сбрасывает бит в нужной позиции.
		*n &^= (1 << pos)
	}
}

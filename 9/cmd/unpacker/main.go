package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var ErrOnlyDigits = errors.New("string can't contain only digits")

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" || input == "quit" {
			os.Exit(1)
		}

		result, err := unpackString(input)
		if err != nil {
			fmt.Printf("ERROR: unpackString failed: %v", err)
		}

		fmt.Println(result)
	}
}

func unpackString(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	var (
		currentRuneIndex = 0
		result           = make([]rune, 0)
		// работаем со слайсом рун для поддержки unicode
		runes      = []rune(input)
		onlyDigits = true
		prevRune   rune
	)

	for i := currentRuneIndex; i < len(runes); i++ {
		r := runes[i]

		if unicode.IsDigit(r) {
			// если перед цифрой был слэш, считаем её литералом
			if prevRune == '\\' {
				result = append(result, r)
				prevRune = r
				continue
			}

			prevRuneLength, err := getFullNumber(runes[i:])
			if err != nil {
				return "", fmt.Errorf("getFullNumber failed: %w", err)
			}

			if r != '0' {
				// добавляем n-1 копий, так как один символ уже записан
				for range prevRuneLength - 1 {
					result = append(result, prevRune)
				}
			} else if len(result) >= 1 {
				// удаляем последний символ при множителе 0
				result = result[:len(result)-1]
			}

			// сдвигаем индекс цикла на длину обработанного числа
			i += len(strconv.Itoa(prevRuneLength)) - 1
		} else if r != '\\' {
			// встретили валидный символ, не являющийся частью escape-последовательности
			onlyDigits = false
			result = append(result, r)
		}

		prevRune = r
	}

	if onlyDigits {
		return "", ErrOnlyDigits
	}

	return string(result), nil
}

func getFullNumber(runes []rune) (int, error) {
	var rightRangeOfNumber int

	// ищем границу, где заканчиваются цифры
	for i, r := range runes {
		rightRangeOfNumber = i + 1

		if !unicode.IsDigit(r) {
			rightRangeOfNumber--
			break
		}
	}

	numStr := string(runes[:rightRangeOfNumber])

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return -1, fmt.Errorf("can't parse '%s' to int", numStr)
	}

	return num, nil
}

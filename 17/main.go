package main

import "fmt"

func main() {
	// отсортированный срез для поиска
	nums := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}

	// цикл для проверки работы поиска
	for numToFind := range 12 {
		index := binarySearch(nums, numToFind)
		if index == -1 {
			fmt.Printf("num %d not found\n", numToFind)
			continue
		}

		fmt.Printf("index of num %d is %d\n", numToFind, index)
	}
}

// binarySearch ищет элемент в отсортированном срезе
func binarySearch(nums []int, numToFind int) int {
	// устанавливаем левую и правую границы
	left, right := 0, len(nums)-1

	// цикл выполняется пока левая граница не станет больше правой
	for left <= right {
		// безопасное вычисление среднего индекса
		mid := left + (right-left)/2

		foundNum := nums[mid]

		// сравниваем искомое число с центральным элементом
		switch {
		// элемент найден
		case foundNum == numToFind:
			return mid
		// искомое число меньше, сдвигаем правую границу
		case foundNum > numToFind:
			right = mid - 1
		// искомое число больше, сдвигаем левую границу
		case foundNum < numToFind:
			left = mid + 1
		}
	}

	// элемент не найден
	return -1
}

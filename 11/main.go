package main

import "fmt"

func main() {
	nums1 := []int{5, 4, 3, 2, 2, 1}
	nums2 := []int{1, 3, 5, 2, 2, 4}

	fmt.Println(intersections(nums1, nums2))
}

func intersections(nums1, nums2 []int) []int {
	// создаём слайс для результата
	result := []int{}
	// если один из слайсов пуст, пересечения нет
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}

	// для экономии памяти мапу будем делать из меньшего слайса
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	// создаём мапу для подсчёта количества каждого числа
	numsCount := make(map[int]int)

	// заполняем мапу числами из меньшего слайса
	for _, num := range nums1 {
		numsCount[num]++
	}

	// проходим по второму слайсу
	for _, num := range nums2 {
		// проверяем, есть ли такое число в мапе и его счётчик больше нуля
		_, ok := numsCount[num]
		// если такого числа нет, переходим к следующему
		if !ok {
			continue
		}

		// если число есть, добавляем его в результат
		result = append(result, num)
		// и уменьшаем его счётчик в мапе
		numsCount[num]--
		// если счётчик обнулился, удаляем ключ из мапы для оптимизации
		if numsCount[num] == 0 {
			delete(numsCount, num)
		}
	}

	return result
}

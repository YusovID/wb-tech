package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4}

	slice, err := deleteElem(slice, 0)
	if err != nil {
		fmt.Printf("can't delete element: %v", err)
	}

	fmt.Println(slice)
}

// deleteElem удаляет элемент по индексу с сохранением порядка
func deleteElem[T any](slice []T, index int) ([]T, error) {
	// получаем индекс последнего элемента
	last := len(slice) - 1
	// проверяем, что индекс находится в допустимых границах
	if index < 0 || index > last {
		return nil, fmt.Errorf("index out of range")
	}

	// сдвигаем часть слайса влево, чтобы затереть удаляемый элемент
	copy(slice[index:], slice[index+1:])
	// очищаем последний элемент, чтобы избежать утечки памяти
	clearElem(slice, last)
	// возвращаем укороченный на один элемент срез
	return slice[:last], nil
}

// clearElem заменяет элемент по индексу на его нулевое значение
func clearElem[T any](slice []T, index int) {
	// создаем нулевое значение для типа T (0, nil, "" и т.д.)
	var zero T
	// присваиваем нулевое значение элементу
	slice[index] = zero
}

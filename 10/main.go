package main

import "fmt"

const (
	// шаг для группировки температур
	groupStep = 10
)

var (
	// исходные данные
	temps = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
)

// temp = temperature
// temps = temperatures
func main() {
	// вызываем функцию группировки
	groupedTemperatures := groupTemps(temps)

	// выводим результат на печать
	for group, temps := range groupedTemperatures {
		fmt.Printf("Group: %d, temperatures: %+v\n", group, temps)
	}
}

// groupTemps группирует слайс температур в map
func groupTemps(temps []float64) map[int][]float64 {
	// создаем map для хранения групп
	groupedTemps := make(map[int][]float64)

	// проходим по каждой температуре
	for _, temp := range temps {
		// получаем ключ группы для текущего значения
		tempGroup := getTempGroup(temp)

		// добавляем значение в слайс соответствующей группы
		groupedTemps[tempGroup] = append(groupedTemps[tempGroup], temp)
	}

	return groupedTemps
}

// getTempGroup вычисляет ключ группы для одной температуры
func getTempGroup(temp float64) int {
	// эта магия работает из-за целочисленного деления
	// при преобразовании в int дробная часть отбрасывается
	// например, -25.4 -> -25; -25 / 10 -> -2; -2 * 10 -> -20
	// а для 19.0 -> 19; 19 / 10 -> 1; 1 * 10 -> 10
	return int(temp) / groupStep * groupStep
}
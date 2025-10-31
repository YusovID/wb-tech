// Основная проблема этой программы заключается в том, что в ней происходит утечка памяти из-за того,
// что глобальная переменная justString указывает на срез строки, созданной в функции createHugeString.
// Не смотря на то, что мы используем срез в 100 байт, исходная строка является слайсом из 1000 байт,
// которые не будут собраны GC, потому что мы до сих пор на него ссылаемся.
// Утечка памяти - это ситуация, когда под программу выделяется большое количество оперативной памяти,
// и она не освобождается до завершения программы, или ее выделяется так много, что программа "съедает" ее,
// и программа завершается аварийно, и устранить эту ошибку, не исправив логику программы, невозможно

// Ответ на вопрос о переменной justString:
// Переменная justString - это глобальная переменная, которая будет использоваться до конца работы программы.
// Внутри функции someFunc в нее записывается срез первых 100 байт строки v, которая возвращается из функции createHugeString.

// Ниже приведен пример правильной реализации

package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

// var justString string

func main() {
	someFunc()

	fmt.Println("memory usage before GC")
	PrintMemUsage()

	runtime.GC()

	fmt.Println("memory usage after GC")
	PrintMemUsage()
}

func someFunc() {
	// для наглядности увеличим размер строки
	v := createHugeString(1 << 20)

	// NOTE: чтобы использовать способы 1-3, разкомментируйте глобальную переменную justString

	// 1
	// создадим временный массив байт, в который заапендим необходимые нам байты,
	// а потом преобразуем в строку и присвоим нашей переменной justString

	// var tempSlice []byte
	// tempSlice = append(tempSlice, []byte(v[:100])...)
	// justString = string(tempSlice)

	// 2
	// перепишем код выше в одну строку

	// justString = string([]byte(v[:100]))

	// 3
	// воспользуемся встроенной функцией Clone

	// justString = strings.Clone(v[:100])

	// 4
	// идиоматичный способ - это не использовать глобальные переменные

	justString := v[:100]
	// потом что-то делаем со строкой и в какой-то момент ее соберет сборщик мусора
	_ = justString

	// если мы будем возвращать эту строку из функции, то все равно понадобятся способы 1-3
}

// желательно вынести в отдельный модуль
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func createHugeString(strLen int) string {
	tempSlice := make([]rune, strLen)

	for i := range tempSlice {
		tempSlice[i] = letters[rand.Intn(len(letters))]
	}

	return string(tempSlice)
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v KB", m.Alloc/1024)
	fmt.Printf("\tTotalAlloc = %v KB", m.TotalAlloc/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

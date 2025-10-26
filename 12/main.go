package main

import "fmt"

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(unique(strs))
}

func unique(strs []string) []string {
	// создаём слайс с аллокацией памяти
	result := make([]string, 0, len(strs))
	// используем мапу как множество для проверки уникальности
	unique := make(map[string]struct{})

	// проходим по исходному слайсу строк
	for _, str := range strs {
		// если такой строки ещё нет в нашем множестве
		if _, ok := unique[str]; !ok {
			// добавляем строку в множество
			unique[str] = struct{}{}
			// и добавляем её в результирующий слайс
			result = append(result, str)
		}
	}

	// возвращаем "копию" слайса, чтобы избежать утечек памяти
	return append([]string(nil), result...)
}

package main

import (
	"context"
	"fmt"
	"level-1/7/generator"
	syncmap "level-1/7/syncMap"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	// Использование context.WithTimeout для ограничения общего времени работы программы.
	// Это обеспечивает graceful shutdown (корректное завершение) всех горутин.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// out - канал, из которого горутины-писатели будут получать данные.
	out := generator.Generate(ctx)

	// syncMap - наша потокобезопасная реализация map.
	syncMap := syncmap.New()

	// wg используется для ожидания завершения всех запущенных горутин.
	wg := &sync.WaitGroup{}

	run(ctx, syncMap, out, wg)

	// Ожидаем, пока счетчик WaitGroup не станет равен нулю.
	// Это гарантирует, что все горутины-читатели и писатели завершили свою работу.
	wg.Wait()

	// Выводим статистику только после того, как все операции с map завершены.
	syncMap.Stat()
}

func run(ctx context.Context, syncMap *syncmap.SyncMap, out chan generator.KeyValue, wg *sync.WaitGroup) {
	// Запускаем горутины для записи.
	writing(ctx, syncMap, out, wg)
	// Добавляем в WaitGroup горутину для чтения и запускаем ее.
	wg.Add(1)
	go reading(ctx, syncMap, wg)
}

// writing запускает 10 горутин, которые конкурентно пишут данные в syncMap.
func writing(ctx context.Context, syncMap *syncmap.SyncMap, out chan generator.KeyValue, wg *sync.WaitGroup) {
	// Запускаем 10 горутин-писателей.
	for range 10 {
		// Увеличиваем счетчик WaitGroup для каждой новой горутины.
		wg.Add(1)
		go func() {
			// Уменьшаем счетчик при завершении горутины.
			defer wg.Done()

			// Бесконечный цикл чтения из канала или ожидания завершения контекста.
			for {
				select {
				// Если контекст завершен, горутина прекращает работу.
				case <-ctx.Done():
					return
				// Читаем новое значение из канала.
				case keyValue, ok := <-out:
					// Если канал закрыт, выходим.
					if !ok {
						return
					}
					fmt.Printf("got new data to write: key = %d, value = %s\n", keyValue.Key, keyValue.Value)
					// Производим запись в потокобезопасную map.
					syncMap.Write(keyValue.Key, keyValue.Value)
				}
			}
		}()
	}
}

// reading - горутина, которая конкурентно читает данные из syncMap.
func reading(ctx context.Context, syncMap *syncmap.SyncMap, wg *sync.WaitGroup) {
	defer wg.Done()

	// Используем тикер, чтобы читать данные с определенной периодичностью.
	ticker := time.NewTicker(100 * time.Millisecond)
	// Освобождаем ресурсы тикера при выходе из функции.
	defer ticker.Stop()

	for {
		select {
		// Если контекст завершен, горутина прекращает работу.
		case <-ctx.Done():
			return
		// По сигналу тикера выполняем чтение.
		case <-ticker.C:
			key := rand.IntN(100)

			// Читаем значение из потокобезопасной map.
			value := syncMap.Read(key)
			if value == nil {
				fmt.Printf("no value for key %d was found\n", key)
				continue
			}
			fmt.Printf("value for key %d was found: %+v\n", key, value)
		}
	}
}

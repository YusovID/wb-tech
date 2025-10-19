package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	TimeToLive  = 5   // секунд
	TimeToSleep = 500 // миллисекунд
)

var (
	nums = []int{1, 2, 3, 4, 5}
)

func main() {
	// создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), TimeToLive*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}

	// запускаем конвейер
	ch1 := writer(ctx, nums)
	ch2 := processor(ctx, ch1, square)
	reader(ctx, ch2, wg)

	// ожидаем завершения
	wg.Wait()
}

// writer отправляет числа в канал
func writer(ctx context.Context, nums []int) chan int {
	writerChan := make(chan int)

	go func() {
		// закрываем канал по завершении
		defer close(writerChan)

		for _, num := range nums {
			select {
			// прерываемся при отмене контекста
			case <-ctx.Done():
				fmt.Printf("writer: context was canceled: %v\n", ctx.Err())
				return

			// отправляем число
			case writerChan <- num:
			}
		}
	}()

	return writerChan
}

// processor читает, обрабатывает и отправляет данные
func processor(ctx context.Context, writerChan chan int, process func(int) int) chan int {
	resultChan := make(chan int)

	go func() {
		// закрываем канал по завершении
		defer close(resultChan)

		for {
			select {
			// прерываемся при отмене контекста
			case <-ctx.Done():
				fmt.Printf("processor: context was canceled while waiting: %v\n", ctx.Err())
				return

			// читаем из входного канала
			case num, ok := <-writerChan:
				// выходим, если входной канал закрыт
				if !ok {
					fmt.Println("writer channel was closed")
					return
				}

				// обрабатываем число
				result := process(num)

				// пытаемся отправить результат
				select {
				// прерываемся при отмене контекста
				case <-ctx.Done():
					fmt.Printf("processor: context was canceled while sending: %v\n", ctx.Err())
					return

				// отправляем результат
				case resultChan <- result:
				}
			}
		}
	}()

	return resultChan
}

// reader читает из канала и выводит результат
func reader(ctx context.Context, resultChan chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		// сигнализируем о завершении
		defer wg.Done()

		for {
			select {
			// прерываемся при отмене контекста
			case <-ctx.Done():
				return
			// читаем из канала
			case num, ok := <-resultChan:
				// выходим, если канал закрыт
				if !ok {
					fmt.Println("result channel was closed")
					return
				}

				// печатаем результат
				fmt.Println(num)
			}
		}
	}()
}

// square - функция обработки данных
func square(num int) int {
	// имитируем долгую операцию
	time.Sleep(TimeToSleep * time.Millisecond)
	return num * num
}

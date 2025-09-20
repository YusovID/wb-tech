// Пакет main является точкой входа в программу.
// Он демонстрирует создание и управление пулом воркеров,
// который обрабатывает числа, генерируемые отдельной горутиной.
// Программа корректно завершается по сигналу SIGINT (Ctrl+C).
package main

import (
	"context"
	"errors"
	"fmt"
	"level-1/4/generator"
	worker_pool "level-1/4/pool"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// NumWorkers определяет количество воркеров в пуле.
const NumWorkers = 10

// Pool определяет интерфейс для пула воркеров.
// Это позволяет легко заменять реализации пула в будущем.
type Pool interface {
	Start(int)
	Process(int) error
	Stop()
}

// main - основная функция программы.
func main() {
	// Создаем корневой контекст с функцией отмены.
	// cancel() будет вызван для оповещения всех горутин о необходимости завершения.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}

	// Запускаем генератор чисел. Он вернет канал, из которого можно читать числа.
	out := generator.Generate(ctx)

	// Создаем и запускаем пул воркеров.
	var pool Pool = worker_pool.New(ctx, NumWorkers, square)
	pool.Start(NumWorkers)

	// Запускаем горутину-диспетчер, которая читает числа от генератора
	// и отправляет их в пул воркеров.
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done(): // Если программа завершается, выходим.
				fmt.Println("Stopping program...")
				return

			case number, ok := <-out:
				if !ok { // Если канал генератора закрыт, выходим.
					return
				}

				// Отправляем число на обработку в пул.
				err := pool.Process(number)
				if err != nil {
					// Если пул уже закрыт, прекращаем отправку.
					if errors.Is(err, worker_pool.ErrClosedPool) {
						fmt.Println("Worker pool is closed.")
						return
					}
				}
			}
		}
	}()

	// Настраиваем обработку сигнала прерывания (Ctrl+C).
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // Добавлен SIGTERM для более универсальной обработки.

	// Ожидаем получения сигнала.
	<-sigChan
	fmt.Println("\nCtrl+C was pressed. Stopping program...")

	// Инициируем отмену контекста. Это сигнал всем горутинам на завершение.
	cancel()

	// Ожидаем завершения горутины-диспетчера.
	wg.Wait()

	// Останавливаем пул воркеров, ожидая завершения всех задач.
	pool.Stop()

	fmt.Println("Program was stopped.")
}

// square является функцией-обработчиком для воркеров.
// Она имитирует полезную работу с помощью time.Sleep,
// а затем вычисляет и выводит квадрат полученного числа.
func square(number int) {
	time.Sleep(1 * time.Second)
	fmt.Printf("square of %d = %d\n", number, number*number)
}

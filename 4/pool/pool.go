// Package worker_pool реализует пул горутин-воркеров.
// Пул управляет жизненным циклом воркеров и распределяет задачи между ними.
// Завершение работы пула контролируется через context.Context.
package worker_pool

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// ErrClosedPool возвращается при попытке отправить задачу в пул,
// который уже находится в процессе остановки (контекст отменен).
var ErrClosedPool = errors.New("pool is closed")

// Pool представляет собой пул воркеров.
type Pool struct {
	ctx     context.Context // Контекст для управления отменой операций и завершением работы воркеров.
	numbers chan int        // Канал для передачи задач (чисел) воркерам.
	handler func(int)       // Функция-обработчик, выполняемая каждым воркером для каждой задачи.
	wg      *sync.WaitGroup // WaitGroup для ожидания завершения всех горутин-воркеров.
}

// New создает и возвращает новый экземпляр Pool.
// Требуется контекст для управления жизненным циклом, количество воркеров для запуска
// и функция handler, которая будет обрабатывать поступающие числа.
func New(ctx context.Context, numWorkers int, handler func(number int)) *Pool {
	return &Pool{
		ctx:     ctx,
		numbers: make(chan int), // Канал небуферизованный.
		handler: handler,
		wg:      &sync.WaitGroup{},
	}
}

// Start запускает указанное количество горутин-воркеров.
// Каждый воркер ожидает получения задач из канала или сигнала о завершении от контекста.
func (p *Pool) Start(numWorkers int) {
	// В Go 1.22 и новее `workerID` создается для каждой итерации, что безопасно для горутин.
	for workerID := range numWorkers {
		p.wg.Add(1)
		go func(workerID int) { // Явно передаем workerID для совместимости и ясности.
			defer p.wg.Done()

			fmt.Printf("Starting worker %d.\n", workerID)
			defer fmt.Printf("Worker %d was stopped.\n", workerID)

			for {
				select {
				case <-p.ctx.Done(): // Если контекст отменен, воркер немедленно завершает работу.
					return

				case number, ok := <-p.numbers: // Читаем задачу из канала.
					if !ok {
						// Эта ветка может быть достигнута, если канал будет закрыт,
						// но в текущей реализации Stop() канал не закрывается.
						// continue здесь для полноты, но основной механизм - ctx.Done().
						continue
					}
					p.handler(number) // Выполняем обработку задачи.
				}
			}
		}(workerID) // Передаем копию workerID в замыкание.
	}
}

// Process отправляет число (задачу) в пул для обработки.
// Метод является неблокирующим по отношению к отмене: если контекст
// был отменен, он немедленно вернет ошибку ErrClosedPool.
// В противном случае он заблокируется до тех пор, пока один из воркеров
// не станет доступен для приема задачи.
func (p *Pool) Process(number int) error {
	select {
	case <-p.ctx.Done(): // Проверяем, не был ли пул остановлен.
		return ErrClosedPool

	case p.numbers <- number: // Отправляем задачу в канал.
		fmt.Printf("Got new number - %d.\n", number)
		return nil
	}
}

// Stop блокирует выполнение до тех пор, пока все воркеры не завершат свою работу.
// Завершение воркеров инициируется отменой контекста, переданного при создании пула.
func (p *Pool) Stop() {
	fmt.Println("Stopping worker pool...")

	// Ожидаем, пока счетчик WaitGroup не станет равен нулю,
	// что означает завершение всех горутин, запущенных с p.wg.Add(1).
	p.wg.Wait()

	fmt.Println("Worker pool was stopped.")
}

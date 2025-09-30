package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	timeToLive  = 5 // секунд
	timeToSleep = 1 // секунд
)

var (
	toStop bool
	mu     *sync.Mutex     = &sync.Mutex{}
	wg     *sync.WaitGroup = &sync.WaitGroup{}
)

func main() {
	runStopViaCondition()

	fmt.Println("---------------------------------------------------------")

	runStopViaNotification()

	fmt.Println("---------------------------------------------------------")

	runStopViaContext()

	fmt.Println("---------------------------------------------------------")

	runStopViaGoexit()

	fmt.Println("---------------------------------------------------------")

	runStopViaClosingChannel()
}

// runStopViaCondition демонстрирует остановку горутины через изменение общего флага.
func runStopViaCondition() {
	fmt.Println("This goroutine will stop via condition")

	wg.Add(1)
	go stopViaCondition()

	time.Sleep(timeToLive * time.Second)

	// Сигнализируем горутине о необходимости завершения, безопасно изменив флаг.
	mu.Lock()
	fmt.Println("changing flag...")
	toStop = true
	mu.Unlock()

	wg.Wait()
}

// stopViaCondition в цикле проверяет значение флага под мьютексом.
func stopViaCondition() {
	defer wg.Done()

	for {
		mu.Lock()
		if toStop {
			fmt.Println("stop working")
			mu.Unlock()
			return // Выход из цикла и завершение горутины.
		}
		mu.Unlock()

		fmt.Println("i'm working")
		time.Sleep(timeToSleep * time.Second)
	}
}

// runStopViaNotification демонстрирует остановку через отправку сигнала в канал.
func runStopViaNotification() {
	fmt.Println("This goroutine will stop via notification")

	doneChan := make(chan struct{})

	wg.Add(1)
	go stopViaNotification(doneChan)

	time.Sleep(timeToLive * time.Second)

	// Отправка сигнала в канал служит уведомлением для завершения.
	fmt.Println("sending notification...")
	doneChan <- struct{}{}

	wg.Wait()
}

// stopViaNotification использует select для неблокирующего ожидания сигнала в канале.
func stopViaNotification(doneChan chan struct{}) {
	defer wg.Done()

	for {
		select {
		case <-doneChan: // Если из канала получены данные,
			fmt.Println("stop working")
			return // горутина завершается.
		default:
			fmt.Println("i'm working")
			time.Sleep(timeToSleep * time.Second)
		}
	}
}

// runStopViaContext демонстрирует каноничный способ остановки с помощью пакета context.
func runStopViaContext() {
	fmt.Println("This goroutine will stop via context")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go stopViaContext(ctx)

	time.Sleep(timeToLive * time.Second)

	// Вызов функции cancel() сигнализирует об отмене всем горутинам, использующим данный контекст.
	fmt.Println("cancelling context...")
	cancel()

	wg.Wait()
}

// stopViaContext ожидает сигнала отмены через канал Done() контекста.
func stopViaContext(ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done(): // Канал закрывается при вызове cancel().
			fmt.Println("stop working")
			return
		default:
			fmt.Println("i'm working")
			time.Sleep(timeToSleep * time.Second)
		}
	}
}

// runStopViaGoexit демонстрирует принудительное завершение горутины изнутри.
func runStopViaGoexit() {
	fmt.Println("This goroutine will stop via Goexit")

	wg.Add(1)
	go stopViaGoexit()

	wg.Wait()
}

// stopViaGoexit завершает свою работу вызовом runtime.Goexit().
func stopViaGoexit() {
	defer wg.Done()
	// Отложенные вызовы (defer) выполняются перед завершением горутины.
	defer fmt.Println("stop working")

	for range 5 {
		fmt.Println("i'm working")
		time.Sleep(timeToSleep * time.Second)
	}

	fmt.Println("calling Goexit...")
	// Немедленно прекращает выполнение текущей горутины.
	runtime.Goexit()

	fmt.Println("this line will never be printed") // потому что мы прекратили выполнение горутины.
}

// runStopViaClosingChannel демонстрирует остановку воркера через закрытие канала с "задачами".
func runStopViaClosingChannel() {
	fmt.Println("This goroutine will stop via closing channel")

	outChan := make(chan struct{})

	wg.Add(1)
	go stopViaClosingChannel(outChan)

	// Имитируем отправку нескольких задач в канал.
	for range 5 {
		outChan <- struct{}{}
	}

	fmt.Println("closing channel...")
	// Закрытие канала сигнализирует воркеру, что новых задач больше не будет.
	close(outChan)

	wg.Wait()
}

// stopViaClosingChannel читает данные из канала, пока он не будет закрыт.
func stopViaClosingChannel(outChan chan struct{}) {
	defer wg.Done()

	// Цикл for range автоматически прекращает работу, когда канал закрывается и опустошается.
	for range outChan {
		fmt.Println("i'm working")
		time.Sleep(timeToSleep * time.Second)
	}

	fmt.Println("stop working")
}

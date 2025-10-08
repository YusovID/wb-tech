package generator

import (
	"context"
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
)

const (
	MaxTimeToSleep = 100 // миллисекунд
)

type KeyValue struct {
	Key   int
	Value string
}

// Generate запускает горутину, которая симулирует поступление данных из внешнего источника.
// Функция возвращает канал, из которого можно читать сгенерированные данные.
func Generate(ctx context.Context) chan KeyValue {
	out := make(chan KeyValue)

	go func() {
		// Бесконечный цикл генерации данных.
		for {
			select {
			// Если основной контекст завершен, горутина-генератор закрывает канал и останавливается.
			case <-ctx.Done():
				close(out)
				return
			default:
				// Создаем новую пару ключ-значение.
				keyValue := KeyValue{
					Key:   rand.IntN(100),
					Value: uuid.New().String(),
				}

				// Отправляем данные в канал.
				out <- keyValue

				// Добавляем случайную задержку для имитации неравномерного потока данных.
				randSleep()
			}
		}
	}()

	return out
}

func randSleep() {
	randTimeToSleep := time.Duration(rand.IntN(MaxTimeToSleep))
	time.Sleep(randTimeToSleep * time.Millisecond)
}

package adapters

import (
	"fmt"
	"level-1/21/messenger"
	"level-1/21/telegram"
)

// TelegramAdapter это адаптер для клиента Telegram
type TelegramAdapter struct {
	// он содержит ссылку на адаптируемый объект (adaptee)
	client *telegram.Telegram
}

// NewTelegramAdapter конструктор, который возвращает общий интерфейс
func NewTelegramAdapter(telegramClient *telegram.Telegram) messenger.Messenger {
	return &TelegramAdapter{client: telegramClient}
}

// SendMessage реализует целевой интерфейс
func (ta *TelegramAdapter) SendMessage(userID, message string) error {
	// и делегирует вызов адаптируемому объекту
	err := ta.client.Send(userID, message)
	if err != nil {
		return fmt.Errorf("telegram adapter: can't send message: %v", err)
	}

	return nil
}

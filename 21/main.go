package main

import (
	"level-1/21/adapters"
	"level-1/21/messenger"
	"level-1/21/telegram"
	"level-1/21/whatsapp"
	"log"
)

func main() {
	// создаем срез типа интерфейса, который будет содержать наши адаптеры
	notifiers := []messenger.Messenger{
		adapters.NewTelegramAdapter(telegram.NewClient()),
		adapters.NewWhatsAppAdapter(whatsapp.New()),
	}

	userID, message := "1234", "Hello, user!"

	// вызываем клиентский код, который ничего не знает о конкретных реализациях
	eventHandler(notifiers, userID, message)
}

// eventHandler это наш клиент, который работает с любым типом, удовлетворяющим интерфейсу
func eventHandler(notifiers []messenger.Messenger, userID, message string) {
	// полиморфно вызываем метод для каждого элемента в срезе
	for _, notifier := range notifiers {
		err := notifier.SendMessage(userID, message)
		if err != nil {
			log.Printf("ERROR: can't send message: %v\n", err)
		}
	}
}

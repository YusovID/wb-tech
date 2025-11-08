package adapters

import (
	"fmt"
	"level-1/21/messenger"
	"level-1/21/whatsapp"
	"strconv"
)

// WhatsAppAdapter это адаптер для клиента WhatsApp
type WhatsAppAdapter struct {
	// он также содержит ссылку на свой адаптируемый объект
	client *whatsapp.WhatsAppClient
}

// NewWhatsAppAdapter конструктор для WhatsApp адаптера
func NewWhatsAppAdapter(whatsAppClient *whatsapp.WhatsAppClient) messenger.Messenger {
	return &WhatsAppAdapter{client: whatsAppClient}
}

// SendMessage реализует тот же самый интерфейс
func (wa *WhatsAppAdapter) SendMessage(userID, message string) error {
	// здесь происходит основная работа адаптера – преобразование данных
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return fmt.Errorf("invalid user id '%s': %v", userID, err)
	}

	// вызов метода адаптируемого объекта с преобразованными данными
	resp, err := wa.client.Message(int64(userIDInt), message)
	if resp.Error != nil {
		return fmt.Errorf("error in response: %v", resp.Error)
	}

	if err != nil {
		return fmt.Errorf("whatsapp adapter: can't send message: %v", err)
	}

	return nil
}

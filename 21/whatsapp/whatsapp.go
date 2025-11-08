package whatsapp

import "log"

// WhatsAppClient это наш второй несовместимый сервис (adaptee)
type WhatsAppClient struct {
}

// Response его собственная структура ответа
type Response struct {
	Body  string
	Error error
}

// New его конструктор
func New() *WhatsAppClient {
	return &WhatsAppClient{}
}

// Message имеет совершенно другую сигнатуру, несовместимую с целевым интерфейсом
func (wa *WhatsAppClient) Message(userID int64, text string) (Response, error) {
	// some implementation

	log.Printf("whatsapp: message was sent to user %d", userID)

	return Response{}, nil
}

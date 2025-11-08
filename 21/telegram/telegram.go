package telegram

import "log"

// Telegram это наш первый несовместимый сервис (adaptee)
type Telegram struct{}

// NewClient его конструктор
func NewClient() *Telegram {
	return &Telegram{}
}

// Send имеет сигнатуру, которая нас устраивает, но он находится во внешней системе
func (t *Telegram) Send(userID string, text string) error {
	// some implementation

	log.Printf("telegram: message was send to user %s", userID)

	return nil
}

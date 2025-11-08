package messenger

// Messenger это наш целевой интерфейс, который ожидают клиенты
type Messenger interface {
	SendMessage(userID, message string) error
}

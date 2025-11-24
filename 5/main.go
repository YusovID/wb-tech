package main

// структура customError соответствует интерфейсу error, т.к. имеет метод Error()
type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

// в функции test мы возвращаем указатель на структуру customError, значение которой равно nil
func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error
	err = test()
	// при проверке, равен ли интерфейс err nil, условие выполнится,
	// т.к. интерфейс содержит динамический тип customError, а не nil.
	// Чтобы условие выполнилось, надо чтобы и динамический тип интерфейса, и его значение были nil,
	// поэтому выведется error
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

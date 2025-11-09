package main

import (
	"fmt"
	"os"
)

// Foo возвращает тип error. Внутри нее инициализируется переменная err типа *os.PathError,
// который реализует интерфейс error. err мы присваивам nil и возвращаем
func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	// мы записываем то, что вернула функция Foo (nil)
	err := Foo()
	fmt.Println(err)        // nil
	fmt.Println(err == nil) // false
	// при сравнении err с nil, мы пытаемся сравнить не значение, а интерфейс,
	// т.к. в интерфейсе хранится о его типе и значении, сравнение с nil даст true только тогда,
	// когда и тип, и значение равны nil. В данном случае тип - *os.PathError, поэтому вывод будет false

	// продемонстрируем на этом примере. Переменная err1 имеет тип interface{}, который инициализируется
	// с zero value (type: nil, value: nil), а значит сравнение с nil даст true
	var err1 interface{}
	fmt.Println(err1 == nil)
}

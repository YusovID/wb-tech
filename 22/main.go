package main

import (
	"errors"
	"fmt"
	"math/big"
)

var (
	// ErrDivisionByZero возвращается при попытке деления на ноль
	ErrDivisionByZero = errors.New("деление на ноль")
	// ErrUnknownType возвращается, если тип числовой переменной не поддерживается
	ErrUnknownType = errors.New("неподдерживаемый числовой тип")
)

// BigNumberPtr определяет обобщенный тип для указателей на числа из пакета math/big
type BigNumberPtr interface {
	*big.Int | *big.Rat | *big.Float
}

func main() {
	testBigInt()
	testBigFloat()
	testBigRat()
}

// testBigInt демонстрирует арифметические операции с типом *big.Int
func testBigInt() {
	fmt.Println("\n--- Тестирование big.Int ---")

	bigIntA, ok := new(big.Int).SetString("2400000000000000000000000", 10)
	if !ok {
		fmt.Println("testBigInt: не удалось создать bigIntA")
		return
	}

	bigIntB, ok := new(big.Int).SetString("760000000000000000000000", 10)
	if !ok {
		fmt.Println("testBigInt: не удалось создать bigIntB")
		return
	}

	fmt.Printf("A: %s\n", bigIntA.String())
	fmt.Printf("B: %s\n", bigIntB.String())
	fmt.Println("------------------------------")

	sumResult, err := Sum(bigIntA, bigIntB)
	if err != nil {
		fmt.Printf("Ошибка сложения: %v\n", err)
		return
	}
	fmt.Printf("Сумма:        %s\n", sumResult.String())

	diffResult, err := Diff(bigIntA, bigIntB)
	if err != nil {
		fmt.Printf("Ошибка вычитания: %v\n", err)
		return
	}
	fmt.Printf("Разность:     %s\n", diffResult.String())

	mulResult, err := Multiply(bigIntA, bigIntB)
	if err != nil {
		fmt.Printf("Ошибка умножения: %v\n", err)
		return
	}
	fmt.Printf("Произведение: %s\n", mulResult.String())

	divResult, err := Divide(bigIntA, bigIntB)
	if err != nil {
		if errors.Is(err, ErrDivisionByZero) {
			fmt.Printf("Попытка деления на ноль: %v\n", err)
		} else {
			fmt.Printf("Ошибка деления: %v\n", err)
		}
		return
	}
	fmt.Printf("Частное:      %s\n", divResult.String())
}

// testBigFloat демонстрирует арифметические операции с типом *big.Float
func testBigFloat() {
	fmt.Println("\n--- Тестирование big.Float ---")

	bigFloatA, ok := new(big.Float).SetString("1234567890.123456789")
	if !ok {
		fmt.Println("testBigFloat: не удалось создать bigFloatA")
		return
	}

	bigFloatB, ok := new(big.Float).SetString("0.000000000987654321")
	if !ok {
		fmt.Println("testBigFloat: не удалось создать bigFloatB")
		return
	}

	fmt.Printf("A: %s\n", bigFloatA.Text('f', -1))
	fmt.Printf("B: %s\n", bigFloatB.Text('f', -1))
	fmt.Println("------------------------------")

	sumResult, err := Sum(bigFloatA, bigFloatB)
	if err != nil {
		fmt.Printf("Ошибка сложения: %v\n", err)
		return
	}
	fmt.Printf("Сумма:        %s\n", sumResult.Text('f', 10))

	diffResult, err := Diff(bigFloatA, bigFloatB)
	if err != nil {
		fmt.Printf("Ошибка вычитания: %v\n", err)
		return
	}
	fmt.Printf("Разность:     %s\n", diffResult.Text('f', 10))

	mulResult, err := Multiply(bigFloatA, bigFloatB)
	if err != nil {
		fmt.Printf("Ошибка умножения: %v\n", err)
		return
	}
	fmt.Printf("Произведение: %s\n", mulResult.Text('f', 10))

	divResult, err := Divide(bigFloatA, bigFloatB)
	if err != nil {
		if errors.Is(err, ErrDivisionByZero) {
			fmt.Printf("Попытка деления на ноль: %v\n", err)
		} else {
			fmt.Printf("Ошибка деления: %v\n", err)
		}
		return
	}
	fmt.Printf("Частное:      %s\n", divResult.Text('f', 10))
}

// testBigRat демонстрирует арифметические операции с типом *big.Rat
func testBigRat() {
	fmt.Println("\n--- Тестирование big.Rat (Рациональные числа) ---")

	bigRatA := big.NewRat(2, 3)
	bigRatB := big.NewRat(1, 6)

	fmt.Printf("A: %s\n", bigRatA.RatString())
	fmt.Printf("B: %s\n", bigRatB.RatString())
	fmt.Println("------------------------------")

	sumResult, err := Sum(bigRatA, bigRatB)
	if err != nil {
		fmt.Printf("Ошибка сложения: %v\n", err)
		return
	}
	fmt.Printf("Сумма (2/3 + 1/6):        %s\n", sumResult.RatString())

	diffResult, err := Diff(bigRatA, bigRatB)
	if err != nil {
		fmt.Printf("Ошибка вычитания: %v\n", err)
		return
	}
	fmt.Printf("Разность (2/3 - 1/6):     %s\n", diffResult.RatString())

	mulResult, err := Multiply(bigRatA, bigRatB)
	if err != nil {
		fmt.Printf("Ошибка умножения: %v\n", err)
		return
	}
	fmt.Printf("Произведение (2/3 * 1/6): %s\n", mulResult.RatString())

	divResult, err := Divide(bigRatA, bigRatB)
	if err != nil {
		if errors.Is(err, ErrDivisionByZero) {
			fmt.Printf("Попытка деления на ноль: %v\n", err)
		} else {
			fmt.Printf("Ошибка деления: %v\n", err)
		}
		return
	}
	fmt.Printf("Частное (2/3 / 1/6):      %s\n", divResult.RatString())
}

// Sum складывает два больших числа одного типа
func Sum[P BigNumberPtr](a, b P) (P, error) {
	// преобразуем обобщенный тип P в any для проверки его конкретного типа
	switch x := any(a).(type) {
	case *big.Int:
		// приводим b к конкретному типу *big.Int, так как он должен совпадать с типом a
		y := any(b).(*big.Int)
		result := new(big.Int)
		result.Add(x, y)
		// преобразуем конкретный результат обратно в обобщенный тип P для возврата
		return any(result).(P), nil
	case *big.Float:
		y := any(b).(*big.Float)
		result := new(big.Float)
		result.Add(x, y)
		return any(result).(P), nil
	case *big.Rat:
		y := any(b).(*big.Rat)
		result := new(big.Rat)
		result.Add(x, y)
		return any(result).(P), nil
	default:
		return nil, ErrUnknownType
	}
}

// Diff вычитает одно большое число из другого
func Diff[P BigNumberPtr](a, b P) (P, error) {
	// преобразуем обобщенный тип P в any для проверки его конкретного типа
	switch x := any(a).(type) {
	case *big.Int:
		// приводим b к конкретному типу *big.Int, так как он должен совпадать с типом a
		y := any(b).(*big.Int)
		result := new(big.Int)
		result.Sub(x, y)
		// преобразуем конкретный результат обратно в обобщенный тип P для возврата
		return any(result).(P), nil
	case *big.Float:
		y := any(b).(*big.Float)
		result := new(big.Float)
		result.Sub(x, y)
		return any(result).(P), nil
	case *big.Rat:
		y := any(b).(*big.Rat)
		result := new(big.Rat)
		result.Sub(x, y)
		return any(result).(P), nil
	default:
		return nil, ErrUnknownType
	}
}

// Multiply перемножает два больших числа
func Multiply[P BigNumberPtr](a, b P) (P, error) {
	// преобразуем обобщенный тип P в any для проверки его конкретного типа
	switch x := any(a).(type) {
	case *big.Int:
		// приводим b к конкретному типу *big.Int, так как он должен совпадать с типом a
		y := any(b).(*big.Int)
		result := new(big.Int)
		result.Mul(x, y)
		// преобразуем конкретный результат обратно в обобщенный тип P для возврата
		return any(result).(P), nil
	case *big.Float:
		y := any(b).(*big.Float)
		result := new(big.Float)
		result.Mul(x, y)
		return any(result).(P), nil
	case *big.Rat:
		y := any(b).(*big.Rat)
		result := new(big.Rat)
		result.Mul(x, y)
		return any(result).(P), nil
	default:
		return nil, ErrUnknownType
	}
}

// IsZeroer определяет интерфейс для типов, которые можно проверить на равенство нулю
type IsZeroer interface {
	Sign() int
}

// Divide делит одно большое число на другое с проверкой делителя
func Divide[P interface {
	*big.Int | *big.Float | *big.Rat
	IsZeroer
}](a, b P) (P, error) {
	// преобразуем b в интерфейс IsZeroer для вызова метода Sign
	if any(b).(IsZeroer).Sign() == 0 {
		var zero P
		return zero, ErrDivisionByZero
	}

	// преобразуем обобщенный тип P в any для проверки его конкретного типа
	switch x := any(a).(type) {
	case *big.Int:
		// приводим b к конкретному типу *big.Int, так как он должен совпадать с типом a
		y := any(b).(*big.Int)
		result := new(big.Int)
		result.Quo(x, y)
		// преобразуем конкретный результат обратно в обобщенный тип P для возврата
		return any(result).(P), nil
	case *big.Float:
		y := any(b).(*big.Float)
		result := new(big.Float)
		result.Quo(x, y)
		return any(result).(P), nil
	case *big.Rat:
		y := any(b).(*big.Rat)
		result := new(big.Rat)
		result.Quo(x, y)
		return any(result).(P), nil
	default:
		return nil, ErrUnknownType
	}
}

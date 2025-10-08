package syncmap

import (
	"fmt"
	"math"
	"sync"
)

// SyncMap - это обертка над стандартной map, обеспечивающая потокобезопасность
// операций чтения и записи с помощью sync.RWMutex.
type SyncMap struct {
	data map[int][]string
	// RWMutex позволяет выполнять множество одновременных чтений или одну запись.
	mu *sync.RWMutex
	// Поля для сбора статистики.
	maxValueSize int
	minValueSize int
	avgValueSize float64
	maxKey       int
	minKey       int
}

// New - конструктор для создания экземпляра SyncMap.
func New() *SyncMap {
	return &SyncMap{
		data:         make(map[int][]string, 0),
		mu:           &sync.RWMutex{},
		maxValueSize: math.MinInt,
		minValueSize: math.MaxInt,
	}
}

// Write - потокобезопасный метод для записи данных в map.
func (s *SyncMap) Write(key int, value string) {
	// Lock() устанавливает эксклюзивную блокировку.
	// В этот момент никакая другая горутина не может ни читать, ни писать.
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = append(s.data[key], value)
}

// Read - потокобезопасный метод для чтения данных из map.
func (s *SyncMap) Read(key int) []string {
	// RLock() устанавливает блокировку на чтение.
	// Другие горутины могут одновременно читать данные, но не могут их изменять.
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.data[key]
	if !ok {
		return nil
	}

	return value
}

// Stat выводит собранную статистику по данным в map.
func (s *SyncMap) Stat() {
	// Сначала вычисляем статистику...
	s.fillStat()

	// ...затем выводим.
	fmt.Println("-------------STATISTIC-----------------")
	fmt.Printf("Size of map: %d\n", len(s.data))
	fmt.Printf("Key %d has max slice of values size: %d\n", s.maxKey, s.maxValueSize)
	fmt.Printf("Key %d has min slice of values size: %d\n", s.minKey, s.minValueSize)
	fmt.Printf("Average slice of values size: %.2f\n", s.avgValueSize)
}

// fillStat - потокобезопасный метод для сбора статистики.
func (s *SyncMap) fillStat() {
	// Используем RLock для безопасного чтения данных во время сбора статистики.
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Проверка на случай, если map осталась пустой, во избежание деления на ноль.
	if len(s.data) == 0 {
		s.avgValueSize = 0
		return
	}

	fullValueSize := 0
	isFirst := true

	// Проходим по всем элементам map для вычисления метрик.
	for key, value := range s.data {
		valLen := len(value)

		// Корректная инициализация начальных значений для min/max.
		if isFirst {
			s.maxValueSize = valLen
			s.minValueSize = valLen
			s.maxKey = key
			s.minKey = key
			isFirst = false
		}

		if len(value) > s.maxValueSize {
			s.maxKey = key
			s.maxValueSize = len(value)
		} else if len(value) < s.minValueSize {
			s.minKey = key
			s.minValueSize = len(value)
		}

		fullValueSize += len(value)
	}

	s.avgValueSize = float64(fullValueSize) / float64(len(s.data))
}

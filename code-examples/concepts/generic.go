package concepts

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) insert(key K, value V) {
	m.data[key] = value
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{data: make(map[K]V)}
}

func foo[T any](val T) {
	fmt.Println(val)
}

func Generics() {
	m1 := NewCustomMap[string, int]()
	m1.insert("one", 1)
	m1.insert("two", 2)
	fmt.Printf("m1: %+v\n", m1)

	m2 := NewCustomMap[int, float64]()
	m2.insert(1, 1.1)
	m2.insert(2, 2.2)
	fmt.Printf("m2: %+v\n", m2)

	foo[int](1)
	foo[string]("hello")
}

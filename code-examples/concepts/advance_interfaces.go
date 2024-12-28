package concepts

import (
	"fmt"
	"strings"
)

type Putter interface {
	Put(int, any) (any, error)
}

type Storage interface {
	Putter
	Get(int) (any, error)
}

type FooStorage struct{}

func (f *FooStorage) Get(id int) (any, error) {
	return nil, nil
}
func (f *FooStorage) Put(id int, any any) (any, error) {
	return nil, nil
}

type CustomServer struct {
	store Storage
}

func updateValue(id int, value any, p Putter) {
	p.Put(id, value)
}

/**
 * Transform function type
 */
type TransformFunc func(string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Lowercase(s string) string {
	return strings.ToLower(s)
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func InterfacesV3() {
	s := &CustomServer{store: &FooStorage{}}

	updateValue(1, "foo", s.store)

	fmt.Println(transformString("some value", Uppercase))
	fmt.Println(transformString("some value", Lowercase))
	fmt.Println(transformString("some value", Prefixer("FOO_")))
}

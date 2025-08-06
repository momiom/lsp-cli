// Package main contains type hierarchy test cases
package main

import "fmt"

// BaseInterface is the root interface
type BaseInterface interface {
	BaseMethod() string
}

// Interface1 is a simple interface with generic
type Interface1[T any] interface {
	Method1(T) error
}

// Interface2 is another interface
type Interface2 interface {
	Method2() int
}

// ExtendedInterface extends multiple interfaces
type ExtendedInterface[T any] interface {
	BaseInterface
	Interface1[T]
	ExtendedMethod() bool
}

// BaseClass is a base struct
type BaseClass[T any] struct {
	ID    int
	Value T
}

// BaseMethod implements BaseInterface
func (b BaseClass[T]) BaseMethod() string {
	return fmt.Sprintf("BaseClass[%v]", b.Value)
}

// SimpleChild embeds BaseClass with concrete type
type SimpleChild struct {
	BaseClass[string]
	Name string
}

// MultipleInterfaces implements multiple interfaces
type MultipleInterfaces struct {
	data string
}

// Method1 implements Interface1
func (m MultipleInterfaces) Method1(v any) error {
	return nil
}

// Method2 implements Interface2
func (m MultipleInterfaces) Method2() int {
	return 42
}

// ComplexChild embeds BaseClass and implements multiple interfaces
type ComplexChild[T any, U comparable] struct {
	BaseClass[T]
	SecondValue U
}

// Method1 implements Interface1
func (c ComplexChild[T, U]) Method1(v U) error {
	return nil
}

// Method2 implements Interface2
func (c ComplexChild[T, U]) Method2() int {
	return len(fmt.Sprint(c.SecondValue))
}

// KitchenSink embeds and implements everything
type KitchenSink[T any] struct {
	BaseClass[T]
	ComplexChild[T, string]
	additionalData map[string]interface{}
}

// ExtendedMethod implements ExtendedInterface
func (k KitchenSink[T]) ExtendedMethod() bool {
	return true
}

// Method1 implements Interface1
func (k KitchenSink[T]) Method1(v T) error {
	return nil
}

// Method2 implements Interface2
func (k KitchenSink[T]) Method2() int {
	return len(k.additionalData)
}

// FunctionType is a function type alias
type FunctionType func(int) string

// DummyFunc1 separates type definitions to avoid gopls merging them
func DummyFunc1() {}

// HandlerFunc is a more complex function type
type HandlerFunc[T any, R any] func(T) (R, error)

// DummyFunc2 separates type definitions to avoid gopls merging them
func DummyFunc2() {}

// ChannelType is a channel type alias
type ChannelType chan<- string

// DummyFunc3 separates type definitions to avoid gopls merging them
func DummyFunc3() {}

// MapType is a map type alias  
type MapType[K comparable, V any] map[K]V

// DummyFunc4 separates type definitions to avoid gopls merging them
func DummyFunc4() {}

// SliceType is a slice type alias
type SliceType[T any] []T
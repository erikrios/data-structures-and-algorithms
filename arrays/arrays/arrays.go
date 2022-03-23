package arrays

import (
	"errors"
	"fmt"
	"strings"
)

type array[T comparable] struct {
	items []T
}

func NewArray[T comparable]() *array[T] {
	return &array[T]{}
}

func (a *array[T]) String() string {
	builder := strings.Builder{}
	builder.WriteByte('[')

	for i := 0; i < len(a.items); i++ {
		builder.WriteString(fmt.Sprintf("%v", a.items[i]))
		builder.WriteByte(',')
	}

	builder.WriteByte(']')

	if builder.Len() > 2 {
		s := builder.String()[:builder.Len()-2]
		s += "]"
		return string(s)
	} else {
		return builder.String()
	}
}

func (a *array[T]) Print() {
	fmt.Println(a)
}

func (a *array[T]) Insert(item T) {
	a.items = append(a.items, item)
}

func (a *array[T]) RemoveAt(index int) (err error) {
	if index < 0 || index > len(a.items)-1 {
		err = errors.New("index out of bounds")
		return
	}

	a.items = append(a.items[:index], a.items[index+1:]...)
	return
}

func (a *array[T]) IndexOf(item T) (index int) {
	index = -1

	for i := 0; i < len(a.items); i++ {
		if a.items[i] == item {
			index = i
			return
		}
	}

	return
}

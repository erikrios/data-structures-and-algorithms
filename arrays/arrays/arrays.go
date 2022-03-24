package arrays

import (
	"errors"
	"fmt"
	"strings"
)

type NumOrFloat interface {
	int64 | float64
}

type array[T NumOrFloat] struct {
	items []T
}

func NewArray[T NumOrFloat]() *array[T] {
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

func (a *array[T]) Max() (max T, err error) {
	if len(a.items) == 0 {
		err = errors.New("array is empty")
		return
	}

	max = a.items[1]

	for i := 0; i < len(a.items); i++ {
		if max < a.items[i] {
			max = a.items[i]
		}
	}

	return
}

func (a *array[T]) Intersect(b *array[T]) (intersects *array[T]) {
	intersects = NewArray[T]()
	uniqueIntersects := make(map[T]interface{})

	for i := 0; i < len(a.items); i++ {
		for j := 0; j < len(b.items); j++ {
			if a.items[i] == b.items[j] {
				uniqueIntersects[a.items[i]] = nil
			}
		}
	}

	for k, _ := range uniqueIntersects {
		intersects.items = append(intersects.items, k)
	}

	return
}

func (a *array[T]) Reverse() {
	for i := 0; i < len(a.items)/2; i++ {
		a.items[i], a.items[len(a.items)-1-i] = a.items[len(a.items)-1-i], a.items[i]
	}
}

func (a *array[T]) InsertAt(item T, index int) {
	if index >= len(a.items) {
		a.Insert(item)
		return
	}

	temp := make([]T, len(a.items[index:]))
	copy(temp, a.items[index:])
	a.items = append(a.items[:index], item)
	a.items = append(a.items, temp...)
}

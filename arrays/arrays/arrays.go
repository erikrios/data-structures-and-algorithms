package arrays

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type array struct {
	items []int
}

func NewArray() *array {
	return &array{}
}

func (a *array) String() string {
	builder := strings.Builder{}
	builder.WriteByte('[')

	for i := 0; i < len(a.items); i++ {
		builder.WriteString(strconv.Itoa(a.items[i]))
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

func (a *array) Print() {
	fmt.Println(a)
}

func (a *array) Insert(item int) {
	a.items = append(a.items, item)
}

func (a *array) RemoveAt(index int) (err error) {
	if index < 0 || index > len(a.items)-1 {
		err = errors.New("index out of bounds")
		return
	}

	a.items[index] = 0
	a.items = append(a.items[:index], a.items[index+1:]...)
	return
}

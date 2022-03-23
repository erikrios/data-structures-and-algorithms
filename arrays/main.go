package main

import (
	"fmt"

	"github.com/erikrios/arrays/arrays"
)

func main() {
	numbers := arrays.NewArray[int]()
	numbers.Insert(1)
	numbers.Insert(2)
	numbers.Insert(3)
	numbers.Insert(4)
	numbers.Print()
	if err := numbers.RemoveAt(3); err != nil {
		fmt.Println(err)
	} else {
		numbers.Print()
	}

	fmt.Println(numbers)

	fmt.Println(numbers.IndexOf(1))
	fmt.Println(numbers.IndexOf(4))
}

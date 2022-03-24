package main

import (
	"fmt"

	"github.com/erikrios/arrays/arrays"
)

func main() {
	numbers := arrays.NewArray[int64]()
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
	fmt.Println(numbers.Max())

	newNumbers := arrays.NewArray[int64]()
	newNumbers.Insert(2)
	newNumbers.Insert(3)

	fmt.Println(numbers.Intersect(newNumbers))

	numbers.Reverse()
	fmt.Println(numbers)
	newNumbers.Reverse()
	fmt.Println(newNumbers)

	newNums := arrays.NewArray[int64]()
	newNums.Insert(10)
	newNums.Insert(3)
	newNums.Insert(2)
	newNums.Insert(18)
	newNums.Insert(12)
	fmt.Println(newNums)

	newNums.InsertAt(5, 2)
	fmt.Println(newNums)
}

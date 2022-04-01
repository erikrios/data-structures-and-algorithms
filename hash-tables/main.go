package main

import (
	"fmt"

	"github.com/hash-tables/exercise"
)

func main() {
	fmt.Println(string(exercise.FirstNonRepeatingChar("a green apple")))
	fmt.Println(string(exercise.FirstRepeatedChar("a green apple")))
}

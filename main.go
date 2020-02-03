package main

import (
	"fmt"
	"summing_nums/tools"
)

func main() {
	//r := tools.GetNumbersWhichSums([]int{1, 2, 3, 6, 9, 12, -3, -6, -9, 5, 7, 4, -1, -2}, 3)
	r := tools.GetNumbersWhichSums([]int{-1, 1, 1, 2, 2}, 3)
	fmt.Printf("%v\n", r)
}

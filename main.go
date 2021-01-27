package main

import (
	"algorithms/arrays"
	"fmt"
)

func main() {
	fmt.Println("Main file")
	fmt.Println(arrays.RemoveIntFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(arrays.MergeSortedArrays([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(arrays.FindSumOfTwo([]int{1, 21, 3, 14, 5, 60, 7, 6}, 27))
}

package main

import (
	"algorithms/arrays"
	"algorithms/lists"
	"algorithms/stack"
	"fmt"
)

func main() {
	fmt.Println("Main file")
	fmt.Println(arrays.RemoveIntFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(arrays.MergeSortedArrays([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(arrays.FindSumOfTwo([]int{1, 21, 3, 14, 5, 60, 7, 6}, 27))

	linkedList := lists.LinkedList{}
	linkedList.Insert(1)
	linkedList.Insert(2)
	linkedList.Insert(3)
	linkedList.Insert(3333)
	linkedList.InsertAtEnd(22)

	linkedList.Display()

	stack := stack.NewStack(10)
	stack.Push(1)
	stack.Push(2)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}

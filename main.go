package main

import (
	"fmt"
	"leetCode/LinkedList"
	"leetCode/Stack"
)

func main() {
	//testMinStack()
	//testReversePolishNotation()
	//testDailyTemperature()
	testFindDuplicatedConstantSpace()
}

func testReversePolishNotation() {
	fmt.Println(Stack.Evaluate([]string{"4", "13", "5", "/", "+"}))
}

func testDailyTemperature() {
	fmt.Println(Stack.DailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}))
}

func testFindDuplicatedConstantSpace() {
	fmt.Println(LinkedList.FindDuplicatesConstantSpace([]int{3, 1, 3, 4, 2}))
}

func testMinStack() {
	obj := Stack.Constructor()
	obj.Push(2)
	obj.Push(0)
	obj.Push(3)
	obj.Push(0)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
}

func testListAddTwoNumbers() {
	list2 := LinkedList.BuildLinkedList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	list2s := LinkedList.BuildLinkedList([]int{5, 6, 4})
	node := list2
	nodeS := list2s
	node = LinkedList.AddTwoNumbers(node, nodeS)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

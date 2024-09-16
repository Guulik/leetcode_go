package main

import (
	"fmt"
	"leetCode/LinkedList"
	"leetCode/Stack"
)

func main() {
	/*//fmt.Println(arrays.ContainsDuplicate([]int{1, 2, 3}))

	//fmt.Println(arrays.TwoSum([]int{3, 2, 4}, 6))

	//fmt.Println(Arrays_and_Hashing.GroupAnagram([]string{"eaat", "teaa", "tan", "ate", "nat", "bat"}))

	//fmt.Println(arrays.ProductExceptSelf([]int{1, 2, 3, 4}))

	//fmt.Println(Arrays_and_Hashing.LongestSequence([]int{100, 4, 200, 1, 3, 2}))

	//fmt.Println(stacking.ValidParentheses("("))
	//list5 := &LinkedList.ListNode{1,
	//	&LinkedList.ListNode{2,
	//		&LinkedList.ListNode{3,
	//			&LinkedList.ListNode{4,
	//				&LinkedList.ListNode{5,
	//					nil,
	//				},
	//			},
	//		},
	//	},
	//}

	//list1 := &LinkedList.ListNode{1, &LinkedList.ListNode{2, &LinkedList.ListNode{4, nil}}}
	//singleList := &LinkedList.ListNode{1, nil}
	//node := LinkedList.MergeSortedLists(list1, list2)
	//for node.Next != nil {
	//	fmt.Println(node.Val)
	//	node = node.Next
	//}

	fmt.Println(Arrays_and_Hashing.IsValidSudoku([][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	},
	))

	//fmt.Println(SlidingWindow.BestTimeToTrade([]int{7, 1, 5, 3, 6, 4}))

	//fmt.Println(TwoPointers.Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))

	//fmt.Println(TwoPointers.ThreeSum([]int{-3, 3, 4, -3, 1, 2}))
	//fmt.Println(TwoPointers.TwoSumSorted([]int{2, 7, 11, 15}, 9))
	//
	//fmt.Println(Arrays_and_Hashing.TwoSum([]int{2, 7, 11, 15}, 9))

	//node := LinkedList.ReverseLinkedList(list1)
	//LinkedList.ReorderList(node)
	//node := singleList
	*/

	//testMinStack()
	//testReversePolishNotation()
	testDailyTemperature()
}

func testReversePolishNotation() {
	fmt.Println(Stack.Evaluate([]string{"4", "13", "5", "/", "+"}))
}

func testDailyTemperature() {
	fmt.Println(Stack.DailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}))
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

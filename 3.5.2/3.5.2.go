package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func createList(sl []int) *ListNode {
	var first, second *ListNode

	for _, value := range sl {
		if first != nil {
			second.Next = &ListNode{Val: value}
			second = second.Next
		} else {
			first = &ListNode{Val: value}
			second = first
		}
	}
	
	return first
}

func (list *ListNode) IsEmpty() bool {
    return list.Next != nil
}

func deleteDuplicates(l *ListNode) *ListNode {
	check := l
	var new_val *ListNode

	for check.Next != nil {
		if check.Val != check.Next.Val {
			check = check.Next
		} else {
			new_val = check.Next.Next
			check.Next = new_val
		}
	}

	return l
}

func main() {
	listHead:=createList([]int{1,2,3,3,5,5,6})
	printer(listHead)
	fmt.Println("\n")
	listHead=deleteDuplicates(listHead)

}

func printer(first *ListNode) {
	for first.IsEmpty() {
		fmt.Printf("%v ", first.Val)
		first = first.Next
	}
	fmt.Printf("%v ", first.Val)
}
package main

import (
	"fmt"
	"reflect"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func reverseArray(arr interface{}) {
	length := reflect.ValueOf(arr).Len()
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		valI := reflect.ValueOf(arr).Index(i).Interface()
		valJ := reflect.ValueOf(arr).Index(j).Interface()
		reflect.ValueOf(arr).Index(i).Set(reflect.ValueOf(valJ))
		reflect.ValueOf(arr).Index(j).Set(reflect.ValueOf(valI))
	}
}

//--------------------------IDEA_1

func reversNumber(arr []int) uint64 {
	var temp uint64 = 0
	for i := len(arr) - 1; i >= 0; i-- {
		temp = uint64(arr[i]) + temp*10
	}

	return temp
}

func idea_1(l1 []int, l2 []int) {
	fmt.Println(reversNumber(l1) + reversNumber(l2))
}

////////////////////////////////////

//--------------------------IDEA_2

// implement LinkedList
type node struct { //define a node struct
	value int
	next  *node
}

func add_node(head *node, value int) *node {
	newNode := &node{value: value}
	if head == nil {
		head = newNode
	} else {
		current := head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	return head
}

////////////////////

/////implement Stack

type Stack []int //stack

func (st *Stack) Push(v int) {
	*st = append(*st, v)
}

func (st *Stack) Pop() int {
	if st.IsEmpty() {
		return 0
	}
	top := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return top
}

func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

////////////////////

func main() {

	var head_1 *node
	head_1 = add_node(head_1, 9)
	head_1 = add_node(head_1, 9)
	head_1 = add_node(head_1, 9)
	head_1 = add_node(head_1, 9)
	head_1 = add_node(head_1, 9)
	head_1 = add_node(head_1, 9)
	head_1 = add_node(head_1, 9)
	// head_1 = add_node(head_1, 4)

	var head_2 *node
	head_2 = add_node(head_2, 9)
	head_2 = add_node(head_2, 9)
	head_2 = add_node(head_2, 9)
	head_2 = add_node(head_2, 9)
	// head_2 = add_node(head_2, 1)

	st_1 := Stack{}
	st_2 := Stack{}

	current_1 := head_1
	for current_1 != nil {
		st_1.Push(current_1.value)
		current_1 = current_1.next
	}

	current_2 := head_2
	for current_2 != nil {
		st_2.Push(current_2.value)
		current_2 = current_2.next
	}

	temp1 := 0
	temp2 := 0
	carry := 0

	var head_3 *node
	for (!st_1.IsEmpty()) || (!st_2.IsEmpty()) {

		if st_1.IsEmpty() {
			temp1 = 0
		} else {
			a := st_1.Pop()
			// fmt.Println(a)
			temp1 = a

		}

		if st_2.IsEmpty() {
			temp2 = 0
		} else {
			temp2 = st_2.Pop()
		}
		res := temp1 + temp2 + carry
		// fmt.Println(res)

		head_3 = add_node(head_3, res%10)
		carry = res / 10
		// fmt.Println(carry)

	}
	if carry > 0 {
		head_3 = add_node(head_3, carry)
	}

	current_3 := head_3
	for current_3 != nil {
		fmt.Println(current_3.value)
		current_3 = current_3.next
	}

}

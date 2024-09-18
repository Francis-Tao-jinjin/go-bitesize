package main

import (
	"fmt"
)

func IndexOf[T comparable](s []T, e T) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next  *List[T]
	value T
}

func (li *List[T]) InsertFront(value T) *List[T] {
	return &List[T]{next: li, value: value}
}

func TryGenerics() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(IndexOf(s, 3))
	fmt.Println(IndexOf(s, 6))

	s2 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(IndexOf(s2, "c"))
	fmt.Println(IndexOf(s2, "f"))

	// li := &List[int]{value: 1}
	// li = li.InsertFront(2)
	// li = li.InsertFront(3)
	// fmt.Println(li)

	// for cur := li; cur != nil; cur = cur.next {
	// 	fmt.Println(cur.value)
	// }

	// this will enter an infinite loop
	// li := List[int]{value: 1}
	// li = *li.InsertFront(2)
	// li = *li.InsertFront(3)
	// for cur := &li; cur != nil; cur = cur.next {
	// 	fmt.Println(cur.value)
	// }

	// this will work as the li is a pointer
	var li *List[int]
	li = li.InsertFront(1)
	li = li.InsertFront(2)
	li = li.InsertFront(3)

	for cur := li; cur != nil; cur = cur.next {
		fmt.Println(cur.value)
	}
}

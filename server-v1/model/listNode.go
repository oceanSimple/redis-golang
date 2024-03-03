package model

type ListNode struct {
	// Next and previous pointers in the doubly-linked list of elements.
	prev, next *ListNode
	// The list to which this element belongs.
	list *List
	// The value stored with this element.
	Value interface{}
}

type ListNodeInterface interface {
	// Next Returns the next listNode or nil
	Next() *ListNode
	// Prev Returns the previous listNode or nil
	Prev() *ListNode
}

func (l *ListNode) Next() *ListNode {
	p := l.next
	// 1. p must belong to the list
	// 2. p must not be the root
	if p.list != nil && p != &p.list.root {
		return p
	}
	return nil
}

func (l *ListNode) Prev() *ListNode {
	p := l.prev
	// 1. p must belong to the list
	// 2. p must not be the root
	if p.list != nil && p != &p.list.root {
		return p
	}
	return nil
}

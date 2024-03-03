package model

import "fmt"

type ListInterface interface {
	// Init create an empty list
	Init() *List
	// Len returns the number of elements of list l
	Len() int
	// Front returns the first element of list l or nil if the list is empty
	Front() *ListNode
	// Back returns the last element of list l or nil if the list is empty
	Back() *ListNode
	// Remove removes e from its list, decrements l.len
	Remove(e *ListNode) any
	// PushFront inserts a new element e with value v at the front of list l and returns e
	PushFront(v any) *ListNode
	// PushBack inserts a new element e with value v at the back of list l and returns e
	PushBack(v any) *ListNode
	// InsertBefore inserts a new element e with value v immediately before mark and returns e
	InsertBefore(v any, mark *ListNode) *ListNode
	// InsertAfter inserts a new element e with value v immediately after mark and returns e
	InsertAfter(v any, mark *ListNode) *ListNode

	// lazyInit call Init() if l is nil
	// it will be called in other methods to make sure the list is initialized
	lazyInit()
	// insert inserts e after at, increments l.len, and returns e
	insert(e, at *ListNode) *ListNode
	// insertValue is a convenience wrapper for insert, it will create a new listNode with value v
	insertValue(v any, at *ListNode) *ListNode
	// remove removes e from its list, decrements l.len
	remove(e *ListNode)
	// move moves e to next of at
	move(e, at *ListNode)
}

type List struct {
	// we stipulate that root's next to store the first element,
	// and root's prev to store the last element
	root ListNode
	// store the number of listNodes
	len int
}

func (l *List) Init() *List {
	// root's next and prev point to itself
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *ListNode {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *ListNode {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) InsertBefore(v any, mark *ListNode) *ListNode {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

func (l *List) InsertAfter(v any, mark *ListNode) *ListNode {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

func (l *List) Remove(e *ListNode) any {
	if e == nil || e.list != l {
		return nil
	}
	l.remove(e)
	return e.Value
}

func (l *List) PushFront(v any) *ListNode {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *List) PushBack(v any) *ListNode {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

func (l *List) remove(e *ListNode) {
	// break the link between e and its neighbors
	e.prev.next = e.next
	e.next.prev = e.prev
	// avoid memory leaks
	e.next = nil
	e.prev = nil
	e.list = nil
	// decrease the number of listNodes
	l.len--
}

func (l *List) move(e, at *ListNode) {
	if e == at {
		return
	}

	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next

	at.next = e
	e.next.prev = e
}

func (l *List) insert(e, at *ListNode) *ListNode {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List) insertValue(v any, at *ListNode) *ListNode {
	return l.insert(&ListNode{Value: v}, at)
}

func (l *List) lazyInit() {
	// if l is nil, call Init() to create an empty list
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) RangePrint() {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

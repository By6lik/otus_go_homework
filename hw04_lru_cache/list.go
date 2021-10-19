package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func (l *list) Len() int {
	return l.size
}

func (l *list) MoveToFront(i *ListItem) {
	if l.size == 1 {
		return
	}
	var varTemp = l.Extract(i)
	l.head.Prev = varTemp
	varTemp.Next = l.head
	l.head = varTemp
	l.size++
}


func (l *list) Extract(i *ListItem) *ListItem {
	if i == nil {
		return i
	}
	if i == l.head {
		l.head = i.Next
		l.head.Prev = nil
		i.Prev = nil

	} else if i == l.tail {
		l.tail = i.Prev
		
		l.tail.Next = nil
		i.Prev = nil
	} else {
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
		i.Prev = nil
		i.Next = nil
	}
	l.size--
	return i
}

func (l *list) Remove(i *ListItem) {
	l.Extract(i)
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := new(ListItem)
	item.Value = v
	if l.head == nil {
		l.size = 1
		l.head = item
		l.tail = item
	} else {

		item.Next = l.head
		l.head.Prev = item
		l.head = item
		l.size++
	}
	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := new(ListItem)
	item.Value = v
	if l.head == nil {
		l.size = 1
		l.head = item
		l.tail = item
	} else {
		item.Prev = l.tail
		l.tail.Next = item
		l.tail = item
		l.size++
	}
	return l.tail
}



type list struct {
	head, tail *ListItem
	size       int
}

func NewList() List {
	return new(list)
}

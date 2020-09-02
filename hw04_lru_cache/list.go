package hw04_lru_cache // nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *ListItem                  // первый Item
	Back() *ListItem                   // последний Item
	PushFront(v interface{}) *ListItem // добавить значение в начало
	PushBack(v interface{}) *ListItem  // добавить значение в конец
	Remove(i *ListItem)                // удалить элемент
	MoveToFront(i *ListItem)           // переместить элемент в начало
}

type ListItem struct {
	Value interface{} // значение
	Next  *ListItem   // следующий элемент
	Prev  *ListItem   // предыдущий элемент
}

type list struct {
	len       int
	firstItem *ListItem
	lastItem  *ListItem
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.firstItem
}

func (l *list) Back() *ListItem {
	return l.lastItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	i := &ListItem{Value: v}

	if l.len > 0 {
		fi := l.Front()
		i.Next = fi
		fi.Prev = i
	} else {
		l.lastItem = i
	}

	l.firstItem = i
	l.len++

	return i
}

func (l *list) PushBack(v interface{}) *ListItem {
	i := &ListItem{Value: v}

	if l.len > 0 {
		li := l.Back()
		i.Prev = li
		li.Next = i
	} else {
		l.firstItem = i
	}

	l.lastItem = i
	l.len++

	return i
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.firstItem == i && l.lastItem == i:
		l.firstItem = nil
		l.lastItem = nil
	case l.firstItem == i:
		l.firstItem = i.Next
		i.Next.Prev = nil
	case l.lastItem == i:
		l.lastItem = i.Prev
		i.Prev.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	i.Prev = nil
	i.Next = nil

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

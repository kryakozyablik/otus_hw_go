package hw04_lru_cache // nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый Item
	Back() *listItem                   // последний Item
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало
}

type listItem struct {
	Value interface{} // значение
	Next  *listItem   // следующий элемент
	Prev  *listItem   // предыдущий элемент
}

type list struct {
	len       int
	firstItem *listItem
	lastItem  *listItem
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *listItem {
	return l.firstItem
}

func (l *list) Back() *listItem {
	return l.lastItem
}

func (l *list) PushFront(v interface{}) *listItem {
	i := &listItem{Value: v}

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

func (l *list) PushBack(v interface{}) *listItem {
	i := &listItem{Value: v}

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

func (l *list) Remove(i *listItem) {
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

func (l *list) MoveToFront(i *listItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

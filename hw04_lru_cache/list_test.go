package hw04_lru_cache // nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, listToSliceInt(l))
	})

	t.Run("step by step test", func(t *testing.T) {
		l := NewList()

		l.PushFront(30) // [30]
		require.Equal(t, []int{30}, listToSliceInt(l))
		require.Equal(t, 30, l.Front().Value.(int))
		require.Equal(t, 30, l.Back().Value.(int))

		l.PushFront(20) // [20, 30]
		require.Equal(t, []int{20, 30}, listToSliceInt(l))
		require.Equal(t, 20, l.Front().Value.(int))
		require.Equal(t, 30, l.Back().Value.(int))

		l.PushFront(10) // [10, 20, 30]
		require.Equal(t, []int{10, 20, 30}, listToSliceInt(l))
		require.Equal(t, 10, l.Front().Value.(int))
		require.Equal(t, 30, l.Back().Value.(int))

		l.PushBack(40) // [10, 20, 30, 40]
		require.Equal(t, []int{10, 20, 30, 40}, listToSliceInt(l))

		l.PushBack(50) // [10, 20, 30, 40, 50]
		require.Equal(t, []int{10, 20, 30, 40, 50}, listToSliceInt(l))

		l.Remove(l.Front()) // [20, 30, 40, 50]
		require.Equal(t, []int{20, 30, 40, 50}, listToSliceInt(l))

		l.Remove(l.Back()) // [20, 30, 40]
		require.Equal(t, []int{20, 30, 40}, listToSliceInt(l))

		l.MoveToFront(l.Back()) // [40, 20, 30]
		require.Equal(t, []int{40, 20, 30}, listToSliceInt(l))

		l.MoveToFront(l.Front()) // [40, 20, 30]
		require.Equal(t, []int{40, 20, 30}, listToSliceInt(l))

		l.Remove(l.Front()) // [20, 30]
		require.Equal(t, []int{20, 30}, listToSliceInt(l))
		require.Equal(t, 20, l.Front().Value.(int))
		require.Equal(t, 30, l.Back().Value.(int))

		l.Remove(l.Front()) // [30]
		require.Equal(t, []int{30}, listToSliceInt(l))
		require.Equal(t, 30, l.Front().Value.(int))
		require.Equal(t, 30, l.Back().Value.(int))

		l.Remove(l.Front())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("interface test", func(t *testing.T) {
		l := NewList()

		l.PushFront("привет")
		require.Equal(t, "привет", l.Front().Value.(string))

		l.PushFront([]int{1, 2, 3})
		require.Equal(t, []int{1, 2, 3}, l.Front().Value.([]int))

		l.PushFront(map[string]int{"привет": 1})
		require.Equal(t, map[string]int{"привет": 1}, l.Front().Value.(map[string]int))

		l.PushFront(false)
		require.Equal(t, false, l.Front().Value.(bool))
	})
}

func listToSliceInt(l List) []int {
	elems := make([]int, 0, l.Len())
	for i := l.Front(); i != nil; i = i.Next {
		elems = append(elems, i.Value.(int))
	}
	return elems
}

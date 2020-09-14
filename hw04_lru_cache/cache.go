package hw04_lru_cache // nolint:golint,stylecheck
import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                              // Очистить кэш
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mux      sync.Mutex
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	ci := cacheItem{key: key, value: value}

	l.mux.Lock()
	defer l.mux.Unlock()

	if item, ok := l.items[key]; ok {
		l.queue.MoveToFront(item)
		item.Value = ci

		return true
	}

	l.items[key] = l.queue.PushFront(ci)

	if l.queue.Len() > l.capacity {
		ListItem := l.queue.Back()
		ci := ListItem.Value.(cacheItem)
		delete(l.items, ci.key)
		l.queue.Remove(ListItem)
	}

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mux.Lock()
	defer l.mux.Unlock()

	if item, ok := l.items[key]; ok {
		ci := item.Value.(cacheItem)

		return ci.value, true
	}

	return nil, false
}

func (l *lruCache) Clear() {
	l.mux.Lock()
	defer l.mux.Unlock()

	l.queue = NewList()
	l.items = make(map[Key]*ListItem)
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{capacity: capacity, queue: NewList(), items: make(map[Key]*ListItem)}
}

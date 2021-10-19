package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}


func (lru *lruCache) Clear() {
}

func (lru *lruCache) Get(key Key) (interface{}, bool) {
	
	item := lru.items[key]



	if item == nil{
		return nil, false
	} 
	a :=item.Value.(*cacheItem)
	lru.queue.MoveToFront(item)
	return a.value,true
}

func (lru *lruCache) Set(key Key,value interface{}) bool {
	item := lru.items[key]
	

	if item != nil{

		a :=item.Value.(*cacheItem)
		a.value = value
		lru.queue.MoveToFront(item)
		return true
	} else {

		if len(lru.items) == lru.capacity {
			m := lru.queue.Back()
			lru.queue.Remove(m)
			delete(lru.items, m.Value.(*cacheItem).key)
			}
		c := cacheItem{key,value}
		lru.items[key] = lru.queue.PushFront(&c)

		return false
	}
}


type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
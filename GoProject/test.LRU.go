package main

import (
	"container/list"
)

type LRUcache struct {
	cap       int
	cache     map[int]*list.Element
	orderlist *list.List
}

type cacheItem struct {
	key   int
	value int
}

func LRUInit(capacity int) *LRUcache {
	return &LRUcache{
		cap:       capacity,
		cache:     make(map[int]*list.Element),
		orderlist: list.New(),
	}
}

func (this *LRUcache) Get(key int) (value int) {
	elem, ok := this.cache[key]
	if ok {
		return elem.Value.(*cacheItem).value
	}
	return -1
}

func (this *LRUcache) Put(key int, value int) {
	elem, ok := this.cache[key]
	if ok {
		elem.Value.(*cacheItem).value = value
		this.orderlist.MoveToFront(elem)
	} else {
		if len(this.cache) >= this.cap {
			trailItem := this.orderlist.Back()
			delete(this.cache, trailItem.Value.(*cacheItem).key)
			this.orderlist.Remove(trailItem)
		}
		newItem := &cacheItem{key: key, value: value}
		newElem := this.orderlist.PushFront(newItem)
		this.cache[key] = newElem
	}

}

// type LRUCache struct {
// 	capacity  int
// 	cache     map[int]*list.Element
// 	orderList *list.List
// }

// type cacheItem struct {
// 	key   int
// 	value int
// }

// func Constructor(capacity int) LRUCache {
// 	return LRUCache{
// 		capacity:  capacity,
// 		cache:     make(map[int]*list.Element),
// 		orderList: list.New(),
// 	}
// }

// func (this *LRUCache) Get(key int) int {
// 	if elem, ok := this.cache[key]; ok {
// 		this.orderList.MoveToFront(elem)
// 		return elem.Value.(*cacheItem).value
// 	}
// 	return -1
// }

// func (this *LRUCache) Put(key int, value int) {
// 	if elem, ok := this.cache[key]; ok {
// 		elem.Value.(*cacheItem).value = value
// 		this.orderList.MoveToFront(elem)
// 	} else {
// 		if len(this.cache) == this.capacity {
// 			// 淘汰最久未使用的元素
// 			tailElem := this.orderList.Back()
// 			delete(this.cache, tailElem.Value.(*cacheItem).key)
// 			this.orderList.Remove(tailElem)
// 		}
// 		// 新元素添加到列表头部
// 		newItem := &cacheItem{key: key, value: value}
// 		newElem := this.orderList.PushFront(newItem)
// 		this.cache[key] = newElem
// 	}
// }

// func TestLRU() {
// 	var lrucache = Constructor(3)
// 	lrucache.Put(1, 123862)
// 	fmt.Printf("1 lrucache.cache: %v\n", lrucache.cache) //1 lrucache.cache: map[1:0xc00010e210]
// 	lrucache.Put(2, 4354343)
// 	fmt.Printf("2 lrucache.cache: %v\n", lrucache.cache) //2 lrucache.cache: map[1:0xc00010e210 2:0xc00010e270]
// 	lrucache.Put(3, 324532)
// 	fmt.Printf("3 lrucache.cache: %v\n", lrucache.cache) //3 lrucache.cache: map[1:0xc00010e210 2:0xc00010e270 3:0xc00010e330]
// 	lrucache.Put(4, 94223)
// 	fmt.Printf("4 lrucache.cache: %v\n", lrucache.cache) //4 lrucache.cache: map[2:0xc00010e270 3:0xc00010e330 4:0xc00010e390] 1:0xc00010e210被淘汰
// }

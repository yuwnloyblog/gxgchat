package caches

import (
	"container/list"
	"sync"
	"time"

	"github.com/hashicorp/golang-lru/simplelru"
)

type lruCacheItem struct {
	value       interface{}
	updatedTime int64
}

type keyItem struct {
	key       interface{}
	addedTime int64
}

type LruCache struct {
	lru                simplelru.LRUCache
	createdList        *list.List
	createdMap         map[interface{}]*list.Element
	lock               sync.RWMutex
	readTimeoutChecker *time.Ticker
	addTimeoutChecker  *time.Ticker
	valueCreator       func(key interface{}) interface{}
}

func NewLruCacheWithAddReadTimeout(size int, onEvict simplelru.EvictCallback, timeoutAfterRead time.Duration, timeoutAfterCreate time.Duration) *LruCache {
	cache := NewLruCache(size, onEvict)
	cache.AddTimeoutAfterRead(timeoutAfterRead)
	cache.AddTimeoutAfterCreate(timeoutAfterCreate)
	return cache
}

func NewLruCache(size int, onEvict simplelru.EvictCallback) *LruCache {
	myLru, _ := simplelru.NewLRU(size, onEvict)
	cache := &LruCache{
		lru:         myLru,
		createdList: list.New(),
		createdMap:  make(map[interface{}]*list.Element),
	}
	return cache
}

func (c *LruCache) AddValueCreator(creator func(interface{}) interface{}) *LruCache {
	c.valueCreator = creator
	return c
}

func (c *LruCache) AddTimeoutAfterCreate(timeout time.Duration) *LruCache {
	if c.addTimeoutChecker != nil {
		c.addTimeoutChecker.Stop()
	}
	c.addTimeoutChecker = time.NewTicker(time.Second)
	go func() {
		for task := range c.addTimeoutChecker.C {
			if time.Now().UnixMilli()-task.UnixMilli() > 500 {
				continue
			}
			timeLine := task.UnixMilli() - int64(timeout)/(1000*1000)
			c.cleanOldestByCreateTime(timeLine)
		}
	}()
	return c
}

func (c *LruCache) cleanOldestByCreateTime(timeLine int64) {
	for {
		ent := c.createdList.Back()
		if ent != nil {
			keyItem := ent.Value.(*keyItem)
			addedTime := keyItem.addedTime
			if addedTime < timeLine {
				c.Remove(keyItem.key)
			} else {
				break
			}
		}
	}
}

func (c *LruCache) AddTimeoutAfterRead(timeout time.Duration) *LruCache {
	if c.readTimeoutChecker != nil {
		c.readTimeoutChecker.Stop()
	}
	c.readTimeoutChecker = time.NewTicker(time.Second)
	go func() {
		for task := range c.readTimeoutChecker.C {
			if time.Now().UnixMilli()-task.UnixMilli() > 500 {
				continue
			}
			timeLine := task.UnixMilli() - int64(timeout)/(1000*1000)
			c.cleanOdlestByReadTime(timeLine)
		}
	}()
	return c
}

func (c *LruCache) cleanOdlestByReadTime(timeLine int64) {
	for {
		itemKey, itemValue, ok := c.lru.GetOldest()
		if ok {
			updatedTime := itemValue.(lruCacheItem).updatedTime
			if updatedTime < timeLine {
				c.Remove(itemKey)
			} else {
				break
			}
		}
	}
}

func (c *LruCache) Add(key, value interface{}) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.innerAdd(key, value)
}

func (c *LruCache) innerAdd(key, value interface{}) bool {
	nowTime := time.Now().UnixMilli()

	if ent, ok := c.createdMap[key]; ok {
		c.createdList.MoveToFront(ent)
		ent.Value.(*keyItem).addedTime = nowTime
	} else {
		ent := &keyItem{
			addedTime: nowTime,
			key:       key,
		}
		entry := c.createdList.PushFront(ent)
		c.createdMap[key] = entry
	}

	return c.lru.Add(key, lruCacheItem{
		value:       value,
		updatedTime: nowTime,
	})
}

func (c *LruCache) Get(key interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.innerGet(key)
}
func (c *LruCache) innerGet(key interface{}) (interface{}, bool) {
	item, ok := c.lru.Get(key)
	if ok {
		cacheItem := item.(lruCacheItem)
		cacheItem.updatedTime = time.Now().UnixMilli()
		return cacheItem.value, ok
	} else {
		return nil, ok
	}
}
func (c *LruCache) GetByDefault(key interface{}, defaultValue interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	val, ok := c.innerGet(key)
	if ok {
		return val, ok
	} else {
		return defaultValue, ok
	}
}
func (c *LruCache) GetByCreator(key interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	val, ok := c.innerGet(key)
	if ok {
		return val, ok
	} else {
		if c.valueCreator != nil {
			newVal := c.valueCreator(key)
			if newVal != nil {
				c.innerAdd(key, newVal)
				return newVal, true
			}
		}
	}
	return nil, ok
}
func (c *LruCache) Contains(key interface{}) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.Contains(key)
}

func (c *LruCache) Peek(key interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	item, ok := c.lru.Peek(key)
	if ok {
		return item.(lruCacheItem).value, ok
	} else {
		return nil, ok
	}
}

func (c *LruCache) Purge() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.lru.Purge()
	c.createdList.Init()
	for k := range c.createdMap {
		delete(c.createdMap, k)
	}
}
func (c *LruCache) Len() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.Len()
}
func (c *LruCache) ReSize(size int) int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.Resize(size)
}

func (c *LruCache) Remove(key interface{}) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	//remove from keyMap
	if ent, ok := c.createdMap[key]; ok {
		c.createdList.Remove(ent)
		keyItem := ent.Value.(*keyItem)
		delete(c.createdMap, keyItem.key)
	}

	return c.lru.Remove(key)
}

func (c *LruCache) Keys() []interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.Keys()
}

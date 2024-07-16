package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "net/http"
    "time"
    "sync"
    "container/list"
)

type CacheItem struct {
    Key        string
    Value      interface{}
    Expiration int64
}

type LRUCache struct {
    capacity int
    items    map[string]*list.Element
    evictList *list.List
    mu       sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        items:    make(map[string]*list.Element),
        evictList: list.New(),
    }
}

func (c *LRUCache) Set(key string, value interface{}, expiration time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()

    if ele, ok := c.items[key]; ok {
        c.evictList.MoveToFront(ele)
        ele.Value.(*CacheItem).Value = value
        ele.Value.(*CacheItem).Expiration = time.Now().Add(expiration).UnixNano()
        return
    }

    item := &CacheItem{
        Key:        key,
        Value:      value,
        Expiration: time.Now().Add(expiration).UnixNano(),
    }

    ele := c.evictList.PushFront(item)
    c.items[key] = ele

    if c.evictList.Len() > c.capacity {
        c.evict()
    }
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    if ele, ok := c.items[key]; ok {
        if time.Now().UnixNano() > ele.Value.(*CacheItem).Expiration {
            c.evictList.Remove(ele)
            delete(c.items, key)
            return nil, false
        }
        c.evictList.MoveToFront(ele)
        return ele.Value.(*CacheItem).Value, true
    }
    return nil, false
}

func (c *LRUCache) evict() {
    ele := c.evictList.Back()
    if ele != nil {
        c.evictList.Remove(ele)
        delete(c.items, ele.Value.(*CacheItem).Key)
    }
}

var cache *LRUCache

func main() {
    cache = NewLRUCache(1024)
    r := gin.Default()
    r.Use(cors.Default())

    r.GET("/get_lru_key/:key", getHandler)
    r.POST("/set", setHandler)

    r.Run(":8080")
}

func getHandler(c *gin.Context) {
    key := c.Param("key")
    value, ok := cache.Get(key)
    if !ok {
        c.JSON(http.StatusNotFound, gin.H{"error": "key not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"value": value})
}

func setHandler(c *gin.Context) {
    var req struct {
        Key        string        `json:"key"`
        Value      interface{}   `json:"value"`
        Expiration time.Duration `json:"expiration"`
    }
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    cache.Set(req.Key, req.Value, req.Expiration*time.Second)
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}

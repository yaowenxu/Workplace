package cache

import (
	"sync"
	"time"
)

type value struct {
	v       []byte
	created time.Time
}

type inMemoryCache struct {
	c     map[string]value
	mutex sync.RWMutex
	Stat
	ttl time.Duration // 时间范围，表示生存时间;
}

func (c *inMemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.c[k] = value{v, time.Now()} // 每次set操作更新 ttl，也就是提前设置好时间;
	c.add(k, v)
	return nil
}

func (c *inMemoryCache) Get(k string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.c[k].v, nil
}

func (c *inMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, exist := c.c[k]
	if exist {
		delete(c.c, k)
		c.del(k, v.v)
	}
	return nil
}

func (c *inMemoryCache) GetStat() Stat {
	return c.Stat
}

func newInMemoryCache(ttl int) *inMemoryCache {
	c := &inMemoryCache{make(map[string]value), sync.RWMutex{}, Stat{}, time.Duration(ttl) * time.Second}
	if ttl > 0 {
		go c.expirer()
	}
	return c
}

func (c *inMemoryCache) expirer() { // 超过ttl时间，进行超时操作;
	for {
		time.Sleep(c.ttl)
		c.mutex.RLock()
		for k, v := range c.c {
			c.mutex.RUnlock()
			if v.created.Add(c.ttl).Before(time.Now()) {
				c.Del(k)
			}
			c.mutex.RLock()
		}
		c.mutex.RUnlock()
	}
}

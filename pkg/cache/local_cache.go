package сache

import (
	"github.com/gofiber/fiber/v2"
	"sync"
	"time"
)

type LocalCacheManagerSync struct {
	data  map[string]interface{}
	mutex sync.RWMutex
}

func NewLocalCacheManagerSync() *LocalCacheManagerSync {
	return &LocalCacheManagerSync{
		data: make(map[string]interface{}),
	}
}

func GetLocalCacheByLocals(c *fiber.Ctx) *LocalCacheManagerSync {
	return c.Locals("cache_manager").(*LocalCacheManagerSync)
}

func (c *LocalCacheManagerSync) Set(key string, value interface{}, expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value

	// Schedule a function to remove the key after expiration
	go func() {
		time.Sleep(expiration)
		c.Delete(key)
	}()
}

func (c *LocalCacheManagerSync) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, exists := c.data[key]
	return value, exists
}

func (c *LocalCacheManagerSync) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
}

func (c *LocalCacheManagerSync) Remember(key string, expiration time.Duration, compute func() interface{}) interface{} {
	if value, exists := c.Get(key); exists {
		return value
	}

	result := compute()
	c.Set(key, result, expiration)
	return result
}

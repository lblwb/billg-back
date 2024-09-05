package сache

import (
	"github.com/dgraph-io/ristretto"
	"github.com/gofiber/fiber/v2"
	"sync"
	"time"
)

type CacheManagerSync struct {
	cache *ristretto.Cache
	mutex sync.RWMutex
}

func NewCacheManagerSync() *CacheManagerSync {
	config := &ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	}
	cache, _ := ristretto.NewCache(config)

	return &CacheManagerSync{
		cache: cache,
	}
}

func GetCacheByLocals(c *fiber.Ctx) *CacheManagerSync {
	return c.Locals("cache_manager").(*CacheManagerSync)
}

func (c *CacheManagerSync) Set(key string, value interface{}, expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache.SetWithTTL(key, value, 1, expiration)
}

func (c *CacheManagerSync) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if value, found := c.cache.Get(key); found {
		return value, true
	}
	return nil, false
}

func (c *CacheManagerSync) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache.Del(key)
}

func (c *CacheManagerSync) Remember(key string, expiration time.Duration, compute func() interface{}) interface{} {
	if value, exists := c.Get(key); exists {
		return value
	}

	// Вызываем функцию compute и сохраняем результат в кэше
	result := compute()
	c.Set(key, result, expiration)
	return result
}

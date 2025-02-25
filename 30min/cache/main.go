/*
Необходимо написать in-memory кэш, который будет по ключу (uuid пользователя) возвращать профиль и список его заказов.

1. У кэша должен быть TTL (2 сек)
2. Кэшем может пользоваться функция(-и), которая работает с заказами (добавляет/обновляет/удаляет). Если TTL истек, то возвращается nil.
	При апдейте TTL снова устанавливается 2 сек. Методы должны быть потокобезопасными
3. Должны быть написаны тестовые сценарии использования данного кэша
(базовые структуры не менять)

Доп задание: автоматическая очистка истекших записей

type Profile struct {
	UUID   string
	Name   string
	Orders []*Order
}

type Order struct {
	UUID      string
	Value     any
	CreatedAt time.Time
	UpdatedAt time.Time
}
*/

package main

import (
	"sync"
	"time"
)

type Profile struct {
	UUID   string
	Name   string
	Orders []*Order
}

type Order struct {
	UUID      string
	Value     any
	CreatedAt time.Time
	UpdatedAt time.Time
}

const TTL = 2 * time.Second

type CacheItem struct {
	*Profile
	expireAt time.Time
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]*CacheItem
	ttl   time.Duration
}

func NewCache() *Cache {
	return &Cache{
		items: map[string]*CacheItem{},
		ttl:   TTL,
	}
}

func (c *Cache) Get(uuid string) (p *Profile, found bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[uuid]
	if !found {
		return
	}

	if item.expireAt.Before(time.Now()) {
		return p, false
	}

	p = item.Profile

	return
}

func (c *Cache) Set(p *Profile) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[p.UUID] = &CacheItem{
		Profile:  p,
		expireAt: time.Now().Add(c.ttl),
	}
}

func (c *Cache) Cleanup() {
	for range time.Tick(c.ttl) {
		println("debug tick")
		for _, v := range c.items {
			if v.expireAt.Before(time.Now()) {
				println("debug delete")
				c.mu.Lock()
				delete(c.items, v.UUID)
				c.mu.Unlock()
			}
		}
	}
}

func main() {
	cache := NewCache()
	go cache.Cleanup()

	const uuidToFind = "aslkfj-123"

	r, f := cache.Get(uuidToFind)
	println(r, f) //false

	cache.Set(
		&Profile{
			UUID:   uuidToFind,
			Name:   "user1",
			Orders: []*Order{{UUID: "312", Value: "test1"}},
		},
	)

	time.Sleep(1 * time.Second)

	r, f = cache.Get(uuidToFind)
	println(r.Name, r.Orders, f) //true

	time.Sleep(3 * time.Second)

	r, f = cache.Get(uuidToFind)
	println(r, f) //false

}

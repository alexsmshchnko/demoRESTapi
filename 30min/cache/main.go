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
	Profile
	expireAt time.Time
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]*CacheItem
	ttl   time.Duration
}

func NewCache() (c *Cache) {
	c = &Cache{
		items: map[string]*CacheItem{},
		ttl:   TTL,
	}

	go c.Cleanup()

	return
}

func (c *Cache) Get(uuid string) (p Profile, found bool) {
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

func (c *Cache) Set(p Profile) {
	c.mu.Lock()
	defer c.mu.Unlock()

	orders := make([]*Order, len(p.Orders))
	for i, v := range p.Orders {
		v := *v
		orders[i] = &v
	}

	profile := p
	profile.Orders = orders

	c.items[p.UUID] = &CacheItem{
		Profile:  profile,
		expireAt: time.Now().Add(c.ttl),
	}
}

func (c *Cache) Cleanup() {
	for range time.Tick(c.ttl) {
		println("debug: tick")
		for _, v := range c.items {
			if v.expireAt.Before(time.Now()) {
				println("debug: delete")
				c.mu.Lock()
				delete(c.items, v.UUID)
				c.mu.Unlock()
			}
		}
	}
}

func main() {
	cache := NewCache()

	const uuidToFind = "aslkfj-123"

	r, f := cache.Get(uuidToFind)
	println(r.Name, r.Orders, f) //false

	p := &Profile{
		UUID:   uuidToFind,
		Name:   "user1",
		Orders: []*Order{{UUID: "321", Value: "test1"}},
	}

	cache.Set(*p)

	//check cache is isolated and safe
	p.Name = "user0"
	p.Orders[0].UUID = "000"

	time.Sleep(1 * time.Second)

	r, f = cache.Get(uuidToFind)
	println(r.Name, r.Orders[0].UUID, f) //true

	time.Sleep(3 * time.Second)

	r, f = cache.Get(uuidToFind)
	println(r.Name, r.Orders, f) //false

}

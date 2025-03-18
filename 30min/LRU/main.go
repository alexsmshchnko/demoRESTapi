package main

import (
	"container/list"
	"fmt"
)

type cache struct {
	cache map[string]*list.Element
	list  *list.List
	cap   int
}

type entry struct {
	key string
	val []byte
}

func (e entry) String() string {
	return fmt.Sprintf("%s: %s", e.key, e.val)
}

func newCache(capacity int) *cache {
	return &cache{
		cache: make(map[string]*list.Element),
		list:  list.New().Init(),
		cap:   capacity,
	}
}

func (c *cache) push(e entry) {
	if c.list.Len() == c.cap {
		b := c.list.Back()
		c.list.Remove(b)
		delete(c.cache, b.Value.(entry).key)
	}
	v := c.list.PushFront(e)
	c.cache[e.key] = v
}

func (c *cache) get(key string) (e entry, f bool) {
	v, f := c.cache[key]
	if f {
		c.list.MoveToFront(v)
		return v.Value.(entry), true
	}

	return
}

func main() {
	c := newCache(2)
	c.push(entry{"gold", []byte("Mr gold")})
	c.push(entry{"black", []byte("Mrs black")})
	c.push(entry{"green", []byte("Mr green")})

	println(c.list.Len()) //2
	res, f := c.get("gold")
	if f {
		fmt.Println(res)
	} else {
		fmt.Println("not found") //not found
	}

	res, f = c.get("green")
	if f {
		fmt.Println(res) //green
	} else {
		fmt.Println("not found")
	}

	c.push(entry{"pink", []byte("Mr pink")})
	println(c.list.Len()) //2

	fmt.Println("front", c.list.Front().Value.(entry)) //pink
	res, f = c.get("green")
	if f {
		fmt.Println(res) //found
	} else {
		fmt.Println("not found")
	}
	fmt.Println("front", c.list.Front().Value.(entry)) //green

	res, f = c.get("black")
	if f {
		fmt.Println(res)
	} else {
		fmt.Println("not found") //not found
	}
}

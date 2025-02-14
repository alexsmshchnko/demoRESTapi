package main

import (
	"fmt"
	"sync"
	"time"

	"math/rand"
)

type total struct {
	mu sync.Mutex
	n  int
}

func (t *total) add(n int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.n += n
}

func produce(id int, ch chan<- int, t *total) {
	waitProducer.Add(1)
	defer waitProducer.Done()
	for i := 0; i < id+2; i++ {
		data := rand.Intn(100)
		// fmt.Printf("producer%d sent %d\n", id, data)
		t.add(1)
		ch <- data
		time.Sleep(1 * time.Millisecond)
	}

}

func consumer(id int, ch <-chan int, t *total) {
	waitConsumer.Add(1)
	defer waitConsumer.Done()
	for data := range ch {
		t.add(-1)
		// v := data + 1
		fmt.Printf("consumer%d received %d ", id, data)
		time.Sleep(1 * time.Millisecond)
	}
}

var (
	waitProducer sync.WaitGroup
	waitConsumer sync.WaitGroup
)

func main() {
	total := &total{}

	ch := make(chan int, 5)

	for i := 0; i < 50; i++ {
		i := i
		// time.Sleep(time.Millisecond)
		go produce(i, ch, total)
	}

	for i := 0; i < 20; i++ {
		i := i
		go consumer(i, ch, total)
	}

	waitProducer.Wait()

	fmt.Printf("length before close: %d\n", len(ch))

	close(ch)

	fmt.Printf("length after close: %d\n", len(ch))

	waitConsumer.Wait()

	_, ok := <-ch
	fmt.Printf("is open after all read? %v\n", ok)

	fmt.Printf("balance: %d\n", total.n)

	fmt.Printf("done")
}

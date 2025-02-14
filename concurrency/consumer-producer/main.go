package main

import (
	"fmt"
	"sync"
	"time"

	"math/rand"
)

func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for i := 0; i < id+2; i++ {
		data := rand.Intn(100)
		fmt.Printf("producer%d sent %d\n", id, data)
		ch <- data
		time.Sleep(100 * time.Millisecond)
	}

}

func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for data := range ch {
		fmt.Printf("consumer%d received %d\n", id, data)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	var (
		wgp sync.WaitGroup
		wgc sync.WaitGroup
		ok  bool
	)

	ch := make(chan int, 5)

	for i := 0; i < 3; i++ {
		i := i
		time.Sleep(time.Millisecond)
		go producer(i, ch, &wgp)
	}

	for i := 0; i < 2; i++ {
		i := i
		go consumer(i, ch, &wgc)
	}

	wgp.Wait()

	_, ok = <-ch
	fmt.Printf("is open %v\n", ok)

	close(ch)

	wgc.Wait()

	_, ok = <-ch
	fmt.Printf("is open %v\n", ok)

	fmt.Printf("done")
}

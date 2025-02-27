package main

import "sync"

func joinChannels(cIn ...<-chan (int)) (c0 chan (int)) {
	c0 = make(chan (int))

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(cIn))

		for i, v := range cIn {
			go func(wg *sync.WaitGroup, cIn <-chan int) {
				for v := range cIn {
					println("idx:", i, " val:", v)
				}
				wg.Done()
			}(wg, v)
		}

		wg.Wait()
		close(c0)
	}()

	return c0
}

func main() {
	c1 := make(chan (int))
	c2 := make(chan (int))
	c3 := make(chan (int))

	go func() {
		defer close(c1)
		for _, v := range []int{1, 2, 3} {
			c1 <- v
		}
	}()

	go func() {
		defer close(c2)
		for _, v := range []int{4, 5, 6} {
			c2 <- v
		}
	}()

	go func() {
		defer close(c3)
		for _, v := range []int{7, 8, 9} {
			c3 <- v
		}
	}()

	for num := range joinChannels(c1, c2, c3) {
		println(num)
	}

}

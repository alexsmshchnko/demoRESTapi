package main

import (
	"sync/atomic"
	"time"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type proc struct {
	max_rps int64
	rps     *int64
	Exec    func(f func())
}

func (p *proc) ticker() {
	t := time.Tick(time.Second)
	for v := range t {
		println("tick", v.Format("04:05.000"), atomic.LoadInt64(p.rps), "->0")

		atomic.StoreInt64(p.rps, 0)
	}
}

func newProc(rps int64) (p *proc) {
	var zero int64
	p = &proc{
		max_rps: rps,
		rps:     &zero,
		Exec: func(f func()) {
			r := atomic.AddInt64(p.rps, 1)
			if r > p.max_rps {
				println("reject: limit exceeded", time.Now().Format("04:05.000"), r)
				return
			}
			println("exec", time.Now().Format("04:05.000"), r)

			f()
		},
	}

	go p.ticker()

	return
}

func run() (err error) {
	const RPS int64 = 2

	p := newProc(RPS)
	f := func() {
		time.Sleep(10 * time.Millisecond)
	}

	time.Sleep(10 * time.Millisecond)
	go p.Exec(f)
	go p.Exec(f)
	go p.Exec(f) //no Exec
	time.Sleep(2 * time.Second)
	go p.Exec(f)
	go p.Exec(f)
	time.Sleep(2 * time.Second)
	go p.Exec(f)
	go p.Exec(f)
	go p.Exec(f) //no Exec
	go p.Exec(f) //no Exec
	time.Sleep(2 * time.Second)

	return
}

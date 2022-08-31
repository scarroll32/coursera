package main

import (
	"fmt"
	"sync"
)

type Philosopher struct {
	num int
}

type Chop struct {
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(Chop)
	},
}

var w sync.WaitGroup

func (p Philosopher) eat(boss chan int) {
	defer w.Done()

	<-boss

	fmt.Printf("starting to eat %d\n", p.num)
	left := pool.Get()
	right := pool.Get()
	pool.Put(left)
	pool.Put(right)
	fmt.Printf("finishing eating %d\n", p.num)

	boss <- 1
}

func main() {
	boss := make(chan int, 2)

	for i := 0; i < 5; i++ {
		pool.Put(new(Chop))
	}

	philos := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{i + 1}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			w.Add(1)
			go philos[i].eat(boss)
		}
	}

	boss <- 1
	boss <- 1

	w.Wait()
}

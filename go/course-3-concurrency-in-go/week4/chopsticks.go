package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	id              int
	leftCS, rightCS *Chopstick
	eatTimes        int
}

var wait sync.Mutex

func (p Philosopher) eat() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		wait.Lock()
		if (p.eatTimes) < 3 {
			p.leftCS.Lock()
			p.rightCS.Lock()
			wait.Unlock()
			p.eatTimes = p.eatTimes + 1
			fmt.Printf("Philosopher %d eating %d times\n", p.id, p.eatTimes)
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("Philosopher finished %d\n", p.id)

			p.leftCS.Unlock()
			p.rightCS.Unlock()

		} else {
			wait.Unlock()
		}
	}
}

var wg sync.WaitGroup

func main() {
	CSticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(Chopstick)
	}

	Philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		Philosophers[i] = &Philosopher{i + 1, CSticks[i], CSticks[(i+1)%5], 0}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Philosophers[i].eat()
	}
	wg.Wait()
}

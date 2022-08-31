package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const philosophersCount = 5
const chopsticksCount = 5
const eatTimes = 3
const maxEatingPhilosophers = 2

// Chopstick struct
type Chopstick struct {
	sync.Mutex
}

// Philosopher struct
type Philosopher struct {
	number                 int
	rChopstick, lChopstick *Chopstick
}

func (ph *Philosopher) eat(hostCh chan *Philosopher, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := eatTimes; i > 0; i-- {
		hostCh <- ph // get permission from host
		ph.rChopstick.Lock()
		ph.lChopstick.Lock()
		fmt.Printf("starting to eat %d\n", ph.number)
		fmt.Printf("finishing eating %d\n", ph.number)
		ph.rChopstick.Unlock()
		ph.lChopstick.Unlock()
	}
}

func main() {
	var chopsticks [chopsticksCount]*Chopstick
	for i := 0; i < chopsticksCount; i++ {
		chopsticks[i] = &Chopstick{}
	}

	var philosophers [philosophersCount]*Philosopher
	for i := 0; i < philosophersCount; i++ {
		chopstickIndex := rand.Intn(chopsticksCount)
		philosophers[i] = &Philosopher{number: i + 1, rChopstick: chopsticks[chopstickIndex], lChopstick: chopsticks[(chopstickIndex+1)%philosophersCount]}
	}

	hostCh := make(chan *Philosopher, maxEatingPhilosophers) // channal with only maxEatingPhilosophers philosophers in it at the same time
	go func() {
		for {
			<-hostCh
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < philosophersCount; i++ {
		wg.Add(1)
		go philosophers[i].eat(hostCh, &wg)
	}
	wg.Wait()
}

package main

import (
	"fmt"
	"time"
)

var x int = 1

func read() {
	fmt.Println(x)
}

func write() {
	x++
}

func main() {
	go read()
	go write()
	time.Sleep(time.Millisecond)
}

/*
Race Condition:
--------------
Race conditions are where 2 threads are accessing memory at the same time, one of which is writing.
Race conditions occur because of unsynchronized access to shared memory.
*/

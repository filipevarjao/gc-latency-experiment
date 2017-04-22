package main

import (
	"fmt"
	"time"
	slab "../go-slab"
)

const (
	windowSize = 200000
	msgCount   = 1000000
)

type message []byte

type channel [windowSize]message

var worst time.Duration

var m []byte

func mkMessage(n int, arena *slab.Arena) message {
	//m := make(message, 1024)
	m = arena.Alloc(1024)
	for i := range m {
		m[i] = byte(n)
	}
	arena.DecRef(m)
	return m
}

func pushMsg(c *channel, highID int) {
	start := time.Now()
	m := mkMessage(highID, slab.NewArena(1, 1024, 2, nil))
	(*c)[highID%windowSize] = m
	elapsed := time.Since(start)
	if elapsed > worst {
		worst = elapsed
	}
}

func main() {
	var c channel
	for i := 0; i < msgCount; i++ {
		pushMsg(&c, i)
	}
	fmt.Println("Worst push time: ", worst)
}

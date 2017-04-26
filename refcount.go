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

func mkMessage(n int, m []byte) message {
	for i := range m {
		m[i] = byte(n)
	}
	return m
}

func pushMsg(c *channel, highID int, msg []byte, arena *slab.Arena) {
	start := time.Now()
	m := mkMessage(highID, msg)
	(*c)[highID%windowSize] = m
	arena.DecRef(msg)
	elapsed := time.Since(start)
	if elapsed > worst {
		worst = elapsed
	}
}

func main() {
	var c channel
	var msg []byte
	arena := slab.NewArena(1, 1024, 2, nil)

	for i := 0; i < msgCount; i++ {
		msg = arena.Alloc(1024)
		pushMsg(&c, i, msg, arena)
	}
	fmt.Println("Worst push time: ", worst)
}

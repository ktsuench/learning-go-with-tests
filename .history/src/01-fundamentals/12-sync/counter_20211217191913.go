package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

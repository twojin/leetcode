package design

import (
	"fmt"
)

type Product struct {
	Val  int
	buf  chan int
	stop bool
}

func (p *Product) producer(count int) {
	p.buf = make(chan int, count)
	i := 0
	for {
		if p.stop {
			break
		} else {
			p.buf <- i
		}
		i++
	}
}

func (p *Product) consumer() {
	for {
		fmt.Println(<-p.buf)

	}
}

func (p *Product) stopP() {
	p.stop = true
}

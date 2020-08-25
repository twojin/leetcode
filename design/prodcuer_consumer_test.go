package design

import (
	"testing"
	"time"
)

func Test_prodcuer_consumer(t *testing.T) {
	p := &Product{}
	go p.producer(10)
	go p.consumer()
	time.Sleep(time.Second * 3)
	p.stopP()
}

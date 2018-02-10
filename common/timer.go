package common

import (
	"fmt"
	"time"
)

type (
	//Timer 定时器
	Timer struct {
		ticker *time.Ticker
	}
)

//Start 开始定时器
func (t *Timer) Start() {

	t.ticker = time.NewTicker(time.Second * 1)

	go func() {
		for _ = range t.ticker.C {
			fmt.Printf("ticked at %v\n", time.Now())
		}
	}()
}

//Stop 关闭定时器
func (t *Timer) Stop() {
	t.ticker.Stop()
}

//NewTimer 实例化一个定时器
func NewTimer() *Timer {
	return &Timer{}
}

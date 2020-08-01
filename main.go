package main

import (
	"fmt"
	"time"
)

func main() {
	race()
	//notRace()
}

func race() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(time.Second * 2, func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(time.Second * 1)
	})
	time.Sleep(10 * time.Second)
}

func notRace() {
	start := time.Now()
	reset := make(chan bool)
	var t *time.Timer
	t = time.AfterFunc(time.Second * 2, func() {
		fmt.Println(time.Now().Sub(start))
		reset <- true
	})

	for time.Since(start) < 30*time.Second {
		<-reset
		t.Reset(time.Second * 1)
	}
}
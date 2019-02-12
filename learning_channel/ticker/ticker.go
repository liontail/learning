package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	ticker := time.NewTicker(time.Millisecond * 500)
	ticker2 := time.NewTicker(time.Millisecond * 1000)
	go func() {
		for t := range ticker.C {
			i++
			fmt.Println("Tick at :", i, t)
		}
	}()

	go func() {
		for t := range ticker2.C {
			fmt.Println("Tick2 at :", i, t)
		}
	}()

	time.Sleep(time.Millisecond * 2100)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

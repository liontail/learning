package main

import (
	"sync"

	"github.com/Sirupsen/logrus"
)

var Deck []int
var countP [4]int
var count int
var mutex = &sync.Mutex{}

// var c chan int

func main() {
	// c = make(chan int)
	for i := 0; i < 52; i++ {
		Deck = append(Deck, i+1)
	}

	for {
		if count > 52 {
			break
		}
		for i := 0; i < 4; i++ {

			go Worker(i)
			// countP[i]++

		}

	}
	// time.Sleep(time.Second * 1)
	for i := 0; i < 4; i++ {
		logrus.Infoln("No: ", i, " Pick Total: ", countP[i])
	}
}

func Worker(no int) {
	mutex.Lock()
	if count < 52 {
		logrus.Infoln(no, ">>> pick:", Deck[count])
		countP[no]++
	}
	mutex.Unlock()
	count = count + 1

}

package main

import (
	"log"
	"time"
)

type Work struct {
	x, y, z int
}

func worker(in <-chan *Work, out chan<- *Work) {
	for work := range in {
		work.z = work.x * work.y
		time.Sleep(time.Duration(work.z) * time.Millisecond)
		out <- work
	}
}

func GetWork(out <-chan *Work) {
	for {
		select {
		case output := <-out:
			log.Println(*output)
		}
	}
}

func SendWork(in chan<- *Work) {
	for i := 0; i < 1000; i++ {
		in <- &Work{
			x: i,
			y: i + 1,
		}
	}
}

func Run() {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < 40; i++ {
		go worker(in, out)
	}
	go SendWork(in)
	GetWork(out)
}

func main() {
	Run()

}

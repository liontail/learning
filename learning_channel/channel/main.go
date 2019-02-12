package main

import (
	"fmt"
	"log"
)
import "time"

func test1() {
	c := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	a := <-c

	b := <-c
	log.Println(a)
	log.Println(b)
}

func testChannel1() {
	done := make(chan bool)
	stop := make(chan bool)
	job := make(chan int)

	go func() {
		for i := 0; i < 2999999; i++ {
			job <- i
		}
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		stop <- true
		time.Sleep(500 * time.Millisecond)
		log.Printf("Done sleeping")
		done <- true
	}()

	go func() {
		for {
			select {
			case a := <-job:
				log.Println(a)
			case <-stop:
				done <- true
				log.Println("----------Stop!!-----------")
			}
		}

	}()

	<-done
}

func sendMessage() {

	messages := make(chan string, 1)
	signals := make(chan bool, 1)
	select {
	case msg := <-messages:
		fmt.Println("Received Message", msg)
	default:
		fmt.Println("No message received")
	}
	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("Sent message", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	case sig := <-signals:
		fmt.Println("Received signal", sig)
	default:
		fmt.Println("No activity")
	}

}

func main() {
	// test1()
	testChannel1()
	// sendMessage()

}

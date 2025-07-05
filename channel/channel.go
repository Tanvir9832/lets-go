package main

import (
	"fmt"
	"time"
)

// Communication between go routines are happends through channels

// Receive here (0)
// func processNum(numChannel chan int) {
// 	for {
// 		fmt.Println(<-numChannel)
// 		time.Sleep(time.Second)
// 	}
// }

// Receive here (1)
// func multiChan(numChanX chan []int) {
// 	arr := <-numChanX
// 	res := arr[0] + arr[1]
// 	numChanX <- []int{res}
// }

// buffer
func sendEmail(bufNumChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range bufNumChan {
		fmt.Println("Email sending to ", email)
		time.Sleep(time.Second)
	}
}

func main() {
	// numChannel := make(chan int)
	// go processNum(numChannel)
	// numChannel <- 5

	// Send from here (0)
	// for {
	// 	numChannel <- rand.Intn(100)
	// }

	// Send from here (1)
	// numChan := make(chan []int)

	// go multiChan(numChan)
	// numChan <- []int{10, 20}
	// fmt.Println(<-numChan)

	//buffer channel (non blocking)
	numChannelBuf := make(chan string, 100)
	done := make(chan bool)
	go sendEmail(numChannelBuf, done)

	for i := range 5 {
		numChannelBuf <- fmt.Sprintf("%d@gmail.com", i) // buffer non blocking
	}

	close(numChannelBuf) //non blocking (closing is must, specially for buffer channel)
	<-done               //Blocking
	fmt.Println("Done sending...")
}

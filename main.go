package main

import (
	cookingsim "lesson12/cooking-sim"
	"time"
)

func main() {
	eod := make(chan struct{})
	go cookingsim.StartCooking(eod)

	time.Sleep(10 * time.Second)

	eod <- struct{}{}

	time.Sleep(1 * time.Second)
}

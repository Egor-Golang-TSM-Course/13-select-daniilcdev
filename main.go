package main

import (
	cookingsim "lesson12/cooking-sim"
	network "lesson12/restaurant-network"
	"time"
)

func main() {
	eod := make(chan struct{})
	go cookingsim.StartCooking(eod)

	time.Sleep(10 * time.Second)

	eod <- struct{}{}

	time.Sleep(time.Second)

	go network.RunNetwork(eod)

	time.Sleep(10 * time.Second)

	eod <- struct{}{}

	time.Sleep(time.Second)
}

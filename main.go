package main

import (
	"bufio"
	"fmt"
	"lesson12/chat"
	cookingsim "lesson12/cooking-sim"
	network "lesson12/restaurant-network"
	"os"
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

	go chat.Chat(eod)

	time.Sleep(10 * time.Second)

	eod <- struct{}{}

	time.Sleep(time.Second)
	fmt.Print("enter any key to exit: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}

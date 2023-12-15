package cookingsim

import (
	"fmt"
	"math/rand"
	"time"
)

func StartCooking(eod <-chan struct{}) {
	names := []string{"John", "Alex", "Fred", "Mark", "Antonio"}

	done := make(chan struct{}, len(names))

	fmt.Println("restaurant opened")

	for _, name := range names {
		c := &Chef{Name: name, CookingTime: time.Duration(100<<10 + (rand.Uint32() % 200))}
		go c.Cook(done)
	}

	<-eod
	fmt.Println("END OF DAY! closing restaurant...")

	for i := 0; i < len(names); i++ {
		done <- struct{}{}
	}

	fmt.Println("restaurant closed")
}

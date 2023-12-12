package cookingsim

import (
	"fmt"
	"math/rand"
	"time"
)

func StartCooking(eod <-chan struct{}) {
	names := []string{"John", "Alex", "Fred", "Mark", "Antonio"}

	done := make(chan struct{}, len(names))

	fmt.Println("restourant opened")

	for _, name := range names {
		c := &Chef{Name: name, CookingTime: time.Duration(100<<10 + (rand.Uint32() % 200))}
		go c.Cook(done)
	}
	for {
		select {
		case <-eod:
			fmt.Println("END OF DAY! clsoing restourant...")

			for i := 0; i < len(names); i++ {
				done <- struct{}{}
			}

			fmt.Println("restourant closed")
		default:
			continue
		}
	}

}

package cookingsim

import (
	"fmt"
	"time"
)

type Chef struct {
	Name        string
	CookingTime time.Duration
}

func (c *Chef) Cook(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Printf("Chef %s - stop working, end of day\n", c.Name)
			return
		default:
			fmt.Printf("Chef %s - cooking started\n", c.Name)
			time.Sleep(c.CookingTime)
			fmt.Printf("Chef %s - cooking done\n", c.Name)
		}
	}
}

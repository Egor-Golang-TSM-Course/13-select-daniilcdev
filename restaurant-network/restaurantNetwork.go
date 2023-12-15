package network

import (
	"fmt"
)

type mainOffice struct {
	workDay     chan struct{}
	reports     chan *restaurant
	restaurants []restaurant
}

func RunNetwork(eod <-chan struct{}) {
	fmt.Println("restaurant network started")

	const nRestaurants = 5

	mo := mainOffice{
		workDay:     make(chan struct{}),
		reports:     make(chan *restaurant, nRestaurants),
		restaurants: make([]restaurant, 0, nRestaurants),
	}

	for i := 0; i < nRestaurants; i++ {
		r := restaurant{name: fmt.Sprintf("Restaurant #%d", i+1)}
		mo.restaurants = append(mo.restaurants, r)
		go r.open(mo.workDay, mo.reports)
	}

	<-eod
	mo.close()
}

func (mo *mainOffice) close() {
	var sum int64 = 0

	for range mo.restaurants {
		mo.workDay <- struct{}{}

		r := <-mo.reports
		sum += r.sales

		fmt.Printf("report received : %s sold %d\n", r.name, r.sales)
	}

	fmt.Println("restaurant network stopped, total sales:", sum)
}

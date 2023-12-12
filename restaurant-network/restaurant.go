package network

import (
	"fmt"
	"math/rand"
)

type restaurant struct {
	name  string
	sales int64
}

func (r *restaurant) open(workDay <-chan struct{}, reports chan<- *restaurant) {
	<-workDay
	r.endOfDayReport(reports)
}

func (r *restaurant) endOfDayReport(report chan<- *restaurant) {
	r.sales = int64(rand.Int63() % 100)
	fmt.Printf("report submitted: %s sold %d\n", r.name, r.sales)
	report <- r
}

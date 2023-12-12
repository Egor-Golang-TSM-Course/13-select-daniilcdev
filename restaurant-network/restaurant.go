package network

import (
	"fmt"
	"math/rand"
)

type restaurant struct {
	name  string
	sales int64
}

func (r *restaurant) open(workDay <-chan struct{}, reports chan<- int64) {
	for {
		select {
		case <-workDay:
			r.endOfDayReport(reports)
			return
		default:
			continue
		}
	}
}

func (r *restaurant) endOfDayReport(report chan<- int64) {
	r.sales = int64(rand.Int63() % 100)
	fmt.Printf("%s submitted report: sales during day - %d\n", r.name, r.sales)
	report <- r.sales
}

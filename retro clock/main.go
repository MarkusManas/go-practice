package main

import (
	"clock/symbols"
	"fmt"
	"time"
)

func ComputeTime(n int) (int, int, int) {
	hours := n / 3600
	r := n % 3600
	minutes := r / 60
	seconds := r % 60

	return hours, minutes, seconds
}
func getCurrentSeconds(t time.Time) int {
	return 60*60*t.Hour() + 60*t.Minute() + t.Second()
}

func main() {
	var h, m, s int

	cs := getCurrentSeconds(time.Now())
	for {
		h, m, s = ComputeTime(cs)
		time_symbols := symbols.ConvertTimeToSymbols(h, m, s)
		symbols.PrintClock(time_symbols)
		cs++
		time.Sleep(1 * time.Second)
		fmt.Println()
	}
}

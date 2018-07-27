package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)  // sends on chan tick every 100ms
	boom := time.After(500 * time.Millisecond) // sends on chan boom one after 500ms
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		}
	}
}

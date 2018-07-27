package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func main() {
	var wg sync.WaitGroup
	ids := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, id := range ids {
		wg.Add(1)            // wait for one more goroutine to finish
		go func(id string) { // run database query's in parallel
			err := UpdateDatabase(id, "new val")
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(id)
	}
	wg.Wait()
}

func UpdateDatabase(id string, val string) error { // Fake DB operation that take some time
	time.Sleep(time.Second)
	return nil
}

// END OMIT

package examples

import (
	"fmt"
	"sync"
)

var (
	counter int
	lock    sync.Mutex
)

func increment() {
	lock.Lock()
	defer lock.Unlock()

	counter++
	fmt.Println("Incremented counter", counter)

}

func decrement() {
	lock.Lock()
	defer lock.Unlock()

	counter--
	fmt.Println("Decremented counter", counter)

}

func ExecuteOne() {

	for i := 0; i < 10; i++ {
		go increment()
		go decrement()
	}

	// This is to make sure that main goroutine does not exit before child routines
	// This input will halt the execution of the main goroutine and meanwhile child rountines will be finished.

	// Note: There are better ways to acheive the same but we are doing it here brute force(in example one) and will improve later

	var input string
	fmt.Scanln(&input)
}

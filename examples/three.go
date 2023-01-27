package examples

import (
	"fmt"
	"sync"
)

// Please note that channel will close for the given producer once it's done producing
// Consumer will stop once the channel is closed(range ch)
func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("Producing ", i, "from producer ", id)
		ch <- i
	}
	close(ch)
	wg.Done()

}

func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("Consumed ", i, " by consumer ", id)
	}
	wg.Done()
}

func ExecuteThree() {
	var wg sync.WaitGroup
	wg.Add(4)

	ch := make(chan int)

	go producer(1, ch, &wg)

	go consumer(1, ch, &wg)
	go consumer(2, ch, &wg)
	go consumer(3, ch, &wg)

	wg.Wait()
}

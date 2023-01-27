package examples

import (
	"fmt"
	"sync"
)

func processElemnt(val int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Processing : ", val)
}

func ExecuteTwo() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup

	for _, val := range arr {
		wg.Add(1)
		go processElemnt(val, &wg)
	}

	wg.Wait()

	fmt.Println("All goroutnies are completed")

}

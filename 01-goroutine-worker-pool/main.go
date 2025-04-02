package main

import (
	"fmt"
	"sync"
)

func generateTasks(n int) <-chan int {
	tasks := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			tasks <- i
		}
		close(tasks)
	}()
	return tasks
}

func workers(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		result := task * task
		fmt.Printf("Worker %d processed task %d, result: %d ", id, task, result)
		results <- result
	}
}

func main() {
	numTasks := 100
	numWorkers := 5

	tasks := generateTasks(numTasks)
	results := make(chan int, numTasks)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go workers(i, tasks, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Collected Results: %d\n", result)
	}
}

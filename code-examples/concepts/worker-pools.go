package concepts

import "fmt"

// Worker pools are a model in which a fixed number of m workers (implemented in Go with goroutines)
// work their way through n tasks in a work queue (implemented in Go with a channel).
// This is very useful for breakable task that require a lot of processing.

// Fibonacci recursive fibonacci function
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fib)
		results <- fib
	}
}

func WorkerPools() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	workersNum := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < workersNum; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}

	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

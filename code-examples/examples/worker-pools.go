package examples

import "fmt"

// recursive fibonacci function
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

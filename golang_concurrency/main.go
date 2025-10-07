package main

import (
	"fmt"
	"sync"
)

// func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for job := range jobs {
// 		result := job * 2
// 		results <- result
// 	}
// }

// func main() {
// 	const numWorkers = 3
// 	jobs := make(chan int, 100)

// 	results := make(chan int, 100)

// 	var wg sync.WaitGroup

// 	for i := 0; i < numWorkers; i++ {
// 		wg.Add(1)
// 		go worker(i, jobs, results, &wg)
// 	}

// 	totalJobs := 10

// 	for j := 1; j <= totalJobs; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		fmt.Println("Result:", result)
// 	}

// 	// Start workers

// }

// Implement the fanout pattern
// queue setup should be from the queue -> routing
// how to do the routing here ->
//

// Approach :
// Input channel -> distribute work to different workers
// Worker Pool -> process in parallel
// Merge Results -> single output (fan-in)

// I will be starting three workers
// I am sending job data into a channel, and the data is consumed from the channel and results are pushed into different channel
// This is fan-out and fan-in approacn

// Implementing the fan-out and fan-in approach

// Basic Implementation
// func BasicWorker(Id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for job := range jobs {
// 		job = job * 2
// 		results <- job
// 	}
// }

// func BasicMain() {

// 	jobs := make(chan int, 100)
// 	results := make(chan int, 100)

// 	workers := 3

// 	var wg sync.WaitGroup

// 	for i := 0; i < workers; i++ {
// 		wg.Add(1)
// 		go worker(i, jobs, results, &wg)
// 	}

// 	for j := 1; j <= 10; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)

// 	// In a serperate go routine

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		fmt.Println("Result:", result)
// 	}

// }

// context Cancellation - stop workers early
// some times, If we want to cancel a wokers when any error occurs in one of them, user hits ctrl + c

// func worker(Id int, ctx context.Context, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Printf("Worker %d stopping\n", Id)
// 			return // exit early
// 		case job, ok := <-jobs:
// 			if !ok {
// 				return // Jobs channel closed
// 			}

// 			// Simulate work
// 			time.Sleep(600 * time.Millisecond)
// 			result <- job * 2
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
// 	defer cancel()

// 	jobs := make(chan int, 10)
// 	result := make(chan int, 10)

// 	var wg sync.WaitGroup

// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go worker(i, ctx, jobs, result, &wg)
// 	}

// 	for i := 1; i <= 10; i++ {
// 		select {
// 		case jobs <- i:
// 		case <-ctx.Done():
// 			fmt.Println("Stopped Job execution")
// 			break
// 		}
// 	}

// 	close(jobs)

// 	go func() {
// 		wg.Wait()
// 		close(result)
// 	}()

// 	for res := range result {
// 		fmt.Println("Got:", res)
// 	}

// 	fmt.Println("Done")

// }

// This is all fine and good, we need to control the backPressure here
// so each channel, To control the back pressure

type Result struct {
	Value int
}

func worker(Id int, jobs <-chan int, out chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		out <- Result{Value: job * 2}
	}
	close(out)
}

func merge(cs []chan Result) <-chan Result {
	var wg sync.WaitGroup

	out := make(chan Result)

	outPut := func(c chan Result) {
		defer wg.Done()
		for m := range c {
			out <- m
		}
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go outPut(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}

func main() {
	numWorkers := 3

	jobs := make(chan int, 10)

	resultsChannel := make([]chan Result, numWorkers)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		resultsChannel[i] = make(chan Result)
		wg.Add(1)
		go worker(i, jobs, resultsChannel[i], &wg)
	}

	for i := 1; i <= 9; i++ {
		// push this job in to the channels
		jobs <- i
	}
	close(jobs)

	merged := merge(resultsChannel)

	for res := range merged {
		fmt.Println("Merged result:", res.Value)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data string
}

type Result struct {
	Job    Job
	Output string
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)

		// Execute task
		task()

		result := Result{
			Job:    job,
			Output: fmt.Sprintf("Processed by worker %d: %s", id, job.Data),
		}

		results <- result
	}

	fmt.Printf("Worker %d finished\n", id)
}

func task() {
	// Simulate work
	time.Sleep(time.Second)
}

func saveToDB(batch []Result) {
	fmt.Printf("Saving batch of %d results to DB: ", len(batch))
	for i, result := range batch {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("Job%d", result.Job.ID)
	}
	fmt.Println()

	// Simulate DB save operation
	time.Sleep(100 * time.Millisecond)
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	for i := 1; i <= numJobs; i++ {
		job := Job{
			ID:   i,
			Data: fmt.Sprintf("job_%d_data", i),
		}
		jobs <- job
	}
	close(jobs)

	// Close results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results and save to DB in batches of 2
	fmt.Println("\nResults:")
	var batch []Result
	batchSize := 2

	for result := range results {
		fmt.Printf("Job %d result: %s\n", result.Job.ID, result.Output)
		batch = append(batch, result)

		if len(batch) == batchSize {
			saveToDB(batch)
			batch = nil // reset batch
		}
	}

	// Save remaining results if any
	if len(batch) > 0 {
		saveToDB(batch)
	}

	fmt.Println("\nAll jobs completed!")
}

// Output:
// % go run main.go

// Results:
// Worker 1 processing job 1
// Worker 3 processing job 3
// Worker 2 processing job 2
// Worker 3 processing job 4
// Worker 1 processing job 5
// Job 3 result: Processed by worker 3: job_3_data
// Job 1 result: Processed by worker 1: job_1_data
// Saving batch of 2 results to DB: Job3, Job1
// Worker 2 processing job 6
// Job 2 result: Processed by worker 2: job_2_data
// Worker 3 processing job 7
// Worker 1 processing job 8
// Worker 2 processing job 9
// Job 4 result: Processed by worker 3: job_4_data
// Saving batch of 2 results to DB: Job2, Job4
// Job 5 result: Processed by worker 1: job_5_data
// Job 6 result: Processed by worker 2: job_6_data
// Saving batch of 2 results to DB: Job5, Job6
// Worker 2 processing job 10
// Worker 3 finished
// Worker 1 finished
// Job 9 result: Processed by worker 2: job_9_data
// Job 7 result: Processed by worker 3: job_7_data
// Saving batch of 2 results to DB: Job9, Job7
// Job 8 result: Processed by worker 1: job_8_data
// Worker 2 finished
// Job 10 result: Processed by worker 2: job_10_data
// Saving batch of 2 results to DB: Job8, Job10

// All jobs completed!

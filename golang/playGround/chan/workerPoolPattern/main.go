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

	// for {
	// 	job, ok := <-jobs
	// 	if !ok {
	// 		break
	// 	}
	// }
	for job := range jobs { // for ... job, ok := <-jobsと同じ。データがjobs channelを通して「関数に」入ってくる(受信)
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)

		// Execute task
		task()

		result := Result{
			Job:    job,
			Output: fmt.Sprintf("Processed by worker %d: %s", id, job.Data),
		}

		results <- result // データが「関数から」results channelに流れていく(送信)
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
		// fmt.Printf("Job %d result: %s\n", result.Job.ID, result.Output)
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
// Worker 2 processing job 2
// Worker 3 processing job 3
// Worker 3 processing job 4
// Worker 2 processing job 5
// Saving batch of 2 results to DB: Job3Worker 1 processing job 6
// , Job2
// Worker 1 processing job 7
// Worker 3 processing job 9
// Saving batch of 2 results to DB: Worker 2 processing job 8
// Job1, Job6
// Saving batch of 2 results to DB: Job5, Job4
// Worker 3 processing job 10
// Saving batch of 2 results to DB: Job9, Job7
// Worker 1 finished
// Worker 2 finished
// Worker 3 finished
// Saving batch of 2 results to DB: Job8, Job10

// All jobs completed!

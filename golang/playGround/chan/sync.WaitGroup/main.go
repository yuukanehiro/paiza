package main

import (
    "fmt"
    "log"
    "runtime/debug"
    "sync"
    "time"
)

type Result struct {
    Job int
    Val int
    Err error
}

func worker(id int, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
    defer wg.Done()

    for job := range jobs {
        // 1ジョブ単位のガード：この中でpanicしても必ず1件返す
        func(job int) {
            defer func() {
                if r := recover(); r != nil {
                    results <- Result{Job: job, Err: fmt.Errorf("worker %d panic: %v", id, r)}
                    log.Printf("[panic recovered] worker=%d job=%d\n%s", id, job, debug.Stack())
                }
            }()

            // 例: job==3 でわざとpanic
            if job == 3 {
                panic("boom")
            }

            fmt.Printf("Worker %d started job %d\n", id, job)
            time.Sleep(time.Second)
            fmt.Printf("Worker %d finished job %d\n", id, job)
            results <- Result{Job: job, Val: job * 2, Err: nil}
        }(job)
    }
}

func main() {
    jobs := make(chan int, 5)
    results := make(chan Result, 5)

    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, jobs, results, &wg)
    }

    // ジョブ投入
    total := 5
    for j := 1; j <= total; j++ {
        jobs <- j
    }
    close(jobs)

    // ワーカー終了後にresultsをcloseしてrangeで受信するパターン
    go func() {
        wg.Wait()
        close(results)
    }()

    // 受信：成功でも失敗でも必ず1件来るので詰まらない
    got := 0
    for r := range results {
        got++
        if r.Err != nil {
            fmt.Printf("Job %d failed: %v\n", r.Job, r.Err)
            continue
        }
        fmt.Printf("Job %d result: %d\n", r.Job, r.Val)
    }
    fmt.Println("received:", got, "messages (should equal jobs)")
}

// Output:
// % go run main.go
// Worker 2 started job 2
// Worker 1 started job 1
// Job 3 failed: worker 3 panic: boom
// 2025/08/13 13:55:35 [panic recovered] worker=3 job=3
// goroutine 36 [running]:
// runtime/debug.Stack()
//         /opt/homebrew/Cellar/go/1.24.4/libexec/src/runtime/debug/stack.go:26 +0x64
// main.worker.func1.1()
//         /Users/kanehiroyuu/Documents/GitHub/paiza/golang/playGround/chan/main.go:26 +0xc0
// panic({0x104f01fa0?, 0x104f1cdc0?})
//         /opt/homebrew/Cellar/go/1.24.4/libexec/src/runtime/panic.go:792 +0x124
// main.worker.func1(0x0?, 0x0?, 0x0?)
//         /Users/kanehiroyuu/Documents/GitHub/paiza/golang/playGround/chan/main.go:32 +0x190
// main.worker(0x3, 0x1400012e000, 0x14000130000, 0x0?)
//         /Users/kanehiroyuu/Documents/GitHub/paiza/golang/playGround/chan/main.go:39 +0x64
// created by main.main in goroutine 1
//         /Users/kanehiroyuu/Documents/GitHub/paiza/golang/playGround/chan/main.go:50 +0x70
// Worker 3 started job 4
// Worker 3 finished job 4
// Worker 2 finished job 2
// Worker 3 started job 5
// Job 4 result: 8
// Job 2 result: 4
// Worker 1 finished job 1
// Job 1 result: 2
// Worker 3 finished job 5
// Job 5 result: 10
// received: 5 messages (should equal jobs)
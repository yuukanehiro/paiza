package main

import (
    "context"
    "fmt"
    "log"
    "runtime/debug"
    "time"

    "golang.org/x/sync/errgroup"
)

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int) error {
    for job := range jobs {
        // panicガード
        defer func() {
            if r := recover(); r != nil {
                log.Printf("[panic recovered] worker=%d job=%d: %v\n%s", id, job, r, debug.Stack())
            }
        }()

        select {
        case <-ctx.Done():
            return ctx.Err() // キャンセル時に終了
        default:
        }

        if job == 3 {
            return fmt.Errorf("worker %d failed on job %d", id, job) // エラー返すと全体キャンセル
        }

        fmt.Printf("Worker %d started job %d\n", id, job)
        time.Sleep(time.Second)
        fmt.Printf("Worker %d finished job %d\n", id, job)
        results <- job * 2
    }
    return nil
}

func main() {
    jobs := make(chan int, 5)
    results := make(chan int, 5)

    ctx := context.Background()
    g, ctx := errgroup.WithContext(ctx)

    // ワーカー起動
    for i := 1; i <= 3; i++ {
        id := i
        g.Go(func() error {
            return worker(ctx, id, jobs, results)
        })
    }

    // ジョブ投入
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // 別ゴルーチンで結果を受け取る
    g.Go(func() error {
        for {
            select {
            case <-ctx.Done():
                return nil // 全体キャンセル時に終了
            case r, ok := <-results:
                if !ok {
                    return nil
                }
                fmt.Println("Result:", r)
            }
        }
    })

    if err := g.Wait(); err != nil {
        log.Println("errgroup returned:", err)
    }
    close(results)
}

# Output:
// Worker 2 started job 2
// Worker 1 started job 1
// Worker 2 finished job 2
// Worker 1 finished job 1
// 2025/08/13 14:03:19 errgroup returned: worker 3 failed on job 3
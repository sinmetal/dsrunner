package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func goQueryKeysOnly(ds *DatastoreStore, goroutine int, endCh chan<- error) {
	go func() {
		for {
			var wg sync.WaitGroup
			for i := 0; i < goroutine; i++ {
				i := i
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					fmt.Printf("%+v goQueryKeysOnly GoRoutine:%d\n", time.Now(), i)
					defer func(n time.Time) {
						fmt.Printf("goQueryKeysOnly: %v\n", time.Since(n))
					}(time.Now())
					ctx := context.Background()

					var cancel context.CancelFunc
					if _, hasDeadline := ctx.Deadline(); !hasDeadline {
						ctx, cancel = context.WithTimeout(ctx, 2*time.Second)
						defer cancel()
					}
					if err := ds.QueryKeysOnly(ctx); err != nil {
						endCh <- err
					}

					time.Sleep(100 * time.Millisecond)
				}(i)
			}
			wg.Wait()
		}
	}()
}

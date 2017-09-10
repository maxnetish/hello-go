package utils

import (
	"time"
	"fmt"
)

func LongJob(jobId int, hardeness int64, iterations int, progress chan string) {
	dur := time.Duration(hardeness) * time.Millisecond
	for iter := 0; iter < iterations; iter++ {
		tn := time.Now()
		time.Sleep(dur)
		progress <- fmt.Sprintf("Async job %v internal step #%v done at %v.", jobId, iter, tn)
	}
	progress <- fmt.Sprintf("Async job %v done", jobId)
	close(progress)
}

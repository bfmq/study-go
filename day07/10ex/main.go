package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	value int64
}

type Result struct {
	Job *Job
	sum int64
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)
var wg sync.WaitGroup

func zhoulin(jc chan<- *Job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &Job{
			value: x,
		}
		jc <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func baodelu(jc <-chan *Job, rc chan<- *Result) {
	defer wg.Done()
	for {
		job := <-jc
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}

		result := &Result{
			Job: job,
			sum: sum,
		}
		rc <- result
	}
}

func main() {
	wg.Add(1)
	go zhoulin(jobChan)

	wg.Add(24)
	for i := 0; i < 24; i++ {
		go baodelu(jobChan, resultChan)
	}

	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.Job.value, result.sum)
	}

	wg.Wait()
}

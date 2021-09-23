package main

var notifyCh = make(chan struct{}, 5)

func worker(w int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * 2
		notifyCh <- struct{}{}
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go func() {
		for j := 1; j < 5; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	go func() {
		for i := 1; i < 5; i++ {
			<-notifyCh
		}
		close(results)
	}()

	for r := 1; r < 5; r++ {
		<-results
	}
	// close(results)
}

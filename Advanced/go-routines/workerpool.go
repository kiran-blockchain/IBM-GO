package main

import (
	"fmt"
	"time"
)

func main() {
	numberOfJobs := 10
	startTime := time.Now()
	jobs := make(chan int, numberOfJobs)
	result := make(chan int, numberOfJobs)

	for x := 1; x < 11; x++ {
		go worker(x, jobs, result)
	}

	for j:=1;j<numberOfJobs;j++ {
		jobs <- j
	}
	close(jobs)
	for a:=1;a<numberOfJobs; a++ {
		<-result
	}
	endTime:= time.Now()
	fmt.Println(endTime.Sub(startTime))
}

func worker(id int, jobs <-chan int, r chan<- int) {

	for j := range jobs {
		fmt.Println("Workers Id",id,"started job" ,j)
		time.Sleep(time.Second)
		fmt.Println("Worker id", id, "finished job",j)
		r <- j*2
	}

}
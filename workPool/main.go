package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id      int32
	RandNum int32
}

type Result struct {
	job      *Job
	sum      int32
	goFunNum int32
}

func main() {
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)
	createPool(4, jobChan, resultChan)

	// 打印结果
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("gofunNUm %v job id:%v randnum:%v result:%d\n", result.goFunNum, result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)

	// 生产者
	for id := 0; ; id++ {
		rNum := rand.Int31()
		jobChan <- &Job{
			Id:      int32(id),
			RandNum: rNum,
		}
	}
}

func createPool(
	num int,
	jobChan chan *Job,
	resultChan chan *Result,
) {
	for i := 0; i < num; i++ {
		chanNum := i
		go func(
			jobChan chan *Job,
			resultChan chan *Result,
			chanNum int,
		) {
			// 消费
			for job := range jobChan {
				rNum := job.RandNum
				var sum int32
				for rNum != 0 {
					sum += rNum % 10
					rNum /= 10
				}
				resultChan <- &Result{
					job:      job,
					sum:      sum,
					goFunNum: int32(chanNum),
				}
			}
		}(jobChan, resultChan, chanNum)
	}
}

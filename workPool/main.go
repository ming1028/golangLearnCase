package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
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
	nowDate := time.Now()
	time.Sleep(time.Second * 2)
	//endDate := time.Now()
	fmt.Println("耗时：", time.Since(nowDate))
	time.Sleep(time.Hour)
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

	go func() {
		fmt.Println(http.ListenAndServe("0.0.0.0:10000", nil))
	}()

	for id := 0; id < 20000; id++ {
		// rNum := rand.Int31()
		jobChan <- &Job{
			Id:      int32(id),
			RandNum: int32(id),
		}
		time.Sleep(time.Second)
	}
	// 生产者
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

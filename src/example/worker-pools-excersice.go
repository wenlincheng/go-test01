package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	计算所输入数字的每一位的和
*/

// 作业
type Job struct {
	id int
	// 要计算的数字
	randomNo int
}

// 结果
type Result struct {
	// 结果对应的任务
	job Job
	// 计算的结果
	sumOfDigits int
}

// 接收作业和写入结果的缓冲信道
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// 工作协程
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	return sum
}

// 创建工作协程的函数
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomNo)}
		results <- output
	}
	wg.Done()
}

// 创建了一个 Go 协程的工作池
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	// 关闭信道不再向信道发送数据
	close(results)
}

// 创建作业
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNo := rand.Intn(999)
		job := Job{i, randomNo}
		jobs <- job
	}
	close(jobs)
}

// 打印结果
func result(done chan bool) {
	for result := range results {
		fmt.Printf("任务id:%d 计算的数字:%d 计算结果:%d\n", result.job.id, result.job.randomNo, result.sumOfDigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	// 创建作业
	noOfJobs := 100
	go allocate(noOfJobs)
	// 打印结果
	done := make(chan bool)
	go result(done)
	// 创建工作池
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	// main 函数会监听 done 信道的通知，等待所有结果打印结束
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Printf("总共耗时：%f\n", diff.Seconds())
}

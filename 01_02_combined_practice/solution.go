package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// 参考答案 - 仅供学习，先自己尝试！
// ============================================

// 练习 1 参考答案: 生产者-消费者模式
func exercise1Solution() {
	// 创建 channel（可以无缓冲或小缓冲）
	jobs := make(chan int, 5)

	// 创建 WaitGroup 等待消费者完成
	var wg sync.WaitGroup

	// 生产者
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("  [生产者] 生成任务 %d\n", i)
			jobs <- i
			time.Sleep(50 * time.Millisecond)
		}
		close(jobs) // ⚠️ 重要：生产完毕后关闭 channel
	}()

	// 启动 2 个消费者
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(consumerID int) {
			defer wg.Done()
			for job := range jobs { // ✅ range 会在 channel 关闭后自动退出
				fmt.Printf("  [消费者 %d] 处理任务 %d\n", consumerID, job)
				time.Sleep(100 * time.Millisecond)
			}
			fmt.Printf("  [消费者 %d] 完成所有任务\n", consumerID)
		}(i)
	}

	wg.Wait() // 等待所有消费者完成
	fmt.Println("  [主程序] 所有任务处理完毕")
}

// 练习 2 参考答案: 工作池模式
func exercise2Solution() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动 workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				result := job * 2
				fmt.Printf("  Worker %d 处理任务 %d，结果 %d\n", workerID, job, result)
				time.Sleep(100 * time.Millisecond)
				results <- result
			}
		}(w)
	}

	// 发送任务
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs) // ⚠️ 关闭 jobs 让 workers 知道没有更多任务
	}()

	// 等待所有 workers 完成，然后关闭 results
	go func() {
		wg.Wait()
		close(results) // ⚠️ 所有 workers 完成后才能关闭 results
	}()

	// 收集结果
	fmt.Print("  结果: ")
	for result := range results {
		fmt.Printf("%d ", result)
	}
	fmt.Println()
}

// 练习 3 参考答案: 并发求和
func exercise3Solution() {
	partialSums := make(chan int, 4) // 缓冲大小 = goroutine 数量

	// 启动 4 个 goroutine 分段计算
	go func() {
		sum := 0
		for i := 1; i <= 25; i++ {
			sum += i
		}
		partialSums <- sum
	}()

	go func() {
		sum := 0
		for i := 26; i <= 50; i++ {
			sum += i
		}
		partialSums <- sum
	}()

	go func() {
		sum := 0
		for i := 51; i <= 75; i++ {
			sum += i
		}
		partialSums <- sum
	}()

	go func() {
		sum := 0
		for i := 76; i <= 100; i++ {
			sum += i
		}
		partialSums <- sum
	}()

	// 汇总结果
	totalSum := 0
	for i := 0; i < 4; i++ {
		totalSum += <-partialSums
	}

	fmt.Printf("  1 到 100 的和 = %d\n", totalSum)
	fmt.Println("  (参考答案: 5050)")
}

// 练习 4 参考答案: 数据管道
func exercise4Solution() {
	// 阶段 1: 生成器
	generator := func() <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 1; i <= 5; i++ {
				out <- i
			}
		}()
		return out
	}

	// 阶段 2: 平方计算
	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for n := range in {
				out <- n * n
			}
		}()
		return out
	}

	// 组装管道
	numbers := generator()
	squares := square(numbers)

	// 阶段 3: 打印结果
	fmt.Print("  管道输出: ")
	for result := range squares {
		fmt.Printf("%d ", result)
	}
	fmt.Println("(1²=1, 2²=4, 3²=9, 4²=16, 5²=25)")
}

// 如果你想直接运行参考答案，取消下面的注释
/*
func main() {
	fmt.Println("=== 参考答案 ===\n")

	fmt.Println("【练习 1】生产者-消费者")
	exercise1Solution()

	fmt.Println("\n【练习 2】工作池")
	exercise2Solution()

	fmt.Println("\n【练习 3】并发求和")
	exercise3Solution()

	fmt.Println("\n【练习 4】数据管道")
	exercise4Solution()
}
*/

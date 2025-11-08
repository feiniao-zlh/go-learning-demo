package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== Goroutine + Channel 综合实战 ===")
	fmt.Println("提示: 这些都是面试高频场景！")

	// 练习 1: 生产者-消费者模式
	fmt.Println("【练习 1】生产者-消费者模式")
	exercise1()

	// 练习 2: 工作池模式 (Worker Pool)
	fmt.Println("\n【练习 2】工作池模式")
	exercise2()

	// 练习 3: 并发求和
	fmt.Println("\n【练习 3】并发求和")
	exercise3()

	// 练习 4: 数据管道 (Pipeline)
	fmt.Println("\n【练习 4】数据管道")
	exercise4()
}

// ============================================
// 练习 1: 生产者-消费者模式 🏭
// 场景: 一个生产者生成数据，多个消费者处理数据
// 面试考点: channel 的创建、发送、接收、关闭
// ============================================
func exercise1() {
	// TODO: 创建一个 int 类型的 channel
	// 提示: jobs := make(chan int, ?)  需要缓冲吗？
	// 你的代码:
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// TODO: 启动生产者 goroutine
	// 它应该：
	// 1. 生成 5 个任务（发送数字 1-5 到 channel）
	// 2. 完成后关闭 channel
	// 提示:
	// go func() {
	//     for i := 1; i <= 5; i++ {
	//         jobs <- i
	//     }
	//     close(jobs)
	// }()
	// 你的代码:
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	// TODO: 启动 2 个消费者 goroutine
	// 每个消费者应该：
	// 1. 从 jobs channel 接收任务
	// 2. 打印 "消费者 X 处理任务 Y"
	// 3. 模拟处理时间（100ms）
	// 提示: 用 range 遍历 channel
	// 你的代码:
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("g%d ===> j%d\n", i, job)
			}
		}(i)
	}

	wg.Wait()
	// TODO: 等待所有消费者完成
	// 问题：怎么知道所有消费者都处理完了？
	// 提示: 可以简单用 time.Sleep，但更好的方式是什么？
	// 你的代码:

}

// ============================================
// 练习 2: 工作池模式 (Worker Pool) 👷
// 场景: 固定数量的 worker 处理大量任务
// 面试考点: 多个 goroutine 共享 channel，WaitGroup 同步
// ============================================
func exercise2() {
	const numWorkers = 3 // 3 个 worker
	const numJobs = 10   // 10 个任务

	// TODO: 创建两个 channel
	// jobs: 用于发送任务（缓冲大小 = numJobs）
	// results: 用于接收结果（缓冲大小 = numJobs）
	// 你的代码:

	// TODO: 启动 3 个 worker goroutine
	// 每个 worker 应该：
	// 1. 从 jobs channel 接收任务 id
	// 2. "处理"任务（可以简单地将 id * 2）
	// 3. 将结果发送到 results channel
	// 4. 打印 "Worker X 处理任务 Y，结果 Z"
	// 提示:
	// for w := 1; w <= numWorkers; w++ {
	//     go func(workerID int) {
	//         for job := range jobs {
	//             // 处理任务
	//             result := job * 2
	//             results <- result
	//         }
	//     }(w)
	// }
	// 你的代码:

	// TODO: 发送任务
	// 发送 10 个任务到 jobs channel，然后关闭它
	// 你的代码:

	// TODO: 接收结果
	// 从 results channel 接收所有结果并打印
	// 提示: 已知有 numJobs 个结果
	// 你的代码:

}

// ============================================
// 练习 3: 并发求和 ➕
// 场景: 多个 goroutine 分段计算，汇总结果
// 面试考点: goroutine 协作，结果汇总
// ============================================
func exercise3() {
	// 目标: 计算 1 到 100 的和，用 4 个 goroutine 分段计算
	// goroutine 1: 1-25
	// goroutine 2: 26-50
	// goroutine 3: 51-75
	// goroutine 4: 76-100

	// TODO: 创建一个 int 类型的 channel 用于接收部分和
	// 你的代码:
	res := make(chan int, 4)
	var wg sync.WaitGroup

	// TODO: 启动 4 个 goroutine，每个计算一段的和
	// 提示:
	// go func(start, end int) {
	//     sum := 0
	//     for i := start; i <= end; i++ {
	//         sum += i
	//     }
	//     // 发送部分和
	// }(1, 25)
	// 你的代码:
	for i := 1; i+24 <= 100; {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			sum := 0
			for j := start; j <= end; j++ {
				sum += j
			}
			fmt.Printf("%d=>%d=%d\n", start, end, sum)
			res <- sum
		}(i, i+24)
		i += 25
	}

	wg.Wait()
	close(res)
	ans := 0
	for re := range res {
		fmt.Println(re)
		ans += re
	}
	fmt.Println(ans)
}

// ============================================
// 练习 4: 数据管道 (Pipeline) 🚰
// 场景: 数据经过多个阶段处理（类似 Unix 管道）
// 面试考点: channel 链式传递，goroutine 协作
// ============================================
func exercise4() {
	// 目标: 构建一个三阶段管道
	// 阶段1: 生成数字 1-5
	// 阶段2: 将每个数字平方
	// 阶段3: 打印结果
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for item := range ch1 {
			ch2 <- item * item
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}

	// TODO: 实现生成器函数
	// 它应该返回一个只读 channel (<-chan int)
	// 在 goroutine 中生成 1-5，然后关闭 channel
	// 提示:
	// generator := func() <-chan int {
	//     out := make(chan int)
	//     go func() {
	//         defer close(out)
	//         for i := 1; i <= 5; i++ {
	//             out <- i
	//         }
	//     }()
	//     return out
	// }
	// 你的代码:

	// TODO: 实现平方函数
	// 它接收一个 channel，返回一个新 channel
	// 从输入 channel 读取数字，计算平方，发送到输出 channel
	// 提示:
	// square := func(in <-chan int) <-chan int {
	//     out := make(chan int)
	//     go func() {
	//         defer close(out)
	//         for n := range in {
	//             out <- n * n
	//         }
	//     }()
	//     return out
	// }
	// 你的代码:

	// TODO: 组装管道并打印结果
	// 提示:
	// numbers := generator()
	// squares := square(numbers)
	// for result := range squares {
	//     fmt.Printf("  %d ", result)
	// }
	// 你的代码:

}

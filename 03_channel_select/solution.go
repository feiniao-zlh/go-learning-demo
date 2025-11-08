package main

import (
	"fmt"
	"time"
)

// 这个文件包含所有练习的参考答案
// 建议: 先自己完成 main.go 中的练习，再查看这个文件

func mainSolution() {
	fmt.Println("=== Select 多路复用 - 参考答案 ===")

	exercise1Solution()
	exercise2Solution()
	exercise3Solution()
	exercise4Solution()
	exercise5Solution()
}

// 练习 1 答案
func exercise1Solution() {
	fmt.Println("\n【练习 1 答案】Select 基础语法")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个 goroutine
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自 ch1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "来自 ch2"
	}()

	// 使用 select
	select {
	case msg1 := <-ch1:
		fmt.Println("收到:", msg1)
	case msg2 := <-ch2:
		fmt.Println("收到:", msg2)
	}

	// 答案: ch1 会先被接收，因为它只等待 100ms，比 ch2 快
}

// 练习 2 答案
func exercise2Solution() {
	fmt.Println("\n【练习 2 答案】多个 Channel 选择")

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// 启动三个 goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(50 * time.Millisecond)
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i <= 30; i += 10 {
			time.Sleep(50 * time.Millisecond)
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		for i := 100; i <= 300; i += 100 {
			time.Sleep(50 * time.Millisecond)
			ch3 <- i
		}
		close(ch3)
	}()

	// 循环接收
	for {
		select {
		case val, ok := <-ch1:
			if ok {
				fmt.Printf("ch1: %d\n", val)
			} else {
				fmt.Println("ch1 关闭")
				ch1 = nil // 关键：防止重复触发
			}
		case val, ok := <-ch2:
			if ok {
				fmt.Printf("ch2: %d\n", val)
			} else {
				fmt.Println("ch2 关闭")
				ch2 = nil
			}
		case val, ok := <-ch3:
			if ok {
				fmt.Printf("ch3: %d\n", val)
			} else {
				fmt.Println("ch3 关闭")
				ch3 = nil
			}
		}

		// 所有 channel 都关闭后退出
		if ch1 == nil && ch2 == nil && ch3 == nil {
			break
		}
	}
}

// 练习 3 答案
func exercise3Solution() {
	fmt.Println("\n【练习 3 答案】超时控制")

	// 场景 1: 快速响应
	fmt.Println("场景 1: 快速响应")
	ch1 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "快速响应的数据"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("✅ 收到数据:", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("❌ 超时了！")
	}

	// 场景 2: 慢速响应
	fmt.Println("\n场景 2: 慢速响应")
	ch2 := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "慢速响应的数据"
		// 注意: 这个数据会发送成功，但没人接收（因为已经超时了）
		// 这个 goroutine 会永久阻塞！这是 goroutine 泄漏的常见原因
	}()

	select {
	case msg := <-ch2:
		fmt.Println("✅ 收到数据:", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("❌ 超时了！")
	}

	time.Sleep(150 * time.Millisecond) // 等待观察
}

// 练习 4 答案
func exercise4Solution() {
	fmt.Println("\n【练习 4 答案】非阻塞操作")

	ch := make(chan int, 2)

	// 非阻塞发送
	fmt.Println("测试非阻塞发送:")
	for i := 1; i <= 3; i++ {
		select {
		case ch <- i:
			fmt.Printf("  ✅ 成功发送: %d\n", i)
		default:
			fmt.Printf("  ❌ 发送失败: %d (channel 已满)\n", i)
		}
	}

	// 非阻塞接收
	fmt.Println("\n测试非阻塞接收:")
	for i := 1; i <= 3; i++ {
		select {
		case val := <-ch:
			fmt.Printf("  ✅ 接收到: %d\n", val)
		default:
			fmt.Println("  ❌ 接收失败 (channel 为空)")
		}
	}
}

// 练习 5 答案
func exercise5Solution() {
	fmt.Println("\n【练习 5 答案】实战 - 多数据源聚合")

	type Result struct {
		source string
		data   string
	}

	results := make(chan Result)

	// 模拟三个 API
	go func() {
		time.Sleep(50 * time.Millisecond)
		results <- Result{"API-1", "快速数据"}
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		results <- Result{"API-2", "最新数据"}
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		results <- Result{"API-3", "备用数据"}
	}()

	// 获取第一个返回的结果
	select {
	case result := <-results:
		fmt.Printf("✅ 使用 %s 的数据: %s\n", result.source, result.data)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("❌ 所有 API 都超时了")
	}

	time.Sleep(200 * time.Millisecond) // 等待其他 goroutine 完成

	// 改进版本：使用 done channel 通知其他 goroutine 停止
	fmt.Println("\n改进版本 - 使用 done channel:")
	exercise5ImprovedSolution()
}

// 练习 5 改进版答案
func exercise5ImprovedSolution() {
	type Result struct {
		source string
		data   string
	}

	results := make(chan Result)
	done := make(chan struct{}) // 用于通知其他 goroutine 停止

	// 模拟三个 API（改进版）
	callAPI := func(name string, delay time.Duration, data string) {
		select {
		case <-time.After(delay):
			select {
			case results <- Result{name, data}:
				// 发送成功
			case <-done:
				// 已经有结果了，停止
				fmt.Printf("  %s 被取消\n", name)
			}
		case <-done:
			// 还在延迟中就被取消了
			fmt.Printf("  %s 被取消\n", name)
		}
	}

	go callAPI("API-1", 50*time.Millisecond, "快速数据")
	go callAPI("API-2", 150*time.Millisecond, "最新数据")
	go callAPI("API-3", 100*time.Millisecond, "备用数据")

	// 获取第一个返回的结果
	select {
	case result := <-results:
		fmt.Printf("✅ 使用 %s 的数据: %s\n", result.source, result.data)
		close(done) // 通知其他 goroutine 停止
	case <-time.After(200 * time.Millisecond):
		fmt.Println("❌ 所有 API 都超时了")
		close(done)
	}

	time.Sleep(100 * time.Millisecond) // 等待观察取消信息
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Channel 基础练习 ===\n")

	// 练习 1: 创建和使用无缓冲 channel
	exercise1()

	// 练习 2: 有缓冲 channel
	exercise2()

	// 练习 3: 关闭 channel
	exercise3()

	// 练习 4: Goroutine 之间通信
	exercise4()
}

// ============================================
// 练习 1: 无缓冲 channel
// ============================================
func exercise1() {
	fmt.Println("【练习 1】无缓冲 channel")

	// TODO: 创建一个无缓冲的 int 类型 channel
	// 提示: ch := make(chan int)
	// 你的代码:

	// TODO: 启动一个 goroutine，向 channel 发送数字 42
	// 提示: go func() { ch <- 42 }()
	// 你的代码:

	// TODO: 从 channel 接收数据并打印
	// 提示: value := <-ch
	// 你的代码:

	fmt.Println()
}

// ============================================
// 练习 2: 有缓冲 channel
// ============================================
func exercise2() {
	fmt.Println("【练习 2】有缓冲 channel")

	// TODO: 创建一个容量为 3 的 int 类型 channel
	// 提示: ch := make(chan int, 3)
	// 你的代码:

	// TODO: 不用 goroutine，直接发送 3 个数字：1, 2, 3
	// 提示: ch <- 1
	// 你的代码:

	// TODO: 接收并打印这 3 个数字
	// 你的代码:

	fmt.Println()
}

// ============================================
// 练习 3: 关闭 channel
// ============================================
func exercise3() {
	fmt.Println("【练习 3】关闭 channel")

	ch := make(chan int, 5)

	// 发送 5 个数字
	for i := 1; i <= 5; i++ {
		ch <- i
	}

	// TODO: 关闭 channel
	// 提示: close(ch)
	// 你的代码:

	// TODO: 用 range 遍历 channel，打印所有值
	// 提示: for value := range ch { ... }
	// 你的代码:

	fmt.Println()
}

// ============================================
// 练习 4: Goroutine 之间通信
// 目标: 两个 goroutine 通过 channel 协作
// ============================================
func exercise4() {
	fmt.Println("【练习 4】Goroutine 通信")

	// TODO: 创建一个 string 类型的 channel
	// 你的代码:

	// TODO: 启动生产者 goroutine
	// 它应该发送 3 条消息："消息1", "消息2", "消息3"，然后关闭 channel
	// 提示:
	// go func() {
	//     ch <- "消息1"
	//     ...
	//     close(ch)
	// }()
	// 你的代码:

	// TODO: 消费者：用 range 接收所有消息并打印
	// 你的代码:

	fmt.Println()
}

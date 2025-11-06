package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Goroutine 练习 ===")

	// 练习 1: 创建你的第一个 goroutine
	exercise1()

	// 练习 2: 创建多个 goroutine
	exercise2()

	// 练习 3: 闭包陷阱
	exercise3()
}

// 练习 1: 创建第一个 goroutine
// 目标: 让主函数和 goroutine 同时打印消息
func exercise1() {
	fmt.Println("【练习 1】创建第一个 goroutine")

	// TODO: 在这里用 go 关键字启动一个 goroutine
	// 让它调用 sayHello() 函数
	// 你的代码:
	go sayHello()

	fmt.Println("主函数继续执行")

	// TODO: 让主函数等待 1 秒，确保 goroutine 执行完
	// 你的代码:
	time.Sleep(time.Second)

	fmt.Println()
}

func sayHello() {
	fmt.Println("  → Hello from goroutine!")
}

// 练习 2: 创建 5 个 goroutine
// 目标: 每个 goroutine 打印自己的编号 (1-5)
func exercise2() {
	fmt.Println("【练习 2】创建多个 goroutine")

	// TODO: 用 for 循环创建 5 个 goroutine
	// 每个打印: "Goroutine X 正在运行"
	// 注意: 要把 i 作为参数传给 goroutine，避免闭包陷阱
	// 你的代码:
	for i := 0; i < 5; i++ {
		//fmt.Println("pre:", i)
		//i := i
		//fmt.Println("changed", i)
		go func() {
			fmt.Printf("Goroutine fun %d 正在运行\n", i)
		}()
		//go fmt.Printf("Goroutine %d 正在运行\n", i)
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println()
}

// 练习 3: 闭包陷阱 ⚠️
// 目标: 对比错误写法和正确写法
func exercise3() {
	fmt.Println("【练习 3】闭包陷阱")

	// ❌ 错误示例（会输出什么？）
	fmt.Println("错误写法:")
	for i := 1; i <= 3; i++ {
		go fmt.Println("pre:", i)
	}
	time.Sleep(100 * time.Millisecond)

	// ✅ 正确示例
	fmt.Println("正确写法:")
	for i := 1; i <= 3; i++ {
		// TODO: 把 i 作为参数传给 goroutine
		// 你的代码:
		go func(i int) {
			fmt.Println("after:", i)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Println()
}

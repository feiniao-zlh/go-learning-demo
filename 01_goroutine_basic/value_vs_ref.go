package main

import (
	"fmt"
	"time"
)

func testValueVsRef() {
	fmt.Println("=== 值传递 vs 引用传递 ===\n")

	// 方式1: 直接传参（值传递）✅
	fmt.Println("方式1: 直接传参")
	for i := 1; i <= 3; i++ {
		go fmt.Printf("  直接传参: i=%d\n", i)
	}
	time.Sleep(100 * time.Millisecond)

	// 方式2: 闭包引用（引用传递）❌
	fmt.Println("\n方式2: 闭包引用")
	for i := 1; i <= 3; i++ {
		go func() {
			fmt.Printf("  闭包引用: i=%d\n", i)
		}()
	}
	time.Sleep(100 * time.Millisecond)

	// 方式3: 闭包+传参（值传递）✅
	fmt.Println("\n方式3: 闭包传参")
	for i := 1; i <= 3; i++ {
		go func(num int) {
			fmt.Printf("  闭包传参: i=%d\n", num)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)

	// 详细解释
	fmt.Println("\n【解释】")
	fmt.Println("方式1: i 在 go 这一行就被求值，复制给 fmt.Printf")
	fmt.Println("方式2: i 在闭包执行时才被读取，此时循环已结束")
	fmt.Println("方式3: i 在 go 这一行被求值，复制给参数 num")
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
🎯 实战项目 3: 超时重试机制
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

💼 真实场景:
你在调用外部服务（数据库、微服务、第三方 API），但是：
- 网络可能不稳定，偶尔会超时
- 服务可能临时不可用
- 我们不能一次失败就放弃，要重试几次
- 但也不能无限重试，要设置总超时时间

这在分布式系统中非常常见！

📚 涉及知识点:
✓ Select - 超时控制
✓ Channel - 结果传递
✓ time.After - 超时检测
✓ 循环重试 - 错误处理

🔨 实现思路:
1. 尝试执行操作
2. 如果失败，等待一段时间后重试
3. 如果超过最大重试次数或总超时时间，返回失败
*/

// RetryConfig 重试配置
type RetryConfig struct {
	MaxRetries     int           // 最大重试次数
	InitialBackoff time.Duration // 初始退避时间
	MaxBackoff     time.Duration // 最大退避时间
	Timeout        time.Duration // 总超时时间
}

// DefaultRetryConfig 默认配置
var DefaultRetryConfig = RetryConfig{
	MaxRetries:     3,
	InitialBackoff: 100 * time.Millisecond,
	MaxBackoff:     2 * time.Second,
	Timeout:        10 * time.Second,
}

// ============================================
// 模拟不稳定的服务
// ============================================

// unstableService 模拟一个不稳定的服务
// 70% 概率失败，30% 概率成功
func unstableService(requestID int) (string, error) {
	// 模拟网络延迟
	delay := time.Duration(50+rand.Intn(200)) * time.Millisecond
	time.Sleep(delay)

	// 70% 失败率
	if rand.Float32() < 0.7 {
		return "", fmt.Errorf("服务暂时不可用")
	}

	return fmt.Sprintf("请求 %d 成功！数据: {id: %d, status: 'ok'}", requestID, requestID), nil
}

// ============================================
// 实现带超时的重试机制
// ============================================

// RetryWithTimeout 带超时的重试函数
func RetryWithTimeout(
	operation func() (string, error), // 要执行的操作
	config RetryConfig,
	requestID int,
) (string, error) {
	//startTime := time.Now()
	//backoff := config.InitialBackoff

	// TODO(human): 实现重试逻辑
	// 你需要实现以下功能：
	//
	// 1. 在循环中重试操作（最多 MaxRetries 次）
	// 2. 每次重试前检查是否已经总超时
	// 3. 如果操作失败，等待 backoff 时间后重试
	// 4. 每次重试后，增加 backoff 时间（指数退避）
	//
	// 提示代码结构：
	//
	// for attempt := 1; attempt <= config.MaxRetries; attempt++ {
	//     // 检查总超时
	//     if time.Since(startTime) > config.Timeout {
	//         return "", fmt.Errorf("总超时 (%v)", config.Timeout)
	//     }
	//
	//     fmt.Printf("  [尝试 %d/%d] 调用服务...\n", attempt, config.MaxRetries)
	//
	//     // 使用 channel + select 实现带超时的操作
	//     resultCh := make(chan string, 1)
	//     errorCh := make(chan error, 1)
	//
	//     go func() {
	//         result, err := operation()
	//         if err != nil {
	//             errorCh <- err
	//         } else {
	//             resultCh <- result
	//         }
	//     }()
	//
	//     // 等待结果或超时
	//     select {
	//     case result := <-resultCh:
	//         // 成功！
	//         fmt.Printf("  ✅ 成功！(第 %d 次尝试)\n", attempt)
	//         return result, nil
	//
	//     case err := <-errorCh:
	//         // 失败，准备重试
	//         fmt.Printf("  ❌ 失败: %v\n", err)
	//
	//         if attempt < config.MaxRetries {
	//             fmt.Printf("  ⏳ 等待 %v 后重试...\n", backoff)
	//             time.Sleep(backoff)
	//
	//             // 指数退避：每次失败后，等待时间翻倍
	//             backoff *= 2
	//             if backoff > config.MaxBackoff {
	//                 backoff = config.MaxBackoff
	//             }
	//         }
	//
	//     case <-time.After(5 * time.Second):
	//         // 单次调用超时
	//         fmt.Printf("  ⏰ 单次调用超时\n")
	//         if attempt >= config.MaxRetries {
	//             return "", fmt.Errorf("达到最大重试次数")
	//         }
	//     }
	// }
	//
	// return "", fmt.Errorf("达到最大重试次数 (%d)", config.MaxRetries)

	// 你的代码写在这里：
	// ==================== 开始实现 ====================

	// ==================== 结束实现 ====================

	return "", fmt.Errorf("请实现重试逻辑")
}

func mainT() {
	fmt.Println("🎯 实战项目 3: 超时重试机制")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	// 场景 1: 不重试（看看成功率多低）
	fmt.Println("📍 场景 1: 不重试（成功率很低）")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	successCount := 0
	for i := 1; i <= 10; i++ {
		result, err := unstableService(i)
		if err != nil {
			fmt.Printf("请求 %d: ❌ %v\n", i, err)
		} else {
			fmt.Printf("请求 %d: ✅ %s\n", i, result)
			successCount++
		}
	}
	fmt.Printf("\n成功率: %d/10 (%.0f%%)\n", successCount, float64(successCount)*10)
	fmt.Println()

	time.Sleep(1 * time.Second)

	// 场景 2: 使用重试机制
	fmt.Println("📍 场景 2: 使用重试机制（提高成功率）")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	config := RetryConfig{
		MaxRetries:     3,
		InitialBackoff: 100 * time.Millisecond,
		MaxBackoff:     1 * time.Second,
		Timeout:        10 * time.Second,
	}

	successCount = 0
	totalDuration := time.Duration(0)

	for i := 1; i <= 5; i++ {
		fmt.Printf("\n🔄 请求 %d:\n", i)
		start := time.Now()

		result, err := RetryWithTimeout(
			func() (string, error) {
				return unstableService(i)
			},
			config,
			i,
		)

		duration := time.Since(start)
		totalDuration += duration

		if err != nil {
			fmt.Printf("  ❌ 最终失败: %v (耗时: %v)\n", err, duration)
		} else {
			fmt.Printf("  ✅ 最终成功: %s (耗时: %v)\n", result, duration)
			successCount++
		}
	}

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📊 统计:")
	fmt.Printf("  成功率: %d/5 (%.0f%%)\n", successCount, float64(successCount)*20)
	fmt.Printf("  平均耗时: %v\n", totalDuration/5)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	// 场景 3: 演示指数退避
	fmt.Println("📍 场景 3: 指数退避演示")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("退避策略: 100ms -> 200ms -> 400ms -> 800ms -> 1s (max)")
	fmt.Println()

	backoff := 100 * time.Millisecond
	maxBackoff := 1 * time.Second

	for i := 1; i <= 5; i++ {
		fmt.Printf("第 %d 次重试: 等待 %v\n", i, backoff)
		backoff *= 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
	}

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("💡 总结:")
	fmt.Println("  • 重试可以大幅提高成功率（30% -> 90%+）")
	fmt.Println("  • 指数退避避免对服务造成压力")
	fmt.Println("  • 超时控制防止无限等待")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// 💡 思考题:
	// 1. 为什么要用"指数退避"而不是固定间隔重试？
	// 2. 如果所有客户端同时重试，会发生什么？（提示：雪崩）
	// 3. 如何添加"抖动"(jitter) 避免重试风暴？
	// 4. 什么时候应该放弃重试？（提示：4xx 错误 vs 5xx 错误）
}

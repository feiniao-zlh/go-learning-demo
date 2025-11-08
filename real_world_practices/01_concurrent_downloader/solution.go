package main

import (
	"fmt"
	"sync"
	"time"
)

// 💡 这个文件包含完整的参考答案
// 类型定义（File、DownloadResult、simulateDownload）在 main.go 中
// 这里只包含实现逻辑

// simulateDownloadWithCancel 模拟可取消的下载
// 会定期检查 done channel，如果关闭了就提前返回
func simulateDownloadWithCancel(file File, done chan struct{}) DownloadResult {
	// 模拟下载分成多个步骤，每个步骤检查是否取消
	steps := 5                             // 假设下载分 5 个步骤
	stepDuration := 100 * time.Millisecond // 每步耗时 100ms

	for i := 0; i < steps; i++ {
		// 检查是否收到取消信号
		select {
		case <-done:
			// 收到取消信号，立即返回
			return DownloadResult{
				FileID:   file.ID,
				FileName: file.Name,
				Success:  false,
				Error:    "下载被取消",
				Duration: time.Duration(i) * stepDuration,
				Timeout:  true,
			}
		default:
			// 继续执行
			time.Sleep(stepDuration)
		}
	}

	// 模拟 20% 失败率
	success := true
	if i := file.ID % 5; i == 0 {
		success = false
	}

	return DownloadResult{
		FileID:   file.ID,
		FileName: file.Name,
		Success:  success,
		Error: func() string {
			if !success {
				return "网络错误"
			}
			return ""
		}(),
		Duration: time.Duration(steps) * stepDuration,
	}
}

// runSolution 运行完整的参考答案
func runSolution() {
	fmt.Println("🎯 实战项目 1: 并发下载器（完整解法）")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	files := []File{
		{1, "image1.jpg", "https://example.com/image1.jpg"},
		{2, "image2.jpg", "https://example.com/image2.jpg"},
		{3, "video1.mp4", "https://example.com/video1.mp4"},
		{4, "document.pdf", "https://example.com/document.pdf"},
		{5, "music.mp3", "https://example.com/music.mp3"},
		{6, "image3.png", "https://example.com/image3.png"},
		{7, "video2.mp4", "https://example.com/video2.mp4"},
		{8, "data.csv", "https://example.com/data.csv"},
		{9, "archive.zip", "https://example.com/archive.zip"},
		{10, "report.xlsx", "https://example.com/report.xlsx"},
	}

	const maxWorkers = 3
	const downloadTimeout = 600 * time.Millisecond // 超时时间设置为 600ms

	startTime := time.Now()

	// ==================== 完整实现 ====================

	var wg sync.WaitGroup
	jobs := make(chan File, len(files))
	results := make(chan DownloadResult, len(files))

	// 🔹 步骤 1: 启动 workers（每个 worker 都有超时检测）
	for i := 1; i <= maxWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			// 🔸 每个 worker 循环处理任务
			for file := range jobs {
				fmt.Printf("⏳ [Worker %d] 开始下载: %s\n", workerID, file.Name)

				// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
				// 🎯 优化版：超时后取消任务
				// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

				resultCh := make(chan DownloadResult, 1)
				done := make(chan struct{}) // ← 新增：用于通知取消

				// 启动 goroutine 执行下载
				go func(f File, doneCh chan struct{}) {
					// 模拟下载过程，但会检查取消信号
					result := simulateDownloadWithCancel(f, doneCh)

					// 检查是否被取消
					select {
					case <-doneCh:
						// 被取消了，不发送结果
						fmt.Printf("  [Worker %d] 任务被取消，停止发送: %s\n", workerID, f.Name)
						return
					case resultCh <- result:
						// 没被取消，发送结果
						//fmt.Printf(" ")
					}
				}(file, done)

				// 等待结果或超时
				var result DownloadResult
				select {
				case result = <-resultCh:
					// ✅ 下载完成

				case <-time.After(downloadTimeout):
					// ⏰ 超时！关闭 done channel 通知 goroutine 停止
					close(done)

					result = DownloadResult{
						FileID:   file.ID,
						FileName: file.Name,
						Success:  false,
						Error:    "下载超时",
						Timeout:  true,
						Duration: downloadTimeout,
					}
					fmt.Printf("⏰ [Worker %d] 超时，已发送取消信号: %s\n", workerID, file.Name)
				}

				// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

				results <- result
			}
		}(i)
	}

	// 🔹 步骤 2: 发送所有任务
	for _, file := range files {
		jobs <- file
	}
	close(jobs) // 关闭 jobs，让 workers 知道没有更多任务了

	// 🔹 步骤 3: 启动一个 goroutine 等待所有 workers 完成
	//           完成后关闭 results channel
	go func() {
		wg.Wait()      // 等待所有 workers 完成
		close(results) // 关闭 results，让主线程知道不会再有结果了
	}()

	// 🔹 步骤 4: 主线程收集所有结果并统计
	successCount := 0
	failCount := 0
	timeoutCount := 0

	// 从 results 读取所有结果
	// 当 results 被关闭且清空后，range 循环会自动退出
	for result := range results {
		if result.Timeout {
			// 超时的情况
			fmt.Printf("❌ [超时] %s (耗时: %v)\n",
				result.FileName, result.Duration)
			timeoutCount++
			failCount++
		} else if result.Success {
			// 成功的情况
			fmt.Printf("✅ [成功] %s (耗时: %v)\n",
				result.FileName, result.Duration)
			successCount++
		} else {
			// 失败的情况（非超时）
			fmt.Printf("❌ [失败] %s - %s\n",
				result.FileName, result.Error)
			failCount++
		}
	}

	// ==================== 结束实现 ====================

	totalTime := time.Since(startTime)

	// 打印统计信息
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📊 下载完成！")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("总文件数: %d\n", len(files))
	fmt.Printf("✅ 成功: %d\n", successCount)
	fmt.Printf("❌ 失败: %d\n", failCount)
	fmt.Printf("⏰ 超时: %d\n", timeoutCount)
	fmt.Printf("⏱️  总耗时: %.2fs\n", totalTime.Seconds())
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// 💡 性能分析
	avgTimePerFile := float64(totalTime.Milliseconds()) / float64(len(files))
	fmt.Printf("\n💡 平均每个文件耗时: %.0fms\n", avgTimePerFile)

	// 如果是串行下载，预估时间
	estimatedSerialTime := time.Duration(250*len(files)) * time.Millisecond
	fmt.Printf("📈 如果串行下载预估: %.2fs\n", estimatedSerialTime.Seconds())
	fmt.Printf("🚀 并发加速比: %.1fx\n",
		float64(estimatedSerialTime)/float64(totalTime))
}

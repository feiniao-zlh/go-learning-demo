package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
🎯 实战项目 1: 并发下载器
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

💼 真实场景:
你在开发一个批量下载工具，需要同时下载多个文件。
- 有 10 个文件需要下载
- 为了快速完成，你想同时下载 3 个（限制并发数）
- 需要显示每个文件的下载进度
- 需要统计总共用了多少时间
- 如果某个下载失败了，不要影响其他文件

📚 涉及知识点:
✓ Goroutine - 并发执行下载任务
✓ Channel - 任务分发和结果收集
✓ Select - 超时控制
✓ WaitGroup - 等待所有任务完成（会在后面学，这里简单体验）

🔨 实现思路:
1. 创建一个任务队列 (jobs channel)
2. 启动固定数量的 worker goroutine (3个)
3. 每个 worker 从队列取任务、执行下载、报告结果
4. 主程序等待所有任务完成并统计
*/

// File 表示要下载的文件
type File struct {
	ID   int
	Name string
	URL  string
}

// DownloadResult 表示下载结果
type DownloadResult struct {
	FileID   int
	FileName string
	Success  bool
	Error    string
	Duration time.Duration
	Timeout  bool // 标记是否超时
}

// simulateDownload 模拟下载过程（实际项目中这里会是真实的 HTTP 请求）
func simulateDownload(file File) DownloadResult {
	// 模拟下载耗时（随机 100-500ms）
	duration := time.Duration(100+rand.Intn(400)) * time.Millisecond
	time.Sleep(duration)

	// 模拟 20% 的失败率
	success := rand.Float32() > 0.2

	result := DownloadResult{
		FileID:   file.ID,
		FileName: file.Name,
		Duration: duration,
		Success:  success,
	}

	if !success {
		result.Error = "网络错误或文件不存在"
	}

	return result
}

func main() {
	// 🔧 如果想看完整答案，将这里改成 true
	runAnswerInstead := false

	if runAnswerInstead {
		runSolution() // 运行 solution.go 中的完整答案
		return
	}

	// ============ 下面是你的练习代码 ============

	fmt.Println("🎯 实战项目 1: 并发下载器")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	// 准备 10 个要下载的文件
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

	const maxWorkers = 5 // 最多同时下载 3 个文件

	startTime := time.Now()

	// TODO(human): 实现并发下载逻辑
	// 你需要实现以下功能：
	//
	// 1. 创建两个 channel:
	//    - jobs: 用于发送下载任务
	//    - results: 用于接收下载结果
	//
	// 2. 启动 3 个 worker goroutine
	//    每个 worker 做什么：
	//    - 从 jobs channel 接收文件
	//    - 调用 simulateDownload(file) 下载
	//    - 将结果发送到 results channel
	//
	// 3. 发送所有任务到 jobs channel，然后关闭它
	//
	// 4. 接收所有结果并打印：
	//    - 成功: "✅ [Worker X] 下载成功: filename (耗时: XXXms)"
	//    - 失败: "❌ [Worker X] 下载失败: filename - 错误信息"
	//
	// 提示代码结构:
	//
	// jobs := make(chan File, len(files))
	// results := make(chan DownloadResult, len(files))
	//
	// // 启动 workers
	// for w := 1; w <= maxWorkers; w++ {
	//     go func(workerID int) {
	//         for file := range jobs {
	//             fmt.Printf("⏳ [Worker %d] 开始下载: %s\n", workerID, file.Name)
	//             result := simulateDownload(file)
	//             results <- result
	//         }
	//     }(w)
	// }
	//
	// // 发送任务
	// for _, file := range files {
	//     jobs <- file
	// }
	// close(jobs)
	//
	// // 收集结果
	// successCount := 0
	// failCount := 0
	// for i := 0; i < len(files); i++ {
	//     result := <-results
	//     if result.Success {
	//         fmt.Printf("✅ 下载成功: %s (耗时: %v)\n", result.FileName, result.Duration)
	//         successCount++
	//     } else {
	//         fmt.Printf("❌ 下载失败: %s - %s\n", result.FileName, result.Error)
	//         failCount++
	//     }
	// }

	// 你的代码写在这里:
	// ==================== 开始实现 ====================

	var wg sync.WaitGroup
	jobs := make(chan File, 10)
	results := make(chan DownloadResult, 10)
	for _, f := range files {
		jobs <- f
	}
	close(jobs)

	for i := 1; i <= maxWorkers; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			for j := range jobs {
				downloadResult := simulateDownload(j)
				results <- downloadResult
				fmt.Printf("%d 完成了任务 %d。任务结果：%v\n", num, j.ID, downloadResult)
			}

		}(i)
	}

	wg.Wait()
	close(results)

	// ==================== 结束实现 ====================

	totalTime := time.Since(startTime)

	// TODO(human): 打印最终统计信息
	// 格式如下：
	//
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// 📊 下载完成！
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// 总文件数: 10
	// 成功: X
	// 失败: X
	// 总耗时: XXXs
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	//
	// 你的代码:
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📊 下载完成！")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	// TODO: 打印统计信息
	fmt.Printf("总耗时: %.2fs\n", totalTime.Seconds())
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// 💡 思考题:
	// 1. 如果只用 1 个 worker，总耗时会是多少？
	// 2. 如果用 10 个 worker（每个文件一个），总耗时会是多少？
	// 3. 为什么我们限制 worker 数量为 3？（提示：现实中的资源限制）
	// 4. 如何添加超时控制（比如单个文件下载超过 1 秒就放弃）？
}

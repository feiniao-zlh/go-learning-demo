package main

import (
	"fmt"
	"sync"
)

/*
🎯 项目 1：并发爬虫（Go Tour 经典练习）
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📚 学习目标：
  ✓ 使用 sync.Mutex 保护共享数据
  ✓ 使用 sync.WaitGroup 等待所有 goroutine
  ✓ Goroutine 递归调用
  ✓ Map 去重

🔨 任务：
  实现 Crawl 函数，并发爬取网页，避免重复访问

📖 参考：https://go.dev/tour/concurrency/10
*/

// Fetcher 返回 URL 的 body 和该页面上的 URL 列表
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// SafeCounter 并发安全的访问记录
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]bool
}

// TODO(human): 实现 SafeCounter 的方法
//
// 提示：
// 1. 实现 Set(key string) - 标记 URL 已访问
// 2. 实现 Visited(key string) bool - 检查 URL 是否已访问
//
// 参考代码结构：
// func (c *SafeCounter) Set(key string) {
//     c.mu.Lock()
//     defer c.mu.Unlock()
//     c.v[key] = true
// }
//
// func (c *SafeCounter) Visited(key string) bool {
//     c.mu.Lock()
//     defer c.mu.Unlock()
//     return c.v[key]
// }

// 你的代码：



// Crawl 并发爬取 URL，深度最多为 depth
func Crawl(url string, depth int, fetcher Fetcher, visited *SafeCounter, wg *sync.WaitGroup) {
	defer wg.Done()

	// TODO(human): 实现爬虫逻辑
	//
	// 提示：
	// 1. 如果 depth <= 0，直接返回
	// 2. 检查 URL 是否已访问（用 visited.Visited）
	//    - 如果已访问，返回
	//    - 如果未访问，标记为已访问（用 visited.Set）
	// 3. 调用 fetcher.Fetch(url) 获取页面内容
	// 4. 打印结果
	// 5. 对每个子 URL，启动新的 goroutine 爬取
	//    - 记得 wg.Add(1)
	//    - 记得 depth-1
	//
	// 参考代码结构：
	// if depth <= 0 {
	//     return
	// }
	//
	// if visited.Visited(url) {
	//     return
	// }
	// visited.Set(url)
	//
	// body, urls, err := fetcher.Fetch(url)
	// if err != nil {
	//     fmt.Println(err)
	//     return
	// }
	// fmt.Printf("found: %s %q\n", url, body)
	//
	// for _, u := range urls {
	//     wg.Add(1)
	//     go Crawl(u, depth-1, fetcher, visited, wg)
	// }

	// 你的代码：


}

func main() {
	// 创建访问记录
	visited := &SafeCounter{v: make(map[string]bool)}

	// 创建 WaitGroup
	var wg sync.WaitGroup

	// 启动第一个爬虫
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, visited, &wg)

	// 等待所有爬虫完成
	wg.Wait()

	fmt.Println("\n爬虫完成！")
}

// ============================================
// 以下是模拟的网页数据（不用修改）
// ============================================

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是一个模拟的数据源
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

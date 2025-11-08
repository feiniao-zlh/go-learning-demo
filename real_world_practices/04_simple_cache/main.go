package main

import (
	"fmt"
	"sync"
	"time"
)

/*
🎯 实战项目 4: 并发安全的缓存系统
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

💼 真实场景:
你在开发一个 Web 应用，需要频繁查询数据库：
- 数据库查询很慢（100ms+）
- 相同的数据会被多次查询
- 多个 goroutine 可能同时访问缓存
- 缓存的数据需要定时过期

解决方案：实现一个并发安全的缓存！

📚 涉及知识点:
✓ sync.RWMutex - 读写锁（多读少写场景）
✓ Map - 存储缓存数据
✓ time.Time - 记录过期时间
✓ Goroutine - 后台清理过期数据

🔨 实现思路:
1. 使用 map 存储缓存数据
2. 使用 RWMutex 保护并发访问
3. 每个缓存项记录过期时间
4. 启动后台 goroutine 定期清理过期数据
*/

// CacheItem 缓存项
type CacheItem struct {
	Value      interface{} // 缓存的值
	ExpireTime time.Time   // 过期时间
}

// IsExpired 检查是否过期
func (item *CacheItem) IsExpired() bool {
	return time.Now().After(item.ExpireTime)
}

// Cache 缓存结构
type Cache struct {
	items map[string]*CacheItem
	mu    sync.RWMutex // 读写锁
}

// NewCache 创建缓存
func NewCache() *Cache {
	cache := &Cache{
		items: make(map[string]*CacheItem),
	}

	// 启动后台清理 goroutine
	go cache.cleanupLoop()

	return cache
}

// Set 设置缓存（带过期时间）
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	// TODO(human): 实现 Set 方法
	// 提示：
	// 1. 使用写锁保护（因为要修改 map）
	// 2. 创建 CacheItem，设置过期时间
	// 3. 存入 map
	//
	// c.mu.Lock()
	// defer c.mu.Unlock()
	//
	// c.items[key] = &CacheItem{
	//     Value:      value,
	//     ExpireTime: time.Now().Add(ttl),
	// }

	// 你的代码：

}

// Get 获取缓存
func (c *Cache) Get(key string) (interface{}, bool) {
	// TODO(human): 实现 Get 方法
	// 提示：
	// 1. 使用读锁保护（因为只读取 map）
	// 2. 检查 key 是否存在
	// 3. 检查是否过期
	// 4. 返回值和是否存在的标志
	//
	// c.mu.RLock()
	// defer c.mu.RUnlock()
	//
	// item, exists := c.items[key]
	// if !exists {
	//     return nil, false
	// }
	//
	// if item.IsExpired() {
	//     return nil, false
	// }
	//
	// return item.Value, true

	// 你的代码：
	return nil, false
}

// Delete 删除缓存
func (c *Cache) Delete(key string) {
	// TODO(human): 实现 Delete 方法
	// 提示：需要写锁
	//
	// c.mu.Lock()
	// defer c.mu.Unlock()
	// delete(c.items, key)

	// 你的代码：

}

// Count 返回缓存项数量
func (c *Cache) Count() int {
	// TODO(human): 实现 Count 方法
	// 提示：需要读锁
	//
	// c.mu.RLock()
	// defer c.mu.RUnlock()
	// return len(c.items)

	// 你的代码：
	return 0
}

// cleanupLoop 后台清理过期数据
func (c *Cache) cleanupLoop() {
	ticker := time.NewTicker(1 * time.Second) // 每秒检查一次
	defer ticker.Stop()

	for range ticker.C {
		c.cleanup()
	}
}

// cleanup 清理过期缓存
func (c *Cache) cleanup() {
	// TODO(human): 实现清理逻辑
	// 提示：
	// 1. 使用写锁（要删除数据）
	// 2. 遍历 map，删除过期的项
	//
	// c.mu.Lock()
	// defer c.mu.Unlock()
	//
	// for key, item := range c.items {
	//     if item.IsExpired() {
	//         delete(c.items, key)
	//     }
	// }

	// 你的代码：

}

// ============================================
// 模拟数据库查询
// ============================================

// 假设这是一个很慢的数据库查询
func queryDatabase(userID string) string {
	fmt.Printf("  🔍 查询数据库: userID=%s (慢...)\n", userID)
	time.Sleep(100 * time.Millisecond) // 模拟数据库延迟
	return fmt.Sprintf("用户%s的数据", userID)
}

// getUserWithCache 使用缓存获取用户数据
func getUserWithCache(cache *Cache, userID string) string {
	// 先查缓存
	if value, exists := cache.Get(userID); exists {
		fmt.Printf("  ✅ 缓存命中: userID=%s\n", userID)
		return value.(string)
	}

	// 缓存未命中，查数据库
	fmt.Printf("  ❌ 缓存未命中: userID=%s\n", userID)
	data := queryDatabase(userID)

	// 存入缓存（1 秒过期）
	cache.Set(userID, data, 1*time.Second)

	return data
}

func mainT() {
	fmt.Println("🎯 实战项目 4: 并发安全的缓存系统")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	cache := NewCache()

	// 场景 1: 基本使用
	fmt.Println("📍 场景 1: 基本使用")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// 设置缓存
	cache.Set("user:1", "Alice", 2*time.Second)
	cache.Set("user:2", "Bob", 2*time.Second)
	fmt.Println("✓ 设置了 2 个缓存项")

	// 读取缓存
	if value, exists := cache.Get("user:1"); exists {
		fmt.Printf("✓ 读取缓存: user:1 = %v\n", value)
	}

	// 统计
	fmt.Printf("✓ 当前缓存数量: %d\n", cache.Count())
	fmt.Println()

	time.Sleep(1 * time.Second)

	// 场景 2: 缓存加速查询
	fmt.Println("📍 场景 2: 缓存加速查询")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// 第一次查询（会查数据库）
	fmt.Println("第 1 次查询:")
	start := time.Now()
	getUserWithCache(cache, "user:100")
	fmt.Printf("  耗时: %v\n\n", time.Since(start))

	// 第二次查询（命中缓存）
	fmt.Println("第 2 次查询（马上）:")
	start = time.Now()
	getUserWithCache(cache, "user:100")
	fmt.Printf("  耗时: %v (快了 100 倍！)\n\n", time.Since(start))

	// 场景 3: 并发访问测试
	fmt.Println("📍 场景 3: 并发访问（测试线程安全）")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	var wg sync.WaitGroup

	// 10 个 goroutine 同时读写缓存
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			key := fmt.Sprintf("concurrent:%d", id%3) // 只有 3 个不同的 key

			// 写入
			cache.Set(key, fmt.Sprintf("数据%d", id), 2*time.Second)

			// 读取
			if value, exists := cache.Get(key); exists {
				fmt.Printf("  Goroutine %d: 读取 %s = %v\n", id, key, value)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("✓ 并发测试完成（没有 panic 说明线程安全）")
	fmt.Println()

	time.Sleep(500 * time.Millisecond)

	// 场景 4: 缓存过期
	fmt.Println("📍 场景 4: 缓存过期机制")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	cache.Set("expire:test", "这条数据会过期", 2*time.Second)
	fmt.Println("✓ 设置缓存 (2秒过期)")
	fmt.Printf("  当前缓存数量: %d\n", cache.Count())

	fmt.Println("\n等待 1 秒...")
	time.Sleep(1 * time.Second)

	if value, exists := cache.Get("expire:test"); exists {
		fmt.Printf("✅ 1 秒后: 数据还在 = %v\n", value)
	} else {
		fmt.Println("❌ 1 秒后: 数据已过期")
	}

	fmt.Println("\n再等待 2 秒（总共 3 秒）...")
	time.Sleep(2 * time.Second)

	if value, exists := cache.Get("expire:test"); exists {
		fmt.Printf("✅ 3 秒后: 数据还在 = %v\n", value)
	} else {
		fmt.Println("❌ 3 秒后: 数据已过期")
	}

	fmt.Printf("  当前缓存数量: %d (后台清理已生效)\n", cache.Count())
	fmt.Println()

	// 场景 5: 性能对比
	fmt.Println("📍 场景 5: 性能对比（缓存 vs 无缓存）")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// 无缓存：查询 10 次
	fmt.Println("无缓存: 查询 10 次")
	start = time.Now()
	for i := 0; i < 10; i++ {
		queryDatabase("user:999")
	}
	noCacheDuration := time.Since(start)
	fmt.Printf("  总耗时: %v\n\n", noCacheDuration)

	// 有缓存：查询 10 次（只有第一次查数据库）
	fmt.Println("有缓存: 查询 10 次")
	start = time.Now()
	for i := 0; i < 10; i++ {
		getUserWithCache(cache, "user:999")
	}
	cacheDuration := time.Since(start)
	fmt.Printf("  总耗时: %v\n", cacheDuration)
	fmt.Printf("  ⚡ 快了 %.1f 倍！\n", float64(noCacheDuration)/float64(cacheDuration))
	fmt.Println()

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("💡 总结:")
	fmt.Println("  • 缓存可以大幅提升性能（10-100 倍）")
	fmt.Println("  • RWMutex 允许多个读操作并发执行")
	fmt.Println("  • 定时过期保证数据不会太旧")
	fmt.Println("  • 后台清理防止内存泄漏")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// 💡 思考题:
	// 1. 为什么用 RWMutex 而不是 Mutex？
	// 2. 如果缓存数据很大，怎么限制内存使用？（提示：LRU）
	// 3. 如果多个 goroutine 同时请求同一个不存在的 key，会发生什么？
	//    （提示：缓存击穿，解决方案：singleflight）
	// 4. 在分布式系统中，如何保证多个服务器的缓存一致性？
}

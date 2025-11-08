# Goroutine + Channel 综合实战

## 🎯 学习目标

巩固 01 和 02 的内容，掌握 goroutine 和 channel 的组合使用模式。

## 📋 练习内容

### 练习 1: 生产者-消费者模式 🏭
**场景**: 一个生产者生成任务，多个消费者并发处理

**关键点**:
- 生产者完成后要 **关闭 channel**
- 消费者用 `range` 遍历 channel（自动处理关闭）
- 使用 `sync.WaitGroup` 等待所有消费者完成

**面试常问**:
- "生产者不关闭 channel 会怎样？" → 消费者的 range 永远不会退出（goroutine 泄漏）
- "多个消费者如何公平分配任务？" → Go 的 channel 自动负载均衡

---

### 练习 2: 工作池模式 (Worker Pool) 👷
**场景**: 固定数量的 worker 处理大量任务

**关键点**:
- 用缓冲 channel 存储任务
- 多个 worker 从同一个 channel 接收任务
- 需要两个 channel：jobs（输入）和 results（输出）
- 注意关闭顺序：先关 jobs，等 workers 完成后再关 results

**面试常问**:
- "为什么要限制 worker 数量？" → 控制并发度，避免资源耗尽
- "如何优雅关闭 worker pool？" → 关闭 jobs → 等待 WaitGroup → 关闭 results

---

### 练习 3: 并发求和 ➕
**场景**: 分治算法 - 多个 goroutine 分段计算，汇总结果

**关键点**:
- 每个 goroutine 负责一段数据
- 通过 channel 发送部分结果
- 主 goroutine 汇总所有结果

**面试常问**:
- "并发一定比串行快吗？" → 不一定，要考虑任务粒度和通信开销
- "如何确保收到所有结果？" → 知道 goroutine 数量 或 使用 WaitGroup + close

---

### 练习 4: 数据管道 (Pipeline) 🚰
**场景**: 数据经过多个阶段处理（类似 Unix 管道：`cat | grep | sort`）

**关键点**:
- 每个阶段是一个函数，输入和输出都是 channel
- 函数返回只读 channel (`<-chan`)
- 使用 `defer close(out)` 确保 channel 关闭
- 阶段之间自动流式传递数据

**面试常问**:
- "Pipeline 的优势？" → 解耦、并发、内存高效（流式处理）
- "如何处理错误？" → 可以增加 error channel 或使用 context

---

## 🔥 常见陷阱

### 1. 忘记关闭 channel
```go
// ❌ 错误：生产者不关闭 channel
go func() {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    // 忘记 close(ch)
}()

for v := range ch {  // ⚠️ 永远不会退出！
    fmt.Println(v)
}
```

### 2. 在错误的 goroutine 中关闭 channel
```go
// ❌ 错误：消费者关闭 channel
for v := range ch {
    fmt.Println(v)
    close(ch)  // ⚠️ 其他消费者会 panic
}
```
**原则**: 谁生产谁关闭（发送方关闭，接收方不关闭）

### 3. 关闭已关闭的 channel
```go
close(ch)
close(ch)  // ⚠️ panic: close of closed channel
```

### 4. 向已关闭的 channel 发送数据
```go
close(ch)
ch <- 1  // ⚠️ panic: send on closed channel
```

### 5. WaitGroup 计数错误
```go
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    go func() {
        wg.Add(1)  // ❌ 应该在启动 goroutine 前 Add
        defer wg.Done()
        // ...
    }()
}
wg.Wait()  // ⚠️ 可能在所有 Add 之前就执行完了
```
**正确**: `wg.Add(1)` 放在 `go` 之前

---

## 🚀 运行方式

### 运行练习（需要你填写 TODO）
```bash
go run main.go
```

### 运行参考答案
打开 `solution.go`，取消最后的 `main` 函数注释，然后：
```bash
go run solution.go
```

### 竞态检测
```bash
go run -race main.go
```

---

## 📊 完成进度

- [ ] 练习 1: 生产者-消费者模式
- [ ] 练习 2: 工作池模式
- [ ] 练习 3: 并发求和
- [ ] 练习 4: 数据管道

完成后，你就可以开始学习 **03_channel_select** 了！

---

## 💡 学习建议

1. **先自己写** - 不要直接看答案，先尝试实现
2. **对比答案** - 完成后对比参考答案，看有什么不同
3. **思考面试题** - 每个练习下面的"面试常问"要能回答出来
4. **避开陷阱** - 确保你理解每个陷阱，最好自己重现一次错误

---

## 🎓 面试高频问题

### Q1: 如何避免 goroutine 泄漏？
- 确保 channel 被正确关闭
- 使用 context 控制 goroutine 生命周期
- 用带缓冲的 channel 避免阻塞

### Q2: WaitGroup 和 channel 关闭的关系？
```go
// 正确模式
go func() {
    wg.Wait()      // 等待所有 worker 完成
    close(results)  // 再关闭结果 channel
}()
```

### Q3: 生产者-消费者 vs 工作池有什么区别？
- **生产者-消费者**: 关注数据流动，1:N 或 N:M
- **工作池**: 关注资源控制，固定数量的 worker 处理任务队列

### Q4: 什么时候用缓冲 channel，什么时候用无缓冲？
- **无缓冲**: 需要同步（发送方会阻塞等待接收方）
- **缓冲**: 解耦、允许突发、提高吞吐量

---

## ⭐ 核心概念总结

| 概念 | 作用 | 关键点 |
|------|------|--------|
| `go func()` | 启动并发 | 传参避免闭包陷阱 |
| `make(chan T)` | 创建 channel | 无缓冲 = 同步 |
| `make(chan T, n)` | 创建缓冲 channel | 缓冲满了才阻塞 |
| `close(ch)` | 关闭 channel | 发送方关闭 |
| `range ch` | 遍历 channel | 自动处理关闭 |
| `sync.WaitGroup` | 等待 goroutine | Add/Done/Wait |

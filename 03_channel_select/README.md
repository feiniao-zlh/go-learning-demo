# 03 - Channel Select 多路复用

## 🎯 学习目标

掌握 Go 中最强大的并发工具之一：`select` 语句。Select 让你能够同时等待多个 channel 操作，是构建复杂并发系统的基础。

## 📚 核心概念

### 1. Select 基本语法
```go
select {
case val := <-ch1:
    // ch1 有数据可读
case ch2 <- value:
    // ch2 可以写入
case <-time.After(time.Second):
    // 超时
default:
    // 所有 case 都阻塞时执行（非阻塞）
}
```

### 2. Select 执行规则
- ✅ **随机选择**: 多个 case 同时就绪时，随机选一个执行
- ✅ **阻塞等待**: 没有 case 就绪且无 default，会阻塞
- ✅ **非阻塞**: 有 default 则不会阻塞
- ✅ **nil channel**: nil channel 永远不会被选中

### 3. 常见使用场景
1. **超时控制**: 配合 `time.After()` 实现操作超时
2. **多数据源**: 同时监听多个数据源，先到先处理
3. **非阻塞操作**: 使用 default 避免阻塞
4. **优雅退出**: 配合 done channel 实现取消机制

## 🏃 开始练习

### 运行练习
```bash
# 进入目录
cd 03_channel_select

# 运行练习（会看到很多 TODO）
go run main.go

# 运行答案（查看正确实现）
go run solution.go
```

### 推荐学习流程
1. **先运行答案**: `go run solution.go` 看看预期效果
2. **阅读注释**: 仔细阅读 main.go 中的提示
3. **独立实现**: 填写所有 TODO，不要偷看答案
4. **验证运行**: 运行你的代码，看是否达到预期
5. **对比答案**: 对比 solution.go，理解差异

## 💡 重点提示

### ⚠️ 常见陷阱

#### 1. 关闭的 channel 会一直触发
```go
// ❌ 错误示例
for {
    select {
    case val := <-ch:
        fmt.Println(val) // ch 关闭后会无限循环打印零值
    }
}

// ✅ 正确做法
for {
    select {
    case val, ok := <-ch:
        if !ok {
            ch = nil // 设为 nil 后不会再被选中
            continue
        }
        fmt.Println(val)
    }
}
```

#### 2. time.After 会造成内存泄漏
```go
// ❌ 在循环中使用 time.After（每次都创建新 timer）
for {
    select {
    case <-ch:
        // ...
    case <-time.After(time.Second): // 泄漏！
    }
}

// ✅ 正确做法
timeout := time.After(time.Second)
for {
    select {
    case <-ch:
        // ...
    case <-timeout:
        return
    }
}
```

#### 3. Goroutine 泄漏
```go
// ❌ 超时后，发送者会永久阻塞
ch := make(chan int)
go func() {
    time.Sleep(2 * time.Second)
    ch <- 1 // 如果主 goroutine 已超时，这里会永久阻塞
}()

select {
case <-ch:
case <-time.After(1 * time.Second):
    return // goroutine 泄漏！
}

// ✅ 使用 buffered channel 或 done channel
ch := make(chan int, 1) // 缓冲 channel，发送不会阻塞
```

## 🎓 面试高频问题

### Q1: Select 和 switch 有什么区别？
**答案**:
- `switch` 是顺序判断条件，找到第一个匹配的执行
- `select` 是同时等待多个 channel 操作，随机选择就绪的执行
- `select` 的 case 必须是 channel 操作

### Q2: 多个 case 同时就绪，select 怎么选？
**答案**: 随机选择一个执行。这是为了避免饥饿问题。

### Q3: Select 中可以有多个 default 吗？
**答案**: 不可以，最多一个 default。

### Q4: 为什么要把关闭的 channel 设为 nil？
**答案**:
- 关闭的 channel 会立即返回零值和 false
- 如果不设为 nil，select 会不断选中它，造成 CPU 空转
- nil channel 永远阻塞，select 不会选中它

### Q5: time.After 返回什么类型？
**答案**: 返回 `<-chan Time`，一个只读的 time channel。在指定时间后，这个 channel 会接收到当前时间。

## 🚀 进阶挑战

完成基础练习后，尝试这些挑战：

1. **实现一个通用的超时函数**
   ```go
   func WithTimeout(f func(), timeout time.Duration) bool {
       // 执行 f，如果超时返回 false
   }
   ```

2. **实现 Merge 函数**
   ```go
   // 合并多个 channel 到一个
   func Merge(channels ...<-chan int) <-chan int {
       // 返回一个 channel，包含所有输入 channel 的数据
   }
   ```

3. **实现 First 函数**
   ```go
   // 返回第一个完成的结果
   func First(funcs ...func() string) string {
       // 并发执行所有函数，返回最快的结果
   }
   ```

## 📖 扩展阅读

- [Go Blog: Go Concurrency Patterns: Timing out, moving on](https://go.dev/blog/concurrency-timeouts)
- [Effective Go: Channels](https://go.dev/doc/effective_go#channels)
- Rob Pike: Go Concurrency Patterns (YouTube)

## ✅ 完成检查清单

- [ ] 理解 select 的基本语法
- [ ] 知道如何处理关闭的 channel
- [ ] 会使用 time.After 实现超时
- [ ] 了解 default 的非阻塞特性
- [ ] 能识别和避免 goroutine 泄漏
- [ ] 完成所有 5 个练习
- [ ] 独立实现至少 2 个进阶挑战

---

**下一步**: 完成本练习后，继续学习 `04_sync_waitgroup`，学习更多的同步工具！

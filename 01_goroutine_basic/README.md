# 练习 1: Goroutine 基础

## 🎯 学习目标
- 学会用 `go` 关键字创建 goroutine
- 理解主 goroutine 和子 goroutine 的关系
- 掌握闭包捕获变量的陷阱

## 📝 任务说明

打开 `main.go`，找到 `TODO` 标记的地方，填写你的代码。

### 练习 1: 创建第一个 goroutine
- 用 `go sayHello()` 启动一个 goroutine
- 用 `time.Sleep(1 * time.Second)` 让主函数等待

### 练习 2: 创建多个 goroutine
```go
for i := 1; i <= 5; i++ {
    go func(id int) {
        fmt.Printf("Goroutine %d 正在运行\n", id)
    }(i)  // 传递参数！
}
```

### 练习 3: 理解闭包陷阱
- 错误写法：直接在 goroutine 里用 `i`
- 正确写法：把 `i` 作为参数传进去

## 🚀 运行
```bash
cd 01_goroutine_basic
go run main.go
```

## 💡 关键知识点

1. **创建 goroutine**: `go function()`
2. **等待完成**: `time.Sleep()` (临时方案)
3. **传递参数**: 避免闭包陷阱

## ❓ 思考题

1. 如果不加 `time.Sleep()`，会发生什么？
2. 练习 3 的错误写法会输出什么？为什么？

---

**完成后告诉我，我会检查你的代码！**

# Go 学习示例仓库

这是一个系统学习 Go 语言并发编程和核心特性的实战仓库，通过**渐进式**的示例和真实项目，帮助你从零基础到能够独立开发并发应用。

## 🎯 仓库定位

- **目标受众**：有其他语言基础，想系统学习 Go 并发编程的开发者
- **学习方式**：理论 + 练习 + 真实项目，边学边练
- **代码风格**：详细注释 + TODO 引导，适合自学

## 📁 目录结构

```
go-learning-demo/
│
├── 01_goroutine_basic/          # Goroutine 基础
│   ├── main.go                  # 练习代码
│   └── value_vs_ref.go          # 值传递 vs 引用传递
│
├── 02_channel_basic/            # Channel 基础
│   └── main.go                  # Channel 创建、发送、接收、关闭
│
├── 03_channel_select/           # Select 多路复用
│   ├── main.go                  # Select 练习
│   └── solution.go              # 参考答案
│
├── 01_02_combined_practice/     # 综合练习
│   ├── main.go                  # Goroutine + Channel 综合应用
│   └── solution.go              # 参考答案
│
├── real_world_practices/        # 真实项目实战
│   ├── README.md                # 实战项目说明
│   │
│   ├── 01_concurrent_downloader/  # 项目 1：并发下载器
│   │   ├── main.go                # 练习代码（Worker Pool 模式）
│   │   └── solution.go            # 完整实现（含超时、取消）
│   │
│   ├── 02_rate_limiter/           # 项目 2：API 限流器
│   │   └── main.go                # 令牌桶算法实现
│   │
│   ├── 03_timeout_retry/          # 项目 3：超时重试机制
│   │   └── main.go                # 指数退避 + 超时控制
│   │
│   ├── 04_simple_cache/           # 项目 4：并发安全缓存
│   │   └── main.go                # RWMutex + 过期清理
│   │
│   └── web_crawler_exercise.go    # 经典爬虫练习（Go Tour）
│
├── LEARNING_PATH.md             # 完整学习路径（22 个主题）
└── CLAUDE.md                    # Claude Code 使用指南
```

## 🚀 快速开始

### 1. 克隆仓库
```bash
git clone <your-repo-url>
cd go-learning-demo
```

### 2. 选择学习路径

#### 路径 A：系统学习（推荐）
按照目录顺序学习：
```bash
# 第 1 天：Goroutine 基础
cd 01_goroutine_basic
go run main.go

# 第 2 天：Channel 基础
cd 02_channel_basic
go run main.go

# 第 3 天：Select 多路复用
cd 03_channel_select
go run main.go
```

#### 路径 B：直接实战
如果你已经了解基础，可以直接做项目：
```bash
cd real_world_practices

# 并发爬虫（最简单，50 行）
go run web_crawler_exercise.go

# 并发下载器（Worker Pool）
cd 01_concurrent_downloader
go run main.go
```

### 3. 查看完整学习路径
```bash
cat LEARNING_PATH.md
```

## 📚 学习内容

### 阶段 1：并发编程核心（已完成）
- ✅ **Goroutine 基础** - 创建、启动、生命周期
- ✅ **Channel 基础** - 无缓冲/有缓冲、关闭、range 遍历
- ✅ **Select 多路复用** - 超时控制、非阻塞操作

### 阶段 2：真实项目实战（进行中）
- 🔄 **并发下载器** - Worker Pool、任务队列、超时控制
- 🔄 **API 限流器** - 令牌桶算法、time.Ticker
- 🔄 **超时重试** - 指数退避、错误处理
- 🔄 **并发缓存** - sync.RWMutex、过期清理

### 阶段 3：高级特性（规划中）
- ⬜ Context 上下文
- ⬜ 接口与多态
- ⬜ 错误处理
- ⬜ 内存管理与性能

详见 [LEARNING_PATH.md](./LEARNING_PATH.md)

## 🎓 使用方法

### 练习代码
每个练习都有 `TODO(human)` 标记，表示需要你实现的部分：

```go
// TODO(human): 实现并发下载逻辑
// 提示：
// 1. 创建 jobs 和 results channel
// 2. 启动 workers
// 3. 发送任务
// 你的代码：
```

### 查看答案
大部分练习都有 `solution.go` 参考答案：

```bash
# 运行练习代码
go run main.go

# 查看答案
go run solution.go
```

### 运行测试
建议使用 race detector 检测并发问题：

```bash
go run -race main.go
```

## 💡 学习建议

### 每个主题的学习流程
1. **阅读代码** - 看注释，理解要实现什么
2. **独立实现** - 根据 TODO 提示完成代码
3. **运行测试** - 看效果是否正确
4. **对比答案** - 查看 solution.go，理解差异
5. **思考问题** - 回答代码中的"思考题"

### 学习节奏
- **快速突击（1-2 周）**：每天 2-3 个主题，重点做实战项目
- **稳扎稳打（1 个月）**：每天 1 个主题，完成所有练习
- **深度学习（2 个月）**：每个主题深入研究，阅读源码

### 推荐工具
- **VSCode** + Go 插件
- **Goland** (JetBrains)
- **go fmt** 格式化代码
- **go vet** 检查代码问题

## 🌟 项目亮点

### 1. 渐进式学习
从最基础的 goroutine 到复杂的 Worker Pool，循序渐进。

### 2. 真实项目
所有实战项目都是工作中会遇到的真实场景：
- 并发下载 → 爬虫、CDN
- 限流器 → API 保护、服务限流
- 超时重试 → 微服务通信
- 并发缓存 → Redis 的简化版

### 3. 详细注释
每行关键代码都有注释，理解为什么这样写。

### 4. 思考题
帮助你深入理解原理，而不是死记代码。

## 📖 推荐阅读

- [Go 官方文档](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go 并发模式](https://go.dev/blog/pipelines)

## 🤝 贡献

这是个人学习仓库，欢迎：
- 提 Issue 指出错误
- 提 PR 改进代码
- 分享学习心得

## 📝 学习记录

### 已完成
- [x] Goroutine 基础
- [x] Channel 基础
- [x] Select 多路复用（进行中）

### 进行中
- [ ] 并发爬虫练习
- [ ] 并发下载器
- [ ] API 限流器

### 下一步
查看 [LEARNING_PATH.md](./LEARNING_PATH.md) 了解完整路线图。

## 📬 联系方式

如果有问题或建议，欢迎：
- 提 Issue
- 发邮件
- 留言讨论

---

**开始学习吧！** 🚀

从 `01_goroutine_basic` 开始，或者直接挑战 `real_world_practices` 中的项目！

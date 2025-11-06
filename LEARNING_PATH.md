# Go 学习路径 - 从基础到面试

## 📚 学习路线（按优先级排序）

### 阶段 1️⃣：并发编程核心（⭐⭐⭐⭐⭐ 面试必考）
这是 Go 的核心特性，也是面试最常考的内容。

- [ ] **01_goroutine_basic** - Goroutine 基础
  - 创建和启动 goroutine
  - goroutine 的生命周期
  - main goroutine 退出问题

- [ ] **02_channel_basic** - Channel 基础
  - 无缓冲 channel
  - 有缓冲 channel
  - channel 的关闭
  - range 遍历 channel

- [ ] **03_channel_select** - Select 多路复用
  - select 语句使用
  - 超时控制
  - default 非阻塞操作

- [ ] **04_sync_waitgroup** - WaitGroup 同步
  - 等待多个 goroutine 完成
  - Add/Done/Wait 使用
  - 常见错误

- [ ] **05_sync_mutex** - Mutex 互斥锁
  - 互斥锁基础
  - 读写锁 RWMutex
  - 竞态条件检测

- [ ] **06_channel_patterns** - Channel 常见模式
  - Worker Pool（工作池）
  - Pipeline（管道）
  - Fan-out/Fan-in
  - 生产者消费者

### 阶段 2️⃣：接口与多态（⭐⭐⭐⭐⭐ 面试高频）

- [ ] **07_interface_basic** - 接口基础
  - 接口定义和实现
  - 空接口 interface{}
  - 类型断言和类型判断

- [ ] **08_interface_advanced** - 接口进阶
  - 接口组合
  - 接口值的内部结构
  - 接口与 nil

### 阶段 3️⃣：错误处理与 Context（⭐⭐⭐⭐ 实战必备）

- [ ] **09_error_handling** - 错误处理
  - error 接口
  - 自定义错误
  - errors.Is 和 errors.As
  - panic 和 recover

- [ ] **10_context** - Context 上下文
  - context 的作用
  - WithCancel/WithTimeout/WithDeadline
  - WithValue 传递元数据
  - context 最佳实践

### 阶段 4️⃣：内存管理与性能（⭐⭐⭐⭐ 面试加分）

- [ ] **11_pointer_value** - 指针与值
  - 指针接收者 vs 值接收者
  - 什么时候用指针
  - 逃逸分析

- [ ] **12_slice_internals** - Slice 内部原理
  - slice 底层结构
  - append 扩容机制
  - slice 陷阱

- [ ] **13_map_internals** - Map 内部原理
  - map 底层实现
  - map 并发安全问题
  - sync.Map 使用

- [ ] **14_memory_gc** - 内存与 GC
  - Go 内存分配
  - GC 工作原理
  - 内存优化技巧

### 阶段 5️⃣：高级特性（⭐⭐⭐ 进阶提升）

- [ ] **15_reflect** - 反射
  - reflect 基础
  - 结构体标签解析
  - 反射的性能影响

- [ ] **16_defer** - Defer 机制
  - defer 执行顺序
  - defer 与闭包
  - defer 性能

- [ ] **17_generic** - 泛型（Go 1.18+）
  - 泛型函数
  - 泛型类型
  - 类型约束

### 阶段 6️⃣：网络编程（⭐⭐⭐⭐ 后端必备）

- [ ] **18_http_server** - HTTP 服务器
  - net/http 基础
  - 路由和处理器
  - 中间件模式

- [ ] **19_http_client** - HTTP 客户端
  - 发送请求
  - 超时控制
  - 连接池管理

- [ ] **20_tcp_socket** - TCP 编程
  - TCP 服务器/客户端
  - 粘包处理

### 阶段 7️⃣：测试与工具（⭐⭐⭐ 工程实践）

- [ ] **21_testing** - 单元测试
  - testing 包使用
  - 表驱动测试
  - Mock 和 stub

- [ ] **22_benchmark** - 性能测试
  - 基准测试
  - pprof 性能分析

## 🎯 面试高频考点速查

### 必考题（100% 会问）
1. **Goroutine 和线程的区别？** → `01_goroutine_basic`
2. **Channel 的底层实现？** → `02_channel_basic`
3. **如何避免 goroutine 泄漏？** → `10_context`
4. **Slice 和 Array 的区别？** → `12_slice_internals`
5. **Map 是否并发安全？** → `13_map_internals`

### 高频题（80% 会问）
1. **Mutex 和 RWMutex 的使用场景？** → `05_sync_mutex`
2. **Context 的作用？** → `10_context`
3. **defer 的执行顺序？** → `16_defer`
4. **接口的动态类型和动态值？** → `07_interface_basic`
5. **指针接收者和值接收者的区别？** → `11_pointer_value`

### 进阶题（50% 会问）
1. **GC 的工作原理？** → `14_memory_gc`
2. **调度器 GMP 模型？** → `01_goroutine_basic`
3. **内存逃逸分析？** → `11_pointer_value`
4. **sync.Pool 的使用？** → `14_memory_gc`

## 📝 学习建议

### 每日学习计划
- **快速突击（1-2周）**：每天 2-3 个主题，重点学习阶段 1、2、3
- **稳扎稳打（1个月）**：每天 1 个主题，完成所有阶段
- **深度学习（2个月+）**：每个主题深入研究，做项目实战

### 学习方法
1. **看代码** → 运行示例代码理解概念
2. **写代码** → 自己实现类似功能
3. **改代码** → 修改参数观察不同结果
4. **讲代码** → 能向别人解释清楚

### 练习建议
每个主题学完后：
1. 不看代码，自己写一遍
2. 思考：这个特性解决什么问题？
3. 思考：面试官可能怎么问？
4. 总结：用自己的话概括核心点

## 🚀 开始学习

从 `01_goroutine_basic` 开始，按顺序学习！

告诉我你准备好了，我会为你创建第一个学习示例！

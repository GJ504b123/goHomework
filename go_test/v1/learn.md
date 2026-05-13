```PlainText
v2/
├── controller/          # 控制器层（命令行版本）
│   └── task_controller.go  # 任务控制器（命令行交互）
├── dao/                 # 数据访问层
│   └── task_dao.go      # 任务数据访问对象（与数据库交互）
├── frontend/            # 前端界面
│   └── index.html       # 前端页面（HTML + JavaScript）
├── model/               # 数据模型层
│   └── task.go          # 任务数据结构定义
├── service/             # 业务逻辑层
│   └── task_service.go  # 任务业务逻辑
├── go.mod               # Go 模块文件（依赖管理）
├── go.sum               # Go 依赖校验文件
├── main.go              # 主程序（HTTP 服务器）
└── 实验报告.md          # 实验报告

```
## model/  定义数据结构
```go
// Task 任务数据结构
type Task struct {
    ID          int       `json:"id"`          // 任务编号
    Name        string    `json:"name"`        // 任务名称
    Description string    `json:"description"` // 任务描述
    Status      string    `json:"status"`      // 任务状态
    Priority    string    `json:"priority"`    // 优先级
    CreatedAt   time.Time `json:"created_at"`  // 创建时间
}
```
## dao/ 数据访问层
与数据库交互，执行sql

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"test/controller"
	"test/dao"
	"test/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 这里改成你的 MySQL 密码！
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_forum?charset=utf8mb4"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("连接失败：", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("数据库不通：", err)
	}

	// 创建任务表
	createTaskTable(db)

	// 初始化各个层
	taskDAO := dao.NewTaskDAO(db)
	taskService := service.NewTaskService(taskDAO)
	taskController := controller.NewTaskController(taskService)

	// 检查是否需要测试模式
	if len(os.Args) > 1 && os.Args[1] == "--test" {
		runTestMode(taskService)
		return
	}

	// 命令行交互
	for {
		printMenu()
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			taskController.CreateTask()
		case "2":
			taskController.GetTaskByID()
		case "3":
			taskController.GetAllTasks()
		case "4":
			taskController.GetTasksByStatus()
		case "5":
			taskController.GetTasksByPriority()
		case "6":
			taskController.UpdateTask()
		case "7":
			taskController.DeleteTask()
		case "8":
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("无效的选择，请重新输入")
		}
	}
}

// runTestMode 运行测试模式
func runTestMode(taskService *service.TaskService) {
	fmt.Println("===== 测试模式 =====")
	
	// 测试创建任务
	fmt.Println("1. 测试创建任务")
	task, err := taskService.CreateTask("测试任务", "这是一个测试任务", "待处理", "高")
	if err != nil {
		fmt.Printf("创建任务失败: %v\n", err)
		return
	}
	fmt.Printf("任务创建成功，ID: %d\n", task.ID)
	
	// 测试查询所有任务
	fmt.Println("\n2. 测试查询所有任务")
	tasks, err := taskService.GetAllTasks()
	if err != nil {
		fmt.Printf("获取任务列表失败: %v\n", err)
		return
	}
	for _, t := range tasks {
		fmt.Printf("ID: %d, 名称: %s, 状态: %s, 优先级: %s\n", t.ID, t.Name, t.Status, t.Priority)
	}
	
	// 测试根据ID查询任务
	fmt.Printf("\n3. 测试根据ID查询任务 (ID: %d)\n", task.ID)
	taskByID, err := taskService.GetTaskByID(task.ID)
	if err != nil {
		fmt.Printf("获取任务失败: %v\n", err)
		return
	}
	fmt.Printf("ID: %d, 名称: %s, 状态: %s, 优先级: %s\n", taskByID.ID, taskByID.Name, taskByID.Status, taskByID.Priority)
	
	// 测试更新任务
	fmt.Println("\n4. 测试更新任务")
	updatedTask, err := taskService.UpdateTask(task.ID, "更新后的测试任务", "这是一个更新后的测试任务", "进行中", "中")
	if err != nil {
		fmt.Printf("更新任务失败: %v\n", err)
		return
	}
	fmt.Printf("任务更新成功，ID: %d, 名称: %s, 状态: %s, 优先级: %s\n", updatedTask.ID, updatedTask.Name, updatedTask.Status, updatedTask.Priority)
	
	// 测试根据状态查询任务
	fmt.Println("\n5. 测试根据状态查询任务 (状态: 进行中)")
	tasksByStatus, err := taskService.GetTasksByStatus("进行中")
	if err != nil {
		fmt.Printf("获取任务列表失败: %v\n", err)
		return
	}
	for _, t := range tasksByStatus {
		fmt.Printf("ID: %d, 名称: %s, 状态: %s, 优先级: %s\n", t.ID, t.Name, t.Status, t.Priority)
	}
	
	// 测试根据优先级查询任务
	fmt.Println("\n6. 测试根据优先级查询任务 (优先级: 中)")
	tasksByPriority, err := taskService.GetTasksByPriority("中")
	if err != nil {
		fmt.Printf("获取任务列表失败: %v\n", err)
		return
	}
	for _, t := range tasksByPriority {
		fmt.Printf("ID: %d, 名称: %s, 状态: %s, 优先级: %s\n", t.ID, t.Name, t.Status, t.Priority)
	}
	
	// 测试删除任务
	fmt.Printf("\n7. 测试删除任务 (ID: %d)\n", task.ID)
	err = taskService.DeleteTask(task.ID)
	if err != nil {
		fmt.Printf("删除任务失败: %v\n", err)
		return
	}
	fmt.Println("任务删除成功")
	
	// 测试删除后查询所有任务
	fmt.Println("\n8. 测试删除后查询所有任务")
	tasksAfterDelete, err := taskService.GetAllTasks()
	if err != nil {
		fmt.Printf("获取任务列表失败: %v\n", err)
		return
	}
	if len(tasksAfterDelete) == 0 {
		fmt.Println("没有任务")
	} else {
		for _, t := range tasksAfterDelete {
			fmt.Printf("ID: %d, 名称: %s, 状态: %s, 优先级: %s\n", t.ID, t.Name, t.Status, t.Priority)
		}
	}
	
	fmt.Println("\n===== 测试完成 =====")
}

// createTaskTable 创建任务表
func createTaskTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		status VARCHAR(50) NOT NULL,
		priority VARCHAR(50) NOT NULL,
		created_at DATETIME NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("创建任务表失败：", err)
	}

	fmt.Println("✅ 任务表创建成功！")
}

// printMenu 打印菜单
func printMenu() {
	fmt.Println("\n===== 任务管理系统 =====")
	fmt.Println("1. 创建任务")
	fmt.Println("2. 根据ID查询任务")
	fmt.Println("3. 查询所有任务")
	fmt.Println("4. 根据状态查询任务")
	fmt.Println("5. 根据优先级查询任务")
	fmt.Println("6. 更新任务")
	fmt.Println("7. 删除任务")
	fmt.Println("8. 退出系统")
	fmt.Print("请输入选择：")
}

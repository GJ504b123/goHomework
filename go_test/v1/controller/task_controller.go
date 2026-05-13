package controller

import (
	"fmt"
	"strconv"
	"test/model"
	"test/service"
)

// TaskController 任务控制器
type TaskController struct {
	TaskService *service.TaskService
}

// NewTaskController 创建任务控制器实例
func NewTaskController(taskService *service.TaskService) *TaskController {
	return &TaskController{TaskService: taskService}
}

// CreateTask 创建任务
func (c *TaskController) CreateTask() {
	var name, description, status, priority string

	fmt.Print("请输入任务名称: ")
	fmt.Scanln(&name)
	
	fmt.Print("请输入任务描述: ")
	fmt.Scanln(&description)
	
	fmt.Print("请输入任务状态: ")
	fmt.Scanln(&status)
	
	fmt.Print("请输入任务优先级: ")
	fmt.Scanln(&priority)

	task, err := c.TaskService.CreateTask(name, description, status, priority)
	if err != nil {
		fmt.Printf("创建任务失败: %v\n", err)
		return
	}

	fmt.Printf("任务创建成功，ID: %d\n", task.ID)
}

// GetTaskByID 根据ID获取任务
func (c *TaskController) GetTaskByID() {
	var idStr string
	fmt.Print("请输入任务ID: ")
	fmt.Scanln(&idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("无效的ID")
		return
	}

	task, err := c.TaskService.GetTaskByID(id)
	if err != nil {
		fmt.Printf("获取任务失败: %v\n", err)
		return
	}

	c.printTask(task)
}

// GetAllTasks 获取所有任务
func (c *TaskController) GetAllTasks() {
	tasks, err := c.TaskService.GetAllTasks()
	if err != nil {
		fmt.Printf("获取任务列表失败: %v\n", err)
		return
	}

	c.printTasks(tasks)
}

// GetTasksByStatus 根据状态获取任务
func (c *TaskController) GetTasksByStatus() {
	var status string
	fmt.Print("请输入任务状态: ")
	fmt.Scanln(&status)

	tasks, err := c.TaskService.GetTasksByStatus(status)
	if err != nil {
		fmt.Printf("获取任务列表失败: %v\n", err)
		return
	}

	c.printTasks(tasks)
}

// GetTasksByPriority 根据优先级获取任务
func (c *TaskController) GetTasksByPriority() {
	var priority string
	fmt.Print("请输入任务优先级: ")
	fmt.Scanln(&priority)

	tasks, err := c.TaskService.GetTasksByPriority(priority)
	if err != nil {
		fmt.Printf("获取任务列表失败: %v\n", err)
		return
	}

	c.printTasks(tasks)
}

// UpdateTask 更新任务
func (c *TaskController) UpdateTask() {
	var idStr, name, description, status, priority string

	fmt.Print("请输入任务ID: ")
	fmt.Scanln(&idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("无效的ID")
		return
	}

	fmt.Print("请输入新的任务名称: ")
	fmt.Scanln(&name)
	
	fmt.Print("请输入新的任务描述: ")
	fmt.Scanln(&description)
	
	fmt.Print("请输入新的任务状态: ")
	fmt.Scanln(&status)
	
	fmt.Print("请输入新的任务优先级: ")
	fmt.Scanln(&priority)

	task, err := c.TaskService.UpdateTask(id, name, description, status, priority)
	if err != nil {
		fmt.Printf("更新任务失败: %v\n", err)
		return
	}

	fmt.Println("任务更新成功")
	c.printTask(task)
}

// DeleteTask 删除任务
func (c *TaskController) DeleteTask() {
	var idStr string
	fmt.Print("请输入任务ID: ")
	fmt.Scanln(&idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("无效的ID")
		return
	}

	err = c.TaskService.DeleteTask(id)
	if err != nil {
		fmt.Printf("删除任务失败: %v\n", err)
		return
	}

	fmt.Println("任务删除成功")
}

// printTask 打印单个任务信息
func (c *TaskController) printTask(task *model.Task) {
	fmt.Printf("ID: %d\n", task.ID)
	fmt.Printf("名称: %s\n", task.Name)
	fmt.Printf("描述: %s\n", task.Description)
	fmt.Printf("状态: %s\n", task.Status)
	fmt.Printf("优先级: %s\n", task.Priority)
	fmt.Printf("创建时间: %s\n", task.CreatedAt)
}

// printTasks 打印任务列表
func (c *TaskController) printTasks(tasks []*model.Task) {
	if len(tasks) == 0 {
		fmt.Println("没有任务")
		return
	}

	for _, task := range tasks {
		c.printTask(task)
		fmt.Println("--------------------")
	}
}

package service

import (
	"test/dao"
	"test/model"
	"time"
)

// TaskService 任务服务
type TaskService struct {
	TaskDAO *dao.TaskDAO
}

// NewTaskService 创建任务服务实例
func NewTaskService(taskDAO *dao.TaskDAO) *TaskService {
	return &TaskService{TaskDAO: taskDAO}
}

// CreateTask 创建任务
func (s *TaskService) CreateTask(name, description, status, priority string) (*model.Task, error) {
	task := &model.Task{
		Name:        name,
		Description: description,
		Status:      status,
		Priority:    priority,
		CreatedAt:   time.Now(),
	}

	err := s.TaskDAO.Create(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// GetTaskByID 根据ID获取任务
func (s *TaskService) GetTaskByID(id int) (*model.Task, error) {
	return s.TaskDAO.GetByID(id)
}

// GetAllTasks 获取所有任务
func (s *TaskService) GetAllTasks() ([]*model.Task, error) {
	return s.TaskDAO.GetAll()
}

// GetTasksByStatus 根据状态获取任务
func (s *TaskService) GetTasksByStatus(status string) ([]*model.Task, error) {
	return s.TaskDAO.GetByStatus(status)
}

// GetTasksByPriority 根据优先级获取任务
func (s *TaskService) GetTasksByPriority(priority string) ([]*model.Task, error) {
	return s.TaskDAO.GetByPriority(priority)
}

// UpdateTask 更新任务
func (s *TaskService) UpdateTask(id int, name, description, status, priority string) (*model.Task, error) {
	task, err := s.TaskDAO.GetByID(id)
	if err != nil {
		return nil, err
	}

	task.Name = name
	task.Description = description
	task.Status = status
	task.Priority = priority

	err = s.TaskDAO.Update(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// DeleteTask 删除任务
func (s *TaskService) DeleteTask(id int) error {
	return s.TaskDAO.Delete(id)
}

package dao

import (
	"database/sql"
	"test/model"
	"time"
)

// TaskDAO 任务数据访问对象
type TaskDAO struct {
	DB *sql.DB
}

// NewTaskDAO 创建任务DAO实例
func NewTaskDAO(db *sql.DB) *TaskDAO {
	return &TaskDAO{DB: db}
}

// Create 创建任务
func (dao *TaskDAO) Create(task *model.Task) error {
	query := `INSERT INTO tasks (name, description, status, priority, created_at) VALUES (?, ?, ?, ?, ?)`
	result, err := dao.DB.Exec(query, task.Name, task.Description, task.Status, task.Priority, task.CreatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = int(id)
	return nil
}

// GetByID 根据ID获取任务
func (dao *TaskDAO) GetByID(id int) (*model.Task, error) {
	query := `SELECT id, name, description, status, priority, created_at FROM tasks WHERE id = ?`
	row := dao.DB.QueryRow(query, id)

	var task model.Task
	var createdAt []byte
	err := row.Scan(&task.ID, &task.Name, &task.Description, &task.Status, &task.Priority, &createdAt)
	if err != nil {
		return nil, err
	}

	// 解析时间字符串
	task.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// GetAll 获取所有任务
func (dao *TaskDAO) GetAll() ([]*model.Task, error) {
	query := `SELECT id, name, description, status, priority, created_at FROM tasks`
	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		var task model.Task
		var createdAt []byte
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Status, &task.Priority, &createdAt)
		if err != nil {
			return nil, err
		}

		// 解析时间字符串
		task.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

// GetByStatus 根据状态获取任务
func (dao *TaskDAO) GetByStatus(status string) ([]*model.Task, error) {
	query := `SELECT id, name, description, status, priority, created_at FROM tasks WHERE status = ?`
	rows, err := dao.DB.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		var task model.Task
		var createdAt []byte
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Status, &task.Priority, &createdAt)
		if err != nil {
			return nil, err
		}

		// 解析时间字符串
		task.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

// GetByPriority 根据优先级获取任务
func (dao *TaskDAO) GetByPriority(priority string) ([]*model.Task, error) {
	query := `SELECT id, name, description, status, priority, created_at FROM tasks WHERE priority = ?`
	rows, err := dao.DB.Query(query, priority)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		var task model.Task
		var createdAt []byte
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Status, &task.Priority, &createdAt)
		if err != nil {
			return nil, err
		}

		// 解析时间字符串
		task.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

// Update 更新任务
func (dao *TaskDAO) Update(task *model.Task) error {
	query := `UPDATE tasks SET name = ?, description = ?, status = ?, priority = ? WHERE id = ?`
	_, err := dao.DB.Exec(query, task.Name, task.Description, task.Status, task.Priority, task.ID)
	return err
}

// Delete 删除任务
func (dao *TaskDAO) Delete(id int) error {
	query := `DELETE FROM tasks WHERE id = ?`
	_, err := dao.DB.Exec(query, id)
	return err
}

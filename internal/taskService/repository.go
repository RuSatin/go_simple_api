package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task
	err := r.db.First(&existingTask, id).Error
	if err != nil {
		return Task{}, err
	}

	if task.Task != "" {
		existingTask.Task = task.Task
	}
	existingTask.IsDone = task.IsDone

	err = r.db.Save(&existingTask).Error
	return existingTask, err
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	return r.db.Delete(&Task{}, id).Error
}

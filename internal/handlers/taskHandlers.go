package handlers

import (
	"context"
	"simple_api/internal/taskService"
	"simple_api/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получаем все задачи из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err // Возвращаем ошибку, если что-то пошло не так
	}

	// Создаем слайс для ответа
	response := tasks.GetTasks200JSONResponse{}

	// Преобразуем задачи из базы данных в формат, ожидаемый клиентом
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// Возвращаем ответ
	return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Получаем данные из тела запроса
	taskRequest := request.Body

	// Создаем задачу для добавления в базу данных
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task, // Используем поле Task, а не Text
		IsDone: *taskRequest.IsDone,
	}

	// Вызываем метод сервиса для создания задачи
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err // Возвращаем ошибку, если что-то пошло не так
	}

	// Создаем ответ
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	// Возвращаем ответ
	return response, nil
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Получаем ID задачи из запроса
	id := request.Id

	// Получаем данные для обновления из тела запроса
	updateData := request.Body

	// Создаем структуру задачи для обновления
	taskToUpdate := taskService.Task{
		Task:   *updateData.Task,
		IsDone: *updateData.IsDone,
	}

	// Обновляем задачу в базе данных
	updatedTask, err := h.Service.UpdateTaskByID(uint(id), taskToUpdate)
	if err != nil {
		return nil, err
	}

	// Создаем ответ
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	err := h.Service.DeleteTaskByID(uint(id))
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}

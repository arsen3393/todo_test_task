package handler

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
	"testSkillsRock/pkg/models"
	"time"
)

var validStatuses = []string{"new", "done", "in_progress"}

func IsValidStatus(status string) bool {
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}

// ****************************************************************************************
// **************************************HANDLERS BLOCK************************************
// ****************************************************************************************

// Получить все задачи
func GetTasks(c fiber.Ctx) error {
	taskModel, ok := c.Locals("TaskModel").(*models.TaskModel)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot cast to models.TaskModel",
		})
	}

	tasks, err := taskModel.Get()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get tasks" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"tasks": tasks,
	})
}

type RequestCreateTask struct {
	Title string `json:"title"`
	Desc  string `json:"description"`
}

func validateTitle(title string) bool {
	if len(title) < 1 {
		return false
	} else {
		return true
	}
}

// Создать Задачу
func CreateTask(c fiber.Ctx) error {
	taskModel, ok := c.Locals("TaskModel").(*models.TaskModel)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot cast to models.TaskModel",
		})
	}

	var request RequestCreateTask
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !validateTitle(request.Title) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title is not valid",
		})
	}

	task := models.Task{
		Title:       request.Title,
		Description: request.Desc,
		Status:      "new",
	}

	err := taskModel.Create(task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task added successfully",
	})
}

func DeleteTask(c fiber.Ctx) error {
	taskModel, ok := c.Locals("TaskModel").(*models.TaskModel)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	TaskId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot convert task id to int",
		})
	}

	err = taskModel.Delete(TaskId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}

type RequestUpdateTask struct {
	Title  string `json:"title"`
	Desc   string `json:"description"`
	Status string `json:"status"`
}

func UpdateTask(c fiber.Ctx) error {
	taskModel, ok := c.Locals("TaskModel").(*models.TaskModel)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	TaskId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot convert task id to int",
		})
	}

	var request RequestUpdateTask

	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	task := models.Task{
		ID:          TaskId,
		Title:       request.Title,
		Description: request.Desc,
		Status:      request.Status,
		UpdatedAt:   time.Now(),
	}
	err = taskModel.Update(task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task updated successfully",
	})
}

package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type TaskModel struct {
	Task    Task
	Connect *DBconnection
}

type Task struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (m *TaskModel) Init(db *DBconnection) {
	m.Connect = db
}

func NewTaskModel(db *DBconnection) *TaskModel {
	m := &TaskModel{}
	m.Init(db)
	return m
}

func (m *TaskModel) Create(t Task) error {
	if !checkStatus(t.Status) {
		return errors.New("invalid status")
	}

	query := `INSERT INTO tasks (title, description, status, created_at, updated_at) 
              VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`

	ctx := context.Background()

	err := m.Connect.db.QueryRow(ctx, query, t.Title, t.Description, t.Status).Scan(&t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *TaskModel) Get() ([]Task, error) {
	query := `SELECT * FROM tasks`
	ctx := context.Background()
	rows, err := m.Connect.db.Query(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return tasks, nil
}

func (m *TaskModel) Delete(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`

	ctx := context.Background()

	if !m.validateID(id) {
		return errors.New("invalid id")
	}

	_, err := m.Connect.db.Exec(ctx, query, id)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (m *TaskModel) Update(t Task) error {
	if !checkStatus(t.Status) {
		return errors.New(fmt.Sprintf("invalid status: %s, Use: %s, %s, %s ", t.Status, "new", "in_progress, done"))
	}

	if !m.validateID(t.ID) {

		return errors.New("invalid id")
	}

	query := `UPDATE tasks
				SET title = $1, description = $2, status = $3, updated_at = NOW() WHERE id = $4`

	ctx := context.Background()

	_, err := m.Connect.db.Exec(ctx, query, t.Title, t.Description, t.Status, t.ID)

	if err != nil {
		return err
	}

	return nil
}

func (m *TaskModel) validateID(id int) bool {
	query := `SELECT id FROM tasks WHERE id = $1`
	ctx := context.Background()

	var taskID int
	err := m.Connect.db.QueryRow(ctx, query, id).Scan(&taskID)

	if err == nil {
		return true
	}

	// тут может быть случай когда queryRow выдаст другую ошибку, из за соединения и т.д. Но чето стало лень это обрабатывать ))
	return false
}

func checkStatus(status string) bool {
	switch status {
	case "new":
		return true
	case "in_progress":
		return true
	case "done":
		return true
	default:
		return false
	}
}

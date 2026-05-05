package repository

import (
	"database/sql"
	"mymodule/model"
)

type TaskRepository interface {
	Create(task model.Task) error
	GetAll() ([]model.Task, error)
	Update(id int) error
	Delete(id int) error
}

// ✅ struct (THIS WAS MISSING / WRONG)
type taskRepo struct {
	db *sql.DB
}

// ✅ constructor (THIS FIXES NewTaskRepo error)
func NewTaskRepo(db *sql.DB) TaskRepository {
	return &taskRepo{db: db}
}

// ✅ CREATE
func (r *taskRepo) Create(task model.Task) error {
	_, err := r.db.Exec(
		"INSERT INTO tasks (name) VALUES ($1)",
		task.Name,
	)
	return err
}

// ✅ GET ALL
func (r *taskRepo) GetAll() ([]model.Task, error) {
	rows, err := r.db.Query("SELECT id, name, is_done FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var t model.Task
		rows.Scan(&t.ID, &t.Name, &t.IsDone)
		tasks = append(tasks, t)
	}

	return tasks, nil
}

// ✅ UPDATE
func (r *taskRepo) Update(id int) error {
	_, err := r.db.Exec(
		"UPDATE tasks SET is_done = TRUE WHERE id=$1",
		id,
	)
	return err
}

// ✅ DELETE
func (r *taskRepo) Delete(id int) error {
	_, err := r.db.Exec(
		"DELETE FROM tasks WHERE id=$1",
		id,
	)
	return err
}

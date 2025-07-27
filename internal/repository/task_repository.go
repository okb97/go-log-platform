package repository

import (
	"log"

	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/model"
)

func GetAllTasks() ([]model.Task, error) {
	rows, err := db.DB.Query("select * from tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			log.Println("Scanエラー:", err)
			continue
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

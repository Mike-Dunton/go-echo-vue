// Package models db interface
package models

import (
	"database/sql"
)

// Task is a struct containing Task data
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	UserID      string `json:"userID`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks returns a list a TaskCollection
func GetTasks(db *sql.DB, userID string) TaskCollection {
	sql := "SELECT id, description FROM tasks where user_id=?"
	rows, err := db.Query(sql, userID)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Description)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

func PutTask(db *sql.DB, description string, userID string) (int64, error) {
	sql := "INSERT INTO tasks(description, user_id) VALUES(?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, sqlExecError := stmt.Exec(description, userID)
	if sqlExecError != nil {
		return 0, sqlExecError
	}

	return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int, userID string) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ? AND user_id=? "
	stmt, prepareError := db.Prepare(sql)
	if prepareError != nil {
		return 0, prepareError
	}
	defer stmt.Close()
	result, sqlExecError := stmt.Exec(id, userID)
	if sqlExecError != nil {
		return 0, sqlExecError
	}
	return result.RowsAffected()
}

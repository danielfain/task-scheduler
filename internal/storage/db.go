package storage

import (
	"database/sql"
	"task-scheduler/internal/scheduler"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./tasks.db?_journal_mode=WAL&_foreign_keys=on")
	if err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func LoadTasks(db *sql.DB, pq *scheduler.TaskQueue) error {
	rows, err := db.Query("SELECT id, expression, command, next_run FROM tasks WHERE is_active = TRUE")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var task scheduler.Task
		if err := rows.Scan(&task.Id, &task.Expression, &task.Command, &task.NextRun); err != nil {
			return err
		}
		pq.Push(&task)
	}
	return rows.Err()
}

func createTables(db *sql.DB) error {
	schema := `
        PRAGMA journal_mode = WAL; -- Enable Write-Ahead Logging for better concurrency

        CREATE TABLE IF NOT EXISTS tasks (
            id          TEXT PRIMARY KEY,
            expression  TEXT NOT NULL,
            command     TEXT NOT NULL,
            next_run    DATETIME NOT NULL,
            is_active   BOOLEAN DEFAULT TRUE,
            metadata    TEXT
        );

        CREATE INDEX IF NOT EXISTS idx_tasks_next_run ON tasks(next_run);

        CREATE TABLE IF NOT EXISTS task_logs (
            id          INTEGER PRIMARY KEY AUTOINCREMENT,
            task_id     TEXT NOT NULL,
            exit_code   INTEGER,
            output      TEXT,
            executed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
        );
	`

	_, err := db.Exec(schema)
	return err
}

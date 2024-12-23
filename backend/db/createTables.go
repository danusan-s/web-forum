package db

import (
	"database/sql"
	"fmt"
)

func CreateTables(db *sql.DB) error {
	createThreadsTableSQL := `
	CREATE TABLE IF NOT EXISTS THREADS (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    author VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

	createUsersTableSQL := `
    CREATE TABLE IF NOT EXISTS USERS (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

	createCommentsTableSQL := `
  CREATE TABLE IF NOT EXISTS COMMENTS (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    thread_id INT NOT NULL,
    content TEXT NOT NULL,
    author VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (thread_id) REFERENCES THREADS(id) ON DELETE CASCADE
);`

	createThreadReactionsTableSQL := `
  CREATE TABLE IF NOT EXISTS THREAD_REACTIONS (
    user_id INT NOT NULL,
    thread_id INT NOT NULL,
    state TINYINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, thread_id),
    FOREIGN KEY (thread_id) REFERENCES THREADS(id) ON DELETE CASCADE
);`

	createCommentReactionsTableSQL := `
  CREATE TABLE IF NOT EXISTS COMMENT_REACTIONS (
    user_id INT NOT NULL,
    comment_id BIGINT UNSIGNED NOT NULL,
    state TINYINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, comment_id),
    FOREIGN KEY (comment_id) REFERENCES COMMENTS(id) ON DELETE CASCADE
);`

	_, err := db.Exec(createThreadsTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create Threads table: %v", err)
	}
	_, err = db.Exec(createUsersTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}
	_, err = db.Exec(createCommentsTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create comments table: %v", err)
	}
	_, err = db.Exec(createThreadReactionsTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create Thread_reactions table: %v", err)
	}
	_, err = db.Exec(createCommentReactionsTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create comment_reactions table: %v", err)
	}

	return nil
}

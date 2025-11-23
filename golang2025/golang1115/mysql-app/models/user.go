package models

import (
	"database/sql"
	"time"
)

// User 用户模型
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAllUsers 获取所有用户
func GetAllUsers(db *sql.DB) ([]User, error) {
	query := "SELECT id, name, email, created_at, updated_at FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// CreateUser 创建新用户
func CreateUser(db *sql.DB, name, email string) (int64, error) {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := db.Exec(query, name, email)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

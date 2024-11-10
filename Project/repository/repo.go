package repository

import (
	"database/sql"
	"fmt"
	"main/database"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository() Repository {
	db := database.DB
	return Repository{DB: db}
}

func (r *Repository) TokenCheckRepo(token string) string {
	if token == "" {
		return "Token required"
	}

	var expiresAt time.Time
	query := "SELECT expires_at FROM tokens WHERE token = $1"
	err := r.DB.QueryRow(query, token).Scan(&expiresAt)
	if err != nil {
		return "Invalid or expired token"
	}

	// Memeriksa apakah token sudah kadaluarsa
	if time.Now().After(expiresAt) {
		return "Token has expired"
	}

	return ""
}

// Fungsi untuk membersihkan token yang sudah kadaluarsa
func (r *Repository) CleanExpiredTokensRepo() string {
	query := "DELETE FROM tokens WHERE expires_at < $1"
	_, err := database.DB.Exec(query, time.Now())
	if err != nil {
		return fmt.Sprintf("Failed to clean expired tokens: %v", err)
	}

	return ""
}

// func (r *Repository) GetUsersRepo(users *[]customers.Customer) {
// 	rows, err := database.DB.Query(`SELECT id, name, active FROM users ORDER BY name`)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var user customers.Customer
// 		if err := rows.Scan(&user.ID, &user.Name); err != nil {
// 			panic(err)
// 		}

// 		*users = append(*users, user)
// 	}
// }

// func (r *Repository) GetTodosRepo(tasks *[]todosModel.Task) {
// 	rows, err := database.DB.Query("SELECT task_id, title, status FROM tasks ORDER BY task_id")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var task todosModel.Task
// 		if err := rows.Scan(&task.Task_id, &task.Title, &task.Status); err != nil {
// 			panic(err)
// 		}

// 		*tasks = append(*tasks, task)
// 	}
// }

// func (r *Repository) GetUserDetailRepo(users *customers.Customer) {
// 	err := database.DB.QueryRow(`SELECT u.id, u.name, u.username, u.password, u.active, t.token FROM users u JOIN tokens t ON t.user_id = u.id WHERE u.id = $1`, users.ID).Scan(
// 		&users.ID, &users.Name, &users.Username, &users.Password, &users.Active, &users.Token,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func (r *Repository) DeleteUserRepo(id int) error {
// 	query := "DELETE FROM users WHERE id = $1"
// 	_, err := database.DB.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *Repository) DeleteTodoRepo(id int) error {
// 	query := "DELETE FROM tasks WHERE task_id = $1"
// 	_, err := database.DB.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

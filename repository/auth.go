package repository

import (
	users "main/model/user"
	"time"

	"github.com/google/uuid"
)

func (r *Repository) LoginRepo(user *users.User) error {
	query := `SELECT u.id, u.username, u.password, u.email, t.token 
	FROM users u
	JOIN tokens t ON t.user_id = u.id
	WHERE u.username = $1 AND u.password = $2`
	err := r.DB.QueryRow(query, user.Username, user.Password).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Token)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RegisterRepo(user users.User) error {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`

	var id int

	err := r.DB.QueryRow(query, user.Username, user.Password, user.Email).Scan(&id)
	if err != nil {
		return err
	}

	user.Token, err = r.GenerateTkn(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GenerateTkn(userID int) (string, error) {
	token := uuid.New().String()
	expiresAt := time.Now().Add(5 * time.Minute)

	query := "INSERT INTO tokens (user_id, token, expires_at) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, userID, token, expiresAt)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *Repository) GetRoleRepo(token string) (string, error) {
	var role string
	query := `SELECT u.role FROM tokens t JOIN costumers u ON u.user_id = t.user_id WHERE t.token = $1;`
	err := r.DB.QueryRow(query, token).Scan(&role)
	if err != nil {
		return "", err
	}
	return role, nil
}

func (r *Repository) GetCustomerByIDRepo(id int) (string, error) {
	var name string
	query := `SELECT u.name
	FROM customers u
	JOIN orders o ON o.customer_id = u.id 
	WHERE o.customer_id = $1
	LIMIT 1
`
	err := r.DB.QueryRow(query, id).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

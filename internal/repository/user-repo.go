package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jurest07/user-service/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) CreateUser(username, email string) (*models.User, error) {
	query := `INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id, username, email, created_at`
	var user models.User
	err := ur.db.QueryRow(query, username, email).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания пользователя: %w", err)
	}

	return &user, nil
}

func (ur *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, username, email, created_at FROM users WHERE email = $1`

	var user models.User
	err := ur.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("пользователя с таким email не найдено: %w", err)
		}
		return nil, fmt.Errorf("ошибка получения пользователя: %w", err)
	}

	return &user, nil
}

func (ur *UserRepo) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, username, email, created_at FROM users WHERE id = $1`

	var user models.User
	err := ur.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("пользователя с таким id не найдено: %w", err)
		}
		return nil, fmt.Errorf("ошибка получения пользователя: %w", err)
	}

	return &user, nil
}

func (ur *UserRepo) UserExists(id int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`
	var exists bool
	err := ur.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка проверки существования: %w", err)
	}
	return exists, nil
}

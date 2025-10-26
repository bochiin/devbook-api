package repository

import (
	"api/src/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (repository UserRepository) Create(user models.User) (uint, error) {

	statment, err := repository.db.Prepare("INSERT INTO users (name, nickname, email, password) values($1, $2, $3, $4) RETURNING id")

	if err != nil {
		return 0, err
	}

	defer statment.Close()

	var id uint

	err = statment.QueryRow(user.Name, user.Nickname, user.Email, user.Password).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

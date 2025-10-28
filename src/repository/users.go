package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repository UserRepository) GetUsers(searchParam string) ([]models.User, error) {
	queryLike := fmt.Sprintf("%%%s%%", searchParam) // %param%

	results, err := repository.db.Query(
		"SELECT id, name, nickname, email, createdin FROM users WHERE LOWER(name) LIKE $1 OR LOWER(nickname) LIKE $2",
		queryLike,
		queryLike,
	)

	if err != nil {
		return nil, err
	}

	defer results.Close()

	var users []models.User

	for results.Next() {
		var user models.User

		if err = results.Scan(
			&user.Id,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedIn,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository UserRepository) GetUser(id uint64) (models.User, error) {

	results, err := repository.db.Query(
		"SELECT id, name, nickname, email, createdin FROM users WHERE id = $1",
		id,
	)

	if err != nil {
		return models.User{}, err
	}

	defer results.Close()

	var user models.User

	if results.Next() {
		if err = results.Scan(
			&user.Id,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedIn,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository UserRepository) UpdateUser(id uint64, user models.User) error {

	statment, err := repository.db.Prepare("UPDATE users SET name = $1, nickname = $2, email = $3 WHERE id = $4")

	if err != nil {
		return err
	}

	defer statment.Close()

	if _, err = statment.Exec(user.Name, user.Nickname, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (repository UserRepository) DeleteUser(id uint64) error {
	statment, err := repository.db.Prepare("DELETE FROM users WHERE id = $1")

	if err != nil {
		return err
	}

	defer statment.Close()

	if _, err = statment.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repository UserRepository) FindByEmail(email string) (models.User, error) {

	result, err := repository.db.Query("SELECT id, password FROM users where email = $1", email)

	if err != nil {
		return models.User{}, err
	}

	defer result.Close()

	var user models.User

	if result.Next() {
		if err = result.Scan(&user.Id, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

package db

import (
	"database/sql"
	"github.com/lukkas-lukkas/golang-todo-list-api/internal/domain/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(u user.User) (user.User, error) {
	stmt, err := r.db.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return user.User{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Name, u.Email, u.Password)
	if err != nil {
		return user.User{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return user.User{}, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return user.User{}, err
	}

	u.ID = int(lastInsertID)

	return u, nil
}

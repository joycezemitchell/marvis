package auth

import (
	"aichat/internal/app/dto"
	"aichat/internal/app/models"
	"aichat/pkg/sqlm"
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
	GetUser(req *dto.LoginRequest) (models.User, error)
}

type repository struct {
	db   *sql.DB
	sqlm sqlm.SQLM
}

func NewRepository(db *sql.DB, sqlm *sqlm.SQLM) Repository {
	return &repository{
		db:   db,
		sqlm: *sqlm,
	}
}

func (r *repository) GetUser(req *dto.LoginRequest) (models.User, error) {
	users := []models.User{}
	user := models.User{}

	rows, _ := r.sqlm.Get(context.Background(), "users", models.User{
		Username: req.Username,
		Password: req.Password,
	})

	jsonData, err := json.Marshal(rows)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		return user, err
	}

	// Return empty user if not found
	if len(users) == 0 {
		return user, nil
	}

	// Return user
	return users[0], nil
}

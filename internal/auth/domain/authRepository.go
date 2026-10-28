package domain

import (
	"database/sql"
	"log"

	"github.com/codacube/go-microservice-banking/internal/errs"
	"github.com/codacube/go-microservice-banking/internal/logger"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindBy(username string, password string) (*Login, *errs.AppError)
	GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, *errs.AppError)
	RefreshTokenExists(refreshToken string) *errs.AppError
}

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func (db AuthRepositoryDb) FindBy(username, password string) (*Login, *errs.AppError) {
	var login Login

	sqlVerify := `SELECT username, u.customer_id, role, group_concat(a.account_id) as account_numbers FROM users u
					LEFT JOIN accounts a ON a.customer_id = u.customer_id
					WHERE username = ? and password = ?
					GROUP BY u.customer_id;`

	err := db.client.Get(&login, sqlVerify, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewAuthenticationError("invalid credentials")
		} else {
			log.Println("Error while verifying login request from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &login, nil
}

func (db AuthRepositoryDb) GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, *errs.AppError) {
	// Generate the refresh token
	var appErr *errs.AppError
	var refreshToken string

	if refreshToken, appErr = authToken.newRefreshToken(); appErr != nil {
		return "", appErr
	}

	// Store it in the store
	sqlInsert := "INSERT INTO refresh_token_store (refresh_token) VALUES (?)"
	_, err := db.client.Exec(sqlInsert, refreshToken)
	if err != nil {
		logger.Error("Unexpected database error: " + err.Error())
		return "", errs.NewUnexpectedError("unexpected database error")
	}

	return refreshToken, nil
}

func (db AuthRepositoryDb) RefreshTokenExists(refreshToken string) *errs.AppError {
	sqlSelect := "SELECT refresh_token FROM refresh_token_store WHERE refresh_token = ?"
	var token string

	err := db.client.Get(&token, sqlSelect, refreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewAuthenticationError("refresh token not registered in the store")
		} else {
			logger.Error("Unexpected database error: " + err.Error())
			return errs.NewUnexpectedError("unexpected database error")
		}
	}

	return nil
}

func NewAuthRepository(client *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client}
}

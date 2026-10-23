package domain

import (
	"strconv"

	"github.com/codacube/go-microservice-banking/errs"
	"github.com/codacube/go-microservice-banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (db AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?);"

	result, err := db.client.Exec(sqlInsert, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while last insert id for the new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountID = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

func (db AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// Atomic transactions, all failures are rolled back and all successes are committed.
	tx, err := db.client.Begin()
	if err != nil {
		logger.Error("Error whilst starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError(("Unexpected database error"))
	}

	// Insert transaction
	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"
	result, _ := tx.Exec(sqlInsert, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)

	// Adjust the account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? WHERE account_id = ?", t.Amount, t.AccountID)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? WHERE account_id = ?", t.Amount, t.AccountID)
	}

	// If an error then Rollback. Changes to both tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transacton: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Everything is good, commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while committing transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	transactionID, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error whilst last insert id for the new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	account, appError := db.FindBy(t.AccountID)
	if appError != nil {
		return nil, appError
	}

	// Update the transaction struct with the new details
	t.TransactionID = strconv.FormatInt(transactionID, 10)
	t.Amount = account.Amount

	return &t, nil
}

func (db AccountRepositoryDb) FindBy(accountID string) (*Account, *errs.AppError) {
	var account Account

	findSQL := "SELECT account_id, customer_id, opening_date, account_type, amount FROM accounts WHERE account_id = ?"

	if err := db.client.Select(&account, findSQL, accountID); err != nil {
		logger.Error("Error whilst fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &account, nil
}

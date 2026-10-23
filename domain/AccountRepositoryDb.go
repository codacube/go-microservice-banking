package domain

import (
	"database/sql"
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
	tx, err := db.client.Begin()
	if err != nil {
		logger.Error("Error whilst starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError(("Unexpected database error"))
	}
	// A deferred function runs when SaveTransaction returns, including on early
	// returns below. Rolling back undoes any partial database changes and releases
	// the account lock. After a successful Commit, Rollback returns sql.ErrTxDone,
	// which is expected because the transaction has already finished.
	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			logger.Error("Error while rolling back bank account transaction: " + err.Error())
		}
	}()

	var balance float64
	// Lock this account row until the transaction finishes. If another withdrawal
	// is running at the same time, it must wait here and then check the updated
	// balance, rather than also approving itself against the old balance.
	err = tx.QueryRow("SELECT amount FROM accounts WHERE account_id = ? FOR UPDATE", t.AccountID).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Account not found while saving transaction")
			return nil, errs.NewNotFoundError("Account not found")
		}
		logger.Error("Error while checking account balance: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	if t.IsWithdrawal() && balance < t.Amount {
		return nil, errs.NewValidationError("Insufficient balance in the account")
	}

	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"
	result, err := tx.Exec(sqlInsert, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error while inserting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	transactionID, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error whilst getting the ID for the new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	if t.IsWithdrawal() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? WHERE account_id = ?", t.Amount, t.AccountID)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? WHERE account_id = ?", t.Amount, t.AccountID)
	}

	if err != nil {
		logger.Error("Error while updating account balance: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.QueryRow("SELECT amount FROM accounts WHERE account_id = ?", t.AccountID).Scan(&balance)
	if err != nil {
		logger.Error("Error while reading updated account balance: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("Error while committing transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	t.TransactionID = strconv.FormatInt(transactionID, 10)
	t.Amount = balance

	return &t, nil
}

func (db AccountRepositoryDb) FindBy(accountID string) (*Account, *errs.AppError) {
	var account Account

	findSQL := "SELECT account_id, customer_id, opening_date, account_type, amount FROM accounts WHERE account_id = ?"

	if err := db.client.Get(&account, findSQL, accountID); err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Account not found")
			return nil, errs.NewNotFoundError("Account not found")
		}
		logger.Error("Error whilst fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &account, nil
}

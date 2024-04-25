package libdb

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Transaction struct {
	tx        *sqlx.Tx
	options   *sql.TxOptions
	activated bool
}

type transactionContextKey struct{}

func transactionFromContext(ctx context.Context) *Transaction {
	tx, ok := ctx.Value(transactionContextKey{}).(*Transaction)
	if !ok {
		return nil
	}

	return tx
}

func deleteTransactionFromContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, transactionContextKey{}, nil)
}

func Begin(ctx context.Context, options *sql.TxOptions) context.Context {
	return context.WithValue(ctx, transactionContextKey{}, &Transaction{
		tx:        nil,
		options:   options,
		activated: false,
	})
}

func Rollback(ctx context.Context) error {
	tx := transactionFromContext(ctx)
	ctx = deleteTransactionFromContext(ctx)

	if tx != nil && tx.IsActivated() {
		tx.activated = false
		return tx.Rollback()
	}

	return nil
}

func Commit(ctx context.Context) error {
	tx := transactionFromContext(ctx)
	ctx = deleteTransactionFromContext(ctx)

	if tx != nil && tx.IsActivated() {
		tx.activated = false
		return tx.Commit()
	}

	return nil
}

func (tx *Transaction) CreateTransaction(ctx context.Context, db *sqlx.DB) error {
	var err error
	tx.tx, err = db.BeginTxx(ctx, tx.options)
	if err != nil {
		return err
	}

	tx.activated = true

	return nil
}

func (tx *Transaction) IsActivated() bool {
	return tx.activated
}

func (tx *Transaction) Rollback() error {
	return tx.tx.Rollback()
}

func (tx *Transaction) Commit() error {
	return tx.tx.Commit()
}

func (tx *Transaction) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return tx.tx.QueryContext(ctx, query, args...)
}

func (tx *Transaction) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return tx.tx.QueryxContext(ctx, query, args...)
}

func (tx *Transaction) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return tx.tx.QueryRowxContext(ctx, query, args...)
}

func (tx *Transaction) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.tx.ExecContext(ctx, query, args...)
}

package libdb

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Connection interface {
	sqlx.QueryerContext
	sqlx.ExecerContext
}

type DB interface {
	Connection(ctx context.Context) (Connection, error)
	Get(ctx context.Context, query squirrel.SelectBuilder, dst interface{}) error
	Select(ctx context.Context, query squirrel.SelectBuilder, dst interface{}) error
	Insert(ctx context.Context, query squirrel.InsertBuilder) error
	Update(ctx context.Context, query squirrel.UpdateBuilder) error
	Delete(ctx context.Context, query squirrel.DeleteBuilder) error
}

type sqlxDB struct {
	db *sqlx.DB
}

func NewSQLXDB(db *sqlx.DB) DB {
	return &sqlxDB{db}
}

func (db *sqlxDB) Connection(ctx context.Context) (Connection, error) {
	tx := transactionFromContext(ctx)
	if tx != nil {
		if !tx.IsActivated() {
			err := tx.CreateTransaction(ctx, db.db)

			if err != nil {
				return nil, err
			}
		}

		return tx, nil
	}

	return db.db, nil
}

func (db *sqlxDB) Get(ctx context.Context, query squirrel.SelectBuilder, dst interface{}) error {
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	conn, err := db.Connection(ctx)
	if err != nil {
		return err
	}

	err = sqlx.GetContext(ctx, conn, dst, sqlQuery, args...)
	if err == sql.ErrNoRows {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (db *sqlxDB) Select(ctx context.Context, query squirrel.SelectBuilder, dst interface{}) error {
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	conn, err := db.Connection(ctx)
	if err != nil {
		return err
	}

	err = sqlx.SelectContext(ctx, conn, dst, sqlQuery, args...)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (db *sqlxDB) Insert(ctx context.Context, query squirrel.InsertBuilder) error {
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	conn, err := db.Connection(ctx)
	if err != nil {
		return err
	}

	_, err = conn.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func (db *sqlxDB) Update(ctx context.Context, query squirrel.UpdateBuilder) error {
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	conn, err := db.Connection(ctx)
	if err != nil {
		return err
	}

	_, err = conn.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func (db *sqlxDB) Delete(ctx context.Context, query squirrel.DeleteBuilder) error {
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	conn, err := db.Connection(ctx)
	if err != nil {
		return err
	}

	_, err = conn.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

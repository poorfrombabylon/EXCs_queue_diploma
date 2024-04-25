package captcha

import (
	"context"
	"excs_queue/internal/libdb"
	"github.com/Masterminds/squirrel"
)

const captchaTable = "captcha"

type CaptchaStorage interface {
	CreateCaptcha(ctx context.Context, count int) error
}

type captchaStorage struct {
	db libdb.DB
}

func NewCaptchaStorage(db libdb.DB) CaptchaStorage {
	return &captchaStorage{
		db: db,
	}
}

func (c *captchaStorage) CreateCaptcha(ctx context.Context, count int) error {
	query := squirrel.Insert(captchaTable).
		Columns("meet").
		Values(count).
		PlaceholderFormat(squirrel.Dollar)

	err := c.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

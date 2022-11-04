package cron

import (
	"context"
	"database/sql"
	"time"
)

type cron struct {
	db *sql.DB
}

func NewCron(db *sql.DB, d time.Duration) *cron {
	return &cron{db: db}
}

func (c *cron) ClearExpiredLink() {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

}

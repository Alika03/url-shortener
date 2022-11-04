package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"shortener/config"
	"sync"
	"time"
)

var (
	dbOnce sync.Once
	db     *sql.DB
)

func GetDbConnection() *sql.DB {
	dbOnce.Do(func() {
		db = getWithTicker()
		db.SetMaxOpenConns(20)
		db.SetMaxIdleConns(20)
		db.SetConnMaxIdleTime(5 * time.Minute)
	})

	return db
}

func getWithTicker() *sql.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := pingDb(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func pingDb(ctx context.Context) (*sql.DB, error) {
	c := time.Tick(2 * time.Second)
	for {
		select {
		case <-c:
			db, err := connectToDb()
			if err != nil {
				log.Println(err)
				continue
			}
			if err := db.Ping(); err != nil {
				log.Println(err)
				continue
			}
			return db, nil
		case <-ctx.Done():
			return nil, errors.New("error: can not get db connection")
		}
	}
}

func connectToDb() (*sql.DB, error) {
	params := config.GetConfig().DbParams
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		params.Username, params.Password, params.Host, params.Port, params.Db,
	)

	return sql.Open("pgx", connString)
}

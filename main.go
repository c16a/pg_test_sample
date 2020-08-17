package main

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"os"
	"strconv"
	"time"
)

var db *pg.DB

func main() {
	Init()
	err := Ping()
	if err != nil {
		panic(err)
	}
}

func Ping() error {
	return db.RunInTransaction(func(tx *pg.Tx) error {
		_, err := tx.Exec("SELECT 1")
		return err
	})
}

func Init() {
	pgString := fmt.Sprintf(
		"postgres://%s:%s@%s",
		os.Getenv("pg_user"), os.Getenv("pg_password"), os.Getenv("pg_url"))
	dbOptions, connError := pg.ParseURL(pgString)

	if connError != nil {
		panic(connError)
	}

	SetDbOptions(dbOptions)
	db = pg.Connect(dbOptions)
}

func SetDbOptions(dbOptions *pg.Options) {
	dbOptions.PoolSize = 150

	if poolTimeout, err := strconv.Atoi(os.Getenv("DB_POOL_TIMEOUT")); err == nil {
		dbOptions.PoolTimeout = time.Duration(poolTimeout) * time.Second
	}

	if dialTimeout, err := strconv.Atoi(os.Getenv("DB_DIAL_TIMEOUT")); err == nil {
		dbOptions.DialTimeout = time.Duration(dialTimeout) * time.Second
	}

	if readTimeout, err := strconv.Atoi(os.Getenv("DB_READ_TIMEOUT")); err == nil {
		dbOptions.ReadTimeout = time.Duration(readTimeout) * time.Second
	}

	if writeTimeout, err := strconv.Atoi(os.Getenv("DB_WRITE_TIMEOUT")); err == nil {
		dbOptions.WriteTimeout = time.Duration(writeTimeout) * time.Second
	}

	if idleTimeout, err := strconv.Atoi(os.Getenv("DB_IDLE_TIMEOUT")); err == nil {
		dbOptions.IdleTimeout = time.Duration(idleTimeout) * time.Second
	}
}

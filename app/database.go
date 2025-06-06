package app

import (
	"database/sql"
	"fmt"
	"log/slog"
	"nesanest-rest-api/helper"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	host := helper.AppConfig.DB.GetHost()
	portStr := helper.AppConfig.DB.GetPort()
	user := helper.AppConfig.DB.GetUsername()
	password := helper.AppConfig.DB.GetPassword()
	dbname := helper.AppConfig.DB.GetDBName()

	port, err := strconv.Atoi(portStr)
	helper.PanicIfError(err)

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	slog.Info("Successfully connect to database")

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

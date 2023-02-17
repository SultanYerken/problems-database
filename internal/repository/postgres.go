package repository

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewPostrgresDB(cfg Config) (*sql.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)
	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println("repo postgres")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("ping:", err)
		return nil, err
	}

	return db, nil
}

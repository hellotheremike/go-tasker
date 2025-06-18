package lib

import (
	"database/sql"
	"log"

	"github.com/hellotheremike/go-tasker/internal/config"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Load().DATABASE_URL)
	if err != nil {
		log.Fatal("Failed to connect:", err)
		
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping:", err)
		db.Close()
		return nil, err
	}

	return db, nil
}


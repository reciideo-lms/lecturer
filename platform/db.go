package platform

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func InitDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	return db
}

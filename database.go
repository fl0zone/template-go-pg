package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB(databaseURL string, createTable bool) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Error connecting to the database: %q", err)
		return &sql.DB{}, err
	}

	_, err = db.Exec("select 1;")
	if err != nil {
		log.Fatalf("Error connecting to the database: %q", err)
		return &sql.DB{}, err
	}

	if createTable {
		createTableQuery := `create table if not exists todo (
			id serial primary key,
			title varchar(100) not null,
			completed boolean default false
		)`
		_, err = db.Exec(createTableQuery)
		if err != nil {
			log.Fatalf("Error creating todo table: %q", err)
			return &sql.DB{}, err
		}
	}

	return db, err
}

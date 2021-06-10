package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/blrobin2/learn-sqlc/postgres"
	_ "github.com/lib/pq"
)

func main() {
	conn, err := sql.Open("postgres", "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db := postgres.New(conn)

	user, err := db.CreateUser(context.Background(), postgres.CreateUserParams{
		Firstname: "John",
		Lastname:  "Doe",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
}

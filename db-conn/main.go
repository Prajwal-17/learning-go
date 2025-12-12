package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// database/sql - to interact with the database
// github.com/lib/pq - importing drivers from pgsql

// --- Docker setup ---
// create a network for conneting to pg admin
// docker create pg_network
// docker run -d -p 5432:5432 testvol:/var/lib/postgresql -e POSTGRES_DB=testdb -e POSTGRES_PASSWORD=pass --network pg_network postgres:latest
// docker run -d --name pgadmin -p 5050:80 -e PGADMIN_DEFAULT_EMAIL=admin@example.com -e PGADMIN_DEFAULT_PASSWORD=secret --network pg_network dpage/pgadmin4

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "testdb"
	password = "golang"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", "postgresql://postgres:golang@localhost:5432/testdb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	createTableSql := `
			CREATE TABLE users (
				id SERIAL PRIMARY KEY,
				name VARCHAR(20) UNIQUE NOT NULL,
				email VARCHAR(40) UNIQUE NOT NULL
			)`

	_, err = db.Exec(createTableSql)

	if err != nil {
		log.Fatalf("FATAL: Failed to execute CREATE TABLE command: %v", err)

	}

	fmt.Println("Successfully connected")
}

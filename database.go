package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

const (
	host     = "192.168.0.7"
	port     = 5432
	user     = "myuser"
	password = "mypass"
	dbname   = "mydb"
)

func connect() *pgx.Conn {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	conn, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn

}

func insertLine(nome string) {

	db := connect()
	defer db.Close(context.Background())
	db.Exec(context.Background(), `insert into usuarios (nome) values ($1)`, nome)

}

func main() {

	insertLine("Marco")

	// res, err := db.Exec(`CREATE TABLE usuarios(
	// 	id serial PRIMARY KEY,
	// 	nome VARCHAR (50))`)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

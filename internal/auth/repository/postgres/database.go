package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// Connect connects to postgresDB
func connect(url string) *pgx.Conn {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	//url := "postgres://postgres:sql@localhost:5432/Travel"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

// Database ...
type Database struct {
	conn *pgx.Conn
}

// NewDatabase ...
func NewDatabase(url string) *Database {
	return &Database{
		conn: connect(url),
	}
}

// Close ...
func (db Database) Close() error {
	return db.conn.Close(context.Background())
}

var errNoRows = pgx.ErrNoRows

// defer conn.Close(context.Background())

// var name string
// var pass string
// err = conn.QueryRow(context.Background(), "select username, pass from users where username=$1", "test").Scan(&name, &pass)
// if err != nil {
// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 	os.Exit(1)
// }

// fmt.Println(name, pass)

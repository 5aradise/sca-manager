package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const driverName = "postgres"

func New(host, user, password, port, name string) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

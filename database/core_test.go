package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"testing"
	"os"
	"fmt"

)

func TestConnectDB(t *testing.T) {
	dbPool, err := pgxpool.Connect(context.Background(),
		fmt.Sprintf("postgres://%v:%v@%v:%v/%v", 
		os.Getenv("DB_NAME"), 
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME")))

	if err != nil {
		t.Fatal("Error while connecting to postgrs", err)
	}

	var temp string
	err = dbPool.QueryRow(context.Background(), "select 'Hello World!'").Scan(&temp)
	if err != nil {
		t.Fatal("Error while copy to variable value", err)
	}

	if temp != "Hello World!" {
		t.Fatal("Incorrect result,got:", temp)
	}
}

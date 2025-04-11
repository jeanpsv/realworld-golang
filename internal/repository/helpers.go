package repository

import (
	"context"
	"database/sql"
	"time"
)

func OpenDB(driverName string, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

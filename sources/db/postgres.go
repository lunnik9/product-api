package db

import (
	"context"

	"github.com/go-pg/pg/v10"
)

// Connect connects to postgres instance
func Connect(url string) (*pg.DB, error) {
	//db := pg.Connect(&pg.Options{
	//	Addr:     address,
	//	User:     user,
	//	Password: password,
	//	Database: database,
	//
	//})
	opts, err := pg.ParseURL(url)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(opts)
	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}

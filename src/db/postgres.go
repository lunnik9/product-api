package db

import (
	"context"

	"github.com/go-pg/pg/v10"
)

// Connect connects to postgres instance
func Connect(address, user, password, database string) (*pg.DB, error) {
	//db := pg.Connect(&pg.Options{
	//	Addr:     address,
	//	User:     user,
	//	Password: password,
	//	Database: database,
	//
	//})
	opts, err := pg.ParseURL("postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95")
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

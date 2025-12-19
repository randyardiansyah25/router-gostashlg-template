package databasefactory

import (
	"errors"
	"os"
)

var AppDb Database

func GetDatabase() (db Database, err error) {
	driverName := os.Getenv("app.database_driver")
	if driverName == "" {
		driverName = DRIVER_MOCK
	}

	if driverName == DRIVER_MYSQL {
		return newMysqlImpl(), nil
	} else if driverName == DRIVER_MOCK {
		return newMockImpl(), nil
	} else {
		return db, errors.New("unimplement database driver")
	}
}

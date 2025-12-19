package employeerepo

import (
	"errors"
	"os"
	"router-gostashlg-template/repository/built_in/databasefactory"
)

func NewEmployeeRepo() (EmployeeRepo, error) {
	driverName := os.Getenv("app.database_driver")
	if driverName == databasefactory.DRIVER_MYSQL {
		return newEmployeeMysqlImpl(), nil
	} else if driverName == databasefactory.DRIVER_MOCK {
		return newEmployeeMockImpl(), nil
	} else {
		return nil, errors.New("unimplemented database driver")
	}

}

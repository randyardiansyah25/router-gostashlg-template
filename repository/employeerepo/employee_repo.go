package employeerepo

import "router-gostashlg-template/entities"

type EmployeeRepo interface {
	GetEmployee() ([]entities.Employee, error)
	GetEmployeeById(id int64) (entities.Employee, error)
}

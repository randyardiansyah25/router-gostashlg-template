package usecase

import (
	"router-gostashlg-template/entities"
	"router-gostashlg-template/entities/app"
	"router-gostashlg-template/entities/brokermessage"
	"router-gostashlg-template/repository/brokerrepo"
	"router-gostashlg-template/repository/employeerepo"
)

type EmployeeUsecase interface {
	GetEmployeeList() ([]entities.Employee, error)
	GetEmployee(id int64) (entities.Employee, error)
}

func NewEmployeeUsecase() EmployeeUsecase {
	return &employeeUsecase{}
}

type employeeUsecase struct{}

func (e *employeeUsecase) GetEmployeeList() (detail []entities.Employee, er error) {
	repo, _ := employeerepo.NewEmployeeRepo()
	detail, er = repo.GetEmployee()

	if er != nil {
		return detail, er
	}

	// ! DO ANOTHER BUSINESS LOGIC HERE

	if len(detail) == 0 {
		return detail, app.ErrNoRecord
	}

	// ! PUBLISH EVENT SAMPLE
	brokerrepo.PublishMessage(brokermessage.Metadata[any]{
		Action:       "sample_action",
		Description:  "Ini sample data saat publish ke rabbit",
		Identity:     app.Identifier,
		RegisteredId: "10010100",
		Channel:      "MOBILE",
		Data:         "Custom data",
	})
	return
}

func (e *employeeUsecase) GetEmployee(id int64) (employee entities.Employee, er error) {
	repo, _ := employeerepo.NewEmployeeRepo()
	employee, er = repo.GetEmployeeById(id)

	// ! DO ANOTHER BUSINESS LOGIC HERE

	return
}

package employeerepo

import (
	"router-gostashlg-template/entities"

	"github.com/go-faker/faker/v4"
)

func newEmployeeMockImpl() EmployeeRepo {
	return &mockImpl{}
}

type mockImpl struct{}

func (m mockImpl) GetEmployee() (list []entities.Employee, er error) {
	for i := 0; i < 100; i++ {
		item := entities.Employee{
			Id:          faker.RandomUnixTime(),
			Name:        faker.Name(),
			Address:     faker.GetRealAddress().Address,
			PhoneNumber: faker.E164PhoneNumber(),
		}

		list = append(list, item)
	}
	return
}
func (m mockImpl) GetEmployeeById(id int64) (employee entities.Employee, er error) {
	return entities.Employee{
		Id:          id,
		Name:        faker.Name(),
		Address:     faker.GetRealAddress().Address,
		PhoneNumber: faker.E164PhoneNumber(),
	}, nil
}

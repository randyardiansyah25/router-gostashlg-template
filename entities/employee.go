package entities

type Employee struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type EmployeeFilter struct {
	Id int64 `form:"id" json:"id"`
}

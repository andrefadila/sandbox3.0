package employee

type Employee struct {
	ID           int
	Name         string
	DepartmentId int
}

type EmployeeService interface {
	GetEmployee(id int) (*Employee, error)
	GetEmployees() ([]*Employee, error)
	CreateEmployee(e *Employee) error
	UpdateEmployee(e *Employee) error
	DeleteEmployee(id int) error
}

package task

type Department struct {
	ID   int
	Name string
}

type Employee struct {
	ID           int
	Name         string
	DepartmentId int
}

type DepartmentService interface {
	GetDepartment(id int) (*Department, error)
	GetDepartments() ([]*Department, error)
	CreateDepartment(d *Department) error
	UpdateDepartment(d *Department) error
	DeleteDepartment(id int) error
}

type EmployeeService interface {
	GetEmployee(id int) (*Employee, error)
	GetEmployees() ([]*Employee, error)
	CreateEmployee(e *Employee) error
	UpdateEmployee(e *Employee) error
	DeleteEmployee(id int) error
}

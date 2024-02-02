package department

type Department struct {
	ID   int
	Name string
}

type DepartmentService interface {
	GetDepartment(id int) (*Department, error)
	GetDepartments() ([]*Department, error)
	CreateDepartment(d *Department) error
	UpdateDepartment(d *Department) error
	DeleteDepartment(id int) error
}

package main

import "fmt"

// Roles
const (
	MANAGER = "manager"
)

// Employee represents an employee
type Employee struct {
	Name      string
	Role      string
	MinSalary int
	MaxSalary int
}

// Builder defines a builder type
type Builder struct {
	employee Employee
}

// Build returns a new employee
func (b *Builder) Build() Employee {
	return b.employee
}

// Name sets an employee name
func (b *Builder) Name(name string) *Builder {
	b.employee.Name = name
	return b
}

// Role sets employees roles, and min / max salary if not set
func (b *Builder) Role(role string) *Builder {
	if role == MANAGER {
		b.employee.MinSalary = 140000
		b.employee.MaxSalary = 250000
	}
	b.employee.Role = role
	return b
}

func main() {
	eBuilder := &Builder{}
	employee := eBuilder.
		Name("John Foo").
		Role("manager").
		Build()

	fmt.Printf("%#v\n", employee)
	fmt.Printf("%+v\n", employee)
}

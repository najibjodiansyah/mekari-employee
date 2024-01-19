package service

import (
	"errors"

	"github.com/najibjodiansyah/mekari-employee/model/domain"
	"github.com/najibjodiansyah/mekari-employee/repository"
)

type EmployeeService interface {
	Find() ([]*domain.Employee, error)
	FindById(id int) (*domain.Employee, error)
	Create(empCreate *domain.Employee) error
	Update(id int, empUpdate domain.Employee) error
	Delete(id int) error
}

type EmployeeServiceImpl struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{
		employeeRepo: employeeRepo,
	}
}

func (u *EmployeeServiceImpl) Find() ([]*domain.Employee, error) {
	employees, err := u.employeeRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (u *EmployeeServiceImpl) FindById(id int) (*domain.Employee, error) {
	employee, err := u.employeeRepo.SelectById(id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (u *EmployeeServiceImpl) Create(empCreate *domain.Employee) error {
	if empCreate.HireDate != nil {
		if !empCreate.ParseRFC3339(*empCreate.HireDate) {
			return errors.New("HireDate format must be RFC3339")
		}
	}
	err := u.employeeRepo.Insert(empCreate)
	if err != nil {
		return err
	}
	return nil
}

func (u *EmployeeServiceImpl) Update(id int, empUpdate domain.Employee) error {
	employee, err := u.employeeRepo.SelectById(id)
	if err != nil {
		return err
	}

	// check datetime format RFC3339
	if empUpdate.HireDate != nil {
		if !employee.ParseRFC3339(*empUpdate.HireDate) {
			return errors.New("HireDate format must be RFC3339")
		}
	}
	if empUpdate.FirstName != "" {
		employee.FirstName = empUpdate.FirstName
	}
	if empUpdate.LastName != "" {
		employee.LastName = empUpdate.LastName
	}
	if empUpdate.Email != "" {
		employee.Email = empUpdate.Email
	}
	if empUpdate.HireDate != nil {
		employee.HireDate = empUpdate.HireDate
	}

	err = u.employeeRepo.Update(employee)
	if err != nil {
		return err
	}
	return nil
}

func (u *EmployeeServiceImpl) Delete(id int) error {
	_, err := u.employeeRepo.SelectById(id)
	if err != nil {
		return err
	}
	err = u.employeeRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

package service

import (
	"context"
	"errors"

	"github.com/najibjodiansyah/mekari-employee/model/domain"
	"github.com/najibjodiansyah/mekari-employee/repository"
)

type EmployeeService interface {
	Find(ctx context.Context) ([]*domain.Employee, error)
	FindById(ctx context.Context, id int) (*domain.Employee, error)
	Create(ctx context.Context, mpCreate *domain.Employee) error
	Update(ctx context.Context, id int, empUpdate domain.Employee) error
	Delete(ctx context.Context, id int) error
}

type EmployeeServiceImpl struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{
		employeeRepo: employeeRepo,
	}
}

func (u *EmployeeServiceImpl) Find(ctx context.Context) ([]*domain.Employee, error) {
	employees, err := u.employeeRepo.SelectAll(ctx)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (u *EmployeeServiceImpl) FindById(ctx context.Context, id int) (*domain.Employee, error) {
	employee, err := u.employeeRepo.SelectById(ctx, id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (u *EmployeeServiceImpl) Create(ctx context.Context, empCreate *domain.Employee) error {
	if empCreate.HireDate != nil {
		if !empCreate.ParseRFC3339(*empCreate.HireDate) {
			return errors.New("HireDate format must be RFC3339")
		}
	}
	err := u.employeeRepo.Insert(ctx, empCreate)
	if err != nil {
		return err
	}
	return nil
}

func (u *EmployeeServiceImpl) Update(ctx context.Context, id int, empUpdate domain.Employee) error {
	employee, err := u.employeeRepo.SelectById(ctx, id)
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

	err = u.employeeRepo.Update(ctx, employee)
	if err != nil {
		return err
	}
	return nil
}

func (u *EmployeeServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := u.employeeRepo.SelectById(ctx, id)
	if err != nil {
		return err
	}
	err = u.employeeRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/najibjodiansyah/mekari-employee/model"
	"github.com/najibjodiansyah/mekari-employee/model/domain"
	"github.com/uptrace/bun"
)

type EmployeeRepository interface {
	SelectAll() ([]*domain.Employee, error)
	SelectById(id int) (*domain.Employee, error)
	Insert(employee *domain.Employee) error
	Update(employee *domain.Employee) error
	Delete(id int) error
}

type EmployeeRepositoryImpl struct {
	db bun.IDB
}

func NewEmployeeRepository(db bun.IDB) EmployeeRepository {
	return &EmployeeRepositoryImpl{
		db: db,
	}
}

func (r *EmployeeRepositoryImpl) SelectAll() ([]*domain.Employee, error) {
	ctx := context.Background()
	var employees []*domain.Employee
	err := r.db.NewSelect().Model(&employees).Scan(ctx)
	if err != nil {
		log.Error(err)
		if err.Error() == model.NoRowsInResultSet {
			return nil, errors.New(model.ErrorNotFound)
		}
		return nil, errors.New(model.ErrorGetData)
	}
	return employees, nil
}

func (r *EmployeeRepositoryImpl) SelectById(id int) (*domain.Employee, error) {
	ctx := context.Background()
	var employee domain.Employee
	err := r.db.NewSelect().Model(&employee).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Error(err)
		if err.Error() == model.NoRowsInResultSet {
			return nil, errors.New(model.ErrorNotFound)
		}
		return nil, errors.New(model.ErrorGetData)
	}
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) Insert(emplpoyee *domain.Employee) error {
	ctx := context.Background()
	_, err := r.db.NewInsert().Model(emplpoyee).Exec(ctx)
	if err != nil {
		log.Error(err)
		return errors.New(model.ErrorInsertData)
	}
	return nil
}

func (r *EmployeeRepositoryImpl) Update(employee *domain.Employee) error {
	ctx := context.Background()
	_, err := r.db.NewUpdate().Model(employee).WherePK().Exec(ctx)
	if err != nil {
		log.Error(err)
		return errors.New(model.ErrorUpdateData)
	}
	return nil
}

func (r *EmployeeRepositoryImpl) Delete(id int) error {
	ctx := context.Background()
	_, err := r.db.NewDelete().Model(&domain.Employee{}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		log.Error(err)
		return errors.New(model.ErrorDeleteData)
	}
	return nil
}

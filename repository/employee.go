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
	SelectAll(ctx context.Context) ([]*domain.Employee, error)
	SelectById(ctx context.Context, id int) (*domain.Employee, error)
	Insert(ctx context.Context, employee *domain.Employee) error
	Update(ctx context.Context, employee *domain.Employee) error
	Delete(ctx context.Context, id int) error
}

type EmployeeRepositoryImpl struct {
	db bun.IDB
}

func NewEmployeeRepository(db bun.IDB) EmployeeRepository {
	return &EmployeeRepositoryImpl{
		db: db,
	}
}

func (r *EmployeeRepositoryImpl) SelectAll(ctx context.Context) ([]*domain.Employee, error) {
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

func (r *EmployeeRepositoryImpl) SelectById(ctx context.Context, id int) (*domain.Employee, error) {
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

func (r *EmployeeRepositoryImpl) Insert(ctx context.Context, emplpoyee *domain.Employee) error {
	_, err := r.db.NewInsert().Model(emplpoyee).Exec(ctx)
	if err != nil {
		log.Error(err)
		return errors.New(model.ErrorInsertData)
	}
	return nil
}

func (r *EmployeeRepositoryImpl) Update(ctx context.Context, employee *domain.Employee) error {
	_, err := r.db.NewUpdate().Model(employee).WherePK().Exec(ctx)
	if err != nil {
		log.Error(err)
		return errors.New(model.ErrorUpdateData)
	}
	return nil
}

func (r *EmployeeRepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := r.db.NewDelete().Model(&domain.Employee{}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		log.Error(err)
		return errors.New(model.ErrorDeleteData)
	}
	return nil
}

package repository

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/najibjodiansyah/mekari-employee/model/domain"
	"github.com/najibjodiansyah/mekari-employee/model/web"
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
		if err.Error() == web.NoRowsInResultSet {
			return nil, errors.New(web.ErrorNotFound)
		}
		return nil, errors.New(web.ErrorGetData)
	}
	return employees, nil
}

func (r *EmployeeRepositoryImpl) SelectById(id int) (*domain.Employee, error) {
	ctx := context.Background()
	var employee domain.Employee
	err := r.db.NewSelect().Model(&employee).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Error(err)
		if err.Error() == web.NoRowsInResultSet {
			return nil, errors.New(web.ErrorNotFound)
		}
		return nil, errors.New(web.ErrorGetData)
	}
	return &employee, nil
}

func (r *EmployeeRepositoryImpl) Insert(emplpoyee *domain.Employee) error {
	ctx := context.Background()
	_, err := r.db.NewInsert().Model(emplpoyee).Exec(ctx)
	if err != nil {
		log.Error(err)
		return errors.New(web.ErrorInsertData)
	}
	return nil
}

func (r *EmployeeRepositoryImpl) Update(employee *domain.Employee) error {
	ctx := context.Background()
	_, err := r.db.NewUpdate().Model(employee).WherePK().Exec(ctx)
	if err != nil {
		log.Error(err)
		return errors.New(web.ErrorUpdateData)
	}
	return nil
}

func (r *EmployeeRepositoryImpl) Delete(id int) error {
	ctx := context.Background()
	_, err := r.db.NewDelete().Model(&domain.Employee{}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		log.Error(err)
		return errors.New(web.ErrorDeleteData)
	}
	return nil
}

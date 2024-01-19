package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/najibjodiansyah/mekari-employee/model/domain"
	"github.com/najibjodiansyah/mekari-employee/service"
	mock_repository "github.com/najibjodiansyah/mekari-employee/tests/mocks/repo"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/golang/mock/gomock"
)

func TestEmployeeService(t *testing.T) {
	Convey("Test Service Employee", t, func() {
		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockEmployeeRepo := mock_repository.NewMockEmployeeRepository(mockCtrl)
		employeeService := service.NewEmployeeService(mockEmployeeRepo)

		Convey("Test Get Employees", func() {
			Convey("Test Get Employees Success", func() {
				hireDate := "2020-01-01 15:00:00"
				employee := []*domain.Employee{
					{
						Id:        1,
						FirstName: "Najib",
						LastName:  "Jodiansyah",
						Email:     "najib@test.com",
						HireDate:  &hireDate,
					},
				}

				mockEmployeeRepo.EXPECT().SelectAll(ctx).Return(employee, nil)
				e, err := employeeService.Find(ctx)
				So(err, ShouldBeNil)
				So(e, ShouldNotBeNil)
			})
			Convey("Test Get Employees Failed", func() {
				mockEmployeeRepo.EXPECT().SelectAll(ctx).Return(nil, errors.New("error"))
				e, err := employeeService.Find(ctx)
				So(err, ShouldNotBeNil)
				So(e, ShouldBeNil)
			})
		})
		Convey("Test Get Employee By Id", func() {
			Convey("Test Get Employee By Id Success", func() {
				hireDate := "2020-01-01 15:00:00"
				employee := &domain.Employee{
					Id:        1,
					FirstName: "Najib",
					LastName:  "Jodiansyah",
					Email:     "najib@test.com",
					HireDate:  &hireDate,
				}

				mockEmployeeRepo.EXPECT().SelectById(ctx, 1).Return(employee, nil)
				e, err := employeeService.FindById(ctx, 1)
				So(err, ShouldBeNil)
				So(e, ShouldNotBeNil)
			})
			Convey("Test Get Employee By Id Failed", func() {
				mockEmployeeRepo.EXPECT().SelectById(ctx, 1).Return(nil, errors.New("error"))
				e, err := employeeService.FindById(ctx, 1)
				So(err, ShouldNotBeNil)
				So(e, ShouldBeNil)
			})
		})
		Convey("Test Create Employee", func() {
			Convey("Test Create Employee Success", func() {
				hireDate := "2020-01-01 15:00:00"
				employee := &domain.Employee{
					Id:        1,
					FirstName: "Najib",
					LastName:  "Jodiansyah",
					Email:     "najib@test.com",
					HireDate:  &hireDate,
				}
				mockEmployeeRepo.EXPECT().Insert(ctx, employee).Return(nil)
				err := employeeService.Create(ctx, employee)
				So(err, ShouldBeNil)
			})
			Convey("Test Create Employee Failed", func() {
				hireDate := "2020-01-01 15:00:00"
				employee := &domain.Employee{
					Id:        1,
					FirstName: "Najib",
					LastName:  "Jodiansyah",
					Email:     "najib@test.com",
					HireDate:  &hireDate,
				}
				mockEmployeeRepo.EXPECT().Insert(ctx, employee).Return(errors.New("error"))
				err := employeeService.Create(ctx, employee)
				So(err, ShouldNotBeNil)
			})
			Convey("Test Create Employee Failed Parse Datetime", func() {
				hireDate := "2020-01-01 15:00:00Z"
				employee := &domain.Employee{
					Id:        1,
					FirstName: "Najib",
					LastName:  "Jodiansyah",
					Email:     "najib@test.com",
					HireDate:  &hireDate,
				}
				err := employeeService.Create(ctx, employee)
				So(err.Error(), ShouldEndWith, "HireDate format must be RFC3339")
			})
		})
		Convey("Test Delete Employee", func() {
			Convey("Test Delete Employee Not Found", func() {
				mockEmployeeRepo.EXPECT().SelectById(ctx, 1).Return(nil, errors.New("error"))
				err := employeeService.Delete(ctx, 1)
				So(err, ShouldNotBeNil)
			})
			Convey("Test Delete Employee Success", func() {
				mockEmployeeRepo.EXPECT().SelectById(ctx, 1).Return(&domain.Employee{}, nil)
				mockEmployeeRepo.EXPECT().Delete(ctx, 1).Return(nil)
				err := employeeService.Delete(ctx, 1)
				So(err, ShouldBeNil)
			})
			Convey("Test Delete Employee Failed", func() {
				mockEmployeeRepo.EXPECT().SelectById(ctx, 1).Return(&domain.Employee{}, nil)
				mockEmployeeRepo.EXPECT().Delete(ctx, 1).Return(errors.New("error"))
				err := employeeService.Delete(ctx, 1)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

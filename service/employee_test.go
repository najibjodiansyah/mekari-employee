package service_test

import (
	"context"
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
				_, err := employeeService.Find(ctx)
				So(err, ShouldBeNil)
			})
		})
	})

}

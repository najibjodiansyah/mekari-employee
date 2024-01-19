package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/najibjodiansyah/mekari-employee/controller"
	mock_service "github.com/najibjodiansyah/mekari-employee/tests/mocks/service"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEmployeeControllerImpl_Get(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	/* 	INTEGRATION TEST SET UP
	- setup db test from env
	*/

	v := validator.New()
	mockEmployeeService := mock_service.NewMockEmployeeService(mockCtrl)
	employeeController := controller.NewEmployeeController(mockEmployeeService, v)
	app := controller.RegisterApp(employeeController)

	Convey("Test Get Employees", t, func() {
		Convey("Test Get Employees Success", func() {
			request := httptest.NewRequest(http.MethodGet, "/v1/employees", nil)
			request.Header.Add("apikey", "MK421")
			mockEmployeeService.EXPECT().Find(request.Context()).Return(nil, nil)
			response, err := app.Test(request)
			So(err, ShouldNotBeNil)
			So(response.StatusCode, ShouldEqual, http.StatusUnauthorized)
		})
	})

}

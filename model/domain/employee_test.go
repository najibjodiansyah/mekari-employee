package domain_test

import (
	"testing"

	"github.com/najibjodiansyah/mekari-employee/model/domain"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEmployeeDomain(t *testing.T) {
	Convey("Unit Test Employee Domain", t, func() {
		Convey("ParseRFC3339", func() {
			Convey("Success", func() {
				val := "2021-01-01 00:00:00"
				employee := domain.Employee{
					HireDate: &val,
				}
				result := employee.ParseRFC3339(*employee.HireDate)
				So(result, ShouldEqual, true)
			})
			Convey("Fail", func() {
				val := "2021-01-01T00:00:00Z"
				employee := domain.Employee{
					HireDate: &val,
				}
				result := employee.ParseRFC3339(*employee.HireDate)
				So(result, ShouldEqual, false)
			})
		})
	})
}

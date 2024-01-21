package controller_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/najibjodiansyah/mekari-employee/controller"
	"github.com/najibjodiansyah/mekari-employee/model/domain"
	"github.com/najibjodiansyah/mekari-employee/pkg/config"
	"github.com/najibjodiansyah/mekari-employee/repository"
	"github.com/najibjodiansyah/mekari-employee/service"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	. "github.com/smartystreets/goconvey/convey"
)

func setUpDBTest() *bun.DB {
	ctx := context.Background()

	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Config.PgCfg.Username,
		config.Config.PgCfg.Password,
		config.Config.PgCfg.Host,
		config.Config.PgCfg.Port,
		config.Config.PgCfg.Database,
	)

	dbConn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dbURI)))

	if err := dbConn.PingContext(ctx); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	db := bun.NewDB(dbConn, pgdialect.New())
	// db.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName(config.Config.PgCfg.Database)))

	// queryLog := bundebug.NewQueryHook(bundebug.WithVerbose(true))
	// db.AddQueryHook(queryLog)

	return db
}

func TruncateDB(db *bun.DB, domain interface{}) {
	fmt.Println("Truncate DB")
	ctx := context.Background()
	_, err := db.NewTruncateTable().Model(domain).Exec(ctx)
	if err != nil {
		log.Fatalf("Error truncate DB: %v", err)
	}
	fmt.Println("Truncate DB Success")
}

func setUpAppTest() (*fiber.App, repository.EmployeeRepository, *bun.DB) {
	db := setUpDBTest()
	v := validator.New()
	repo := repository.NewEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	ctrl := controller.NewEmployeeController(svc, v)
	app := controller.RegisterApp(ctrl)
	return app, repo, db
}

func TestEmployeeController(t *testing.T) {
	/* 	INTEGRATION TEST SET UP
	- setup db test and table
	*/
	os.Setenv("MKR_ENV", "test")
	pathGroup := "/v1/employees"
	Convey("Test Middleware", t, func() {
		app, _, _ := setUpAppTest()
		Convey("Test Success", func() {
			request := httptest.NewRequest(http.MethodGet, pathGroup, nil)
			response, err := app.Test(request)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 401)
		})
		Convey("Test Failed", func() {
			request := httptest.NewRequest(http.MethodGet, pathGroup, nil)
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 200)
		})
	})

	Convey("Test Get Employees", t, func() {
		os.Setenv("MKR_ENV", "test")
		app, repo, db := setUpAppTest()
		Convey("Test Get Employees Success", func() {
			var e *domain.Employee = &domain.Employee{
				FirstName: "najib",
				LastName:  "jodi",
				Email:     "mail@mail.com",
			}
			ctx := context.Background()

			err := repo.Insert(ctx, e)
			if err != nil {
				log.Fatalf("Error insert DB: %v", err)
			}

			request := httptest.NewRequest(http.MethodGet, pathGroup, nil)
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request, 1)
			response.Body.Close()
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 200)
			TruncateDB(db, &domain.Employee{})
		})
		// Convey("Test Get Employees Failed Not Found", func() {
		// 	request := httptest.NewRequest(http.MethodGet, pathGroup, nil)
		// 	request.Header.Add("apikey", "MK421")
		// 	response, err := app.Test(request, 1)
		// 	response.Body.Close()
		// 	So(err, ShouldBeNil)
		// 	So(response.StatusCode, ShouldEqual, 404)
		// 	TruncateDB(db, &domain.Employee{})
		// })
		Convey("Test Get Employees Internal Server Error", func() {
			db.Close()
			request := httptest.NewRequest(http.MethodGet, pathGroup, nil)
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request, 1)
			response.Body.Close()
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 500)
		})
	})

	Convey("Test Get by id Employees", t, func() {
		app, repo, db := setUpAppTest()
		Convey("Test Get Employees Success", func() {
			var e *domain.Employee = &domain.Employee{
				FirstName: "najib",
				LastName:  "jodi",
				Email:     "najib@mail.com",
			}
			ctx := context.Background()

			err := repo.Insert(ctx, e)
			if err != nil {
				log.Fatalf("Error insert DB: %v", err)
			}

			request := httptest.NewRequest(http.MethodGet, pathGroup+"/1", nil)
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request, 1)
			response.Body.Close()
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 200)
			TruncateDB(db, &domain.Employee{})
		})
		Convey("Test Get Employees Failed type id", func() {
			request := httptest.NewRequest(http.MethodGet, pathGroup+"/a", nil)
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request, 1)
			response.Body.Close()
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
		})
		Convey("Test Get Employees Failed Not Found", func() {
			request := httptest.NewRequest(http.MethodGet, pathGroup+"/1", nil)
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request, 100)
			response.Body.Close()
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 404)
		})
		Convey("Test Get Employees Internal Server Error", func() {
			db.Close()
			request := httptest.NewRequest(http.MethodGet, pathGroup+"/1", nil)
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request, 1)
			response.Body.Close()
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 500)
		})
	})

	Convey("Test Post Employees", t, func() {
		app, _, db := setUpAppTest()
		Convey("Test Post Employees Success", func() {
			reqBody := strings.NewReader(`{"first_name":"najib","last_name":"jodi","email":"najib@gmail.com", "hire_date":"2021-01-01 00:00:00"}`)
			request := httptest.NewRequest(http.MethodPost, pathGroup, reqBody)
			request.Header.Add("Content-Type", "application/json")
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 201)
			TruncateDB(db, &domain.Employee{})
		})
		Convey("Test Post Employees Failed Bad Request", func() {
			reqBody := strings.NewReader(`{"firstname":"najib","lastname":"jodi","email":"najib@gmail.com", "hiredate":"2021-01-01 00:00:00"}`)
			request := httptest.NewRequest(http.MethodPost, pathGroup, reqBody)
			request.Header.Add("Content-Type", "application/json")
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
		})
		Convey("Test Post Employees Failed unprocessable entity", func() {
			reqBody := strings.NewReader(`{first_name:"najib","last_name":"jodi","email":"najib@gmail.com", "hiredate":"2021-01-01 00:00:00"}`)
			request := httptest.NewRequest(http.MethodPost, pathGroup, reqBody)
			request.Header.Add("Content-Type", "application/json")
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
		})
		// should be in service ;ayer
		Convey("Test Post Employees Failed invalid type hire date", func() {
			reqBody := strings.NewReader(`{"first_name":"najib","last_name":"jodi","email":"najib@gmail.com", "hire_date":"2021-01-01T00:00:00Z"}`)
			request := httptest.NewRequest(http.MethodPost, pathGroup, reqBody)
			request.Header.Add("Content-Type", "application/json")
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
		})
		Convey("Test Post Employees Internal Server Error", func() {
			db.Close()
			reqBody := strings.NewReader(`{"first_name":"najib","last_name":"jodi","email":"najib@gmail.com", "hire_date":"2021-01-01 00:00:00"}`)
			request := httptest.NewRequest(http.MethodPost, pathGroup, reqBody)
			request.Header.Add("Content-Type", "application/json")
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 500)
		})
	})

	// Convey("Test Update Employees", t, func() {
	// 	Convey("Test Update Employees Success", func() {

	// 	})
	// 	Convey("Test Update Employees Failed Bad Request", func() {

	// 	})
	// 	Convey("Test Update Employees Failed Not Found", func() {

	// 	})
	// 	Convey("Test Update Employees Failed unprocessable entity", func() {

	// 	})
	// 	// should be in service ;ayer
	// 	Convey("Test Update Employees Failed invalid type hire date", func() {

	// 	})
	// 	Convey("Test Update Employees Internal Server Error", func() {

	// 	})
	// })

	Convey("Test Delete Employees", t, func() {
		app, repo, db := setUpAppTest()
		Convey("Test Delete Employees Success", func() {
			var e *domain.Employee = &domain.Employee{
				FirstName: "najib",
				LastName:  "jodi",
				Email:     "najib@mail.com",
			}
			ctx := context.Background()

			err := repo.Insert(ctx, e)
			if err != nil {
				log.Fatalf("Error insert DB: %v", err)
			}

			request := httptest.NewRequest(http.MethodDelete, pathGroup+"/1", nil)
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 200)
			TruncateDB(db, &domain.Employee{})
		})
		Convey("Test Delete Employees Failed type id", func() {
			request := httptest.NewRequest(http.MethodDelete, pathGroup+"/a", nil)
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
		})
		Convey("Test Delete Employees Failed Not Found", func() {
			request := httptest.NewRequest(http.MethodDelete, pathGroup+"/100", nil)
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 404)
		})
		Convey("Test Delete Employees Internal Server Error", func() {
			db.Close()
			request := httptest.NewRequest(http.MethodDelete, pathGroup+"/1", nil)
			request.Header.Add("apikey", "MK421")
			response, err := app.Test(request)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 500)
		})
	})

}

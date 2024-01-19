package controller

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/najibjodiansyah/mekari-employee/pkg/postgres"
	"github.com/najibjodiansyah/mekari-employee/repository"
	"github.com/najibjodiansyah/mekari-employee/service"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bunotel"
)

type EmployeeController interface {
	Get(c *fiber.Ctx) error
	Post(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Put(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

func NewEmployeeController(employeeService service.EmployeeService, validator *validator.Validate) EmployeeController {
	return &EmployeeControllerImpl{
		UserService: employeeService,
		Validate:    validator,
	}
}

func EmployeeApi() {
	ctx := context.Background()
	dbConn := postgres.NewPostgresConn("postgres://postgres:mysecret@postgres:5432/mekari_employee?sslmode=disable")
	defer dbConn.Close()

	if err := dbConn.PingContext(ctx); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	v := validator.New()
	db := bun.NewDB(dbConn, pgdialect.New())
	db.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName("mekari_employee")))
	repo := repository.NewEmployeeRepository(db)
	service := service.NewEmployeeService(repo)
	controller := NewEmployeeController(service, v)
	app := fiber.New(
		fiber.Config{
			Prefork: true,
		},
	)

	// queryLog := bundebug.NewQueryHook(bundebug.WithVerbose(true))
	// db.AddQueryHook(queryLog)

	app.Use(
		logger.New(logger.Config{
			Format: "${pid} ${locals:requestid} [${ip}]:${port} ${status} - ${method} ${path}​\n",
		}),
		func(c *fiber.Ctx) error {
			if c.Get("apikey") != "MK421" {
				return c.Status(401).JSON(&fiber.Map{
					"message": "Unauthorized",
				})
			}
			c.Set("Content-Type", "application/json")
			return c.Next()
		})
	userRoute := app.Group("/v1/employees")
	userRoute.Get("/", controller.Get)
	userRoute.Post("/", controller.Post)
	userRoute.Get("/:id", controller.GetById)
	userRoute.Put("/:id", controller.Put)
	userRoute.Delete("/:id", controller.Delete)

	log.Fatal(app.Listen(":3000"))
}

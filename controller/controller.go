package controller

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/najibjodiansyah/mekari-employee/repository"
	"github.com/najibjodiansyah/mekari-employee/service"
	"github.com/uptrace/bun"
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
	}
}

func RegisterApp(controller EmployeeController) *fiber.App {
	app := fiber.New(
		fiber.Config{
			Prefork: true,
		},
	)
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

	return app
}

func EmployeeApi(db *bun.DB) {
	v := validator.New()
	repo := repository.NewEmployeeRepository(db)
	service := service.NewEmployeeService(repo)
	controller := NewEmployeeController(service, v)

	app := RegisterApp(controller)

	log.Fatal(app.Listen(":3000"))
}

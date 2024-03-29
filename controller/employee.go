package controller

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/najibjodiansyah/mekari-employee/model/domain"
	"github.com/najibjodiansyah/mekari-employee/model/web"
	"github.com/najibjodiansyah/mekari-employee/service"

	"github.com/gofiber/fiber/v2"
)

type EmployeeControllerImpl struct {
	UserService service.EmployeeService
	Validate    *validator.Validate
}

func (uc *EmployeeControllerImpl) Get(c *fiber.Ctx) error {
	users, err := uc.UserService.Find()
	if err != nil {
		if err.Error() == web.ErrorNotFound {
			return c.Status(404).JSON(&fiber.Map{
				"message": err.Error(),
				"data":    users,
			})
		}
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
			"data":    users,
		})
	}
	return c.JSON(&fiber.Map{
		"message": web.Success,
		"data":    users,
	})
}

func (uc *EmployeeControllerImpl) Post(c *fiber.Ctx) error {
	var user *domain.Employee
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": web.ErrorUnprocessableEntity + err.Error(),
		})
	}
	if err := uc.Validate.Struct(user); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": web.ErrorBadRequest + err.Error(),
		})
	}
	err := uc.UserService.Create(user)
	if err != nil {
		if err.Error() == "HireDate format must be RFC3339" {
			return c.Status(400).JSON(&fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(500).JSON(&fiber.Map{
			"message": web.ErrorInternal + err.Error(),
		})
	}
	return c.Status(201).JSON(&fiber.Map{
		"message": web.Created,
	})
}

func (uc *EmployeeControllerImpl) GetById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": web.ErrorBadRequest + err.Error(),
		})
	}
	user, err := uc.UserService.FindById(id)
	if err != nil {
		if err.Error() == web.ErrorNotFound {
			return c.Status(404).JSON(&fiber.Map{
				"message": err.Error(),
				"data":    user,
			})
		}
		return c.Status(500).JSON(&fiber.Map{
			"message": web.ErrorInternal + err.Error(),
			"data":    user,
		})
	}
	return c.JSON(&fiber.Map{
		"message": web.Success,
		"data":    user,
	})
}

func (uc *EmployeeControllerImpl) Put(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": web.ErrorBadRequest + err.Error(),
		})
	}
	var user domain.Employee
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": web.ErrorUnprocessableEntity + err.Error(),
		})
	}
	if err := uc.Validate.Struct(user); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": web.ErrorBadRequest + err.Error(),
		})
	}
	err = uc.UserService.Update(id, user)
	if err != nil {
		if err.Error() == "HireDate format must be RFC3339" {
			return c.Status(400).JSON(&fiber.Map{
				"message": err.Error(),
			})
		}
		if err.Error() == web.ErrorNotFound {
			return c.Status(404).JSON(&fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(500).JSON(&fiber.Map{
			"message": web.ErrorInternal + err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": web.Updated,
	})
}

func (uc *EmployeeControllerImpl) Delete(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": web.ErrorBadRequest + err.Error(),
		})
	}
	err = uc.UserService.Delete(id)
	if err != nil {
		if err.Error() == web.ErrorNotFound {
			return c.Status(404).JSON(&fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(500).JSON(&fiber.Map{
			"message": web.ErrorInternal + err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": web.Deleted,
	})
}

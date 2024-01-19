package domain

import (
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/uptrace/bun"
)

type Employee struct {
	bun.BaseModel `bun:"table:employees,alias:e"`

	Id        int     `bun:"id,pk,autoincrement" json:"id"`
	FirstName string  `bun:"first_name" json:"first_name" validate:"required"`
	LastName  string  `bun:"last_name" json:"last_name" validate:"required"`
	Email     string  `bun:"email" json:"email" validate:"required,email"`
	HireDate  *string `bun:"hire_date" json:"hire_date"`
}

func (e *Employee) ParseRFC3339(value string) bool {
	format := "2006-01-02 15:04:05"
	_, err := time.Parse(format, value)
	if err != nil {
		log.Error(err)
		return false
	} else {
		return true
	}
}

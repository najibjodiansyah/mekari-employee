package main

import (
	"github.com/najibjodiansyah/mekari-employee/controller"
	"github.com/najibjodiansyah/mekari-employee/pkg/postgres"
)

func main() {
	controller.EmployeeApi(postgres.DB())
}

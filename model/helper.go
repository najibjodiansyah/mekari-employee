package model

const (
	ErrorNotFound            = "no employee data found"
	ErrorInternal            = "internal server error"
	ErrorBadRequest          = "bad Request"
	ErrorUnprocessableEntity = "unprocessable entity"

	Success = "Success"
	Created = "Success Created"
	Updated = "Success Updated"
	Deleted = "Success Deleted"

	NoRowsInResultSet = "sql: no rows in result set"
	ErrorGetData      = "error get data"
	ErrorInsertData   = "error insert data"
	ErrorUpdateData   = "error update data"
	ErrorDeleteData   = "error delete data"
)

type CreateEmployeeRequest struct {
	FirstName string  `json:"first_name" validate:"required"`
	LastName  string  `json:"last_name" validate:"required"`
	Email     string  `json:"email" validate:"required,email"`
	HireDate  *string `json:"hire_date"`
}

type UpdateEmployeeRequest struct {
	Id        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	HireDate  *string `json:"hire_date"`
}

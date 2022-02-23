package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/venkat/customer/domain"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// CustomerHandler  represent the httphandler for customer
type CustomerHandler struct {
	CUsecase domain.CustomerUsecase
}

func NewCustomerHandler(e *echo.Echo, us domain.CustomerUsecase) {
	handler := &CustomerHandler{
		CUsecase: us,
	}
	e.GET("/customers", handler.Fetch)
	e.POST("/customers", handler.Store)
	e.PUT("/customers/:id", handler.Update)
	e.GET("/customers/:id", handler.GetByID)
	e.DELETE("/customers/:id", handler.Delete)
}
func (a *CustomerHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := a.CUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (a *CustomerHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}

	id := int64(idP)
	ctx := c.Request().Context()

	art, err := a.CUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)
}

// Store will store the customer by given request body
func (a *CustomerHandler) Store(c echo.Context) (err error) {
	var customer domain.Customer
	err = c.Bind(&customer)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	err = a.CUsecase.Store(ctx, &customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, customer)
}

// Store will update the customer by given request body
func (a *CustomerHandler) Update(c echo.Context) (err error) {
	var updatedFields map[string]interface{}
	err = c.Bind(&updatedFields)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}
	ctx := c.Request().Context()
	err = a.CUsecase.Update(ctx, int64(idP), updatedFields)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, "Data updated successfully")
}

func (a *CustomerHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}

	id := int64(idP)
	ctx := c.Request().Context()

	err = a.CUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

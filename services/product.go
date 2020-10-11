package services

import (
	"net/http"

	"{{<service_name>}}/models"

	"github.com/labstack/echo"
)


func ProductHandler(c echo.Context) error {


	product, err := models.NewProduct()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(product); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}


	// todo: check for role and allow user creation of certain roles
	if err := product.Save(c); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}



	// todo: change
	response := newProductResponse{
		Product:        *product,
	}

	return c.JSON(http.StatusOK, response)
}

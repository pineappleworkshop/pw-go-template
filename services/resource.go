package services

import (
	"net/http"
	"{{<service_name>}}/models"
	"github.com/labstack/echo"
)

func ResourceHandler(c echo.Context) error {
	resource, err := models.NewResource()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(resource); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := resource.Save(c); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := newResourceResponse{
		Resource:        *resource,
	}

	return c.JSON(http.StatusOK, response)
}

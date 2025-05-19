/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package cats

import (
	"net/http"

	"github.com/WahidinAlambiyah/golang/pkg/api/handlers"
	"github.com/WahidinAlambiyah/golang/pkg/api/helpers"
	"github.com/WahidinAlambiyah/golang/pkg/db/models"
	"github.com/WahidinAlambiyah/golang/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {

	id := c.Param("id")

	if id == "" {
		return helpers.Error(c, constants.ERROR_ID_NOT_FOUND, nil)
	}

	m, err := models.CatModel().Find(id)

	if err != nil {
		return helpers.Error(c, err, nil)
	}

	return c.JSON(http.StatusOK, handlers.Success(m.MapToForm()))

}

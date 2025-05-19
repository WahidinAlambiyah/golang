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

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {

	ms, err := models.CatModel().FindAll()

	if err != nil {
		return helpers.Error(c, err, nil)
	}

	var payload []*models.CatForm

	for _, m := range ms {
		f := m.MapToForm()
		payload = append(payload, f)
	}

	return c.JSON(http.StatusOK, handlers.Success(payload))

}

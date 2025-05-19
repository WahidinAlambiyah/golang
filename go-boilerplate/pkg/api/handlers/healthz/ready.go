/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package healthz

import (
	"net/http"

	"github.com/WahidinAlambiyah/golang/pkg/api/handlers"
	"github.com/WahidinAlambiyah/golang/pkg/api/helpers"
	"github.com/WahidinAlambiyah/golang/pkg/clients/dbc"

	"github.com/labstack/echo/v4"
)

func Ready(c echo.Context) error {
	dbClient := dbc.GetDBClient()
	if err := dbClient.Ping(); err != nil {
		return helpers.Error(c, err, nil)
	}

	payload := map[string]string{
		"message": "ready",
	}
	return c.JSON(http.StatusOK, handlers.Success(payload))
}

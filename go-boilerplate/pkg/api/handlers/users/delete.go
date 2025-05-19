/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package users

import (
	"net/http"

	"github.com/WahidinAlambiyah/golang/pkg/api/handlers"
	"github.com/WahidinAlambiyah/golang/pkg/clients/kratos"
	"github.com/WahidinAlambiyah/golang/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	id, err := handlers.GetUUIDParam(c.Param("id"))
	if err != nil {
		c.Logger().Error(constants.ERROR_ID_NOT_FOUND)
		return constants.ERROR_ID_NOT_FOUND
	}
	kratosCli := kratos.GetClient()
	if err := kratosCli.DeleteIdentity(id.String()); err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, handlers.Accepted())
}

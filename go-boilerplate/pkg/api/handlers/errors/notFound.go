/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package errors

import (
	"net/http"

	"github.com/WahidinAlambiyah/golang/pkg/api/handlers"
	"github.com/WahidinAlambiyah/golang/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func NotFound(c echo.Context) error {
	return c.JSON(
		http.StatusNotFound,
		handlers.BuildResponse(
			constants.STATUS_CODE_ROUTE_NOT_FOUND,
			constants.MSG_ROUTE_NOT_FOUND,
			[]string{},
			nil))
}

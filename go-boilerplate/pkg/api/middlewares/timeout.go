/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package middlewares

import (
	"time"

	"github.com/WahidinAlambiyah/golang/pkg/clients/service"
	"github.com/WahidinAlambiyah/golang/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TimeoutMiddleware() echo.MiddlewareFunc {
	serviceCli := service.GetClient()
	config := serviceCli.GetConfig()
	timeoutDuration := utils.IntFromStr(config.RequestTimeoutDuration)

	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Duration(timeoutDuration) * time.Second,
	})
}

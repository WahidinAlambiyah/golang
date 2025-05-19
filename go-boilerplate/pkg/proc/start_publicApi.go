/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package proc

import (
	"github.com/WahidinAlambiyah/golang/pkg/api/routers"
	"github.com/WahidinAlambiyah/golang/pkg/clients/service"
)

func StartPublicApi() {
	serviceCli := service.GetClient()
	config := serviceCli.GetConfig()
	routers.InitPublicAPIRouter()
	routers.PublicAPIRouter().Start(config.Host, config.PublicApiPort)
}

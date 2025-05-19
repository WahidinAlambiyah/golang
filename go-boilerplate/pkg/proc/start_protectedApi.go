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

func StartProtectedApi() {
	serviceCli := service.GetClient()
	config := serviceCli.GetConfig()
	routers.InitProtectedAPIRouter()
	routers.ProtectedAPIRouter().Start(config.Host, config.ProtectedApiPort)
}

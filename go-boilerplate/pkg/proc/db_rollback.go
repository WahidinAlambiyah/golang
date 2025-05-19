/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package proc

import (
	"github.com/WahidinAlambiyah/golang/pkg/clients/dbc"
	"github.com/WahidinAlambiyah/golang/pkg/clients/logger"
	"github.com/WahidinAlambiyah/golang/pkg/db/migrations"
)

func DBRollback() {
	logger.SetLogger(string(logger.DebugLvl))

	dbClient := dbc.GetDBClient()

	dbClient.InitDBConnection()

	migrations.Init(dbClient.DB)

	if err := migrations.Rollback(); err != nil {
		logger.Error("Failed to rollback migrations: %s", err)
	}

}

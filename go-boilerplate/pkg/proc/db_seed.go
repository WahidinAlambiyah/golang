/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package proc

import (
	"github.com/WahidinAlambiyah/golang/pkg/clients/dbc"
	"github.com/WahidinAlambiyah/golang/pkg/clients/logger"
	"github.com/WahidinAlambiyah/golang/pkg/db/seeds"
)

func DBSeed() {
	logger.SetLogger(string(logger.DebugLvl))

	dbClient := dbc.GetDBClient()

	dbClient.InitDBConnection()

	seeds.Init(dbClient.DB)

	if err := seeds.Apply(); err != nil {
		logger.Error("Failed to apply seeds: %s", err)
	}

	logger.Info("Seeds applied successfully")

}

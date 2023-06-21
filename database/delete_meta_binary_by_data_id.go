package database

import (
	"github.com/PretendoNetwork/team-kirby-clash-deluxe-secure/globals"
)

func DeleteMetaBinaryByDataID(dataID uint32) error {
	_, err := Postgres.Exec(`DELETE FROM meta_binaries WHERE data_id=$1`, dataID)
	if err != nil {
		globals.Logger.Critical(err.Error())
		return err
	}

	return nil
}

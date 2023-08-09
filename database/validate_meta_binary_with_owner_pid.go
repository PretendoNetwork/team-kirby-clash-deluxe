package database

import (
	"database/sql"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe-secure/globals"
)

func ValidateMetaBinaryByOwnerPID(dataID uint32, pid uint32) uint32 {
	var ownerPID uint32

	err := Postgres.QueryRow(`SELECT owner_pid FROM meta_binaries WHERE data_id=$1`, dataID).Scan(ownerPID)

	if err == sql.ErrNoRows {
		return nex.Errors.DataStore.NotFound
	}

	if err != nil && err != sql.ErrNoRows {
		globals.Logger.Critical(err.Error())
		return nex.Errors.DataStore.Unknown
	}

	if ownerPID != pid {
		return nex.Errors.DataStore.PermissionDenied
	}

	return 0
}

package database

import (
	"database/sql"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"
)

func ValidateMetaBinaryByOwnerPID(dataID uint32, pid uint32) *nex.Error {
	var ownerPID uint32

	err := Postgres.QueryRow(`SELECT owner_pid FROM meta_binaries WHERE data_id=$1`, dataID).Scan(&ownerPID)

	if err == sql.ErrNoRows {
		return nex.NewError(nex.ResultCodes.DataStore.NotFound, "Meta binary not found")
	}

	if err != nil && err != sql.ErrNoRows {
		globals.Logger.Critical(err.Error())
		return nex.NewError(nex.ResultCodes.DataStore.Unknown, err.Error())
	}

	if ownerPID != pid {
		return nex.NewError(nex.ResultCodes.DataStore.PermissionDenied, "Caller PID does not own the meta binary")
	}

	return nil
}

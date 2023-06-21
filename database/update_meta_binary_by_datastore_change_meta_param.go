package database

import (
	"time"

	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe-secure/globals"
	"github.com/lib/pq"
)

func UpdateMetaBinaryByDataStoreChangeMetaParam(dataStoreChangeMetaParam *datastore.DataStoreChangeMetaParam) error {
	// TODO - Check DataStoreChangeMetaParam flags for changes

	now := time.Now().Unix()

	_, err := Postgres.Exec(`
		UPDATE meta_binaries SET (
			name,
			data_type,
			meta_binary,
			period,
			tags,
			updated_time
		) = ($1, $2, $3, $4, $5, $6) WHERE data_id=$7`,
		dataStoreChangeMetaParam.Name,
		dataStoreChangeMetaParam.DataType,
		dataStoreChangeMetaParam.MetaBinary,
		dataStoreChangeMetaParam.Period,
		pq.Array(dataStoreChangeMetaParam.Tags),
		now,
		uint32(dataStoreChangeMetaParam.DataID),
	)
	if err != nil {
		globals.Logger.Critical(err.Error())
		return err
	}

	return nil
}

package database

func DeleteMetaBinaryByDataID(dataID uint32) error {
	_, err := Postgres.Exec(`DELETE FROM meta_binaries WHERE data_id=$1`, dataID)
	if err != nil {
		return err
	}

	return nil
}

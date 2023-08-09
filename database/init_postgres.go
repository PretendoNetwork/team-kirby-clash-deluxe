package database

import "github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"

func initPostgres() {
	var err error

	_, err = Postgres.Exec(`CREATE TABLE IF NOT EXISTS meta_binaries (
		data_id serial PRIMARY KEY,
		owner_pid integer,
		name text,
		data_type integer,
		meta_binary bytea,
		permission integer,
		del_permission integer,
		flag integer,
		period integer,
		tags text[],
		persistence_slot_id integer,
		extra_data text[],
		creation_time bigint,
		updated_time bigint,
		referred_time bigint,
		expire_time bigint
	)`)
	if err != nil {
		globals.Logger.Critical(err.Error())
		return
	}

	globals.Logger.Success("Postgres tables created")
}

package nex

import (
	datastore "github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe-secure/globals"
	nex_datastore "github.com/PretendoNetwork/team-kirby-clash-deluxe-secure/nex/datastore"
)

func registerNEXProtocols() {
	dataStoreProtocol := datastore.NewProtocol(globals.NEXServer)

	dataStoreProtocol.SearchObjectLight(nex_datastore.SearchObjectLight)
	dataStoreProtocol.PostMetaBinary(nex_datastore.PostMetaBinary)
	dataStoreProtocol.ChangeMeta(nex_datastore.ChangeMeta)
}

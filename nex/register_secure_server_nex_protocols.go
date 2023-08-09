package nex

import (
	datastore "github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"
	nex_datastore "github.com/PretendoNetwork/team-kirby-clash-deluxe/nex/datastore"
)

func registerSecureServerNEXProtocols() {
	dataStoreProtocol := datastore.NewProtocol(globals.SecureServer)

	dataStoreProtocol.SearchObjectLight(nex_datastore.SearchObjectLight)
	dataStoreProtocol.PostMetaBinary(nex_datastore.PostMetaBinary)
	dataStoreProtocol.ChangeMeta(nex_datastore.ChangeMeta)
}

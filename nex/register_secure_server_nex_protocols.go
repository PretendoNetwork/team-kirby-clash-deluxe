package nex

import (
	datastore "github.com/PretendoNetwork/nex-protocols-go/datastore"
	service_item_team_kirby_clash_deluxe "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe"

	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"
	nex_datastore "github.com/PretendoNetwork/team-kirby-clash-deluxe/nex/datastore"
	nex_service_item_team_kirby_clash_deluxe "github.com/PretendoNetwork/team-kirby-clash-deluxe/nex/service-item/team-kirby-clash-deluxe"
)

func registerSecureServerNEXProtocols() {
	dataStoreProtocol := datastore.NewProtocol(globals.SecureServer)

	dataStoreProtocol.SearchObjectLight(nex_datastore.SearchObjectLight)
	dataStoreProtocol.PostMetaBinary(nex_datastore.PostMetaBinary)
	dataStoreProtocol.ChangeMeta(nex_datastore.ChangeMeta)

	serviceItemTeamKirbyClashDeluxeProtocol := service_item_team_kirby_clash_deluxe.NewProtocol(globals.SecureServer)

	serviceItemTeamKirbyClashDeluxeProtocol.GetSupportID(nex_service_item_team_kirby_clash_deluxe.GetSupportID)
	serviceItemTeamKirbyClashDeluxeProtocol.ListServiceItemRequest(nex_service_item_team_kirby_clash_deluxe.ListServiceItemRequest)
	serviceItemTeamKirbyClashDeluxeProtocol.ListServiceItemResponse(nex_service_item_team_kirby_clash_deluxe.ListServiceItemResponse)
}

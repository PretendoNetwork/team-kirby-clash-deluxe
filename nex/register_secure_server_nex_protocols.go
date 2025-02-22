package nex

import (
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	service_item_team_kirby_clash_deluxe "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/team-kirby-clash-deluxe"

	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"
	nex_datastore "github.com/PretendoNetwork/team-kirby-clash-deluxe/nex/datastore"
	nex_service_item_team_kirby_clash_deluxe "github.com/PretendoNetwork/team-kirby-clash-deluxe/nex/service-item/team-kirby-clash-deluxe"
)

func registerSecureServerNEXProtocols() {
	dataStoreProtocol := datastore.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(dataStoreProtocol)

	dataStoreProtocol.SetHandlerSearchObjectLight(nex_datastore.SearchObjectLight)
	dataStoreProtocol.SetHandlerPostMetaBinary(nex_datastore.PostMetaBinary)
	dataStoreProtocol.SetHandlerChangeMeta(nex_datastore.ChangeMeta)

	serviceItemTeamKirbyClashDeluxeProtocol := service_item_team_kirby_clash_deluxe.NewProtocol(globals.SecureEndpoint)
	globals.SecureEndpoint.RegisterServiceProtocol(serviceItemTeamKirbyClashDeluxeProtocol)

	serviceItemTeamKirbyClashDeluxeProtocol.GetSupportID = nex_service_item_team_kirby_clash_deluxe.GetSupportID
	serviceItemTeamKirbyClashDeluxeProtocol.ListServiceItemRequest = nex_service_item_team_kirby_clash_deluxe.ListServiceItemRequest
	serviceItemTeamKirbyClashDeluxeProtocol.ListServiceItemResponse = nex_service_item_team_kirby_clash_deluxe.ListServiceItemResponse
}

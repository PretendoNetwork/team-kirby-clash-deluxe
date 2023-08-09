package nex_service_item_team_kirby_clash_deluxe

import (
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"

	"github.com/PretendoNetwork/nex-go"
	service_item_team_kirby_clash_deluxe "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

func ListServiceItemResponse(err error, client *nex.Client, callID uint32, requestID uint32) uint32 {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nex.Errors.ServiceItem.InvalidArgument
	}

	// * Stubbed
	catalog := service_item_team_kirby_clash_deluxe_types.NewServiceItemCatalog()
	catalog.Balance = service_item_team_kirby_clash_deluxe_types.NewServiceItemAmount()

	nullableCatalog := make([]*service_item_team_kirby_clash_deluxe_types.ServiceItemCatalog, 1)
	nullableCatalog[0] = catalog

	listServiceItemResponse := service_item_team_kirby_clash_deluxe_types.NewServiceItemListServiceItemResponse()
	listServiceItemResponse.HTTPStatus = 200
	listServiceItemResponse.CorrelationID = "8c6d0df0-e506-4f32-b730-1ccfe2476d7f,28b43a34-4709-49f9-b4f3-7ff76e8f79ab"
	listServiceItemResponse.NullableCatalog = nullableCatalog

	// * Stubbed
	rmcResponseStream := nex.NewStreamOut(globals.SecureServer)

	rmcResponseStream.WriteStructure(listServiceItemResponse)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCResponse(service_item_team_kirby_clash_deluxe.ProtocolID, callID)
	rmcResponse.SetSuccess(service_item_team_kirby_clash_deluxe.MethodListServiceItemResponse, rmcResponseBody)

	rmcResponseBytes := rmcResponse.Bytes()

	responsePacket, _ := nex.NewPacketV1(client, nil)

	responsePacket.SetVersion(1)
	responsePacket.SetSource(0xA1)
	responsePacket.SetDestination(0xAF)
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	globals.SecureServer.Send(responsePacket)

	return 0
}

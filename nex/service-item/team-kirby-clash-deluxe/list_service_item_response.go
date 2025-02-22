package nex_service_item_team_kirby_clash_deluxe

import (
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	service_item_team_kirby_clash_deluxe "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/team-kirby-clash-deluxe"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/team-kirby-clash-deluxe/types"
)

func ListServiceItemResponse(err error, packet nex.PacketInterface, callID uint32, requestID types.UInt32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.ServiceItem.InvalidArgument, err.Error())
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	// * Stubbed
	catalog := service_item_team_kirby_clash_deluxe_types.NewServiceItemCatalog()
	catalog.Balance = service_item_team_kirby_clash_deluxe_types.NewServiceItemAmount()

	nullableCatalog := make([]service_item_team_kirby_clash_deluxe_types.ServiceItemCatalog, 1)
	nullableCatalog[0] = catalog

	listServiceItemResponse := service_item_team_kirby_clash_deluxe_types.NewServiceItemListServiceItemResponse()
	listServiceItemResponse.HTTPStatus = 200
	listServiceItemResponse.CorrelationID = "8c6d0df0-e506-4f32-b730-1ccfe2476d7f,28b43a34-4709-49f9-b4f3-7ff76e8f79ab"
	listServiceItemResponse.NullableCatalog = nullableCatalog

	// * Stubbed
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	listServiceItemResponse.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = service_item_team_kirby_clash_deluxe.ProtocolID
	rmcResponse.MethodID = service_item_team_kirby_clash_deluxe.MethodListServiceItemResponse
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

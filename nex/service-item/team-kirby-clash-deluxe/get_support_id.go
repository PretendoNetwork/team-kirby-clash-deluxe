package nex_service_item_team_kirby_clash_deluxe

import (
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	service_item_team_kirby_clash_deluxe "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/team-kirby-clash-deluxe"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/team-kirby-clash-deluxe/types"
)

func GetSupportID(err error, packet nex.PacketInterface, callID uint32, getSupportIDParam service_item_team_kirby_clash_deluxe_types.ServiceItemGetSupportIDParam) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.ServiceItem.InvalidArgument, err.Error())
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	// * Stubbed
	var supportID types.String = "1"

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	supportID.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = service_item_team_kirby_clash_deluxe.ProtocolID
	rmcResponse.MethodID = service_item_team_kirby_clash_deluxe.MethodGetSupportID
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

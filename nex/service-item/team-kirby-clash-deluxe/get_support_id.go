package nex_service_item_team_kirby_clash_deluxe

import (
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"

	"github.com/PretendoNetwork/nex-go"
	service_item_team_kirby_clash_deluxe "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

func GetSupportID(err error, client *nex.Client, callID uint32, getSupportIDParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetSupportIDParam) uint32 {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nex.Errors.ServiceItem.InvalidArgument
	}

	supportID := "1"

	// * Stubbed
	rmcResponseStream := nex.NewStreamOut(globals.SecureServer)

	rmcResponseStream.WriteString(supportID)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCResponse(service_item_team_kirby_clash_deluxe.ProtocolID, callID)
	rmcResponse.SetSuccess(service_item_team_kirby_clash_deluxe.MethodGetSupportID, rmcResponseBody)

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

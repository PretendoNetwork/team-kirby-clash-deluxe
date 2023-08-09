package nex_service_item_team_kirby_clash_deluxe

import (
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"

	"github.com/PretendoNetwork/nex-go"
	// service_item_team_kirby_clash_deluxe "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
	// notifications "github.com/PretendoNetwork/nex-protocols-go/notifications"
	// notifications_types "github.com/PretendoNetwork/nex-protocols-go/notifications/types"
)

func ListServiceItemRequest(err error, client *nex.Client, callID uint32, listServiceItemParam *service_item_team_kirby_clash_deluxe_types.ServiceItemListServiceItemParam) uint32 {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nex.Errors.ServiceItem.InvalidArgument
	}

	return nex.Errors.ServiceItem.Unknown

	/* TODO - WHY THIS DOES NOT WORK?
	// * Stubbed
	var requestID uint32 = 1

	rmcResponseStream := nex.NewStreamOut(globals.SecureServer)

	rmcResponseStream.WriteUInt32LE(requestID)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCResponse(service_item_team_kirby_clash_deluxe.ProtocolID, callID)
	rmcResponse.SetSuccess(service_item_team_kirby_clash_deluxe.MethodListServiceItemRequest, rmcResponseBody)

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

	rmcMessage := nex.NewRMCRequest()
	rmcMessage.SetProtocolID(notifications.ProtocolID)
	rmcMessage.SetCallID(0xffff0000 + callID)
	rmcMessage.SetMethodID(notifications.MethodProcessNotificationEvent)

	category := notifications.NotificationCategories.ServiceItemRequestCompleted
	subtype := notifications.NotificationSubTypes.ServiceItemRequestCompleted.None

	oEvent := notifications_types.NewNotificationEvent()
	oEvent.PIDSource = client.PID()
	oEvent.Type = notifications.BuildNotificationType(category, subtype)
	oEvent.Param1 = requestID
	oEvent.Param2 = 1 // * Unknown

	stream := nex.NewStreamOut(globals.SecureServer)
	oEventBytes := oEvent.Bytes(stream)
	rmcMessage.SetParameters(oEventBytes)

	rmcMessageBytes := rmcMessage.Bytes()

	var messagePacket nex.PacketInterface

	messagePacket, _ = nex.NewPacketV1(client, nil)

	messagePacket.SetVersion(1)
	messagePacket.SetSource(0xA1)
	messagePacket.SetDestination(0xAF)
	messagePacket.SetType(nex.DataPacket)
	messagePacket.SetPayload(rmcMessageBytes)

	messagePacket.AddFlag(nex.FlagNeedsAck)
	messagePacket.AddFlag(nex.FlagReliable)

	globals.SecureServer.Send(messagePacket)

	return 0
	*/
}

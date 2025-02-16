package nex_service_item_team_kirby_clash_deluxe

import (
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	notifications "github.com/PretendoNetwork/nex-protocols-go/v2/notifications"
	notifications_types "github.com/PretendoNetwork/nex-protocols-go/v2/notifications/types"
	service_item_team_kirby_clash_deluxe "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/team-kirby-clash-deluxe"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/team-kirby-clash-deluxe/types"
)

func ListServiceItemRequest(err error, packet nex.PacketInterface, callID uint32, listServiceItemParam service_item_team_kirby_clash_deluxe_types.ServiceItemListServiceItemParam) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.ServiceItem.InvalidArgument, err.Error())
	}

	// return nil, nex.NewError(nex.ResultCodes.ServiceItem.Unknown, "Stubbed")

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	// * Stubbed
	var requestID uint32 = 1

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	rmcResponseStream.WriteUInt32LE(requestID)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = service_item_team_kirby_clash_deluxe.ProtocolID
	rmcResponse.MethodID = service_item_team_kirby_clash_deluxe.MethodListServiceItemRequest
	rmcResponse.CallID = callID

	category := notifications.NotificationCategories.ServiceItemRequestCompleted
	subtype := notifications.NotificationSubTypes.ServiceItemRequestCompleted.None

	oEvent := notifications_types.NewNotificationEvent()
	oEvent.PIDSource = connection.PID()
	oEvent.Type = types.NewUInt32(notifications.BuildNotificationType(category, subtype))
	oEvent.Param1 = types.NewUInt32(requestID)
	oEvent.Param2 = 1 // * Unknown

	go common_globals.SendNotificationEvent(globals.SecureEndpoint, oEvent, []uint64{uint64(connection.PID())})

	return rmcResponse, nil
}

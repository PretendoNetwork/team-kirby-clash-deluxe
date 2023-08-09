package nex_datastore

import (
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/database"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"

	"github.com/PretendoNetwork/nex-go"
	datastore "github.com/PretendoNetwork/nex-protocols-go/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

func ChangeMeta(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreChangeMetaParam) uint32 {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nex.Errors.DataStore.InvalidArgument
	}

	if validateErr := database.ValidateMetaBinaryByOwnerPID(uint32(param.DataID), client.PID()); validateErr != 0 {
		return validateErr
	}

	err = database.UpdateMetaBinaryByDataStoreChangeMetaParam(param)
	if err != nil {
		globals.Logger.Critical(err.Error())
		return nex.Errors.DataStore.Unknown
	}

	rmcResponse := nex.NewRMCResponse(datastore.ProtocolID, callID)
	rmcResponse.SetSuccess(datastore.MethodChangeMeta, nil)

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

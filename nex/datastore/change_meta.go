package nex_datastore

import (
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/database"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"

	"github.com/PretendoNetwork/nex-go/v2"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func ChangeMeta(err error, packet nex.PacketInterface, callID uint32, param datastore_types.DataStoreChangeMetaParam) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.InvalidArgument, err.Error())
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	if nexError := database.ValidateMetaBinaryByOwnerPID(uint32(param.DataID), uint32(connection.PID())); nexError != nil {
		return nil, nexError
	}

	err = database.UpdateMetaBinaryByDataStoreChangeMetaParam(param)
	if err != nil {
		globals.Logger.Critical(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.Unknown, err.Error())
	}

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodChangeMeta
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

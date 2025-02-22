package nex_datastore

import (
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/database"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"

	"github.com/PretendoNetwork/nex-go/v2"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func PostMetaBinary(err error, packet nex.PacketInterface, callID uint32, param datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.InvalidArgument, err.Error())
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	metaBinary := database.GetMetaBinaryByOwnerPID(uint32(connection.PID()))

	// * Meta binary already exists
	if metaBinary.DataID != 0 && param.PersistenceInitParam.DeleteLastObject {
		// * Delete existing object before uploading new one
		err = database.DeleteMetaBinaryByDataID(metaBinary.DataID)
		if err != nil {
			globals.Logger.Critical(err.Error())
			return nil, nex.NewError(nex.ResultCodes.DataStore.Unknown, err.Error())
		}
	}

	dataID, err := database.InsertMetaBinaryByDataStorePreparePostParamWithOwnerPID(param, uint32(connection.PID()))
	if err != nil {
		globals.Logger.Critical(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.Unknown, err.Error())
	}

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	rmcResponseStream.WriteUInt64LE(uint64(dataID))

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodPostMetaBinary
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

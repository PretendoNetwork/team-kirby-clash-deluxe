package nex_datastore

import (
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/database"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/types"

	"github.com/PretendoNetwork/nex-go/v2"
	nex_types "github.com/PretendoNetwork/nex-go/v2/types"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func SearchObjectLight(err error, packet nex.PacketInterface, callID uint32, param datastore_types.DataStoreSearchParam) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.InvalidArgument, err.Error())
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	metaBinaries := make([]*types.MetaBinary, 0)

	if param.SearchTarget == 10 { // * Search for meta binary of this client
		metaBinary := database.GetMetaBinaryByOwnerPID(uint32(connection.PID()))
		if metaBinary.DataID != 0 {
			metaBinaries = append(metaBinaries, metaBinary)
		}
	} else if len(param.DataTypes) > 0 { // * The data type is given inside the DataTypes param
		metaBinaries = database.GetMetaBinariesByDataStoreSearchParam(param)
	}

	pSearchResult := datastore_types.NewDataStoreSearchResult()

	pSearchResult.TotalCount = nex_types.NewUInt32(uint32(len(metaBinaries)))
	pSearchResult.Result = make([]datastore_types.DataStoreMetaInfo, 0, len(metaBinaries))

	if param.TotalCountEnabled == false {
		pSearchResult.TotalCountType = 3 // * Disabled
		pSearchResult.TotalCount = 0
	}

	for i := 0; i < len(metaBinaries); i++ {
		metaBinary := metaBinaries[i]
		result := datastore_types.NewDataStoreMetaInfo()

		result.DataID = nex_types.NewUInt64(uint64(metaBinary.DataID))
		result.OwnerID = nex_types.NewPID(uint64(metaBinary.OwnerPID))
		result.Size = 0
		result.Name = nex_types.NewString(metaBinary.Name)
		result.DataType = nex_types.NewUInt16(metaBinary.DataType)
		result.MetaBinary = nex_types.NewQBuffer(metaBinary.Buffer)
		result.Permission = datastore_types.NewDataStorePermission()
		result.Permission.Permission = nex_types.NewUInt8(metaBinary.Permission)
		result.Permission.RecipientIDs = make([]nex_types.PID, 0)
		result.DelPermission = datastore_types.NewDataStorePermission()
		result.DelPermission.Permission = nex_types.NewUInt8(metaBinary.DeletePermission)
		result.DelPermission.RecipientIDs = make([]nex_types.PID, 0)
		result.CreatedTime = metaBinary.CreationTime
		result.UpdatedTime = metaBinary.UpdatedTime
		result.Period = nex_types.NewUInt16(metaBinary.Period)
		result.Status = 0      // TODO - Figure this out
		result.ReferredCnt = 0 // TODO - Figure this out
		result.ReferDataID = 0 // TODO - Figure this out
		result.Flag = nex_types.NewUInt32(metaBinary.Flag)
		result.ReferredTime = metaBinary.ReferredTime
		result.ExpireTime = metaBinary.ExpireTime

		tags := make([]nex_types.String, len(metaBinaries[i].Tags))
		for j, tag := range metaBinaries[i].Tags {
			tags[j] = nex_types.NewString(tag)
		}
		result.Tags = tags

		result.Ratings = make([]datastore_types.DataStoreRatingInfoWithSlot, 0)

		pSearchResult.Result = append(pSearchResult.Result, result)
	}

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	pSearchResult.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodSearchObjectLight
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

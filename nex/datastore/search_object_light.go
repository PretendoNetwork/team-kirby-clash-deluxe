package nex_datastore

import (
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/database"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/types"

	"github.com/PretendoNetwork/nex-go"
	datastore "github.com/PretendoNetwork/nex-protocols-go/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
)

func SearchObjectLight(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam) uint32 {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nex.Errors.DataStore.InvalidArgument
	}

	metaBinaries := make([]*types.MetaBinary, 0)

	if param.SearchTarget == 10 { // * Search for meta binary of this client
		metaBinary := database.GetMetaBinaryByOwnerPID(client.PID())
		if metaBinary.DataID != 0 {
			metaBinaries = append(metaBinaries, metaBinary)
		}
	} else if len(param.DataTypes) > 0 { // * The data type is given inside the DataTypes param
		metaBinaries = database.GetMetaBinariesByDataStoreSearchParam(param)
	}

	pSearchResult := datastore_types.NewDataStoreSearchResult()

	pSearchResult.TotalCount = uint32(len(metaBinaries))
	pSearchResult.Result = make([]*datastore_types.DataStoreMetaInfo, 0, len(metaBinaries))

	if param.TotalCountEnabled == false {
		pSearchResult.TotalCountType = 3 // * Disabled
		pSearchResult.TotalCount = 0
	}

	for i := 0; i < len(metaBinaries); i++ {
		metaBinary := metaBinaries[i]
		result := datastore_types.NewDataStoreMetaInfo()

		result.DataID = uint64(metaBinary.DataID)
		result.OwnerID = metaBinary.OwnerPID
		result.Size = 0
		result.Name = metaBinary.Name
		result.DataType = metaBinary.DataType
		result.MetaBinary = metaBinary.Buffer
		result.Permission = datastore_types.NewDataStorePermission()
		result.Permission.Permission = metaBinary.Permission
		result.Permission.RecipientIDs = make([]uint32, 0)
		result.DelPermission = datastore_types.NewDataStorePermission()
		result.DelPermission.Permission = metaBinary.DeletePermission
		result.DelPermission.RecipientIDs = make([]uint32, 0)
		result.CreatedTime = metaBinary.CreationTime
		result.UpdatedTime = metaBinary.UpdatedTime
		result.Period = metaBinary.Period
		result.Status = 0      // TODO - Figure this out
		result.ReferredCnt = 0 // TODO - Figure this out
		result.ReferDataID = 0 // TODO - Figure this out
		result.Flag = metaBinary.Flag
		result.ReferredTime = metaBinary.ReferredTime
		result.ExpireTime = metaBinary.ExpireTime
		result.Tags = metaBinary.Tags
		result.Ratings = make([]*datastore_types.DataStoreRatingInfoWithSlot, 0)

		pSearchResult.Result = append(pSearchResult.Result, result)
	}

	rmcResponseStream := nex.NewStreamOut(globals.SecureServer)

	rmcResponseStream.WriteStructure(pSearchResult)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCResponse(datastore.ProtocolID, callID)
	rmcResponse.SetSuccess(datastore.MethodSearchObjectLight, rmcResponseBody)

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

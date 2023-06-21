package nex_datastore

import (
	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe-secure/database"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe-secure/globals"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe-secure/types"
)

func SearchObjectLight(err error, client *nex.Client, callID uint32, param *datastore.DataStoreSearchParam) {
	metaBinaries := make([]*types.MetaBinary, 0)

	if param.SearchTarget == 10 { // * Search for meta binary of this client
		metaBinary := database.GetMetaBinaryByOwnerPID(client.PID())
		if metaBinary.DataID != 0 {
			metaBinaries = append(metaBinaries, metaBinary)
		}
	} else if len(param.DataTypes) > 0 { // * The data type is given inside the DataTypes param
		metaBinaries = database.GetMetaBinariesByDataStoreSearchParam(param)
	}

	pSearchResult := datastore.NewDataStoreSearchResult()

	pSearchResult.TotalCount = uint32(len(metaBinaries))
	pSearchResult.Result = make([]*datastore.DataStoreMetaInfo, 0, len(metaBinaries))

	if param.TotalCountEnabled == false {
		pSearchResult.TotalCountType = 3 // Disabled
		pSearchResult.TotalCount = 0
	}

	for i := 0; i < len(metaBinaries); i++ {
		metaBinary := metaBinaries[i]
		result := datastore.NewDataStoreMetaInfo()

		result.DataID = uint64(metaBinary.DataID)
		result.OwnerID = metaBinary.OwnerPID
		result.Size = 0
		result.Name = metaBinary.Name
		result.DataType = metaBinary.DataType
		result.MetaBinary = metaBinary.Buffer
		result.Permission = datastore.NewDataStorePermission()
		result.Permission.Permission = metaBinary.Permission
		result.Permission.RecipientIds = make([]uint32, 0)
		result.DelPermission = datastore.NewDataStorePermission()
		result.DelPermission.Permission = metaBinary.DeletePermission
		result.DelPermission.RecipientIds = make([]uint32, 0)
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
		result.Ratings = make([]*datastore.DataStoreRatingInfoWithSlot, 0)

		pSearchResult.Result = append(pSearchResult.Result, result)
	}

	rmcResponseStream := nex.NewStreamOut(globals.NEXServer)

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

	globals.NEXServer.Send(responsePacket)
}

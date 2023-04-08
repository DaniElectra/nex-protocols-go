package datastore_super_mario_maker

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetObjectInfos sets the GetObjectInfos handler function
func (protocol *DataStoreSuperMarioMakerProtocol) GetObjectInfos(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64)) {
	protocol.GetObjectInfosHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleGetObjectInfos(packet nex.PacketInterface) {
	if protocol.GetObjectInfosHandler == nil {
		globals.Logger.Warning("DataStoreSMM::GetObjectInfos not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDs := parametersStream.ReadListUInt64LE()

	go protocol.GetObjectInfosHandler(nil, client, callID, dataIDs)
}
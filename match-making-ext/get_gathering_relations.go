package match_making_ext

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetGatheringRelations sets the GetGatheringRelations handler function
func (protocol *MatchMakingExtProtocol) GetGatheringRelations(handler func(err error, client *nex.Client, callID uint32, id uint32, descr string)) {
	protocol.GetGatheringRelationsHandler = handler
}

func (protocol *MatchMakingExtProtocol) HandleGetGatheringRelations(packet nex.PacketInterface) {
	if protocol.GetGatheringRelationsHandler == nil {
		globals.Logger.Warning("MatchMakingExt::GetGatheringRelations not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id := parametersStream.ReadUInt32LE()

	var err error
	var descr string
	descr, err = parametersStream.ReadString()
	if err != nil {
		go protocol.GetGatheringRelationsHandler(err, client, callID, 0, "")
	}

	go protocol.GetGatheringRelationsHandler(nil, client, callID, id, descr)
}
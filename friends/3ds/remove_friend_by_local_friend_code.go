package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveFriendByLocalFriendCode sets the RemoveFriendByLocalFriendCode handler function
func (protocol *Friends3DSProtocol) RemoveFriendByLocalFriendCode(handler func(err error, client *nex.Client, callID uint32, lfc uint64)) {
	protocol.RemoveFriendByLocalFriendCodeHandler = handler
}

func (protocol *Friends3DSProtocol) HandleRemoveFriendByLocalFriendCode(packet nex.PacketInterface) {
	if protocol.RemoveFriendByLocalFriendCodeHandler == nil {
		globals.Logger.Warning("Friends3DS::RemoveFriendByLocalFriendCode not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc := parametersStream.ReadUInt64LE()

	go protocol.RemoveFriendByLocalFriendCodeHandler(nil, client, callID, lfc)
}
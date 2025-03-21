// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetFriendPicture(packet nex.PacketInterface) {
	if protocol.GetFriendPicture == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends3DS::GetFriendPicture not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var unknown types.List[types.UInt32]

	err := unknown.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetFriendPicture(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, unknown)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetFriendPicture(nil, packet, callID, unknown)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

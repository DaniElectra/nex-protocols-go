// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteCommonData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.DeleteCommonData == nil {
		globals.Logger.Warning("Ranking2::DeleteCommonData not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nexUniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		_, errorCode = protocol.DeleteCommonData(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DeleteCommonData(nil, packet, callID, nexUniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}

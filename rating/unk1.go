// Package protocol implements the Rating protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// TODO - Find name if possible
func (protocol *Protocol) handleUnk1(packet nex.PacketInterface) {
	if protocol.Unk1 == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Rating::Unk1 not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.Unk1(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

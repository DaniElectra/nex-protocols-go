// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleInitiateProbe(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.InitiateProbe == nil {
		globals.Logger.Warning("NATTraversal::InitiateProbe not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	urlStationToProbe, err := parametersStream.ReadStationURL()
	if err != nil {
		_, errorCode = protocol.InitiateProbe(fmt.Errorf("Failed to read urlStationToProbe from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.InitiateProbe(nil, packet, callID, urlStationToProbe)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}

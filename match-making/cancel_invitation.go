// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleCancelInvitation(packet nex.PacketInterface) {
	if protocol.CancelInvitation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::CancelInvitation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var idGathering types.UInt32
	var lstPrincipals types.List[types.PID]
	var strMessage types.String

	var err error

	err = idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CancelInvitation(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, idGathering, lstPrincipals, strMessage)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = lstPrincipals.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CancelInvitation(fmt.Errorf("Failed to read lstPrincipals from parameters. %s", err.Error()), packet, callID, idGathering, lstPrincipals, strMessage)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CancelInvitation(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, idGathering, lstPrincipals, strMessage)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.CancelInvitation(nil, packet, callID, idGathering, lstPrincipals, strMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

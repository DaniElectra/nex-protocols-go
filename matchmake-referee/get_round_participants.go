// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetRoundParticipants(packet nex.PacketInterface) {
	if protocol.GetRoundParticipants == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeReferee::GetRoundParticipants not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var roundID types.UInt64

	err := roundID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRoundParticipants(fmt.Errorf("Failed to read roundID from parameters. %s", err.Error()), packet, callID, roundID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRoundParticipants(nil, packet, callID, roundID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

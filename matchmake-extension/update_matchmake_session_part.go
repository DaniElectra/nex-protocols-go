// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
)

func (protocol *Protocol) handleUpdateMatchmakeSessionPart(packet nex.PacketInterface) {
	if protocol.UpdateMatchmakeSessionPart == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::UpdateMatchmakeSessionPart not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	updateMatchmakeSessionParam := match_making_types.NewUpdateMatchmakeSessionParam()

	err := updateMatchmakeSessionParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateMatchmakeSessionPart(fmt.Errorf("Failed to read updateMatchmakeSessionParam from parameters. %s", err.Error()), packet, callID, updateMatchmakeSessionParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateMatchmakeSessionPart(nil, packet, callID, updateMatchmakeSessionParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

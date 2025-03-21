// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
)

func (protocol *Protocol) handleBrowseMatchmakeSessionWithHostURLs(packet nex.PacketInterface) {
	if protocol.BrowseMatchmakeSessionWithHostURLs == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::BrowseMatchmakeSessionWithHostURLs not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	searchCriteria := match_making_types.NewMatchmakeSessionSearchCriteria()
	var resultRange types.ResultRange

	var err error

	err = searchCriteria.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.BrowseMatchmakeSessionWithHostURLs(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), packet, callID, searchCriteria, resultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.BrowseMatchmakeSessionWithHostURLs(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, searchCriteria, resultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.BrowseMatchmakeSessionWithHostURLs(nil, packet, callID, searchCriteria, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleFindMatchmakeSessionBySingleGatheringID(packet nex.PacketInterface) {
	if protocol.FindMatchmakeSessionBySingleGatheringID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::FindMatchmakeSessionBySingleGatheringID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var gid types.UInt32

	err := gid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindMatchmakeSessionBySingleGatheringID(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, gid)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.FindMatchmakeSessionBySingleGatheringID(nil, packet, callID, gid)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

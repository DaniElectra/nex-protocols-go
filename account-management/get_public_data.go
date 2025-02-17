// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetPublicData(packet nex.PacketInterface) {
	if protocol.GetPublicData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::GetPublicData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var idPrincipal types.PID

	err := idPrincipal.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetPublicData(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), packet, callID, idPrincipal)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetPublicData(nil, packet, callID, idPrincipal)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

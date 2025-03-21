// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/v2/messaging/types"
)

func (protocol *Protocol) handleRetrieveAllMessagesWithinRange(packet nex.PacketInterface) {
	if protocol.RetrieveAllMessagesWithinRange == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Messaging::RetrieveAllMessagesWithinRange not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	recipient := messaging_types.NewMessageRecipient()
	var resultRange types.ResultRange

	var err error

	err = recipient.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RetrieveAllMessagesWithinRange(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), packet, callID, recipient, resultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RetrieveAllMessagesWithinRange(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, recipient, resultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RetrieveAllMessagesWithinRange(nil, packet, callID, recipient, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

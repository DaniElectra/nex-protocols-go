// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetMetasMultipleParam(packet nex.PacketInterface) {
	if protocol.GetMetasMultipleParam == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::GetMetasMultipleParam not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params types.List[datastore_types.DataStoreGetMetaParam]

	err := params.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetMetasMultipleParam(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, params)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetMetasMultipleParam(nil, packet, callID, params)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

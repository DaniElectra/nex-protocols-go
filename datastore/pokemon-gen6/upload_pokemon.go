// Package protocol implements the DataStorePokemonGen6 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	datastore_pokemon_gen6_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/pokemon-gen6/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUploadPokemon(packet nex.PacketInterface) {
	if protocol.UploadPokemon == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStorePokemonGen6::UploadPokemon not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	param := datastore_pokemon_gen6_types.NewGlobalTradeStationUploadPokemonParam()

	err := param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UploadPokemon(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, param)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UploadPokemon(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

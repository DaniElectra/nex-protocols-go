// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/wii-sports-club/types"
)

func (protocol *Protocol) handleGetPrepurchaseInfoRequest(packet nex.PacketInterface) {
	if protocol.GetPrepurchaseInfoRequest == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ServiceItemWiiSportsClub::GetPrepurchaseInfoRequest not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	getPrepurchaseInfoParam := service_item_wii_sports_club_types.NewServiceItemGetPrepurchaseInfoParam()

	err := getPrepurchaseInfoParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetPrepurchaseInfoRequest(fmt.Errorf("Failed to read getPrepurchaseInfoParam from parameters. %s", err.Error()), packet, callID, getPrepurchaseInfoParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetPrepurchaseInfoRequest(nil, packet, callID, getPrepurchaseInfoParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

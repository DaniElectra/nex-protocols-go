// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// GetLawMessageRequest sets the GetLawMessageRequest handler function
func (protocol *Protocol) GetLawMessageRequest(handler func(err error, client *nex.Client, callID uint32, getLawMessageParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetLawMessageParam)) {
	protocol.getLawMessageRequestHandler = handler
}

func (protocol *Protocol) handleGetLawMessageRequest(packet nex.PacketInterface) {
	if protocol.getLawMessageRequestHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetLawMessageRequest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getLawMessageParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemGetLawMessageParam())
	if err != nil {
		go protocol.getLawMessageRequestHandler(fmt.Errorf("Failed to read getLawMessageParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getLawMessageRequestHandler(nil, client, callID, getLawMessageParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemGetLawMessageParam))
}
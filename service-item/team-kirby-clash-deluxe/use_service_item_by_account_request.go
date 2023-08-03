// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// UseServiceItemByAccountRequest sets the UseServiceItemByAccountRequest handler function
func (protocol *Protocol) UseServiceItemByAccountRequest(handler func(err error, client *nex.Client, callID uint32, useServiceItemByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemUseServiceItemByAccountParam)) {
	protocol.useServiceItemByAccountRequestHandler = handler
}

func (protocol *Protocol) handleUseServiceItemByAccountRequest(packet nex.PacketInterface) {
	if protocol.useServiceItemByAccountRequestHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::UseServiceItemByAccountRequest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	useServiceItemByAccountParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemUseServiceItemByAccountParam())
	if err != nil {
		go protocol.useServiceItemByAccountRequestHandler(fmt.Errorf("Failed to read useServiceItemByAccountParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.useServiceItemByAccountRequestHandler(nil, client, callID, useServiceItemByAccountParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemUseServiceItemByAccountParam))
}
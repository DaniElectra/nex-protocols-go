// Package friends implements the Friends QRV protocol
package friends

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateDetails sets the UpdateDetails handler function
func (protocol *FriendsProtocol) UpdateDetails(handler func(err error, client *nex.Client, callID uint32, uiPlayer uint32, uiDetails uint32)) {
	protocol.updateDetailsHandler = handler
}

func (protocol *FriendsProtocol) handleUpdateDetails(packet nex.PacketInterface) {
	if protocol.updateDetailsHandler == nil {
		globals.Logger.Warning("Friends::UpdateDetails not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiPlayer, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.updateDetailsHandler(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	uiDetails, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.updateDetailsHandler(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.updateDetailsHandler(nil, client, callID, uiPlayer, uiDetails)
}
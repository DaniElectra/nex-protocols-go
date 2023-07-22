// Package friends implements the Friends QRV protocol
package friends

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriendByName sets the AddFriendByName handler function
func (protocol *FriendsProtocol) AddFriendByName(handler func(err error, client *nex.Client, callID uint32, strPlayerName string, uiDetails uint32, strMessage string)) {
	protocol.addFriendByNameHandler = handler
}

func (protocol *FriendsProtocol) handleAddFriendByName(packet nex.PacketInterface) {
	if protocol.addFriendByNameHandler == nil {
		globals.Logger.Warning("Friends::AddFriendByName not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strPlayerName, err := parametersStream.ReadString()
	if err != nil {
		go protocol.addFriendByNameHandler(fmt.Errorf("Failed to read strPlayerName from parameters. %s", err.Error()), client, callID, "", 0, "")
		return
	}

	uiDetails, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.addFriendByNameHandler(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), client, callID, "", 0, "")
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.addFriendByNameHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, "", 0, "")
		return
	}

	go protocol.addFriendByNameHandler(nil, client, callID, strPlayerName, uiDetails, strMessage)
}
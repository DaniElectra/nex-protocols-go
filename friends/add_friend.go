// Package friends implements the Friends QRV protocol
package friends

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriend sets the AddFriend handler function
func (protocol *FriendsProtocol) AddFriend(handler func(err error, client *nex.Client, callID uint32, uiPlayer uint32, uiDetails uint32, strMessage string)) {
	protocol.addFriendHandler = handler
}

func (protocol *FriendsProtocol) handleAddFriend(packet nex.PacketInterface) {
	if protocol.addFriendHandler == nil {
		globals.Logger.Warning("Friends::AddFriend not implemented")
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
		go protocol.addFriendHandler(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), client, callID, 0, 0, "")
		return
	}

	uiDetails, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.addFriendHandler(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), client, callID, 0, 0, "")
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.addFriendHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, 0, "")
		return
	}

	go protocol.addFriendHandler(nil, client, callID, uiPlayer, uiDetails, strMessage)
}
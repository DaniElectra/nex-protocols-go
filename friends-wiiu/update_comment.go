// Package friends_wiiu implements the Friends WiiU NEX protocol
package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateComment sets the UpdateComment handler function
func (protocol *FriendsWiiUProtocol) UpdateComment(handler func(err error, client *nex.Client, callID uint32, comment *friends_wiiu_types.Comment)) {
	protocol.UpdateCommentHandler = handler
}

func (protocol *FriendsWiiUProtocol) handleUpdateComment(packet nex.PacketInterface) {
	if protocol.UpdateCommentHandler == nil {
		globals.Logger.Warning("FriendsWiiU::UpdateComment not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	comment, err := parametersStream.ReadStructure(friends_wiiu_types.NewComment())
	if err != nil {
		go protocol.UpdateCommentHandler(fmt.Errorf("Failed to read comment from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.UpdateCommentHandler(nil, client, callID, comment.(*friends_wiiu_types.Comment))
}
// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendUserProfiles sets the GetFriendUserProfiles handler function
func (protocol *Protocol) GetFriendUserProfiles(handler func(err error, client *nex.Client, callID uint32, pids []uint64)) {
	protocol.getFriendUserProfilesHandler = handler
}

func (protocol *Protocol) handleGetFriendUserProfiles(packet nex.PacketInterface) {
	if protocol.getFriendUserProfilesHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMonsterHunterXX::GetFriendUserProfiles not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.getFriendUserProfilesHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getFriendUserProfilesHandler(nil, client, callID, pids)
}
// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindBySQLQuery sets the FindBySQLQuery handler function
func (protocol *MatchMakingProtocol) FindBySQLQuery(handler func(err error, client *nex.Client, callID uint32, strQuery string, resultRange *nex.ResultRange)) {
	protocol.findBySQLQueryHandler = handler
}

func (protocol *MatchMakingProtocol) handleFindBySQLQuery(packet nex.PacketInterface) {
	if protocol.findBySQLQueryHandler == nil {
		globals.Logger.Warning("MatchMaking::FindBySQLQuery not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strQuery, err := parametersStream.ReadString()
	if err != nil {
		go protocol.findBySQLQueryHandler(fmt.Errorf("Failed to read strQuery from parameters. %s", err.Error()), client, callID, "", nil)
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findBySQLQueryHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, "", nil)
	}

	go protocol.findBySQLQueryHandler(nil, client, callID, strQuery, resultRange.(*nex.ResultRange))
}
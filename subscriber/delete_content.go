// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteContent sets the DeleteContent handler function
func (protocol *Protocol) DeleteContent(handler func(err error, client *nex.Client, callID uint32, unknown1 []string, unknown2 uint64)) {
	protocol.deleteContentHandler = handler
}

func (protocol *Protocol) handleDeleteContent(packet nex.PacketInterface) {
	if protocol.deleteContentHandler == nil {
		globals.Logger.Warning("Subscriber::DeleteContent not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown1, err := parametersStream.ReadListString()
	if err != nil {
		go protocol.deleteContentHandler(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	unknown2, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.deleteContentHandler(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	go protocol.deleteContentHandler(nil, client, callID, unknown1, unknown2)
}
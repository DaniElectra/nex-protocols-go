// Package protocol implements the StorageManager protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the StorageManager protocol
	ProtocolID = 0x6E

	// MethodAcquireCardID is the method ID for the method AcquireCardID
	MethodAcquireCardID = 0x4

	// MethodActivateWithCardID is the method ID for the method ActivateWithCardID
	MethodActivateWithCardID = 0x5
)

// Protocol stores all the RMC method handlers for the StorageManager protocol and listens for requests
type Protocol struct {
	Server             nex.ServerInterface
	AcquireCardID      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	ActivateWithCardID func(err error, packet nex.PacketInterface, callID uint32, unknown uint8, cardID uint64) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			switch message.MethodID {
			case MethodAcquireCardID:
				protocol.handleAcquireCardID(packet)
			case MethodActivateWithCardID:
				protocol.handleActivateWithCardID(packet)
			default:
				globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported StorageManager method ID: %#v\n", message.MethodID)
			}
		}
	})
}

// NewProtocol returns a new StorageManager protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}

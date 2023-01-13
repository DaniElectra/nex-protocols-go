package nexproto

import (
	nex "github.com/PretendoNetwork/nex-go"
	hpp "github.com/PretendoNetwork/hpp-go"
)

func respondNotImplemented(input interface{}, protocolID uint8) {
	switch packet := input.(type) {
		case nex.PacketInterface:
			client := packet.Sender()
			request := packet.RMCRequest()

			rmcResponse := nex.NewRMCResponse(protocolID, request.CallID())
			rmcResponse.SetError(nex.Errors.Core.NotImplemented)

			rmcResponseBytes := rmcResponse.Bytes()

			var responsePacket nex.PacketInterface
			if packet.Version() == 1 {
				responsePacket, _ = nex.NewPacketV1(client, nil)
			} else {
				responsePacket, _ = nex.NewPacketV0(client, nil)
			}

			responsePacket.SetVersion(packet.Version())
			responsePacket.SetSource(packet.Destination())
			responsePacket.SetDestination(packet.Source())
			responsePacket.SetType(nex.DataPacket)
			responsePacket.SetPayload(rmcResponseBytes)

			responsePacket.AddFlag(nex.FlagNeedsAck)
			responsePacket.AddFlag(nex.FlagReliable)

			client.Server().Send(responsePacket)
		case *hpp.HppRequest:
			request := packet.RMCRequest()

			rmcResponse := hpp.NewRMCResponse(request.CallID())
			rmcResponse.SetError(hpp.Errors.Core.NotImplemented)

			responsePacket := hpp.NewHppResponse(rmcResponse, packet.PID())

			packet.Server().Send(responsePacket)
	}
}

func respondNotImplementedCustom(packet nex.PacketInterface, customID uint16) {
	client := packet.Sender()
	request := packet.RMCRequest()

	rmcResponse := nex.NewRMCResponse(0x7f, request.CallID())
	rmcResponse.SetCustomID(customID)
	rmcResponse.SetError(nex.Errors.Core.NotImplemented)

	rmcResponseBytes := rmcResponse.Bytes()

	var responsePacket nex.PacketInterface
	if packet.Version() == 1 {
		responsePacket, _ = nex.NewPacketV1(client, nil)
	} else {
		responsePacket, _ = nex.NewPacketV0(client, nil)
	}

	responsePacket.SetVersion(packet.Version())
	responsePacket.SetSource(packet.Destination())
	responsePacket.SetDestination(packet.Source())
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	client.Server().Send(responsePacket)
}

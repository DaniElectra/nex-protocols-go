package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateProfile sets the UpdateProfile handler function
func (protocol *Friends3DSProtocol) UpdateProfile(handler func(err error, client *nex.Client, callID uint32, profileData *MyProfile)) {
	protocol.UpdateProfileHandler = handler
}

func (protocol *Friends3DSProtocol) HandleUpdateProfile(packet nex.PacketInterface) {
	if protocol.UpdateProfileHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateProfile not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	profileData, err := parametersStream.ReadStructure(NewMyProfile())
	if err != nil {
		go protocol.UpdateProfileHandler(err, client, callID, nil)
		return
	}

	go protocol.UpdateProfileHandler(nil, client, callID, profileData.(*MyProfile))
}
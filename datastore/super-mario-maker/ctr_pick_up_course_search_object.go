package datastore_super_mario_maker

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CTRPickUpCourseSearchObject sets the CTRPickUpCourseSearchObject handler function
func (protocol *DataStoreSuperMarioMakerProtocol) CTRPickUpCourseSearchObject(handler func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *datastore.DataStoreSearchParam, extraData []string)) {
	protocol.CTRPickUpCourseSearchObjectHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleCTRPickUpCourseSearchObject(packet nex.PacketInterface) {
	if protocol.CTRPickUpCourseSearchObjectHandler == nil {
		globals.Logger.Warning("DataStoreSMM::CTRPickUpCourseSearchObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore.NewDataStoreSearchParam())
	if err != nil {
		go protocol.CTRPickUpCourseSearchObjectHandler(err, client, callID, nil, []string{})
		return
	}

	extraData := parametersStream.ReadListString()

	go protocol.CTRPickUpCourseSearchObjectHandler(nil, client, callID, param.(*datastore.DataStoreSearchParam), extraData)
}
// Package types implements all the types used by the MessageDelivery protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MessageRecipient is a type within the MessageDelivery protocol
type MessageRecipient struct {
	types.Structure
	IDRecipient     types.UInt32 // * NEX <4.0.0
	UIRecipientType types.UInt32
	PrincipalID     types.PID    // * NEX 4.0.0
	GatheringID     types.UInt32 // * NEX 4.0.0
}

// WriteTo writes the MessageRecipient to the given writable
func (mr MessageRecipient) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.Messaging

	contentWritable := writable.CopyNew()

	if !libraryVersion.GreaterOrEqual("4.0.0") {
		mr.IDRecipient.WriteTo(contentWritable)
	}

	mr.UIRecipientType.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("4.0.0") {
		mr.PrincipalID.WriteTo(contentWritable)
		mr.GatheringID.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	mr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MessageRecipient from the given readable
func (mr *MessageRecipient) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.Messaging

	var err error

	err = mr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient header. %s", err.Error())
	}

	if !libraryVersion.GreaterOrEqual("4.0.0") {
		err = mr.IDRecipient.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MessageRecipient.IDRecipient. %s", err.Error())
		}
	}

	err = mr.UIRecipientType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.UIRecipientType. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		err = mr.PrincipalID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MessageRecipient.PrincipalID. %s", err.Error())
		}

		err = mr.GatheringID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MessageRecipient.GatheringID. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of MessageRecipient
func (mr MessageRecipient) Copy() types.RVType {
	copied := NewMessageRecipient()

	copied.StructureVersion = mr.StructureVersion
	copied.IDRecipient = mr.IDRecipient.Copy().(types.UInt32)
	copied.UIRecipientType = mr.UIRecipientType.Copy().(types.UInt32)
	copied.PrincipalID = mr.PrincipalID.Copy().(types.PID)
	copied.GatheringID = mr.GatheringID.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given MessageRecipient contains the same data as the current MessageRecipient
func (mr MessageRecipient) Equals(o types.RVType) bool {
	if _, ok := o.(MessageRecipient); !ok {
		return false
	}

	other := o.(MessageRecipient)

	if mr.StructureVersion != other.StructureVersion {
		return false
	}

	if !mr.IDRecipient.Equals(other.IDRecipient) {
		return false
	}

	if !mr.UIRecipientType.Equals(other.UIRecipientType) {
		return false
	}

	if !mr.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	return mr.GatheringID.Equals(other.GatheringID)
}

// CopyRef copies the current value of the MessageRecipient
// and returns a pointer to the new copy
func (mr MessageRecipient) CopyRef() types.RVTypePtr {
	copied := mr.Copy().(MessageRecipient)
	return &copied
}

// Deref takes a pointer to the MessageRecipient
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (mr *MessageRecipient) Deref() types.RVType {
	return *mr
}

// String returns the string representation of the MessageRecipient
func (mr MessageRecipient) String() string {
	return mr.FormatToString(0)
}

// FormatToString pretty-prints the MessageRecipient using the provided indentation level
func (mr MessageRecipient) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MessageRecipient{\n")
	b.WriteString(fmt.Sprintf("%sIDRecipient: %s,\n", indentationValues, mr.IDRecipient))
	b.WriteString(fmt.Sprintf("%sUIRecipientType: %s,\n", indentationValues, mr.UIRecipientType))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, mr.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sGatheringID: %s,\n", indentationValues, mr.GatheringID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMessageRecipient returns a new MessageRecipient
func NewMessageRecipient() MessageRecipient {
	return MessageRecipient{
		IDRecipient:     types.NewUInt32(0),
		UIRecipientType: types.NewUInt32(0),
		PrincipalID:     types.NewPID(0),
		GatheringID:     types.NewUInt32(0),
	}

}

// Package types implements all the types used by the NintendoNotifications protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NintendoNotificationEventGeneral is a type within the NintendoNotifications protocol
type NintendoNotificationEventGeneral struct {
	types.Structure
	types.Data
	U32Param  types.UInt32
	U64Param1 types.UInt64
	U64Param2 types.UInt64
	StrParam  types.String
}

// ObjectID returns the object identifier of the type
func (nneg NintendoNotificationEventGeneral) ObjectID() types.RVType {
	return nneg.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (nneg NintendoNotificationEventGeneral) DataObjectID() types.RVType {
	return types.NewString("NintendoNotificationEventGeneral")
}

// WriteTo writes the NintendoNotificationEventGeneral to the given writable
func (nneg NintendoNotificationEventGeneral) WriteTo(writable types.Writable) {
	nneg.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	nneg.U32Param.WriteTo(contentWritable)
	nneg.U64Param1.WriteTo(contentWritable)
	nneg.U64Param2.WriteTo(contentWritable)
	nneg.StrParam.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nneg.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NintendoNotificationEventGeneral from the given readable
func (nneg *NintendoNotificationEventGeneral) ExtractFrom(readable types.Readable) error {
	var err error

	err = nneg.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventGeneral.Data. %s", err.Error())
	}

	err = nneg.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventGeneral header. %s", err.Error())
	}

	err = nneg.U32Param.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventGeneral.U32Param. %s", err.Error())
	}

	err = nneg.U64Param1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventGeneral.U64Param1. %s", err.Error())
	}

	err = nneg.U64Param2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventGeneral.U64Param2. %s", err.Error())
	}

	err = nneg.StrParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventGeneral.StrParam. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoNotificationEventGeneral
func (nneg NintendoNotificationEventGeneral) Copy() types.RVType {
	copied := NewNintendoNotificationEventGeneral()

	copied.StructureVersion = nneg.StructureVersion
	copied.Data = nneg.Data.Copy().(types.Data)
	copied.U32Param = nneg.U32Param.Copy().(types.UInt32)
	copied.U64Param1 = nneg.U64Param1.Copy().(types.UInt64)
	copied.U64Param2 = nneg.U64Param2.Copy().(types.UInt64)
	copied.StrParam = nneg.StrParam.Copy().(types.String)

	return copied
}

// Equals checks if the given NintendoNotificationEventGeneral contains the same data as the current NintendoNotificationEventGeneral
func (nneg NintendoNotificationEventGeneral) Equals(o types.RVType) bool {
	if _, ok := o.(NintendoNotificationEventGeneral); !ok {
		return false
	}

	other := o.(NintendoNotificationEventGeneral)

	if nneg.StructureVersion != other.StructureVersion {
		return false
	}

	if !nneg.Data.Equals(other.Data) {
		return false
	}

	if !nneg.U32Param.Equals(other.U32Param) {
		return false
	}

	if !nneg.U64Param1.Equals(other.U64Param1) {
		return false
	}

	if !nneg.U64Param2.Equals(other.U64Param2) {
		return false
	}

	return nneg.StrParam.Equals(other.StrParam)
}

// CopyRef copies the current value of the NintendoNotificationEventGeneral
// and returns a pointer to the new copy
func (nneg NintendoNotificationEventGeneral) CopyRef() types.RVTypePtr {
	copied := nneg.Copy().(NintendoNotificationEventGeneral)
	return &copied
}

// Deref takes a pointer to the NintendoNotificationEventGeneral
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (nneg *NintendoNotificationEventGeneral) Deref() types.RVType {
	return *nneg
}

// String returns the string representation of the NintendoNotificationEventGeneral
func (nneg NintendoNotificationEventGeneral) String() string {
	return nneg.FormatToString(0)
}

// FormatToString pretty-prints the NintendoNotificationEventGeneral using the provided indentation level
func (nneg NintendoNotificationEventGeneral) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoNotificationEventGeneral{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, nneg.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sU32Param: %s,\n", indentationValues, nneg.U32Param))
	b.WriteString(fmt.Sprintf("%sU64Param1: %s,\n", indentationValues, nneg.U64Param1))
	b.WriteString(fmt.Sprintf("%sU64Param2: %s,\n", indentationValues, nneg.U64Param2))
	b.WriteString(fmt.Sprintf("%sStrParam: %s,\n", indentationValues, nneg.StrParam))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoNotificationEventGeneral returns a new NintendoNotificationEventGeneral
func NewNintendoNotificationEventGeneral() NintendoNotificationEventGeneral {
	return NintendoNotificationEventGeneral{
		U32Param:  types.NewUInt32(0),
		U64Param1: types.NewUInt64(0),
		U64Param2: types.NewUInt64(0),
		StrParam:  types.NewString(""),
	}

}

// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetPrepurchaseInfoParam is a type within the ServiceItem protocol
type ServiceItemGetPrepurchaseInfoParam struct {
	types.Structure
	ItemCode    types.String
	ReferenceID types.String
	Limitation  ServiceItemLimitation
	Language    types.String
	UniqueID    types.UInt32
}

// WriteTo writes the ServiceItemGetPrepurchaseInfoParam to the given writable
func (sigpip ServiceItemGetPrepurchaseInfoParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sigpip.ItemCode.WriteTo(contentWritable)
	sigpip.ReferenceID.WriteTo(contentWritable)
	sigpip.Limitation.WriteTo(contentWritable)
	sigpip.Language.WriteTo(contentWritable)
	sigpip.UniqueID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sigpip.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetPrepurchaseInfoParam from the given readable
func (sigpip *ServiceItemGetPrepurchaseInfoParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = sigpip.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam header. %s", err.Error())
	}

	err = sigpip.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.ItemCode. %s", err.Error())
	}

	err = sigpip.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.ReferenceID. %s", err.Error())
	}

	err = sigpip.Limitation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.Limitation. %s", err.Error())
	}

	err = sigpip.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.Language. %s", err.Error())
	}

	err = sigpip.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.UniqueID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetPrepurchaseInfoParam
func (sigpip ServiceItemGetPrepurchaseInfoParam) Copy() types.RVType {
	copied := NewServiceItemGetPrepurchaseInfoParam()

	copied.StructureVersion = sigpip.StructureVersion
	copied.ItemCode = sigpip.ItemCode.Copy().(types.String)
	copied.ReferenceID = sigpip.ReferenceID.Copy().(types.String)
	copied.Limitation = sigpip.Limitation.Copy().(ServiceItemLimitation)
	copied.Language = sigpip.Language.Copy().(types.String)
	copied.UniqueID = sigpip.UniqueID.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given ServiceItemGetPrepurchaseInfoParam contains the same data as the current ServiceItemGetPrepurchaseInfoParam
func (sigpip ServiceItemGetPrepurchaseInfoParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetPrepurchaseInfoParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetPrepurchaseInfoParam)

	if sigpip.StructureVersion != other.StructureVersion {
		return false
	}

	if !sigpip.ItemCode.Equals(other.ItemCode) {
		return false
	}

	if !sigpip.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !sigpip.Limitation.Equals(other.Limitation) {
		return false
	}

	if !sigpip.Language.Equals(other.Language) {
		return false
	}

	return sigpip.UniqueID.Equals(other.UniqueID)
}

// CopyRef copies the current value of the ServiceItemGetPrepurchaseInfoParam
// and returns a pointer to the new copy
func (sigpip ServiceItemGetPrepurchaseInfoParam) CopyRef() types.RVTypePtr {
	copied := sigpip.Copy().(ServiceItemGetPrepurchaseInfoParam)
	return &copied
}

// Deref takes a pointer to the ServiceItemGetPrepurchaseInfoParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sigpip *ServiceItemGetPrepurchaseInfoParam) Deref() types.RVType {
	return *sigpip
}

// String returns the string representation of the ServiceItemGetPrepurchaseInfoParam
func (sigpip ServiceItemGetPrepurchaseInfoParam) String() string {
	return sigpip.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetPrepurchaseInfoParam using the provided indentation level
func (sigpip ServiceItemGetPrepurchaseInfoParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetPrepurchaseInfoParam{\n")
	b.WriteString(fmt.Sprintf("%sItemCode: %s,\n", indentationValues, sigpip.ItemCode))
	b.WriteString(fmt.Sprintf("%sReferenceID: %s,\n", indentationValues, sigpip.ReferenceID))
	b.WriteString(fmt.Sprintf("%sLimitation: %s,\n", indentationValues, sigpip.Limitation.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sLanguage: %s,\n", indentationValues, sigpip.Language))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, sigpip.UniqueID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetPrepurchaseInfoParam returns a new ServiceItemGetPrepurchaseInfoParam
func NewServiceItemGetPrepurchaseInfoParam() ServiceItemGetPrepurchaseInfoParam {
	return ServiceItemGetPrepurchaseInfoParam{
		ItemCode:    types.NewString(""),
		ReferenceID: types.NewString(""),
		Limitation:  NewServiceItemLimitation(),
		Language:    types.NewString(""),
		UniqueID:    types.NewUInt32(0),
	}

}

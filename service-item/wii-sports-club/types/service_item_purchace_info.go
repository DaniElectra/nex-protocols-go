// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemPurchaceInfo is a type within the ServiceItem protocol
type ServiceItemPurchaceInfo struct {
	types.Structure
	TransactionID    types.String
	ExtTransactionID types.String
	ItemCode         types.String
	PostBalance      ServiceItemAmount
}

// WriteTo writes the ServiceItemPurchaceInfo to the given writable
func (sipi ServiceItemPurchaceInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sipi.TransactionID.WriteTo(contentWritable)
	sipi.ExtTransactionID.WriteTo(contentWritable)
	sipi.ItemCode.WriteTo(contentWritable)
	sipi.PostBalance.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sipi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemPurchaceInfo from the given readable
func (sipi *ServiceItemPurchaceInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = sipi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo header. %s", err.Error())
	}

	err = sipi.TransactionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo.TransactionID. %s", err.Error())
	}

	err = sipi.ExtTransactionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo.ExtTransactionID. %s", err.Error())
	}

	err = sipi.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo.ItemCode. %s", err.Error())
	}

	err = sipi.PostBalance.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo.PostBalance. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemPurchaceInfo
func (sipi ServiceItemPurchaceInfo) Copy() types.RVType {
	copied := NewServiceItemPurchaceInfo()

	copied.StructureVersion = sipi.StructureVersion
	copied.TransactionID = sipi.TransactionID.Copy().(types.String)
	copied.ExtTransactionID = sipi.ExtTransactionID.Copy().(types.String)
	copied.ItemCode = sipi.ItemCode.Copy().(types.String)
	copied.PostBalance = sipi.PostBalance.Copy().(ServiceItemAmount)

	return copied
}

// Equals checks if the given ServiceItemPurchaceInfo contains the same data as the current ServiceItemPurchaceInfo
func (sipi ServiceItemPurchaceInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPurchaceInfo); !ok {
		return false
	}

	other := o.(*ServiceItemPurchaceInfo)

	if sipi.StructureVersion != other.StructureVersion {
		return false
	}

	if !sipi.TransactionID.Equals(other.TransactionID) {
		return false
	}

	if !sipi.ExtTransactionID.Equals(other.ExtTransactionID) {
		return false
	}

	if !sipi.ItemCode.Equals(other.ItemCode) {
		return false
	}

	return sipi.PostBalance.Equals(other.PostBalance)
}

// String returns the string representation of the ServiceItemPurchaceInfo
func (sipi ServiceItemPurchaceInfo) String() string {
	return sipi.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemPurchaceInfo using the provided indentation level
func (sipi ServiceItemPurchaceInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaceInfo{\n")
	b.WriteString(fmt.Sprintf("%sTransactionID: %s,\n", indentationValues, sipi.TransactionID))
	b.WriteString(fmt.Sprintf("%sExtTransactionID: %s,\n", indentationValues, sipi.ExtTransactionID))
	b.WriteString(fmt.Sprintf("%sItemCode: %s,\n", indentationValues, sipi.ItemCode))
	b.WriteString(fmt.Sprintf("%sPostBalance: %s,\n", indentationValues, sipi.PostBalance.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPurchaceInfo returns a new ServiceItemPurchaceInfo
func NewServiceItemPurchaceInfo() ServiceItemPurchaceInfo {
	return ServiceItemPurchaceInfo{
		TransactionID:    types.NewString(""),
		ExtTransactionID: types.NewString(""),
		ItemCode:         types.NewString(""),
		PostBalance:      NewServiceItemAmount(),
	}

}

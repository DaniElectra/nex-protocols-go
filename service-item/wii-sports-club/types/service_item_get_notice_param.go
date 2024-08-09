// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetNoticeParam is a type within the ServiceItem protocol
type ServiceItemGetNoticeParam struct {
	types.Structure
	NoticeType types.UInt32
}

// WriteTo writes the ServiceItemGetNoticeParam to the given writable
func (signp ServiceItemGetNoticeParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	signp.NoticeType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	signp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetNoticeParam from the given readable
func (signp *ServiceItemGetNoticeParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = signp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetNoticeParam header. %s", err.Error())
	}

	err = signp.NoticeType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetNoticeParam.NoticeType. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetNoticeParam
func (signp ServiceItemGetNoticeParam) Copy() types.RVType {
	copied := NewServiceItemGetNoticeParam()

	copied.StructureVersion = signp.StructureVersion
	copied.NoticeType = signp.NoticeType.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given ServiceItemGetNoticeParam contains the same data as the current ServiceItemGetNoticeParam
func (signp ServiceItemGetNoticeParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetNoticeParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetNoticeParam)

	if signp.StructureVersion != other.StructureVersion {
		return false
	}

	return signp.NoticeType.Equals(other.NoticeType)
}

// String returns the string representation of the ServiceItemGetNoticeParam
func (signp ServiceItemGetNoticeParam) String() string {
	return signp.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetNoticeParam using the provided indentation level
func (signp ServiceItemGetNoticeParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetNoticeParam{\n")
	b.WriteString(fmt.Sprintf("%sNoticeType: %s,\n", indentationValues, signp.NoticeType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetNoticeParam returns a new ServiceItemGetNoticeParam
func NewServiceItemGetNoticeParam() ServiceItemGetNoticeParam {
	return ServiceItemGetNoticeParam{
		NoticeType: types.NewUInt32(0),
	}

}

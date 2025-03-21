// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemEndChallengeParam is a type within the ServiceItem protocol
type ServiceItemEndChallengeParam struct {
	types.Structure
	ChallengeScheduleID types.UInt32
	UserInfo            ServiceItemUserInfo
}

// WriteTo writes the ServiceItemEndChallengeParam to the given writable
func (siecp ServiceItemEndChallengeParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siecp.ChallengeScheduleID.WriteTo(contentWritable)
	siecp.UserInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siecp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemEndChallengeParam from the given readable
func (siecp *ServiceItemEndChallengeParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = siecp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEndChallengeParam header. %s", err.Error())
	}

	err = siecp.ChallengeScheduleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEndChallengeParam.ChallengeScheduleID. %s", err.Error())
	}

	err = siecp.UserInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEndChallengeParam.UserInfo. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemEndChallengeParam
func (siecp ServiceItemEndChallengeParam) Copy() types.RVType {
	copied := NewServiceItemEndChallengeParam()

	copied.StructureVersion = siecp.StructureVersion
	copied.ChallengeScheduleID = siecp.ChallengeScheduleID.Copy().(types.UInt32)
	copied.UserInfo = siecp.UserInfo.Copy().(ServiceItemUserInfo)

	return copied
}

// Equals checks if the given ServiceItemEndChallengeParam contains the same data as the current ServiceItemEndChallengeParam
func (siecp ServiceItemEndChallengeParam) Equals(o types.RVType) bool {
	if _, ok := o.(ServiceItemEndChallengeParam); !ok {
		return false
	}

	other := o.(ServiceItemEndChallengeParam)

	if siecp.StructureVersion != other.StructureVersion {
		return false
	}

	if !siecp.ChallengeScheduleID.Equals(other.ChallengeScheduleID) {
		return false
	}

	return siecp.UserInfo.Equals(other.UserInfo)
}

// CopyRef copies the current value of the ServiceItemEndChallengeParam
// and returns a pointer to the new copy
func (siecp ServiceItemEndChallengeParam) CopyRef() types.RVTypePtr {
	copied := siecp.Copy().(ServiceItemEndChallengeParam)
	return &copied
}

// Deref takes a pointer to the ServiceItemEndChallengeParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (siecp *ServiceItemEndChallengeParam) Deref() types.RVType {
	return *siecp
}

// String returns the string representation of the ServiceItemEndChallengeParam
func (siecp ServiceItemEndChallengeParam) String() string {
	return siecp.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemEndChallengeParam using the provided indentation level
func (siecp ServiceItemEndChallengeParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemEndChallengeParam{\n")
	b.WriteString(fmt.Sprintf("%sChallengeScheduleID: %s,\n", indentationValues, siecp.ChallengeScheduleID))
	b.WriteString(fmt.Sprintf("%sUserInfo: %s,\n", indentationValues, siecp.UserInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemEndChallengeParam returns a new ServiceItemEndChallengeParam
func NewServiceItemEndChallengeParam() ServiceItemEndChallengeParam {
	return ServiceItemEndChallengeParam{
		ChallengeScheduleID: types.NewUInt32(0),
		UserInfo:            NewServiceItemUserInfo(),
	}

}

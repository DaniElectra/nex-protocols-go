// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RankingChangeAttributesParam is a type within the Ranking protocol
type RankingChangeAttributesParam struct {
	types.Structure
	ModificationFlag types.UInt8
	Groups           types.List[types.UInt8]
	Param            types.UInt64
}

// WriteTo writes the RankingChangeAttributesParam to the given writable
func (rcap RankingChangeAttributesParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rcap.ModificationFlag.WriteTo(contentWritable)
	rcap.Groups.WriteTo(contentWritable)
	rcap.Param.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rcap.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RankingChangeAttributesParam from the given readable
func (rcap *RankingChangeAttributesParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = rcap.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingChangeAttributesParam header. %s", err.Error())
	}

	err = rcap.ModificationFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingChangeAttributesParam.ModificationFlag. %s", err.Error())
	}

	err = rcap.Groups.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingChangeAttributesParam.Groups. %s", err.Error())
	}

	err = rcap.Param.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingChangeAttributesParam.Param. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RankingChangeAttributesParam
func (rcap RankingChangeAttributesParam) Copy() types.RVType {
	copied := NewRankingChangeAttributesParam()

	copied.StructureVersion = rcap.StructureVersion
	copied.ModificationFlag = rcap.ModificationFlag.Copy().(types.UInt8)
	copied.Groups = rcap.Groups.Copy().(types.List[types.UInt8])
	copied.Param = rcap.Param.Copy().(types.UInt64)

	return copied
}

// Equals checks if the given RankingChangeAttributesParam contains the same data as the current RankingChangeAttributesParam
func (rcap RankingChangeAttributesParam) Equals(o types.RVType) bool {
	if _, ok := o.(RankingChangeAttributesParam); !ok {
		return false
	}

	other := o.(RankingChangeAttributesParam)

	if rcap.StructureVersion != other.StructureVersion {
		return false
	}

	if !rcap.ModificationFlag.Equals(other.ModificationFlag) {
		return false
	}

	if !rcap.Groups.Equals(other.Groups) {
		return false
	}

	return rcap.Param.Equals(other.Param)
}

// CopyRef copies the current value of the RankingChangeAttributesParam
// and returns a pointer to the new copy
func (rcap RankingChangeAttributesParam) CopyRef() types.RVTypePtr {
	copied := rcap.Copy().(RankingChangeAttributesParam)
	return &copied
}

// Deref takes a pointer to the RankingChangeAttributesParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rcap *RankingChangeAttributesParam) Deref() types.RVType {
	return *rcap
}

// String returns the string representation of the RankingChangeAttributesParam
func (rcap RankingChangeAttributesParam) String() string {
	return rcap.FormatToString(0)
}

// FormatToString pretty-prints the RankingChangeAttributesParam using the provided indentation level
func (rcap RankingChangeAttributesParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingChangeAttributesParam{\n")
	b.WriteString(fmt.Sprintf("%sModificationFlag: %s,\n", indentationValues, rcap.ModificationFlag))
	b.WriteString(fmt.Sprintf("%sGroups: %s,\n", indentationValues, rcap.Groups))
	b.WriteString(fmt.Sprintf("%sParam: %s,\n", indentationValues, rcap.Param))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingChangeAttributesParam returns a new RankingChangeAttributesParam
func NewRankingChangeAttributesParam() RankingChangeAttributesParam {
	return RankingChangeAttributesParam{
		ModificationFlag: types.NewUInt8(0),
		Groups:           types.NewList[types.UInt8](),
		Param:            types.NewUInt64(0),
	}

}

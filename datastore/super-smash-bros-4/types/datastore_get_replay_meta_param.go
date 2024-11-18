// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreGetReplayMetaParam is a type within the DataStoreSuperSmashBros.4 protocol
type DataStoreGetReplayMetaParam struct {
	types.Structure
	ReplayID types.UInt64
	MetaType types.UInt8
}

// WriteTo writes the DataStoreGetReplayMetaParam to the given writable
func (dsgrmp DataStoreGetReplayMetaParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsgrmp.ReplayID.WriteTo(contentWritable)
	dsgrmp.MetaType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsgrmp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreGetReplayMetaParam from the given readable
func (dsgrmp *DataStoreGetReplayMetaParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsgrmp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetReplayMetaParam header. %s", err.Error())
	}

	err = dsgrmp.ReplayID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetReplayMetaParam.ReplayID. %s", err.Error())
	}

	err = dsgrmp.MetaType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetReplayMetaParam.MetaType. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetReplayMetaParam
func (dsgrmp DataStoreGetReplayMetaParam) Copy() types.RVType {
	copied := NewDataStoreGetReplayMetaParam()

	copied.StructureVersion = dsgrmp.StructureVersion
	copied.ReplayID = dsgrmp.ReplayID.Copy().(types.UInt64)
	copied.MetaType = dsgrmp.MetaType.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given DataStoreGetReplayMetaParam contains the same data as the current DataStoreGetReplayMetaParam
func (dsgrmp DataStoreGetReplayMetaParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreGetReplayMetaParam); !ok {
		return false
	}

	other := o.(*DataStoreGetReplayMetaParam)

	if dsgrmp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsgrmp.ReplayID.Equals(other.ReplayID) {
		return false
	}

	return dsgrmp.MetaType.Equals(other.MetaType)
}

// CopyRef copies the current value of the DataStoreGetReplayMetaParam
// and returns a pointer to the new copy
func (dsgrmp DataStoreGetReplayMetaParam) CopyRef() types.RVTypePtr {
	copied := dsgrmp.Copy().(DataStoreGetReplayMetaParam)
	return &copied
}

// Deref takes a pointer to the DataStoreGetReplayMetaParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dsgrmp *DataStoreGetReplayMetaParam) Deref() types.RVType {
	return *dsgrmp
}

// String returns the string representation of the DataStoreGetReplayMetaParam
func (dsgrmp DataStoreGetReplayMetaParam) String() string {
	return dsgrmp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreGetReplayMetaParam using the provided indentation level
func (dsgrmp DataStoreGetReplayMetaParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetReplayMetaParam{\n")
	b.WriteString(fmt.Sprintf("%sReplayID: %s,\n", indentationValues, dsgrmp.ReplayID))
	b.WriteString(fmt.Sprintf("%sMetaType: %s,\n", indentationValues, dsgrmp.MetaType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetReplayMetaParam returns a new DataStoreGetReplayMetaParam
func NewDataStoreGetReplayMetaParam() DataStoreGetReplayMetaParam {
	return DataStoreGetReplayMetaParam{
		ReplayID: types.NewUInt64(0),
		MetaType: types.NewUInt8(0),
	}

}

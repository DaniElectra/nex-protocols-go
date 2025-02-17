// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreNotification is a type within the DataStore protocol
type DataStoreNotification struct {
	types.Structure
	NotificationID types.UInt64
	DataID         types.UInt64
}

// WriteTo writes the DataStoreNotification to the given writable
func (dsn DataStoreNotification) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsn.NotificationID.WriteTo(contentWritable)
	dsn.DataID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsn.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreNotification from the given readable
func (dsn *DataStoreNotification) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsn.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotification header. %s", err.Error())
	}

	err = dsn.NotificationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotification.NotificationID. %s", err.Error())
	}

	err = dsn.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreNotification.DataID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreNotification
func (dsn DataStoreNotification) Copy() types.RVType {
	copied := NewDataStoreNotification()

	copied.StructureVersion = dsn.StructureVersion
	copied.NotificationID = dsn.NotificationID.Copy().(types.UInt64)
	copied.DataID = dsn.DataID.Copy().(types.UInt64)

	return copied
}

// Equals checks if the given DataStoreNotification contains the same data as the current DataStoreNotification
func (dsn DataStoreNotification) Equals(o types.RVType) bool {
	if _, ok := o.(DataStoreNotification); !ok {
		return false
	}

	other := o.(DataStoreNotification)

	if dsn.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsn.NotificationID.Equals(other.NotificationID) {
		return false
	}

	return dsn.DataID.Equals(other.DataID)
}

// CopyRef copies the current value of the DataStoreNotification
// and returns a pointer to the new copy
func (dsn DataStoreNotification) CopyRef() types.RVTypePtr {
	copied := dsn.Copy().(DataStoreNotification)
	return &copied
}

// Deref takes a pointer to the DataStoreNotification
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dsn *DataStoreNotification) Deref() types.RVType {
	return *dsn
}

// String returns the string representation of the DataStoreNotification
func (dsn DataStoreNotification) String() string {
	return dsn.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreNotification using the provided indentation level
func (dsn DataStoreNotification) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreNotification{\n")
	b.WriteString(fmt.Sprintf("%sNotificationID: %s,\n", indentationValues, dsn.NotificationID))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsn.DataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreNotification returns a new DataStoreNotification
func NewDataStoreNotification() DataStoreNotification {
	return DataStoreNotification{
		NotificationID: types.NewUInt64(0),
		DataID:         types.NewUInt64(0),
	}

}

// Code generated by protoc-gen-gogo.
// source: cockroach/sql/structured.proto
// DO NOT EDIT!

package sql

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "gogoproto"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// These mirror the types supported by the sql/parser. See
// sql/parser/types.go.
type ColumnType_Kind int32

const (
	ColumnType_BIT       ColumnType_Kind = 0
	ColumnType_BOOL      ColumnType_Kind = 1
	ColumnType_INT       ColumnType_Kind = 2
	ColumnType_FLOAT     ColumnType_Kind = 3
	ColumnType_DECIMAL   ColumnType_Kind = 4
	ColumnType_DATE      ColumnType_Kind = 5
	ColumnType_TIME      ColumnType_Kind = 6
	ColumnType_TIMESTAMP ColumnType_Kind = 7
	ColumnType_CHAR      ColumnType_Kind = 8
	ColumnType_TEXT      ColumnType_Kind = 9
	ColumnType_BLOB      ColumnType_Kind = 10
)

var ColumnType_Kind_name = map[int32]string{
	0:  "BIT",
	1:  "BOOL",
	2:  "INT",
	3:  "FLOAT",
	4:  "DECIMAL",
	5:  "DATE",
	6:  "TIME",
	7:  "TIMESTAMP",
	8:  "CHAR",
	9:  "TEXT",
	10: "BLOB",
}
var ColumnType_Kind_value = map[string]int32{
	"BIT":       0,
	"BOOL":      1,
	"INT":       2,
	"FLOAT":     3,
	"DECIMAL":   4,
	"DATE":      5,
	"TIME":      6,
	"TIMESTAMP": 7,
	"CHAR":      8,
	"TEXT":      9,
	"BLOB":      10,
}

func (x ColumnType_Kind) Enum() *ColumnType_Kind {
	p := new(ColumnType_Kind)
	*p = x
	return p
}
func (x ColumnType_Kind) String() string {
	return proto.EnumName(ColumnType_Kind_name, int32(x))
}
func (x *ColumnType_Kind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ColumnType_Kind_value, data, "ColumnType_Kind")
	if err != nil {
		return err
	}
	*x = ColumnType_Kind(value)
	return nil
}

type ColumnType struct {
	Kind ColumnType_Kind `protobuf:"varint,1,opt,name=kind,enum=cockroach.sql.ColumnType_Kind" json:"kind"`
	// BIT, INT, FLOAT, DECIMAL, CHAR and BINARY
	Width int32 `protobuf:"varint,2,opt,name=width" json:"width"`
	// FLOAT and DECIMAL.
	Precision int32 `protobuf:"varint,3,opt,name=precision" json:"precision"`
}

func (m *ColumnType) Reset()         { *m = ColumnType{} }
func (m *ColumnType) String() string { return proto.CompactTextString(m) }
func (*ColumnType) ProtoMessage()    {}

func (m *ColumnType) GetKind() ColumnType_Kind {
	if m != nil {
		return m.Kind
	}
	return ColumnType_BIT
}

func (m *ColumnType) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ColumnType) GetPrecision() int32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

type ColumnDescriptor struct {
	Name     string     `protobuf:"bytes,1,opt,name=name" json:"name"`
	ID       ColumnID   `protobuf:"varint,2,opt,name=id,casttype=ColumnID" json:"id"`
	Type     ColumnType `protobuf:"bytes,3,opt,name=type" json:"type"`
	Nullable bool       `protobuf:"varint,4,opt,name=nullable" json:"nullable"`
}

func (m *ColumnDescriptor) Reset()         { *m = ColumnDescriptor{} }
func (m *ColumnDescriptor) String() string { return proto.CompactTextString(m) }
func (*ColumnDescriptor) ProtoMessage()    {}

func (m *ColumnDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ColumnDescriptor) GetID() ColumnID {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ColumnDescriptor) GetType() ColumnType {
	if m != nil {
		return m.Type
	}
	return ColumnType{}
}

func (m *ColumnDescriptor) GetNullable() bool {
	if m != nil {
		return m.Nullable
	}
	return false
}

type IndexDescriptor struct {
	Name   string  `protobuf:"bytes,1,opt,name=name" json:"name"`
	ID     IndexID `protobuf:"varint,2,opt,name=id,casttype=IndexID" json:"id"`
	Unique bool    `protobuf:"varint,3,opt,name=unique" json:"unique"`
	// An ordered list of column names of which the index is comprised. This list
	// parallels the column_ids list. If duplicating the storage of the column
	// names here proves to be prohibitive, we could clear this field before
	// saving and reconstruct it after loading.
	ColumnNames []string `protobuf:"bytes,4,rep,name=column_names" json:"column_names,omitempty"`
	// An ordered list of column ids of which the index is comprised. This list
	// parallels the column_names list.
	ColumnIDs []ColumnID `protobuf:"varint,5,rep,name=column_ids,casttype=ColumnID" json:"column_ids,omitempty"`
	// An ordered list of implicit column ids associated with the index. For
	// non-unique indexes, these columns will be appended to the key. For unique
	// indexes these columns will be stored in the value. The extra column IDs is
	// computed as PrimaryIndex.column_ids - column_ids. For the primary index
	// the list will be empty.
	ImplicitColumnIDs []ColumnID `protobuf:"varint,6,rep,name=implicit_column_ids,casttype=ColumnID" json:"implicit_column_ids,omitempty"`
}

func (m *IndexDescriptor) Reset()         { *m = IndexDescriptor{} }
func (m *IndexDescriptor) String() string { return proto.CompactTextString(m) }
func (*IndexDescriptor) ProtoMessage()    {}

func (m *IndexDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IndexDescriptor) GetID() IndexID {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *IndexDescriptor) GetUnique() bool {
	if m != nil {
		return m.Unique
	}
	return false
}

func (m *IndexDescriptor) GetColumnNames() []string {
	if m != nil {
		return m.ColumnNames
	}
	return nil
}

func (m *IndexDescriptor) GetColumnIDs() []ColumnID {
	if m != nil {
		return m.ColumnIDs
	}
	return nil
}

func (m *IndexDescriptor) GetImplicitColumnIDs() []ColumnID {
	if m != nil {
		return m.ImplicitColumnIDs
	}
	return nil
}

// A TableDescriptor represents a table and is stored in a structured metadata
// key. The TableDescriptor has a globally-unique ID, while its member
// {Column,Index}Descriptors have locally-unique IDs.
type TableDescriptor struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name"`
	// The alias for the table. This is only used during query
	// processing and not stored persistently.
	Alias   string             `protobuf:"bytes,2,opt,name=alias" json:"alias"`
	ID      ID                 `protobuf:"varint,3,opt,name=id,casttype=ID" json:"id"`
	Columns []ColumnDescriptor `protobuf:"bytes,4,rep,name=columns" json:"columns"`
	// next_column_id is used to ensure that deleted column ids are not reused.
	NextColumnID ColumnID        `protobuf:"varint,5,opt,name=next_column_id,casttype=ColumnID" json:"next_column_id"`
	PrimaryIndex IndexDescriptor `protobuf:"bytes,6,opt,name=primary_index" json:"primary_index"`
	// indexes are all the secondary indexes.
	Indexes []IndexDescriptor `protobuf:"bytes,7,rep,name=indexes" json:"indexes"`
	// next_index_id is used to ensure that deleted index ids are not reused.
	NextIndexID IndexID              `protobuf:"varint,8,opt,name=next_index_id,casttype=IndexID" json:"next_index_id"`
	Privileges  *PrivilegeDescriptor `protobuf:"bytes,9,opt,name=privileges" json:"privileges,omitempty"`
}

func (m *TableDescriptor) Reset()         { *m = TableDescriptor{} }
func (m *TableDescriptor) String() string { return proto.CompactTextString(m) }
func (*TableDescriptor) ProtoMessage()    {}

func (m *TableDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TableDescriptor) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *TableDescriptor) GetID() ID {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *TableDescriptor) GetColumns() []ColumnDescriptor {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *TableDescriptor) GetNextColumnID() ColumnID {
	if m != nil {
		return m.NextColumnID
	}
	return 0
}

func (m *TableDescriptor) GetPrimaryIndex() IndexDescriptor {
	if m != nil {
		return m.PrimaryIndex
	}
	return IndexDescriptor{}
}

func (m *TableDescriptor) GetIndexes() []IndexDescriptor {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *TableDescriptor) GetNextIndexID() IndexID {
	if m != nil {
		return m.NextIndexID
	}
	return 0
}

func (m *TableDescriptor) GetPrivileges() *PrivilegeDescriptor {
	if m != nil {
		return m.Privileges
	}
	return nil
}

// DatabaseDescriptor represents a namespace (aka database) and is stored
// in a structured metadata key. The DatabaseDescriptor has a globally-unique
// ID shared with the TableDescriptor ID.
// Permissions are applied to all tables in the namespace.
type DatabaseDescriptor struct {
	Name       string               `protobuf:"bytes,1,opt,name=name" json:"name"`
	ID         ID                   `protobuf:"varint,2,opt,name=id,casttype=ID" json:"id"`
	Privileges *PrivilegeDescriptor `protobuf:"bytes,3,opt,name=privileges" json:"privileges,omitempty"`
}

func (m *DatabaseDescriptor) Reset()         { *m = DatabaseDescriptor{} }
func (m *DatabaseDescriptor) String() string { return proto.CompactTextString(m) }
func (*DatabaseDescriptor) ProtoMessage()    {}

func (m *DatabaseDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatabaseDescriptor) GetID() ID {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *DatabaseDescriptor) GetPrivileges() *PrivilegeDescriptor {
	if m != nil {
		return m.Privileges
	}
	return nil
}

func init() {
	proto.RegisterEnum("cockroach.sql.ColumnType_Kind", ColumnType_Kind_name, ColumnType_Kind_value)
}
func (m *ColumnType) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ColumnType) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintStructured(data, i, uint64(m.Kind))
	data[i] = 0x10
	i++
	i = encodeVarintStructured(data, i, uint64(m.Width))
	data[i] = 0x18
	i++
	i = encodeVarintStructured(data, i, uint64(m.Precision))
	return i, nil
}

func (m *ColumnDescriptor) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ColumnDescriptor) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintStructured(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x10
	i++
	i = encodeVarintStructured(data, i, uint64(m.ID))
	data[i] = 0x1a
	i++
	i = encodeVarintStructured(data, i, uint64(m.Type.Size()))
	n1, err := m.Type.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x20
	i++
	if m.Nullable {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	return i, nil
}

func (m *IndexDescriptor) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *IndexDescriptor) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintStructured(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x10
	i++
	i = encodeVarintStructured(data, i, uint64(m.ID))
	data[i] = 0x18
	i++
	if m.Unique {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	if len(m.ColumnNames) > 0 {
		for _, s := range m.ColumnNames {
			data[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if len(m.ColumnIDs) > 0 {
		for _, num := range m.ColumnIDs {
			data[i] = 0x28
			i++
			i = encodeVarintStructured(data, i, uint64(num))
		}
	}
	if len(m.ImplicitColumnIDs) > 0 {
		for _, num := range m.ImplicitColumnIDs {
			data[i] = 0x30
			i++
			i = encodeVarintStructured(data, i, uint64(num))
		}
	}
	return i, nil
}

func (m *TableDescriptor) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TableDescriptor) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintStructured(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x12
	i++
	i = encodeVarintStructured(data, i, uint64(len(m.Alias)))
	i += copy(data[i:], m.Alias)
	data[i] = 0x18
	i++
	i = encodeVarintStructured(data, i, uint64(m.ID))
	if len(m.Columns) > 0 {
		for _, msg := range m.Columns {
			data[i] = 0x22
			i++
			i = encodeVarintStructured(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x28
	i++
	i = encodeVarintStructured(data, i, uint64(m.NextColumnID))
	data[i] = 0x32
	i++
	i = encodeVarintStructured(data, i, uint64(m.PrimaryIndex.Size()))
	n2, err := m.PrimaryIndex.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Indexes) > 0 {
		for _, msg := range m.Indexes {
			data[i] = 0x3a
			i++
			i = encodeVarintStructured(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x40
	i++
	i = encodeVarintStructured(data, i, uint64(m.NextIndexID))
	if m.Privileges != nil {
		data[i] = 0x4a
		i++
		i = encodeVarintStructured(data, i, uint64(m.Privileges.Size()))
		n3, err := m.Privileges.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *DatabaseDescriptor) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DatabaseDescriptor) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintStructured(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x10
	i++
	i = encodeVarintStructured(data, i, uint64(m.ID))
	if m.Privileges != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintStructured(data, i, uint64(m.Privileges.Size()))
		n4, err := m.Privileges.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeFixed64Structured(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Structured(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStructured(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ColumnType) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovStructured(uint64(m.Kind))
	n += 1 + sovStructured(uint64(m.Width))
	n += 1 + sovStructured(uint64(m.Precision))
	return n
}

func (m *ColumnDescriptor) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovStructured(uint64(l))
	n += 1 + sovStructured(uint64(m.ID))
	l = m.Type.Size()
	n += 1 + l + sovStructured(uint64(l))
	n += 2
	return n
}

func (m *IndexDescriptor) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovStructured(uint64(l))
	n += 1 + sovStructured(uint64(m.ID))
	n += 2
	if len(m.ColumnNames) > 0 {
		for _, s := range m.ColumnNames {
			l = len(s)
			n += 1 + l + sovStructured(uint64(l))
		}
	}
	if len(m.ColumnIDs) > 0 {
		for _, e := range m.ColumnIDs {
			n += 1 + sovStructured(uint64(e))
		}
	}
	if len(m.ImplicitColumnIDs) > 0 {
		for _, e := range m.ImplicitColumnIDs {
			n += 1 + sovStructured(uint64(e))
		}
	}
	return n
}

func (m *TableDescriptor) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovStructured(uint64(l))
	l = len(m.Alias)
	n += 1 + l + sovStructured(uint64(l))
	n += 1 + sovStructured(uint64(m.ID))
	if len(m.Columns) > 0 {
		for _, e := range m.Columns {
			l = e.Size()
			n += 1 + l + sovStructured(uint64(l))
		}
	}
	n += 1 + sovStructured(uint64(m.NextColumnID))
	l = m.PrimaryIndex.Size()
	n += 1 + l + sovStructured(uint64(l))
	if len(m.Indexes) > 0 {
		for _, e := range m.Indexes {
			l = e.Size()
			n += 1 + l + sovStructured(uint64(l))
		}
	}
	n += 1 + sovStructured(uint64(m.NextIndexID))
	if m.Privileges != nil {
		l = m.Privileges.Size()
		n += 1 + l + sovStructured(uint64(l))
	}
	return n
}

func (m *DatabaseDescriptor) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovStructured(uint64(l))
	n += 1 + sovStructured(uint64(m.ID))
	if m.Privileges != nil {
		l = m.Privileges.Size()
		n += 1 + l + sovStructured(uint64(l))
	}
	return n
}

func sovStructured(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStructured(x uint64) (n int) {
	return sovStructured(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ColumnType) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Kind |= (ColumnType_Kind(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Width", wireType)
			}
			m.Width = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Width |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
			}
			m.Precision = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Precision |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipStructured(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStructured
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *ColumnDescriptor) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (ColumnID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthStructured
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Type.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nullable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Nullable = bool(v != 0)
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipStructured(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStructured
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *IndexDescriptor) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (IndexID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unique", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Unique = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ColumnNames = append(m.ColumnNames, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnIDs", wireType)
			}
			var v ColumnID
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (ColumnID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ColumnIDs = append(m.ColumnIDs, v)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImplicitColumnIDs", wireType)
			}
			var v ColumnID
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (ColumnID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ImplicitColumnIDs = append(m.ImplicitColumnIDs, v)
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipStructured(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStructured
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *TableDescriptor) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alias = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (ID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthStructured
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, ColumnDescriptor{})
			if err := m.Columns[len(m.Columns)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextColumnID", wireType)
			}
			m.NextColumnID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NextColumnID |= (ColumnID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimaryIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthStructured
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PrimaryIndex.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Indexes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthStructured
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Indexes = append(m.Indexes, IndexDescriptor{})
			if err := m.Indexes[len(m.Indexes)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextIndexID", wireType)
			}
			m.NextIndexID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NextIndexID |= (IndexID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Privileges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthStructured
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Privileges == nil {
				m.Privileges = &PrivilegeDescriptor{}
			}
			if err := m.Privileges.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipStructured(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStructured
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *DatabaseDescriptor) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (ID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Privileges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthStructured
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Privileges == nil {
				m.Privileges = &PrivilegeDescriptor{}
			}
			if err := m.Privileges.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipStructured(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStructured
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipStructured(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStructured
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStructured(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStructured = fmt.Errorf("proto: negative length found during unmarshaling")
)

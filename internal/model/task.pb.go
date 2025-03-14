// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task.proto

package model

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Task
type Task struct {
	// id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"index:,priority:3,composite:idx_tasks_status_deleted_at_id;primaryKey;autoIncrement;comment:id"`
	// task id
	TaskId string `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty" gorm:"unique;size:128;comment:task id"`
	// task pid
	TaskPid string `protobuf:"bytes,3,opt,name=task_pid,json=taskPid,proto3" json:"task_pid,omitempty" gorm:"index;size:128;comment:task pid"`
	// uid
	Uid string `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty" gorm:"index;size:128;comment:uid"`
	// other id
	OtherId string `protobuf:"bytes,5,opt,name=other_id,json=otherId,proto3" json:"other_id,omitempty" gorm:"index;size:128;comment:other id"`
	// cover
	Cover string `protobuf:"bytes,6,opt,name=cover,proto3" json:"cover,omitempty" gorm:"comment:cover"`
	// task name
	TaskName string `protobuf:"bytes,7,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty" gorm:"size:128;comment:task name"`
	// other type
	OtherType int32 `protobuf:"varint,9,opt,name=other_type,json=otherType,proto3" json:"other_type,omitempty" gorm:"index;size:4;comment:other type"`
	// other tag
	OtherTag int32 `protobuf:"varint,10,opt,name=other_tag,json=otherTag,proto3" json:"other_tag,omitempty" gorm:"index;size:4;comment:other tag"`
	// other form
	OtherForm int32 `protobuf:"varint,11,opt,name=other_form,json=otherForm,proto3" json:"other_form,omitempty" gorm:"index;size:4;comment:other form"`
	// other group
	OtherGroup int32 `protobuf:"varint,12,opt,name=other_group,json=otherGroup,proto3" json:"other_group,omitempty" gorm:"index;size:4;comment:other group"`
	// task type
	TaskType string `protobuf:"bytes,13,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty" gorm:"index;size:50;comment:task type"`
	// ciphertext
	Ciphertext string `protobuf:"bytes,14,opt,name=ciphertext,proto3" json:"ciphertext,omitempty" gorm:"comment:ciphertext"`
	// rewrite_hls
	RewriteHls []byte `protobuf:"bytes,15,opt,name=rewrite_hls,json=rewriteHls,proto3" json:"rewrite_hls,omitempty" gorm:"comment:rewrite hls"`
	// bstatus
	Bstatus int32 `protobuf:"varint,16,opt,name=bstatus,proto3" json:"bstatus,omitempty" gorm:"index;size:4;comment:business status"`
	// raw
	Raw []byte `protobuf:"bytes,21,opt,name=raw,proto3" json:"raw,omitempty" gorm:"comment:raw"`
	// message
	Message []byte `protobuf:"bytes,22,opt,name=message,proto3" json:"message,omitempty" gorm:"comment:message"`
	// message
	Statistics []byte `protobuf:"bytes,23,opt,name=statistics,proto3" json:"statistics,omitempty" gorm:"comment:statistics"`
	// status
	Status int32 `protobuf:"varint,24,opt,name=status,proto3" json:"status,omitempty" gorm:"index:,priority:1,composite:idx_tasks_status_deleted_at_id;default:1;size:2;comment:status 1 normal"`
	// created_at
	CreatedAt int64 `protobuf:"varint,25,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" gorm:"index;comment:created at"`
	// updated_at
	UpdatedAt int64 `protobuf:"varint,26,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" gorm:"index;comment:updated at"`
	// deleted_at
	DeletedAt int64 `protobuf:"varint,27,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty" gorm:"index:,priority:2,composite:idx_tasks_status_deleted_at_id;comment:deleted at"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Task.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return m.Size()
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Task) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *Task) GetTaskPid() string {
	if m != nil {
		return m.TaskPid
	}
	return ""
}

func (m *Task) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Task) GetOtherId() string {
	if m != nil {
		return m.OtherId
	}
	return ""
}

func (m *Task) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *Task) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *Task) GetOtherType() int32 {
	if m != nil {
		return m.OtherType
	}
	return 0
}

func (m *Task) GetOtherTag() int32 {
	if m != nil {
		return m.OtherTag
	}
	return 0
}

func (m *Task) GetOtherForm() int32 {
	if m != nil {
		return m.OtherForm
	}
	return 0
}

func (m *Task) GetOtherGroup() int32 {
	if m != nil {
		return m.OtherGroup
	}
	return 0
}

func (m *Task) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *Task) GetCiphertext() string {
	if m != nil {
		return m.Ciphertext
	}
	return ""
}

func (m *Task) GetRewriteHls() []byte {
	if m != nil {
		return m.RewriteHls
	}
	return nil
}

func (m *Task) GetBstatus() int32 {
	if m != nil {
		return m.Bstatus
	}
	return 0
}

func (m *Task) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Task) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Task) GetStatistics() []byte {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *Task) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Task) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Task) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Task) GetDeletedAt() int64 {
	if m != nil {
		return m.DeletedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Task)(nil), "Task")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 796 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0x4f, 0x6f, 0x23, 0x35,
	0x18, 0xc6, 0x3b, 0xed, 0xb6, 0x21, 0x6e, 0xf9, 0x23, 0x0b, 0x8a, 0xb7, 0x88, 0x99, 0xc1, 0xab,
	0x85, 0xc0, 0x2e, 0x2d, 0xe9, 0x2e, 0x02, 0x25, 0x42, 0xa8, 0x41, 0xec, 0x12, 0x81, 0x10, 0x1a,
	0xe5, 0x80, 0xb8, 0x44, 0xce, 0xd8, 0x9d, 0x98, 0xcd, 0x8c, 0x07, 0xdb, 0x43, 0x1a, 0x3e, 0x00,
	0x67, 0x3e, 0x16, 0xc7, 0x3d, 0x72, 0x1a, 0x41, 0x7b, 0xe3, 0x38, 0x27, 0x8e, 0xc8, 0xf6, 0xa4,
	0x4d, 0xa3, 0x6c, 0xda, 0x3d, 0x65, 0xde, 0xc9, 0xf3, 0xfc, 0xde, 0xf7, 0x7d, 0xac, 0x38, 0x00,
	0x68, 0xa2, 0x9e, 0x1d, 0xe6, 0x52, 0x68, 0x71, 0xf0, 0x71, 0xc2, 0xf5, 0xb8, 0x18, 0x1d, 0xc6,
	0x22, 0x3d, 0x4a, 0x44, 0x22, 0x8e, 0xec, 0xeb, 0x51, 0x71, 0x6a, 0x2b, 0x5b, 0xd8, 0x27, 0x27,
	0xc7, 0xff, 0xee, 0x81, 0x3b, 0x03, 0xa2, 0x9e, 0xc1, 0x19, 0xd8, 0xe4, 0x14, 0x79, 0xa1, 0xd7,
	0xda, 0xea, 0xf1, 0xaa, 0x0c, 0x58, 0x22, 0x64, 0xda, 0xc1, 0x3c, 0xa3, 0xec, 0xac, 0xf3, 0x30,
	0x97, 0x5c, 0x48, 0xae, 0x67, 0x9d, 0x47, 0x0f, 0x63, 0x91, 0xe6, 0x42, 0x71, 0xcd, 0x3a, 0x9c,
	0x9e, 0x0d, 0x4d, 0x67, 0x35, 0x54, 0x9a, 0xe8, 0x42, 0x0d, 0x29, 0x9b, 0x30, 0xcd, 0xe8, 0x90,
	0xe8, 0x21, 0xa7, 0xdd, 0x5c, 0xf2, 0x94, 0xc8, 0xd9, 0xb7, 0x6c, 0xd6, 0x25, 0x85, 0x16, 0xfd,
	0x2c, 0x96, 0x2c, 0x65, 0x99, 0xee, 0xc6, 0x22, 0x35, 0x9f, 0x1d, 0x4e, 0x71, 0xb4, 0xc9, 0x29,
	0xfc, 0x0a, 0x34, 0x0c, 0x66, 0xc8, 0x29, 0xda, 0x0c, 0xbd, 0x56, 0xb3, 0xf7, 0x51, 0x55, 0x06,
	0xef, 0xbb, 0xfe, 0x45, 0xc6, 0x7f, 0x29, 0x58, 0x57, 0xf1, 0xdf, 0x58, 0xa7, 0x7d, 0xfc, 0xf9,
	0xa5, 0xd7, 0x18, 0x42, 0x03, 0xd8, 0x31, 0x4f, 0x7d, 0x0a, 0xbf, 0x06, 0xaf, 0x58, 0x48, 0xce,
	0x29, 0xda, 0x5a, 0xa6, 0xd8, 0x2d, 0x5e, 0x00, 0xc9, 0x0d, 0xc5, 0x0e, 0xf0, 0x03, 0xa7, 0xf0,
	0x33, 0xb0, 0x55, 0x70, 0x8a, 0xee, 0x58, 0xc2, 0xfd, 0xaa, 0x0c, 0xde, 0x5b, 0x4b, 0x28, 0x8c,
	0xd9, 0x38, 0x4c, 0x7f, 0xa1, 0xc7, 0x4c, 0x9a, 0x2d, 0xb6, 0x6f, 0xd9, 0xdf, 0x1a, 0xec, 0x16,
	0x0d, 0xfb, 0xd8, 0xa7, 0xf0, 0x10, 0x6c, 0xc7, 0xe2, 0x57, 0x26, 0xd1, 0x8e, 0x65, 0xa0, 0xaa,
	0x0c, 0xde, 0x74, 0x8c, 0xb9, 0xc9, 0x7e, 0x8d, 0x23, 0x27, 0x83, 0x3d, 0xd0, 0xb4, 0x6b, 0x67,
	0x24, 0x65, 0xa8, 0xb1, 0x3c, 0xf5, 0xea, 0x8d, 0x8d, 0x16, 0x47, 0x36, 0xae, 0xef, 0x49, 0xca,
	0x60, 0x1f, 0x00, 0x37, 0xba, 0x9e, 0xe5, 0x0c, 0x35, 0x43, 0xaf, 0xb5, 0xfd, 0x82, 0xe1, 0x1f,
	0x2f, 0x8d, 0x6e, 0x0c, 0x38, 0x6a, 0xda, 0x62, 0x30, 0xcb, 0x19, 0x7c, 0x02, 0x9a, 0x35, 0x8a,
	0x24, 0x08, 0x58, 0xd2, 0x87, 0x55, 0x19, 0xdc, 0xbf, 0x99, 0x44, 0x12, 0x1c, 0xb9, 0x04, 0x07,
	0x24, 0xb9, 0x1a, 0xe9, 0x54, 0xc8, 0x14, 0xed, 0xbe, 0xc4, 0x48, 0xc6, 0x30, 0x1f, 0xe9, 0x89,
	0x90, 0x29, 0xfc, 0x0e, 0xec, 0x3a, 0x54, 0x22, 0x45, 0x91, 0xa3, 0x3d, 0xcb, 0x7a, 0x50, 0x95,
	0xc1, 0x07, 0x37, 0xb2, 0xac, 0x03, 0x47, 0x6e, 0x94, 0xa7, 0xa6, 0x80, 0x4f, 0xeb, 0xbc, 0x6d,
	0x54, 0xaf, 0xae, 0x39, 0xe7, 0x4f, 0x3f, 0xb9, 0x1e, 0xba, 0x8b, 0xca, 0x86, 0x6e, 0x93, 0xfa,
	0x02, 0x80, 0x98, 0xe7, 0x63, 0x26, 0x35, 0x3b, 0xd3, 0xe8, 0x35, 0x4b, 0x7a, 0xb7, 0x2a, 0x83,
	0xbb, 0x4b, 0xa7, 0x7d, 0xa9, 0xc1, 0xd1, 0x82, 0x01, 0x7e, 0x09, 0x76, 0x25, 0x9b, 0x4a, 0xae,
	0xd9, 0x70, 0x3c, 0x51, 0xe8, 0xf5, 0xd0, 0x6b, 0xed, 0xf5, 0xfc, 0xaa, 0x0c, 0x0e, 0xae, 0xfb,
	0x6b, 0x51, 0x38, 0x9e, 0x28, 0x1c, 0x81, 0xba, 0xfa, 0x66, 0xa2, 0x60, 0x1f, 0x34, 0x46, 0xee,
	0x67, 0x8b, 0xde, 0xb0, 0x91, 0x1c, 0x55, 0x65, 0xf0, 0x60, 0x4d, 0x24, 0xa3, 0x42, 0xf1, 0x8c,
	0x29, 0x15, 0x3a, 0x17, 0x8e, 0xe6, 0x7e, 0xd8, 0x02, 0x5b, 0x92, 0x4c, 0xd1, 0x5b, 0x76, 0x86,
	0xfd, 0xaa, 0x0c, 0xe0, 0xd2, 0x0c, 0x64, 0x8a, 0x23, 0x23, 0x81, 0x8f, 0x41, 0x23, 0x65, 0x4a,
	0x91, 0x84, 0xa1, 0x7d, 0xab, 0x3e, 0xa8, 0xca, 0x60, 0xff, 0xba, 0xba, 0x16, 0xe0, 0x68, 0x2e,
	0x35, 0x51, 0x99, 0x4e, 0x5c, 0x69, 0x1e, 0x2b, 0xf4, 0xb6, 0x35, 0xae, 0x88, 0xea, 0x4a, 0x83,
	0xa3, 0x05, 0x03, 0xfc, 0xdd, 0x03, 0x3b, 0xf5, 0xa6, 0xc8, 0x6e, 0x9a, 0x55, 0x65, 0xf0, 0xf3,
	0xea, 0xeb, 0xad, 0x7d, 0xfb, 0xeb, 0x8d, 0xb2, 0x53, 0x52, 0x4c, 0x74, 0xa7, 0xed, 0x82, 0x3a,
	0xee, 0x2e, 0xce, 0x51, 0xa8, 0xb0, 0x1d, 0x66, 0x42, 0xa6, 0x64, 0x82, 0xa3, 0xba, 0x3b, 0xec,
	0x01, 0x10, 0x4b, 0x46, 0x1c, 0x02, 0xdd, 0xb5, 0x57, 0xed, 0xbd, 0xaa, 0x0c, 0x82, 0xc5, 0xd4,
	0x2f, 0x0f, 0xde, 0x29, 0x43, 0xa2, 0x71, 0xd4, 0xac, 0x8b, 0x13, 0x6d, 0x18, 0x45, 0x4e, 0xe7,
	0x8c, 0x83, 0xf5, 0x8c, 0x5a, 0xe9, 0x18, 0x75, 0x71, 0xa2, 0xe1, 0x14, 0x80, 0xab, 0x55, 0xd0,
	0x3b, 0x96, 0xf1, 0x63, 0x55, 0x06, 0x83, 0xd5, 0x99, 0x1c, 0xdf, 0x3e, 0x93, 0x79, 0xe3, 0xfa,
	0xad, 0x6b, 0x5c, 0x17, 0x27, 0xba, 0x77, 0xef, 0xbf, 0x7f, 0x7c, 0xef, 0xcf, 0x73, 0xdf, 0x7b,
	0x7e, 0xee, 0x7b, 0x7f, 0x9f, 0xfb, 0xde, 0x1f, 0x17, 0xfe, 0xc6, 0xf3, 0x0b, 0x7f, 0xe3, 0xaf,
	0x0b, 0x7f, 0xe3, 0xa7, 0xed, 0x54, 0x50, 0x36, 0x19, 0xed, 0xd8, 0x3f, 0xa6, 0x47, 0xff, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xa3, 0x30, 0xa2, 0x1e, 0xd5, 0x06, 0x00, 0x00,
}

func (m *Task) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Task) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Task) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DeletedAt != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.DeletedAt))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xd8
	}
	if m.UpdatedAt != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xd0
	}
	if m.CreatedAt != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xc8
	}
	if m.Status != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xc0
	}
	if len(m.Statistics) > 0 {
		i -= len(m.Statistics)
		copy(dAtA[i:], m.Statistics)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Statistics)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xba
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb2
	}
	if len(m.Raw) > 0 {
		i -= len(m.Raw)
		copy(dAtA[i:], m.Raw)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Raw)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	if m.Bstatus != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.Bstatus))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if len(m.RewriteHls) > 0 {
		i -= len(m.RewriteHls)
		copy(dAtA[i:], m.RewriteHls)
		i = encodeVarintTask(dAtA, i, uint64(len(m.RewriteHls)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.Ciphertext) > 0 {
		i -= len(m.Ciphertext)
		copy(dAtA[i:], m.Ciphertext)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Ciphertext)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.TaskType) > 0 {
		i -= len(m.TaskType)
		copy(dAtA[i:], m.TaskType)
		i = encodeVarintTask(dAtA, i, uint64(len(m.TaskType)))
		i--
		dAtA[i] = 0x6a
	}
	if m.OtherGroup != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.OtherGroup))
		i--
		dAtA[i] = 0x60
	}
	if m.OtherForm != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.OtherForm))
		i--
		dAtA[i] = 0x58
	}
	if m.OtherTag != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.OtherTag))
		i--
		dAtA[i] = 0x50
	}
	if m.OtherType != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.OtherType))
		i--
		dAtA[i] = 0x48
	}
	if len(m.TaskName) > 0 {
		i -= len(m.TaskName)
		copy(dAtA[i:], m.TaskName)
		i = encodeVarintTask(dAtA, i, uint64(len(m.TaskName)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Cover) > 0 {
		i -= len(m.Cover)
		copy(dAtA[i:], m.Cover)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Cover)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.OtherId) > 0 {
		i -= len(m.OtherId)
		copy(dAtA[i:], m.OtherId)
		i = encodeVarintTask(dAtA, i, uint64(len(m.OtherId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TaskPid) > 0 {
		i -= len(m.TaskPid)
		copy(dAtA[i:], m.TaskPid)
		i = encodeVarintTask(dAtA, i, uint64(len(m.TaskPid)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TaskId) > 0 {
		i -= len(m.TaskId)
		copy(dAtA[i:], m.TaskId)
		i = encodeVarintTask(dAtA, i, uint64(len(m.TaskId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTask(dAtA []byte, offset int, v uint64) int {
	offset -= sovTask(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedTask(r randyTask, easy bool) *Task {
	this := &Task{}
	this.Id = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Id *= -1
	}
	this.TaskId = string(randStringTask(r))
	this.TaskPid = string(randStringTask(r))
	this.Uid = string(randStringTask(r))
	this.OtherId = string(randStringTask(r))
	this.Cover = string(randStringTask(r))
	this.TaskName = string(randStringTask(r))
	this.OtherType = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.OtherType *= -1
	}
	this.OtherTag = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.OtherTag *= -1
	}
	this.OtherForm = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.OtherForm *= -1
	}
	this.OtherGroup = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.OtherGroup *= -1
	}
	this.TaskType = string(randStringTask(r))
	this.Ciphertext = string(randStringTask(r))
	v1 := r.Intn(100)
	this.RewriteHls = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.RewriteHls[i] = byte(r.Intn(256))
	}
	this.Bstatus = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Bstatus *= -1
	}
	v2 := r.Intn(100)
	this.Raw = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Raw[i] = byte(r.Intn(256))
	}
	v3 := r.Intn(100)
	this.Message = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.Message[i] = byte(r.Intn(256))
	}
	v4 := r.Intn(100)
	this.Statistics = make([]byte, v4)
	for i := 0; i < v4; i++ {
		this.Statistics[i] = byte(r.Intn(256))
	}
	this.Status = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Status *= -1
	}
	this.CreatedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.CreatedAt *= -1
	}
	this.UpdatedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.UpdatedAt *= -1
	}
	this.DeletedAt = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.DeletedAt *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTask interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTask(r randyTask) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTask(r randyTask) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneTask(r)
	}
	return string(tmps)
}
func randUnrecognizedTask(r randyTask, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTask(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTask(dAtA []byte, r randyTask, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTask(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateTask(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateTask(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTask(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTask(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTask(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTask(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Task) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTask(uint64(m.Id))
	}
	l = len(m.TaskId)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.TaskPid)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.OtherId)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Cover)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.TaskName)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.OtherType != 0 {
		n += 1 + sovTask(uint64(m.OtherType))
	}
	if m.OtherTag != 0 {
		n += 1 + sovTask(uint64(m.OtherTag))
	}
	if m.OtherForm != 0 {
		n += 1 + sovTask(uint64(m.OtherForm))
	}
	if m.OtherGroup != 0 {
		n += 1 + sovTask(uint64(m.OtherGroup))
	}
	l = len(m.TaskType)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Ciphertext)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.RewriteHls)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Bstatus != 0 {
		n += 2 + sovTask(uint64(m.Bstatus))
	}
	l = len(m.Raw)
	if l > 0 {
		n += 2 + l + sovTask(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 2 + l + sovTask(uint64(l))
	}
	l = len(m.Statistics)
	if l > 0 {
		n += 2 + l + sovTask(uint64(l))
	}
	if m.Status != 0 {
		n += 2 + sovTask(uint64(m.Status))
	}
	if m.CreatedAt != 0 {
		n += 2 + sovTask(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 2 + sovTask(uint64(m.UpdatedAt))
	}
	if m.DeletedAt != 0 {
		n += 2 + sovTask(uint64(m.DeletedAt))
	}
	return n
}

func sovTask(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTask(x uint64) (n int) {
	return sovTask(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Task) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Task: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Task: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaskId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskPid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaskPid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OtherId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cover", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cover = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaskName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherType", wireType)
			}
			m.OtherType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OtherType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherTag", wireType)
			}
			m.OtherTag = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OtherTag |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherForm", wireType)
			}
			m.OtherForm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OtherForm |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherGroup", wireType)
			}
			m.OtherGroup = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OtherGroup |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaskType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ciphertext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ciphertext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewriteHls", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewriteHls = append(m.RewriteHls[:0], dAtA[iNdEx:postIndex]...)
			if m.RewriteHls == nil {
				m.RewriteHls = []byte{}
			}
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bstatus", wireType)
			}
			m.Bstatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bstatus |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Raw = append(m.Raw[:0], dAtA[iNdEx:postIndex]...)
			if m.Raw == nil {
				m.Raw = []byte{}
			}
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message[:0], dAtA[iNdEx:postIndex]...)
			if m.Message == nil {
				m.Message = []byte{}
			}
			iNdEx = postIndex
		case 23:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statistics", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statistics = append(m.Statistics[:0], dAtA[iNdEx:postIndex]...)
			if m.Statistics == nil {
				m.Statistics = []byte{}
			}
			iNdEx = postIndex
		case 24:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 25:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 26:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 27:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedAt", wireType)
			}
			m.DeletedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeletedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTask(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTask
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTask
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTask
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTask
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTask
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTask
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTask        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTask          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTask = fmt.Errorf("proto: unexpected end of group")
)

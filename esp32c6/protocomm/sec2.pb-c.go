package protocomm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type S2SessionCmd0 struct {
	Unused [8]uint8
}

type S2SessionResp0 struct {
	Unused [8]uint8
}

type S2SessionCmd1 struct {
	Unused [8]uint8
}

type S2SessionResp1 struct {
	Unused [8]uint8
}

type Sec2Payload struct {
	Unused [8]uint8
}
type X_Sec2MsgType c.Int

const (
	SEC2_MSG_TYPE__S2Session_Command0  X_Sec2MsgType = 0
	SEC2_MSG_TYPE__S2Session_Response0 X_Sec2MsgType = 1
	SEC2_MSG_TYPE__S2Session_Command1  X_Sec2MsgType = 2
	SEC2_MSG_TYPE__S2Session_Response1 X_Sec2MsgType = 3
	X_SEC2_MSG_TYPE_IS_INT_SIZE        X_Sec2MsgType = 2147483647
)

type Sec2MsgType X_Sec2MsgType
type Sec2PayloadPayloadCase c.Int

const (
	SEC2_PAYLOAD__PAYLOAD__NOT_SET            Sec2PayloadPayloadCase = 0
	SEC2_PAYLOAD__PAYLOAD_SC0                 Sec2PayloadPayloadCase = 20
	SEC2_PAYLOAD__PAYLOAD_SR0                 Sec2PayloadPayloadCase = 21
	SEC2_PAYLOAD__PAYLOAD_SC1                 Sec2PayloadPayloadCase = 22
	SEC2_PAYLOAD__PAYLOAD_SR1                 Sec2PayloadPayloadCase = 23
	X_SEC2_PAYLOAD__PAYLOAD__CASE_IS_INT_SIZE Sec2PayloadPayloadCase = 2147483647
)

/* S2SessionCmd0 methods */
// llgo:link (*S2SessionCmd0).S2SessionCmd0Init C.s2_session_cmd0__init
func (recv_ *S2SessionCmd0) S2SessionCmd0Init() {
}

// llgo:link (*S2SessionCmd0).S2SessionCmd0GetPackedSize C.s2_session_cmd0__get_packed_size
func (recv_ *S2SessionCmd0) S2SessionCmd0GetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*S2SessionCmd0).S2SessionCmd0Pack C.s2_session_cmd0__pack
func (recv_ *S2SessionCmd0) S2SessionCmd0Pack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*S2SessionCmd0).S2SessionCmd0PackToBuffer C.s2_session_cmd0__pack_to_buffer
func (recv_ *S2SessionCmd0) S2SessionCmd0PackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname S2SessionCmd0Unpack C.s2_session_cmd0__unpack
func S2SessionCmd0Unpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *S2SessionCmd0

// llgo:link (*S2SessionCmd0).S2SessionCmd0FreeUnpacked C.s2_session_cmd0__free_unpacked
func (recv_ *S2SessionCmd0) S2SessionCmd0FreeUnpacked(allocator *ProtobufCAllocator) {
}

/* S2SessionResp0 methods */
// llgo:link (*S2SessionResp0).S2SessionResp0Init C.s2_session_resp0__init
func (recv_ *S2SessionResp0) S2SessionResp0Init() {
}

// llgo:link (*S2SessionResp0).S2SessionResp0GetPackedSize C.s2_session_resp0__get_packed_size
func (recv_ *S2SessionResp0) S2SessionResp0GetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*S2SessionResp0).S2SessionResp0Pack C.s2_session_resp0__pack
func (recv_ *S2SessionResp0) S2SessionResp0Pack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*S2SessionResp0).S2SessionResp0PackToBuffer C.s2_session_resp0__pack_to_buffer
func (recv_ *S2SessionResp0) S2SessionResp0PackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname S2SessionResp0Unpack C.s2_session_resp0__unpack
func S2SessionResp0Unpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *S2SessionResp0

// llgo:link (*S2SessionResp0).S2SessionResp0FreeUnpacked C.s2_session_resp0__free_unpacked
func (recv_ *S2SessionResp0) S2SessionResp0FreeUnpacked(allocator *ProtobufCAllocator) {
}

/* S2SessionCmd1 methods */
// llgo:link (*S2SessionCmd1).S2SessionCmd1Init C.s2_session_cmd1__init
func (recv_ *S2SessionCmd1) S2SessionCmd1Init() {
}

// llgo:link (*S2SessionCmd1).S2SessionCmd1GetPackedSize C.s2_session_cmd1__get_packed_size
func (recv_ *S2SessionCmd1) S2SessionCmd1GetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*S2SessionCmd1).S2SessionCmd1Pack C.s2_session_cmd1__pack
func (recv_ *S2SessionCmd1) S2SessionCmd1Pack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*S2SessionCmd1).S2SessionCmd1PackToBuffer C.s2_session_cmd1__pack_to_buffer
func (recv_ *S2SessionCmd1) S2SessionCmd1PackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname S2SessionCmd1Unpack C.s2_session_cmd1__unpack
func S2SessionCmd1Unpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *S2SessionCmd1

// llgo:link (*S2SessionCmd1).S2SessionCmd1FreeUnpacked C.s2_session_cmd1__free_unpacked
func (recv_ *S2SessionCmd1) S2SessionCmd1FreeUnpacked(allocator *ProtobufCAllocator) {
}

/* S2SessionResp1 methods */
// llgo:link (*S2SessionResp1).S2SessionResp1Init C.s2_session_resp1__init
func (recv_ *S2SessionResp1) S2SessionResp1Init() {
}

// llgo:link (*S2SessionResp1).S2SessionResp1GetPackedSize C.s2_session_resp1__get_packed_size
func (recv_ *S2SessionResp1) S2SessionResp1GetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*S2SessionResp1).S2SessionResp1Pack C.s2_session_resp1__pack
func (recv_ *S2SessionResp1) S2SessionResp1Pack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*S2SessionResp1).S2SessionResp1PackToBuffer C.s2_session_resp1__pack_to_buffer
func (recv_ *S2SessionResp1) S2SessionResp1PackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname S2SessionResp1Unpack C.s2_session_resp1__unpack
func S2SessionResp1Unpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *S2SessionResp1

// llgo:link (*S2SessionResp1).S2SessionResp1FreeUnpacked C.s2_session_resp1__free_unpacked
func (recv_ *S2SessionResp1) S2SessionResp1FreeUnpacked(allocator *ProtobufCAllocator) {
}

/* Sec2Payload methods */
// llgo:link (*Sec2Payload).Sec2PayloadInit C.sec2_payload__init
func (recv_ *Sec2Payload) Sec2PayloadInit() {
}

// llgo:link (*Sec2Payload).Sec2PayloadGetPackedSize C.sec2_payload__get_packed_size
func (recv_ *Sec2Payload) Sec2PayloadGetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*Sec2Payload).Sec2PayloadPack C.sec2_payload__pack
func (recv_ *Sec2Payload) Sec2PayloadPack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*Sec2Payload).Sec2PayloadPackToBuffer C.sec2_payload__pack_to_buffer
func (recv_ *Sec2Payload) Sec2PayloadPackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname Sec2PayloadUnpack C.sec2_payload__unpack
func Sec2PayloadUnpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *Sec2Payload

// llgo:link (*Sec2Payload).Sec2PayloadFreeUnpacked C.sec2_payload__free_unpacked
func (recv_ *Sec2Payload) Sec2PayloadFreeUnpacked(allocator *ProtobufCAllocator) {
}

// llgo:type C
type S2SessionCmd0Closure func(*S2SessionCmd0, c.Pointer)

// llgo:type C
type S2SessionResp0Closure func(*S2SessionResp0, c.Pointer)

// llgo:type C
type S2SessionCmd1Closure func(*S2SessionCmd1, c.Pointer)

// llgo:type C
type S2SessionResp1Closure func(*S2SessionResp1, c.Pointer)

// llgo:type C
type Sec2PayloadClosure func(*Sec2Payload, c.Pointer)

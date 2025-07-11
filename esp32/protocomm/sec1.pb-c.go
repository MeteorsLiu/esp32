package protocomm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SessionCmd1 struct {
	Unused [8]uint8
}

type SessionResp1 struct {
	Unused [8]uint8
}

type SessionCmd0 struct {
	Unused [8]uint8
}

type SessionResp0 struct {
	Unused [8]uint8
}

type Sec1Payload struct {
	Unused [8]uint8
}
type X_Sec1MsgType c.Int

const (
	SEC1_MSG_TYPE__Session_Command0  X_Sec1MsgType = 0
	SEC1_MSG_TYPE__Session_Response0 X_Sec1MsgType = 1
	SEC1_MSG_TYPE__Session_Command1  X_Sec1MsgType = 2
	SEC1_MSG_TYPE__Session_Response1 X_Sec1MsgType = 3
	X_SEC1_MSG_TYPE_IS_INT_SIZE      X_Sec1MsgType = 2147483647
)

type Sec1MsgType X_Sec1MsgType
type Sec1PayloadPayloadCase c.Int

const (
	SEC1_PAYLOAD__PAYLOAD__NOT_SET            Sec1PayloadPayloadCase = 0
	SEC1_PAYLOAD__PAYLOAD_SC0                 Sec1PayloadPayloadCase = 20
	SEC1_PAYLOAD__PAYLOAD_SR0                 Sec1PayloadPayloadCase = 21
	SEC1_PAYLOAD__PAYLOAD_SC1                 Sec1PayloadPayloadCase = 22
	SEC1_PAYLOAD__PAYLOAD_SR1                 Sec1PayloadPayloadCase = 23
	X_SEC1_PAYLOAD__PAYLOAD__CASE_IS_INT_SIZE Sec1PayloadPayloadCase = 2147483647
)

/* SessionCmd1 methods */
// llgo:link (*SessionCmd1).SessionCmd1Init C.session_cmd1__init
func (recv_ *SessionCmd1) SessionCmd1Init() {
}

// llgo:link (*SessionCmd1).SessionCmd1GetPackedSize C.session_cmd1__get_packed_size
func (recv_ *SessionCmd1) SessionCmd1GetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*SessionCmd1).SessionCmd1Pack C.session_cmd1__pack
func (recv_ *SessionCmd1) SessionCmd1Pack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*SessionCmd1).SessionCmd1PackToBuffer C.session_cmd1__pack_to_buffer
func (recv_ *SessionCmd1) SessionCmd1PackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname SessionCmd1Unpack C.session_cmd1__unpack
func SessionCmd1Unpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *SessionCmd1

// llgo:link (*SessionCmd1).SessionCmd1FreeUnpacked C.session_cmd1__free_unpacked
func (recv_ *SessionCmd1) SessionCmd1FreeUnpacked(allocator *ProtobufCAllocator) {
}

/* SessionResp1 methods */
// llgo:link (*SessionResp1).SessionResp1Init C.session_resp1__init
func (recv_ *SessionResp1) SessionResp1Init() {
}

// llgo:link (*SessionResp1).SessionResp1GetPackedSize C.session_resp1__get_packed_size
func (recv_ *SessionResp1) SessionResp1GetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*SessionResp1).SessionResp1Pack C.session_resp1__pack
func (recv_ *SessionResp1) SessionResp1Pack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*SessionResp1).SessionResp1PackToBuffer C.session_resp1__pack_to_buffer
func (recv_ *SessionResp1) SessionResp1PackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname SessionResp1Unpack C.session_resp1__unpack
func SessionResp1Unpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *SessionResp1

// llgo:link (*SessionResp1).SessionResp1FreeUnpacked C.session_resp1__free_unpacked
func (recv_ *SessionResp1) SessionResp1FreeUnpacked(allocator *ProtobufCAllocator) {
}

/* SessionCmd0 methods */
// llgo:link (*SessionCmd0).SessionCmd0Init C.session_cmd0__init
func (recv_ *SessionCmd0) SessionCmd0Init() {
}

// llgo:link (*SessionCmd0).SessionCmd0GetPackedSize C.session_cmd0__get_packed_size
func (recv_ *SessionCmd0) SessionCmd0GetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*SessionCmd0).SessionCmd0Pack C.session_cmd0__pack
func (recv_ *SessionCmd0) SessionCmd0Pack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*SessionCmd0).SessionCmd0PackToBuffer C.session_cmd0__pack_to_buffer
func (recv_ *SessionCmd0) SessionCmd0PackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname SessionCmd0Unpack C.session_cmd0__unpack
func SessionCmd0Unpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *SessionCmd0

// llgo:link (*SessionCmd0).SessionCmd0FreeUnpacked C.session_cmd0__free_unpacked
func (recv_ *SessionCmd0) SessionCmd0FreeUnpacked(allocator *ProtobufCAllocator) {
}

/* SessionResp0 methods */
// llgo:link (*SessionResp0).SessionResp0Init C.session_resp0__init
func (recv_ *SessionResp0) SessionResp0Init() {
}

// llgo:link (*SessionResp0).SessionResp0GetPackedSize C.session_resp0__get_packed_size
func (recv_ *SessionResp0) SessionResp0GetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*SessionResp0).SessionResp0Pack C.session_resp0__pack
func (recv_ *SessionResp0) SessionResp0Pack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*SessionResp0).SessionResp0PackToBuffer C.session_resp0__pack_to_buffer
func (recv_ *SessionResp0) SessionResp0PackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname SessionResp0Unpack C.session_resp0__unpack
func SessionResp0Unpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *SessionResp0

// llgo:link (*SessionResp0).SessionResp0FreeUnpacked C.session_resp0__free_unpacked
func (recv_ *SessionResp0) SessionResp0FreeUnpacked(allocator *ProtobufCAllocator) {
}

/* Sec1Payload methods */
// llgo:link (*Sec1Payload).Sec1PayloadInit C.sec1_payload__init
func (recv_ *Sec1Payload) Sec1PayloadInit() {
}

// llgo:link (*Sec1Payload).Sec1PayloadGetPackedSize C.sec1_payload__get_packed_size
func (recv_ *Sec1Payload) Sec1PayloadGetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*Sec1Payload).Sec1PayloadPack C.sec1_payload__pack
func (recv_ *Sec1Payload) Sec1PayloadPack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*Sec1Payload).Sec1PayloadPackToBuffer C.sec1_payload__pack_to_buffer
func (recv_ *Sec1Payload) Sec1PayloadPackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname Sec1PayloadUnpack C.sec1_payload__unpack
func Sec1PayloadUnpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *Sec1Payload

// llgo:link (*Sec1Payload).Sec1PayloadFreeUnpacked C.sec1_payload__free_unpacked
func (recv_ *Sec1Payload) Sec1PayloadFreeUnpacked(allocator *ProtobufCAllocator) {
}

// llgo:type C
type SessionCmd1Closure func(*SessionCmd1, c.Pointer)

// llgo:type C
type SessionResp1Closure func(*SessionResp1, c.Pointer)

// llgo:type C
type SessionCmd0Closure func(*SessionCmd0, c.Pointer)

// llgo:type C
type SessionResp0Closure func(*SessionResp0, c.Pointer)

// llgo:type C
type Sec1PayloadClosure func(*Sec1Payload, c.Pointer)

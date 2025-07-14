package protocomm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type S0SessionCmd struct {
	Unused [8]uint8
}

type S0SessionResp struct {
	Unused [8]uint8
}

type Sec0Payload struct {
	Unused [8]uint8
}
type X_Sec0MsgType c.Int

const (
	SEC0_MSG_TYPE__S0_Session_Command  X_Sec0MsgType = 0
	SEC0_MSG_TYPE__S0_Session_Response X_Sec0MsgType = 1
	X_SEC0_MSG_TYPE_IS_INT_SIZE        X_Sec0MsgType = 2147483647
)

type Sec0MsgType X_Sec0MsgType
type Sec0PayloadPayloadCase c.Int

const (
	SEC0_PAYLOAD__PAYLOAD__NOT_SET            Sec0PayloadPayloadCase = 0
	SEC0_PAYLOAD__PAYLOAD_SC                  Sec0PayloadPayloadCase = 20
	SEC0_PAYLOAD__PAYLOAD_SR                  Sec0PayloadPayloadCase = 21
	X_SEC0_PAYLOAD__PAYLOAD__CASE_IS_INT_SIZE Sec0PayloadPayloadCase = 2147483647
)

/* S0SessionCmd methods */
// llgo:link (*S0SessionCmd).S0SessionCmdInit C.s0_session_cmd__init
func (recv_ *S0SessionCmd) S0SessionCmdInit() {
}

// llgo:link (*S0SessionCmd).S0SessionCmdGetPackedSize C.s0_session_cmd__get_packed_size
func (recv_ *S0SessionCmd) S0SessionCmdGetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*S0SessionCmd).S0SessionCmdPack C.s0_session_cmd__pack
func (recv_ *S0SessionCmd) S0SessionCmdPack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*S0SessionCmd).S0SessionCmdPackToBuffer C.s0_session_cmd__pack_to_buffer
func (recv_ *S0SessionCmd) S0SessionCmdPackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname S0SessionCmdUnpack C.s0_session_cmd__unpack
func S0SessionCmdUnpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *S0SessionCmd

// llgo:link (*S0SessionCmd).S0SessionCmdFreeUnpacked C.s0_session_cmd__free_unpacked
func (recv_ *S0SessionCmd) S0SessionCmdFreeUnpacked(allocator *ProtobufCAllocator) {
}

/* S0SessionResp methods */
// llgo:link (*S0SessionResp).S0SessionRespInit C.s0_session_resp__init
func (recv_ *S0SessionResp) S0SessionRespInit() {
}

// llgo:link (*S0SessionResp).S0SessionRespGetPackedSize C.s0_session_resp__get_packed_size
func (recv_ *S0SessionResp) S0SessionRespGetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*S0SessionResp).S0SessionRespPack C.s0_session_resp__pack
func (recv_ *S0SessionResp) S0SessionRespPack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*S0SessionResp).S0SessionRespPackToBuffer C.s0_session_resp__pack_to_buffer
func (recv_ *S0SessionResp) S0SessionRespPackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname S0SessionRespUnpack C.s0_session_resp__unpack
func S0SessionRespUnpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *S0SessionResp

// llgo:link (*S0SessionResp).S0SessionRespFreeUnpacked C.s0_session_resp__free_unpacked
func (recv_ *S0SessionResp) S0SessionRespFreeUnpacked(allocator *ProtobufCAllocator) {
}

/* Sec0Payload methods */
// llgo:link (*Sec0Payload).Sec0PayloadInit C.sec0_payload__init
func (recv_ *Sec0Payload) Sec0PayloadInit() {
}

// llgo:link (*Sec0Payload).Sec0PayloadGetPackedSize C.sec0_payload__get_packed_size
func (recv_ *Sec0Payload) Sec0PayloadGetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*Sec0Payload).Sec0PayloadPack C.sec0_payload__pack
func (recv_ *Sec0Payload) Sec0PayloadPack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*Sec0Payload).Sec0PayloadPackToBuffer C.sec0_payload__pack_to_buffer
func (recv_ *Sec0Payload) Sec0PayloadPackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname Sec0PayloadUnpack C.sec0_payload__unpack
func Sec0PayloadUnpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *Sec0Payload

// llgo:link (*Sec0Payload).Sec0PayloadFreeUnpacked C.sec0_payload__free_unpacked
func (recv_ *Sec0Payload) Sec0PayloadFreeUnpacked(allocator *ProtobufCAllocator) {
}

// llgo:type C
type S0SessionCmdClosure func(*S0SessionCmd, c.Pointer)

// llgo:type C
type S0SessionRespClosure func(*S0SessionResp, c.Pointer)

// llgo:type C
type Sec0PayloadClosure func(*Sec0Payload, c.Pointer)

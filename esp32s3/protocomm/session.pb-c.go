package protocomm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SessionData struct {
	Unused [8]uint8
}
type X_SecSchemeVersion c.Int

const (
	SEC_SCHEME_VERSION__SecScheme0   X_SecSchemeVersion = 0
	SEC_SCHEME_VERSION__SecScheme1   X_SecSchemeVersion = 1
	SEC_SCHEME_VERSION__SecScheme2   X_SecSchemeVersion = 2
	X_SEC_SCHEME_VERSION_IS_INT_SIZE X_SecSchemeVersion = 2147483647
)

type SecSchemeVersion X_SecSchemeVersion
type SessionDataProtoCase c.Int

const (
	SESSION_DATA__PROTO__NOT_SET            SessionDataProtoCase = 0
	SESSION_DATA__PROTO_SEC0                SessionDataProtoCase = 10
	SESSION_DATA__PROTO_SEC1                SessionDataProtoCase = 11
	SESSION_DATA__PROTO_SEC2                SessionDataProtoCase = 12
	X_SESSION_DATA__PROTO__CASE_IS_INT_SIZE SessionDataProtoCase = 2147483647
)

/* SessionData methods */
// llgo:link (*SessionData).SessionDataInit C.session_data__init
func (recv_ *SessionData) SessionDataInit() {
}

// llgo:link (*SessionData).SessionDataGetPackedSize C.session_data__get_packed_size
func (recv_ *SessionData) SessionDataGetPackedSize() c.SizeT {
	return 0
}

// llgo:link (*SessionData).SessionDataPack C.session_data__pack
func (recv_ *SessionData) SessionDataPack(out *c.Uint8T) c.SizeT {
	return 0
}

// llgo:link (*SessionData).SessionDataPackToBuffer C.session_data__pack_to_buffer
func (recv_ *SessionData) SessionDataPackToBuffer(buffer *ProtobufCBuffer) c.SizeT {
	return 0
}

//go:linkname SessionDataUnpack C.session_data__unpack
func SessionDataUnpack(allocator *ProtobufCAllocator, len c.SizeT, data *c.Uint8T) *SessionData

// llgo:link (*SessionData).SessionDataFreeUnpacked C.session_data__free_unpacked
func (recv_ *SessionData) SessionDataFreeUnpacked(allocator *ProtobufCAllocator) {
}

// llgo:type C
type SessionDataClosure func(*SessionData, c.Pointer)

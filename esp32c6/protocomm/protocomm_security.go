package protocomm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type ProtocommSecuritySessionEventT c.Int

const (
	PROTOCOMM_SECURITY_SESSION_SETUP_OK                ProtocommSecuritySessionEventT = 0
	PROTOCOMM_SECURITY_SESSION_INVALID_SECURITY_PARAMS ProtocommSecuritySessionEventT = 1
	PROTOCOMM_SECURITY_SESSION_CREDENTIALS_MISMATCH    ProtocommSecuritySessionEventT = 2
)

/**
 * @brief   Protocomm Security 1 parameters: Proof Of Possession
 */

type ProtocommSecurity1Params struct {
	Data *c.Uint8T
	Len  c.Uint16T
}
type ProtocommSecurity1ParamsT ProtocommSecurity1Params
type ProtocommSecurityPopT ProtocommSecurity1ParamsT

/**
 * @brief Protocomm Security 2 parameters: Salt and Verifier
 *
 */

type ProtocommSecurity2Params struct {
	Salt        *c.Char
	SaltLen     c.Uint16T
	Verifier    *c.Char
	VerifierLen c.Uint16T
}
type ProtocommSecurity2ParamsT ProtocommSecurity2Params
type ProtocommSecurityHandleT c.Pointer

/**
 * @brief   Protocomm security object structure.
 *
 * The member functions are used for implementing secure
 * protocomm sessions.
 *
 * @note    This structure should not have any dynamic
 *          members to allow re-entrancy
 */

type ProtocommSecurity struct {
	Ver                   c.Int
	PatchVer              c.Uint8T
	Init                  c.Pointer
	Cleanup               c.Pointer
	NewTransportSession   c.Pointer
	CloseTransportSession c.Pointer
	SecurityReqHandler    c.Pointer
	Encrypt               c.Pointer
	Decrypt               c.Pointer
}
type ProtocommSecurityT ProtocommSecurity

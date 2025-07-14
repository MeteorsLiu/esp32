package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Control / Configuration Registers */
/** Type of m_prime register
 *  Represents M'
 */

type RsaMPrimeRegT struct {
	Val c.Uint32T
}

/** Type of mode register
 *  Configures RSA length
 */

type RsaModeRegT struct {
	Val c.Uint32T
}

/** Type of set_start_modexp register
 *  Starts modular exponentiation
 */

type RsaSetStartModexpRegT struct {
	Val c.Uint32T
}

/** Type of set_start_modmult register
 *  Starts modular multiplication
 */

type RsaSetStartModmultRegT struct {
	Val c.Uint32T
}

/** Type of set_start_mult register
 *  Starts multiplication
 */

type RsaSetStartMultRegT struct {
	Val c.Uint32T
}

/** Type of query_idle register
 *  Represents the RSA status
 */

type RsaQueryIdleRegT struct {
	Val c.Uint32T
}

/** Type of constant_time register
 *  Configures the constant_time option
 */

type RsaConstantTimeRegT struct {
	Val c.Uint32T
}

/** Type of search_enable register
 *  Configures the search option
 */

type RsaSearchEnableRegT struct {
	Val c.Uint32T
}

/** Type of search_pos register
 *  Configures the search position
 */

type RsaSearchPosRegT struct {
	Val c.Uint32T
}

/** Group: Status Register */
/** Type of query_clean register
 *  RSA clean register
 */

type RsaQueryCleanRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Registers */
/** Type of int_clr register
 *  Clears RSA interrupt
 */

type RsaIntClrRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  Enables the RSA interrupt
 */

type RsaIntEnaRegT struct {
	Val c.Uint32T
}

/** Group: Version Control Register */
/** Type of date register
 *  Version control register
 */

type RsaDateRegT struct {
	Val c.Uint32T
}

type RsaDevT struct {
	M               [4]c.Uint32T
	Reserved010     [124]c.Uint32T
	Z               [4]c.Uint32T
	Reserved210     [124]c.Uint32T
	Y               [4]c.Uint32T
	Reserved410     [124]c.Uint32T
	X               [4]c.Uint32T
	Reserved610     [124]c.Uint32T
	MPrime          RsaMPrimeRegT
	Mode            RsaModeRegT
	QueryClean      RsaQueryCleanRegT
	SetStartModexp  RsaSetStartModexpRegT
	SetStartModmult RsaSetStartModmultRegT
	SetStartMult    RsaSetStartMultRegT
	QueryIdle       RsaQueryIdleRegT
	IntClr          RsaIntClrRegT
	ConstantTime    RsaConstantTimeRegT
	SearchEnable    RsaSearchEnableRegT
	SearchPos       RsaSearchPosRegT
	IntEna          RsaIntEnaRegT
	Date            RsaDateRegT
}

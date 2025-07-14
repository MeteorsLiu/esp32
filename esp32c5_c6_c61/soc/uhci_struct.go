package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Register */
/** Type of conf0 register
 *  a
 */

type UhciConf0RegT struct {
	Val c.Uint32T
}

/** Type of conf1 register
 *  a
 */

type UhciConf1RegT struct {
	Val c.Uint32T
}

/** Type of escape_conf register
 *  a
 */

type UhciEscapeConfRegT struct {
	Val c.Uint32T
}

/** Type of hung_conf register
 *  a
 */

type UhciHungConfRegT struct {
	Val c.Uint32T
}

/** Type of ack_num register
 *  a
 */

type UhciAckNumRegT struct {
	Val c.Uint32T
}

/** Type of quick_sent register
 *  a
 */

type UhciQuickSentRegT struct {
	Val c.Uint32T
}

/** Type of reg_qn_word0 register
 *  a
 */

type UhciRegQnWord0RegT struct {
	Val c.Uint32T
}

/** Type of reg_qn_word1 register
 *  a
 */

type UhciRegQnWord1RegT struct {
	Val c.Uint32T
}

/** Type of esc_conf0 register
 *  a
 */

type UhciEscConf0RegT struct {
	Val c.Uint32T
}

/** Type of esc_conf1 register
 *  a
 */

type UhciEscConf1RegT struct {
	Val c.Uint32T
}

/** Type of esc_conf2 register
 *  a
 */

type UhciEscConf2RegT struct {
	Val c.Uint32T
}

/** Type of esc_conf3 register
 *  a
 */

type UhciEscConf3RegT struct {
	Val c.Uint32T
}

/** Type of pkt_thres register
 *  a
 */

type UhciPktThresRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Register */
/** Type of int_raw register
 *  a
 */

type UhciIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  a
 */

type UhciIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  a
 */

type UhciIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  a
 */

type UhciIntClrRegT struct {
	Val c.Uint32T
}

/** Group: UHCI Status Register */
/** Type of state0 register
 *  a
 */

type UhciState0RegT struct {
	Val c.Uint32T
}

/** Type of state1 register
 *  a
 */

type UhciState1RegT struct {
	Val c.Uint32T
}

/** Type of rx_head register
 *  a
 */

type UhciRxHeadRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  a
 */

type UhciDateRegT struct {
	Val c.Uint32T
}

type UhciDevS struct {
	Conf0      UhciConf0RegT
	IntRaw     UhciIntRawRegT
	IntSt      UhciIntStRegT
	IntEna     UhciIntEnaRegT
	IntClr     UhciIntClrRegT
	Conf1      UhciConf1RegT
	State0     UhciState0RegT
	State1     UhciState1RegT
	EscapeConf UhciEscapeConfRegT
	HungConf   UhciHungConfRegT
	AckNum     UhciAckNumRegT
	RxHead     UhciRxHeadRegT
	QuickSent  UhciQuickSentRegT
	QData      [7]struct {
		Word0 UhciRegQnWord0RegT
		Word1 UhciRegQnWord1RegT
	}
	EscConf0 UhciEscConf0RegT
	EscConf1 UhciEscConf1RegT
	EscConf2 UhciEscConf2RegT
	EscConf3 UhciEscConf3RegT
	PktThres UhciPktThresRegT
	Date     UhciDateRegT
}
type UhciDevT UhciDevS

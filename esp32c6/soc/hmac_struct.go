package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Register */
/** Type of set_start register
 *  Process control register 0.
 */

type HmacSetStartRegT struct {
	Val c.Uint32T
}

/** Type of set_para_purpose register
 *  Configure purpose.
 */

type HmacSetParaPurposeRegT struct {
	Val c.Uint32T
}

/** Type of set_para_key register
 *  Configure key.
 */

type HmacSetParaKeyRegT struct {
	Val c.Uint32T
}

/** Type of set_para_finish register
 *  Finish initial configuration.
 */

type HmacSetParaFinishRegT struct {
	Val c.Uint32T
}

/** Type of set_message_one register
 *  Process control register 1.
 */

type HmacSetMessageOneRegT struct {
	Val c.Uint32T
}

/** Type of set_message_ing register
 *  Process control register 2.
 */

type HmacSetMessageIngRegT struct {
	Val c.Uint32T
}

/** Type of set_message_end register
 *  Process control register 3.
 */

type HmacSetMessageEndRegT struct {
	Val c.Uint32T
}

/** Type of set_result_finish register
 *  Process control register 4.
 */

type HmacSetResultFinishRegT struct {
	Val c.Uint32T
}

/** Type of set_invalidate_jtag register
 *  Invalidate register 0.
 */

type HmacSetInvalidateJtagRegT struct {
	Val c.Uint32T
}

/** Type of set_invalidate_ds register
 *  Invalidate register 1.
 */

type HmacSetInvalidateDsRegT struct {
	Val c.Uint32T
}

/** Type of set_message_pad register
 *  Process control register 5.
 */

type HmacSetMessagePadRegT struct {
	Val c.Uint32T
}

/** Type of one_block register
 *  Process control register 6.
 */

type HmacOneBlockRegT struct {
	Val c.Uint32T
}

/** Type of soft_jtag_ctrl register
 *  Jtag register 0.
 */

type HmacSoftJtagCtrlRegT struct {
	Val c.Uint32T
}

/** Type of wr_jtag register
 *  Jtag register 1.
 */

type HmacWrJtagRegT struct {
	Val c.Uint32T
}

/** Group: Status Register */
/** Type of query_error register
 *  Error register.
 */

type HmacQueryErrorRegT struct {
	Val c.Uint32T
}

/** Type of query_busy register
 *  Busy register.
 */

type HmacQueryBusyRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  Date register.
 */

type HmacDateRegT struct {
	Val c.Uint32T
}

type HmacDevT struct {
	Reserved000       [16]c.Uint32T
	SetStart          HmacSetStartRegT
	SetParaPurpose    HmacSetParaPurposeRegT
	SetParaKey        HmacSetParaKeyRegT
	SetParaFinish     HmacSetParaFinishRegT
	SetMessageOne     HmacSetMessageOneRegT
	SetMessageIng     HmacSetMessageIngRegT
	SetMessageEnd     HmacSetMessageEndRegT
	SetResultFinish   HmacSetResultFinishRegT
	SetInvalidateJtag HmacSetInvalidateJtagRegT
	SetInvalidateDs   HmacSetInvalidateDsRegT
	QueryError        HmacQueryErrorRegT
	QueryBusy         HmacQueryBusyRegT
	Reserved070       [4]c.Uint32T
	WrMessage         [16]c.Uint32T
	RdResult          [8]c.Uint32T
	Reserved0e0       [4]c.Uint32T
	SetMessagePad     HmacSetMessagePadRegT
	OneBlock          HmacOneBlockRegT
	SoftJtagCtrl      HmacSoftJtagCtrlRegT
	WrJtag            HmacWrJtagRegT
	Reserved100       [63]c.Uint32T
	Date              HmacDateRegT
}

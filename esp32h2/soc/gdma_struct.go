package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Interrupt Registers */
/** Type of in_int_raw_chn register
 *  Raw status interrupt of channel 0
 */

type GdmaInIntRawChnRegT struct {
	Val c.Uint32T
}

/** Type of in_int_st_chn register
 *  Masked interrupt of channel 0
 */

type GdmaInIntStChnRegT struct {
	Val c.Uint32T
}

/** Type of in_int_ena_chn register
 *  Interrupt enable bits of channel 0
 */

type GdmaInIntEnaChnRegT struct {
	Val c.Uint32T
}

/** Type of in_int_clr_chn register
 *  Interrupt clear bits of channel 0
 */

type GdmaInIntClrChnRegT struct {
	Val c.Uint32T
}

/** Type of out_int_raw_chn register
 *  Raw status interrupt of channel 0
 */

type GdmaOutIntRawChnRegT struct {
	Val c.Uint32T
}

/** Type of out_int_st_chn register
 *  Masked interrupt of channel 0
 */

type GdmaOutIntStChnRegT struct {
	Val c.Uint32T
}

/** Type of out_int_ena_chn register
 *  Interrupt enable bits of channel 0
 */

type GdmaOutIntEnaChnRegT struct {
	Val c.Uint32T
}

/** Type of out_int_clr_chn register
 *  Interrupt clear bits of channel 0
 */

type GdmaOutIntClrChnRegT struct {
	Val c.Uint32T
}

/** Group: Debug Registers */
/** Type of ahb_test register
 *  reserved
 */

type GdmaAhbTestRegT struct {
	Val c.Uint32T
}

/** Group: Configuration Registers */
/** Type of misc_conf register
 *  MISC register
 */

type GdmaMiscConfRegT struct {
	Val c.Uint32T
}

/** Type of in_conf0_chn register
 *  Configure 0 register of Rx channel 0
 */

type GdmaInConf0ChnRegT struct {
	Val c.Uint32T
}

/** Type of in_conf1_chn register
 *  Configure 1 register of Rx channel 0
 */

type GdmaInConf1ChnRegT struct {
	Val c.Uint32T
}

/** Type of in_pop_chn register
 *  Pop control register of Rx channel 0
 */

type GdmaInPopChnRegT struct {
	Val c.Uint32T
}

/** Type of in_link_chn register
 *  Link descriptor configure and control register of Rx channel 0
 */

type GdmaInLinkChnRegT struct {
	Val c.Uint32T
}

/** Type of out_conf0_chn register
 *  Configure 0 register of Tx channel 0
 */

type GdmaOutConf0ChnRegT struct {
	Val c.Uint32T
}

/** Type of out_conf1_chn register
 *  Configure 1 register of Tx channel 0
 */

type GdmaOutConf1ChnRegT struct {
	Val c.Uint32T
}

/** Type of out_push_chn register
 *  Push control register of Rx channel 0
 */

type GdmaOutPushChnRegT struct {
	Val c.Uint32T
}

/** Type of out_link_chn register
 *  Link descriptor configure and control register of Tx channel 0
 */

type GdmaOutLinkChnRegT struct {
	Val c.Uint32T
}

/** Group: Version Registers */
/** Type of date register
 *  Version control register
 */

type GdmaDateRegT struct {
	Val c.Uint32T
}

/** Group: Status Registers */
/** Type of infifo_status_chn register
 *  Receive FIFO status of Rx channel 0
 */

type GdmaInfifoStatusChnRegT struct {
	Val c.Uint32T
}

/** Type of in_state_chn register
 *  Receive status of Rx channel 0
 */

type GdmaInStateChnRegT struct {
	Val c.Uint32T
}

/** Type of in_suc_eof_des_addr_chn register
 *  Inlink descriptor address when EOF occurs of Rx channel 0
 */

type GdmaInSucEofDesAddrChnRegT struct {
	Val c.Uint32T
}

/** Type of in_err_eof_des_addr_chn register
 *  Inlink descriptor address when errors occur of Rx channel 0
 */

type GdmaInErrEofDesAddrChnRegT struct {
	Val c.Uint32T
}

/** Type of in_dscr_chn register
 *  Current inlink descriptor address of Rx channel 0
 */

type GdmaInDscrChnRegT struct {
	Val c.Uint32T
}

/** Type of in_dscr_bf0_chn register
 *  The last inlink descriptor address of Rx channel 0
 */

type GdmaInDscrBf0ChnRegT struct {
	Val c.Uint32T
}

/** Type of in_dscr_bf1_chn register
 *  The second-to-last inlink descriptor address of Rx channel 0
 */

type GdmaInDscrBf1ChnRegT struct {
	Val c.Uint32T
}

/** Type of outfifo_status_chn register
 *  Transmit FIFO status of Tx channel 0
 */

type GdmaOutfifoStatusChnRegT struct {
	Val c.Uint32T
}

/** Type of out_state_chn register
 *  Transmit status of Tx channel 0
 */

type GdmaOutStateChnRegT struct {
	Val c.Uint32T
}

/** Type of out_eof_des_addr_chn register
 *  Outlink descriptor address when EOF occurs of Tx channel 0
 */

type GdmaOutEofDesAddrChnRegT struct {
	Val c.Uint32T
}

/** Type of out_eof_bfr_des_addr_chn register
 *  The last outlink descriptor address when EOF occurs of Tx channel 0
 */

type GdmaOutEofBfrDesAddrChnRegT struct {
	Val c.Uint32T
}

/** Type of out_dscr_chn register
 *  Current inlink descriptor address of Tx channel 0
 */

type GdmaOutDscrChnRegT struct {
	Val c.Uint32T
}

/** Type of out_dscr_bf0_chn register
 *  The last inlink descriptor address of Tx channel 0
 */

type GdmaOutDscrBf0ChnRegT struct {
	Val c.Uint32T
}

/** Type of out_dscr_bf1_chn register
 *  The second-to-last inlink descriptor address of Tx channel 0
 */

type GdmaOutDscrBf1ChnRegT struct {
	Val c.Uint32T
}

/** Group: Priority Registers */
/** Type of in_pri_chn register
 *  Priority register of Rx channel 0
 */

type GdmaInPriChnRegT struct {
	Val c.Uint32T
}

/** Type of out_pri_chn register
 *  Priority register of Tx channel 0.
 */

type GdmaOutPriChnRegT struct {
	Val c.Uint32T
}

/** Group: Peripheral Select Registers */
/** Type of in_peri_sel_chn register
 *  Peripheral selection of Rx channel 0
 */

type GdmaInPeriSelChnRegT struct {
	Val c.Uint32T
}

/** Type of out_peri_sel_chn register
 *  Peripheral selection of Tx channel 0
 */

type GdmaOutPeriSelChnRegT struct {
	Val c.Uint32T
}

type GdmaInIntChnRegT struct {
	Raw GdmaInIntRawChnRegT
	St  GdmaInIntStChnRegT
	Ena GdmaInIntEnaChnRegT
	Clr GdmaInIntClrChnRegT
}

type GdmaOutIntChnRegT struct {
	Raw GdmaOutIntRawChnRegT
	St  GdmaOutIntStChnRegT
	Ena GdmaOutIntEnaChnRegT
	Clr GdmaOutIntClrChnRegT
}

type GdmaInChnRegT struct {
	InConf0         GdmaInConf0ChnRegT
	InConf1         GdmaInConf1ChnRegT
	InfifoStatus    GdmaInfifoStatusChnRegT
	InPop           GdmaInPopChnRegT
	InLink          GdmaInLinkChnRegT
	InState         GdmaInStateChnRegT
	InSucEofDesAddr GdmaInSucEofDesAddrChnRegT
	InErrEofDesAddr GdmaInErrEofDesAddrChnRegT
	InDscr          GdmaInDscrChnRegT
	InDscrBf0       GdmaInDscrBf0ChnRegT
	InDscrBf1       GdmaInDscrBf1ChnRegT
	InPri           GdmaInPriChnRegT
	InPeriSel       GdmaInPeriSelChnRegT
}

type GdmaOutChnRegT struct {
	OutConf0         GdmaOutConf0ChnRegT
	OutConf1         GdmaOutConf1ChnRegT
	OutfifoStatus    GdmaOutfifoStatusChnRegT
	OutPush          GdmaOutPushChnRegT
	OutLink          GdmaOutLinkChnRegT
	OutState         GdmaOutStateChnRegT
	OutEofDesAddr    GdmaOutEofDesAddrChnRegT
	OutEofBfrDesAddr GdmaOutEofBfrDesAddrChnRegT
	OutDscr          GdmaOutDscrChnRegT
	OutDscrBf0       GdmaOutDscrBf0ChnRegT
	OutDscrBf1       GdmaOutDscrBf1ChnRegT
	OutPri           GdmaOutPriChnRegT
	OutPeriSel       GdmaOutPeriSelChnRegT
}

type GdmaChnRegT struct {
	In          GdmaInChnRegT
	ReservedIn  [11]c.Uint32T
	Out         GdmaOutChnRegT
	ReservedOut [11]c.Uint32T
}

type GdmaDevS struct {
	InIntr      [3]GdmaInIntChnRegT
	OutIntr     [3]GdmaOutIntChnRegT
	AhbTest     GdmaAhbTestRegT
	MiscConf    GdmaMiscConfRegT
	Date        GdmaDateRegT
	Reserved06c c.Uint32T
	Channel     [3]GdmaChnRegT
}
type GdmaDevT GdmaDevS

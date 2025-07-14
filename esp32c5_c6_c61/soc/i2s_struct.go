package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Interrupt registers */
/** Type of int_raw register
 *  I2S interrupt raw register, valid in level.
 */

type I2sIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  I2S interrupt status register.
 */

type I2sIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  I2S interrupt enable register.
 */

type I2sIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  I2S interrupt clear register.
 */

type I2sIntClrRegT struct {
	Val c.Uint32T
}

/** Group: RX Control and configuration registers */
/** Type of rx_conf register
 *  I2S RX configure register
 */

type I2sRxConfRegT struct {
	Val c.Uint32T
}

/** Type of rx_conf1 register
 *  I2S RX configure register 1
 */

type I2sRxConf1RegT struct {
	Val c.Uint32T
}

/** Type of tx_pcm2pdm_conf register
 *  I2S TX PCM2PDM configuration register
 */

type I2sTxPcm2pdmConfRegT struct {
	Val c.Uint32T
}

/** Type of tx_pcm2pdm_conf1 register
 *  I2S TX PCM2PDM configuration register
 */

type I2sTxPcm2pdmConf1RegT struct {
	Val c.Uint32T
}

/** Type of rx_tdm_ctrl register
 *  I2S TX TDM mode control register
 */

type I2sRxTdmCtrlRegT struct {
	Val c.Uint32T
}

/** Type of rx_eof_num register
 *  I2S RX data number control register.
 */

type I2sRxEofNumRegT struct {
	Val c.Uint32T
}

/** Group: TX Control and configuration registers */
/** Type of tx_conf register
 *  I2S TX configure register
 */

type I2sTxConfRegT struct {
	Val c.Uint32T
}

/** Type of tx_conf1 register
 *  I2S TX configure register 1
 */

type I2sTxConf1RegT struct {
	Val c.Uint32T
}

/** Type of tx_tdm_ctrl register
 *  I2S TX TDM mode control register
 */

type I2sTxTdmCtrlRegT struct {
	Val c.Uint32T
}

/** Group: RX clock and timing registers */
/** Type of rx_timing register
 *  I2S RX timing control register
 */

type I2sRxTimingRegT struct {
	Val c.Uint32T
}

/** Group: TX clock and timing registers */
/** Type of tx_timing register
 *  I2S TX timing control register
 */

type I2sTxTimingRegT struct {
	Val c.Uint32T
}

/** Group: Control and configuration registers */
/** Type of lc_hung_conf register
 *  I2S HUNG configure register.
 */

type I2sLcHungConfRegT struct {
	Val c.Uint32T
}

/** Type of conf_single_data register
 *  I2S signal data register
 */

type I2sConfSingleDataRegT struct {
	Val c.Uint32T
}

/** Group: TX status registers */
/** Type of state register
 *  I2S TX status register
 */

type I2sStateRegT struct {
	Val c.Uint32T
}

/** Group: ETM registers */
/** Type of etm_conf register
 *  I2S ETM configure register
 */

type I2sEtmConfRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version control register
 */

type I2sDateRegT struct {
	Val c.Uint32T
}

type I2sDevT struct {
	Reserved000    [3]c.Uint32T
	IntRaw         I2sIntRawRegT
	IntSt          I2sIntStRegT
	IntEna         I2sIntEnaRegT
	IntClr         I2sIntClrRegT
	Reserved01c    c.Uint32T
	RxConf         I2sRxConfRegT
	TxConf         I2sTxConfRegT
	RxConf1        I2sRxConf1RegT
	TxConf1        I2sTxConf1RegT
	Reserved030    [4]c.Uint32T
	TxPcm2pdmConf  I2sTxPcm2pdmConfRegT
	TxPcm2pdmConf1 I2sTxPcm2pdmConf1RegT
	Reserved048    [2]c.Uint32T
	RxTdmCtrl      I2sRxTdmCtrlRegT
	TxTdmCtrl      I2sTxTdmCtrlRegT
	RxTiming       I2sRxTimingRegT
	TxTiming       I2sTxTimingRegT
	LcHungConf     I2sLcHungConfRegT
	RxEofNum       I2sRxEofNumRegT
	ConfSingleData I2sConfSingleDataRegT
	State          I2sStateRegT
	EtmConf        I2sEtmConfRegT
	Reserved074    [3]c.Uint32T
	Date           I2sDateRegT
}

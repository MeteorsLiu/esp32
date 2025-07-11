package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: FIFO R/W registers */
/** Type of chndata register
 *  Read and write data for channel n via APB FIFO
 */

type RmtChndataRegT struct {
	Val c.Uint32T
}

/** Group: Configuration registers */
/** Type of chnconf0 register
 *  Channel n configuration register 0
 */

type RmtChnconf0RegT struct {
	Val c.Uint32T
}

/** Type of chnconf1 register
 *  Channel n configuration register 1
 */

type RmtChnconf1RegT struct {
	Val c.Uint32T
}

/** Type of apb_conf register
 *  RMT APB configuration register
 */

type RmtApbConfRegT struct {
	Val c.Uint32T
}

/** Type of ref_cnt_rst register
 *  RMT clock divider reset register
 */

type RmtRefCntRstRegT struct {
	Val c.Uint32T
}

/** Type of chn_rx_carrier_rm register
 *  Channel n carrier remove register
 */

type RmtChnRxCarrierRmRegT struct {
	Val c.Uint32T
}

/** Group: Status registers */
/** Type of chnstatus register
 *  Channel n status register
 */

type RmtChnstatusRegT struct {
	Val c.Uint32T
}

/** Type of chnaddr register
 *  Channel n address register
 */

type RmtChnaddrRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of int_raw register
 *  Raw interrupt status register
 */

type RmtIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  Masked interrupt status register
 */

type RmtIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  Interrupt enable register
 */

type RmtIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Interrupt clear register
 */

type RmtIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Carrier wave duty cycle registers */
/** Type of chncarrier_duty register
 *  Channel n duty cycle configuration register
 */

type RmtChncarrierDutyRegT struct {
	Val c.Uint32T
}

/** Group: Tx event configuration registers */
/** Type of chn_tx_lim register
 *  Channel n Tx event configuration register
 */

type RmtChnTxLimRegT struct {
	Val c.Uint32T
}

/** Type of tx_sim register
 *  Enable RMT simultaneous transmission
 */

type RmtTxSimRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version control register
 */

type RmtDateRegT struct {
	Val c.Uint32T
}

type RmtDevT struct {
	Chndata [4]RmtChndataRegT
	ConfCh  [4]struct {
		Conf0 RmtChnconf0RegT
		Conf1 RmtChnconf1RegT
	}
	Chnstatus      [4]RmtChnstatusRegT
	Chnaddr        [4]RmtChnaddrRegT
	IntRaw         RmtIntRawRegT
	IntSt          RmtIntStRegT
	IntEna         RmtIntEnaRegT
	IntClr         RmtIntClrRegT
	ChncarrierDuty [4]RmtChncarrierDutyRegT
	ChnTxLim       [4]RmtChnTxLimRegT
	ApbConf        RmtApbConfRegT
	TxSim          RmtTxSimRegT
	RefCntRst      RmtRefCntRstRegT
	ChnRxCarrierRm [4]RmtChnRxCarrierRmRegT
	Reserved09c    [24]c.Uint32T
	Date           RmtDateRegT
}

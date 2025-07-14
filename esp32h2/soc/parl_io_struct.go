package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: PARL_IO RX  Mode Configuration */
/** Type of rx_mode_cfg register
 *  Parallel RX Sampling mode configuration register.
 */

type ParlIoRxModeCfgRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO RX Data Configuration */
/** Type of rx_data_cfg register
 *  Parallel RX data configuration register.
 */

type ParlIoRxDataCfgRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO RX General Configuration */
/** Type of rx_genrl_cfg register
 *  Parallel RX general configuration register.
 */

type ParlIoRxGenrlCfgRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO RX Start Configuration */
/** Type of rx_start_cfg register
 *  Parallel RX Start configuration register.
 */

type ParlIoRxStartCfgRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO TX Data Configuration */
/** Type of tx_data_cfg register
 *  Parallel TX data configuration register.
 */

type ParlIoTxDataCfgRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO TX Start Configuration */
/** Type of tx_start_cfg register
 *  Parallel TX Start configuration register.
 */

type ParlIoTxStartCfgRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO TX General Configuration */
/** Type of tx_genrl_cfg register
 *  Parallel TX general configuration register.
 */

type ParlIoTxGenrlCfgRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO FIFO Configuration */
/** Type of fifo_cfg register
 *  Parallel IO FIFO configuration register.
 */

type ParlIoFifoCfgRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Register Update Configuration */
/** Type of reg_update register
 *  Parallel IO FIFO configuration register.
 */

type ParlIoRegUpdateRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Status */
/** Type of st register
 *  Parallel IO module status register0.
 */

type ParlIoStRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Interrupt Configuration and Status */
/** Type of int_ena register
 *  Parallel IO interrupt enable signal configuration register.
 */

type ParlIoIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_raw register
 *  Parallel IO interrupt raw signal status register.
 */

type ParlIoIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  Parallel IO interrupt signal status register.
 */

type ParlIoIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Parallel IO interrupt  clear signal configuration register.
 */

type ParlIoIntClrRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Rx Status0 */
/** Type of rx_st0 register
 *  Parallel IO RX status register0
 */

type ParlIoRxSt0RegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Rx Status1 */
/** Type of rx_st1 register
 *  Parallel IO RX status register1
 */

type ParlIoRxSt1RegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Tx Status0 */
/** Type of tx_st0 register
 *  Parallel IO TX status register0
 */

type ParlIoTxSt0RegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Rx Clock Configuration */
/** Type of rx_clk_cfg register
 *  Parallel IO RX clk configuration register
 */

type ParlIoRxClkCfgRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Tx Clock Configuration */
/** Type of tx_clk_cfg register
 *  Parallel IO TX clk configuration register
 */

type ParlIoTxClkCfgRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Clock Configuration */
/** Type of clk register
 *  Parallel IO clk configuration register
 */

type ParlIoClkRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Version Register */
/** Type of version register
 *  Version register.
 */

type ParlIoVersionRegT struct {
	Val c.Uint32T
}

type ParlIoDevT struct {
	RxModeCfg   ParlIoRxModeCfgRegT
	RxDataCfg   ParlIoRxDataCfgRegT
	RxGenrlCfg  ParlIoRxGenrlCfgRegT
	RxStartCfg  ParlIoRxStartCfgRegT
	TxDataCfg   ParlIoTxDataCfgRegT
	TxStartCfg  ParlIoTxStartCfgRegT
	TxGenrlCfg  ParlIoTxGenrlCfgRegT
	FifoCfg     ParlIoFifoCfgRegT
	RegUpdate   ParlIoRegUpdateRegT
	St          ParlIoStRegT
	IntEna      ParlIoIntEnaRegT
	IntRaw      ParlIoIntRawRegT
	IntSt       ParlIoIntStRegT
	IntClr      ParlIoIntClrRegT
	RxSt0       ParlIoRxSt0RegT
	RxSt1       ParlIoRxSt1RegT
	TxSt0       ParlIoTxSt0RegT
	RxClkCfg    ParlIoRxClkCfgRegT
	TxClkCfg    ParlIoTxClkCfgRegT
	Reserved04c [53]c.Uint32T
	Clk         ParlIoClkRegT
	Reserved124 [182]c.Uint32T
	Version     ParlIoVersionRegT
}

package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Registers */
/** Type of mode register
 *  TWAI mode register.
 */

type TwaiModeRegT struct {
	Val c.Uint32T
}

/** Type of cmd register
 *  TWAI command register.
 */

type TwaiCmdRegT struct {
	Val c.Uint32T
}

/** Type of bus_timing_0 register
 *  Bit timing configuration register 0.
 */

type TwaiBusTiming0RegT struct {
	Val c.Uint32T
}

/** Type of bus_timing_1 register
 *  Bit timing configuration register 1.
 */

type TwaiBusTiming1RegT struct {
	Val c.Uint32T
}

/** Type of err_warning_limit register
 *  TWAI error threshold configuration register.
 */

type TwaiErrWarningLimitRegT struct {
	Val c.Uint32T
}

/** Type of clock_divider register
 *  Clock divider register.
 */

type TwaiClockDividerRegT struct {
	Val c.Uint32T
}

/** Type of sw_standby_cfg register
 *  Software configure standby pin directly.
 */

type TwaiSwStandbyCfgRegT struct {
	Val c.Uint32T
}

/** Type of hw_cfg register
 *  Hardware configure standby pin.
 */

type TwaiHwCfgRegT struct {
	Val c.Uint32T
}

/** Type of hw_standby_cnt register
 *  Configure standby counter.
 */

type TwaiHwStandbyCntRegT struct {
	Val c.Uint32T
}

/** Type of idle_intr_cnt register
 *  Configure idle interrupt counter.
 */

type TwaiIdleIntrCntRegT struct {
	Val c.Uint32T
}

/** Type of eco_cfg register
 *  ECO configuration register.
 */

type TwaiEcoCfgRegT struct {
	Val c.Uint32T
}

/** Group: Status Registers */
/** Type of status register
 *  TWAI status register.
 */

type TwaiStatusRegT struct {
	Val c.Uint32T
}

/** Type of arb_lost_cap register
 *  TWAI arbiter lost capture register.
 */

type TwaiArbLostCapRegT struct {
	Val c.Uint32T
}

/** Type of err_code_cap register
 *  TWAI error info capture register.
 */

type TwaiErrCodeCapRegT struct {
	Val c.Uint32T
}

/** Type of rx_err_cnt register
 *  Rx error counter register.
 */

type TwaiRxErrCntRegT struct {
	Val c.Uint32T
}

/** Type of tx_err_cnt register
 *  Tx error counter register.
 */

type TwaiTxErrCntRegT struct {
	Val c.Uint32T
}

/** Type of rx_message_counter register
 *  Received message counter register.
 */

type TwaiRxMessageCounterRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Registers */
/** Type of interrupt register
 *  Interrupt signals' register.
 */

type TwaiInterruptRegT struct {
	Val c.Uint32T
}

/** Type of interrupt_enable register
 *  Interrupt enable register.
 */

type TwaiInterruptEnableRegT struct {
	Val c.Uint32T
}

/** Group: Data Registers */
/** Type of buffer register
 *  TX RX Buffer.
 */

type TwaiTxRxBufferRegT struct {
	Val c.Uint32T
}

type AcceptanceFilterRegT struct {
	Acr [4]struct {
		Val c.Uint32T
	}
	Amr [4]struct {
		Val c.Uint32T
	}
	Reserved60 c.Uint32T
	Reserved64 c.Uint32T
	Reserved68 c.Uint32T
	Reserved6c c.Uint32T
	Reserved70 c.Uint32T
}

type TwaiDevS struct {
	Mode             TwaiModeRegT
	Cmd              TwaiCmdRegT
	Status           TwaiStatusRegT
	Interrupt        TwaiInterruptRegT
	InterruptEnable  TwaiInterruptEnableRegT
	Reserved014      c.Uint32T
	BusTiming0       TwaiBusTiming0RegT
	BusTiming1       TwaiBusTiming1RegT
	Reserved020      [3]c.Uint32T
	ArbLostCap       TwaiArbLostCapRegT
	ErrCodeCap       TwaiErrCodeCapRegT
	ErrWarningLimit  TwaiErrWarningLimitRegT
	RxErrCnt         TwaiRxErrCntRegT
	TxErrCnt         TwaiTxErrCntRegT
	RxMessageCounter TwaiRxMessageCounterRegT
	Reserved078      c.Uint32T
	ClockDivider     TwaiClockDividerRegT
	SwStandbyCfg     TwaiSwStandbyCfgRegT
	HwCfg            TwaiHwCfgRegT
	HwStandbyCnt     TwaiHwStandbyCntRegT
	IdleIntrCnt      TwaiIdleIntrCntRegT
	EcoCfg           TwaiEcoCfgRegT
}
type TwaiDevT TwaiDevS

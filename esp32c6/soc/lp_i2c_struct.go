package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Timing registers */
/** Type of scl_low_period register
 *  Configures the low level width of the SCL
 *  Clock
 */

type LpI2cSclLowPeriodRegT struct {
	Val c.Uint32T
}

/** Type of sda_hold register
 *  Configures the hold time after a negative SCL edge.
 */

type LpI2cSdaHoldRegT struct {
	Val c.Uint32T
}

/** Type of sda_sample register
 *  Configures the sample time after a positive SCL edge.
 */

type LpI2cSdaSampleRegT struct {
	Val c.Uint32T
}

/** Type of scl_high_period register
 *  Configures the high level width of SCL
 */

type LpI2cSclHighPeriodRegT struct {
	Val c.Uint32T
}

/** Type of scl_start_hold register
 *  Configures the delay between the SDA and SCL negative edge for a start condition
 */

type LpI2cSclStartHoldRegT struct {
	Val c.Uint32T
}

/** Type of scl_rstart_setup register
 *  Configures the delay between the positive
 *  edge of SCL and the negative edge of SDA
 */

type LpI2cSclRstartSetupRegT struct {
	Val c.Uint32T
}

/** Type of scl_stop_hold register
 *  Configures the delay after the SCL clock
 *  edge for a stop condition
 */

type LpI2cSclStopHoldRegT struct {
	Val c.Uint32T
}

/** Type of scl_stop_setup register
 *  Configures the delay between the SDA and
 *  SCL positive edge for a stop condition
 */

type LpI2cSclStopSetupRegT struct {
	Val c.Uint32T
}

/** Type of scl_st_time_out register
 *  SCL status time out register
 */

type LpI2cSclStTimeOutRegT struct {
	Val c.Uint32T
}

/** Type of scl_main_st_time_out register
 *  SCL main status time out register
 */

type LpI2cSclMainStTimeOutRegT struct {
	Val c.Uint32T
}

/** Group: Configuration registers */
/** Type of ctr register
 *  Transmission setting
 */

type LpI2cCtrRegT struct {
	Val c.Uint32T
}

/** Type of to register
 *  Setting time out control for receiving data.
 */

type LpI2cToRegT struct {
	Val c.Uint32T
}

/** Type of fifo_conf register
 *  FIFO configuration register.
 */

type LpI2cFifoConfRegT struct {
	Val c.Uint32T
}

/** Type of filter_cfg register
 *  SCL and SDA filter configuration register
 */

type LpI2cFilterCfgRegT struct {
	Val c.Uint32T
}

/** Type of clk_conf register
 *  I2C CLK configuration register
 */

type LpI2cClkConfRegT struct {
	Val c.Uint32T
}

/** Type of scl_sp_conf register
 *  Power configuration register
 */

type LpI2cSclSpConfRegT struct {
	Val c.Uint32T
}

/** Group: Status registers */
/** Type of sr register
 *  Describe I2C work status.
 */

type LpI2cSrRegT struct {
	Val c.Uint32T
}

/** Type of fifo_st register
 *  FIFO status register.
 */

type LpI2cFifoStRegT struct {
	Val c.Uint32T
}

/** Type of data register
 *  Rx FIFO read data.
 */

type LpI2cDataRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of int_raw register
 *  Raw interrupt status
 */

type LpI2cIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Interrupt clear bits
 */

type LpI2cIntClrRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  Interrupt enable bits
 */

type LpI2cIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_status register
 *  Status of captured I2C communication events
 */

type LpI2cIntStatusRegT struct {
	Val c.Uint32T
}

/** Group: Command registers */
/** Type of command register
 *  I2C command register
 */

type LpI2cCommandRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version register
 */

type LpI2cDateRegT struct {
	Val c.Uint32T
}

/** Group: Address register */
/** Type of txfifo_start_addr register
 *  I2C TXFIFO base address register
 */

type LpI2cTxfifoStartAddrRegT struct {
	Val c.Uint32T
}

/** Type of rxfifo_start_addr register
 *  I2C RXFIFO base address register
 */

type LpI2cRxfifoStartAddrRegT struct {
	Val c.Uint32T
}

type LpI2cDevT struct {
	SclLowPeriod     LpI2cSclLowPeriodRegT
	Ctr              LpI2cCtrRegT
	Sr               LpI2cSrRegT
	To               LpI2cToRegT
	Reserved010      c.Uint32T
	FifoSt           LpI2cFifoStRegT
	FifoConf         LpI2cFifoConfRegT
	Data             LpI2cDataRegT
	IntRaw           LpI2cIntRawRegT
	IntClr           LpI2cIntClrRegT
	IntEna           LpI2cIntEnaRegT
	IntStatus        LpI2cIntStatusRegT
	SdaHold          LpI2cSdaHoldRegT
	SdaSample        LpI2cSdaSampleRegT
	SclHighPeriod    LpI2cSclHighPeriodRegT
	Reserved03c      c.Uint32T
	SclStartHold     LpI2cSclStartHoldRegT
	SclRstartSetup   LpI2cSclRstartSetupRegT
	SclStopHold      LpI2cSclStopHoldRegT
	SclStopSetup     LpI2cSclStopSetupRegT
	FilterCfg        LpI2cFilterCfgRegT
	ClkConf          LpI2cClkConfRegT
	Command          [8]LpI2cCommandRegT
	SclStTimeOut     LpI2cSclStTimeOutRegT
	SclMainStTimeOut LpI2cSclMainStTimeOutRegT
	SclSpConf        LpI2cSclSpConfRegT
	Reserved084      [29]c.Uint32T
	Date             LpI2cDateRegT
	Reserved0fc      c.Uint32T
	TxfifoStartAddr  LpI2cTxfifoStartAddrRegT
	Reserved104      [31]c.Uint32T
	RxfifoStartAddr  LpI2cRxfifoStartAddrRegT
}

package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const TWAI_ALERT_TX_IDLE = 0x00000001
const TWAI_ALERT_TX_SUCCESS = 0x00000002
const TWAI_ALERT_RX_DATA = 0x00000004
const TWAI_ALERT_BELOW_ERR_WARN = 0x00000008
const TWAI_ALERT_ERR_ACTIVE = 0x00000010
const TWAI_ALERT_RECOVERY_IN_PROGRESS = 0x00000020
const TWAI_ALERT_BUS_RECOVERED = 0x00000040
const TWAI_ALERT_ARB_LOST = 0x00000080
const TWAI_ALERT_ABOVE_ERR_WARN = 0x00000100
const TWAI_ALERT_BUS_ERROR = 0x00000200
const TWAI_ALERT_TX_FAILED = 0x00000400
const TWAI_ALERT_RX_QUEUE_FULL = 0x00000800
const TWAI_ALERT_ERR_PASS = 0x00001000
const TWAI_ALERT_BUS_OFF = 0x00002000
const TWAI_ALERT_RX_FIFO_OVERRUN = 0x00004000
const TWAI_ALERT_TX_RETRIED = 0x00008000
const TWAI_ALERT_PERIPH_RESET = 0x00010000
const TWAI_ALERT_ALL = 0x0001FFFF
const TWAI_ALERT_NONE = 0x00000000
const TWAI_ALERT_AND_LOG = 0x00020000

type TwaiObjT struct {
	Unused [8]uint8
}
type TwaiHandleT *TwaiObjT
type TwaiStateT c.Int

const (
	TWAI_STATE_STOPPED    TwaiStateT = 0
	TWAI_STATE_RUNNING    TwaiStateT = 1
	TWAI_STATE_BUS_OFF    TwaiStateT = 2
	TWAI_STATE_RECOVERING TwaiStateT = 3
)

/**
 * @brief   Structure for general configuration of the TWAI driver
 *
 * @note    Macro initializers are available for this structure
 */

type TwaiGeneralConfigT struct {
	ControllerId  c.Int
	Mode          TwaiModeT
	TxIo          GpioNumT
	RxIo          GpioNumT
	ClkoutIo      GpioNumT
	BusOffIo      GpioNumT
	TxQueueLen    c.Uint32T
	RxQueueLen    c.Uint32T
	AlertsEnabled c.Uint32T
	ClkoutDivider c.Uint32T
	IntrFlags     c.Int
	GeneralFlags  struct {
		SleepAllowPd c.Uint32T
	}
}

/**
 * @brief   Structure to store status information of TWAI driver
 */

type TwaiStatusInfoT struct {
	State          TwaiStateT
	MsgsToTx       c.Uint32T
	MsgsToRx       c.Uint32T
	TxErrorCounter c.Uint32T
	RxErrorCounter c.Uint32T
	TxFailedCount  c.Uint32T
	RxMissedCount  c.Uint32T
	RxOverrunCount c.Uint32T
	ArbLostCount   c.Uint32T
	BusErrorCount  c.Uint32T
}

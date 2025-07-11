package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const PCNT_MODE_MAX = 3
const PCNT_COUNT_MAX = 3

type PcntIsrHandleT IntrHandleT
type PcntPortT c.Int

const (
	PCNT_PORT_0   PcntPortT = 0
	PCNT_PORT_MAX PcntPortT = 1
)

type PcntUnitT c.Int

const (
	PCNT_UNIT_0   PcntUnitT = 0
	PCNT_UNIT_1   PcntUnitT = 1
	PCNT_UNIT_2   PcntUnitT = 2
	PCNT_UNIT_3   PcntUnitT = 3
	PCNT_UNIT_4   PcntUnitT = 4
	PCNT_UNIT_5   PcntUnitT = 5
	PCNT_UNIT_6   PcntUnitT = 6
	PCNT_UNIT_7   PcntUnitT = 7
	PCNT_UNIT_MAX PcntUnitT = 8
)

type PcntChannelT c.Int

const (
	PCNT_CHANNEL_0   PcntChannelT = 0
	PCNT_CHANNEL_1   PcntChannelT = 1
	PCNT_CHANNEL_MAX PcntChannelT = 2
)

type PcntEvtTypeT c.Int

const (
	PCNT_EVT_THRES_1 PcntEvtTypeT = 4
	PCNT_EVT_THRES_0 PcntEvtTypeT = 8
	PCNT_EVT_L_LIM   PcntEvtTypeT = 16
	PCNT_EVT_H_LIM   PcntEvtTypeT = 32
	PCNT_EVT_ZERO    PcntEvtTypeT = 64
	PCNT_EVT_MAX     PcntEvtTypeT = 65
)

type PcntCtrlModeT PcntChannelLevelActionT
type PcntCountModeT PcntChannelEdgeActionT

/**
 * @brief Pulse Counter configuration for a single channel
 */

type PcntConfigT struct {
	PulseGpioNum c.Int
	CtrlGpioNum  c.Int
	LctrlMode    PcntCtrlModeT
	HctrlMode    PcntCtrlModeT
	PosMode      PcntCountModeT
	NegMode      PcntCountModeT
	CounterHLim  c.Int16T
	CounterLLim  c.Int16T
	Unit         PcntUnitT
	Channel      PcntChannelT
}

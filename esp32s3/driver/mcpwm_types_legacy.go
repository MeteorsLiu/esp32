package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type McpwmIoSignalsT c.Int

const (
	MCPWM0A       McpwmIoSignalsT = 0
	MCPWM0B       McpwmIoSignalsT = 1
	MCPWM1A       McpwmIoSignalsT = 2
	MCPWM1B       McpwmIoSignalsT = 3
	MCPWM2A       McpwmIoSignalsT = 4
	MCPWM2B       McpwmIoSignalsT = 5
	MCPWM_SYNC_0  McpwmIoSignalsT = 6
	MCPWM_SYNC_1  McpwmIoSignalsT = 7
	MCPWM_SYNC_2  McpwmIoSignalsT = 8
	MCPWM_FAULT_0 McpwmIoSignalsT = 9
	MCPWM_FAULT_1 McpwmIoSignalsT = 10
	MCPWM_FAULT_2 McpwmIoSignalsT = 11
	MCPWM_CAP_0   McpwmIoSignalsT = 84
	MCPWM_CAP_1   McpwmIoSignalsT = 85
	MCPWM_CAP_2   McpwmIoSignalsT = 86
)

/**
 * @brief pin number for MCPWM
 */

type McpwmPinConfigT struct {
	Mcpwm0aOutNum    c.Int
	Mcpwm0bOutNum    c.Int
	Mcpwm1aOutNum    c.Int
	Mcpwm1bOutNum    c.Int
	Mcpwm2aOutNum    c.Int
	Mcpwm2bOutNum    c.Int
	McpwmSync0InNum  c.Int
	McpwmSync1InNum  c.Int
	McpwmSync2InNum  c.Int
	McpwmFault0InNum c.Int
	McpwmFault1InNum c.Int
	McpwmFault2InNum c.Int
	McpwmCap0InNum   c.Int
	McpwmCap1InNum   c.Int
	McpwmCap2InNum   c.Int
}
type McpwmUnitT c.Int

const (
	MCPWM_UNIT_0   McpwmUnitT = 0
	MCPWM_UNIT_1   McpwmUnitT = 1
	MCPWM_UNIT_MAX McpwmUnitT = 2
)

type McpwmTimerT c.Int

const (
	MCPWM_TIMER_0   McpwmTimerT = 0
	MCPWM_TIMER_1   McpwmTimerT = 1
	MCPWM_TIMER_2   McpwmTimerT = 2
	MCPWM_TIMER_MAX McpwmTimerT = 3
)

type McpwmGeneratorT c.Int

const (
	MCPWM_GEN_A   McpwmGeneratorT = 0
	MCPWM_GEN_B   McpwmGeneratorT = 1
	MCPWM_GEN_MAX McpwmGeneratorT = 2
)

type McpwmOperatorT McpwmGeneratorT
type McpwmCarrierOutIvtT c.Int

const (
	MCPWM_CARRIER_OUT_IVT_DIS McpwmCarrierOutIvtT = 0
	MCPWM_CARRIER_OUT_IVT_EN  McpwmCarrierOutIvtT = 1
)

type McpwmFaultSignalT c.Int

const (
	MCPWM_SELECT_F0 McpwmFaultSignalT = 0
	MCPWM_SELECT_F1 McpwmFaultSignalT = 1
	MCPWM_SELECT_F2 McpwmFaultSignalT = 2
)

type McpwmSyncSignalT c.Int

const (
	MCPWM_SELECT_NO_INPUT    McpwmSyncSignalT = 0
	MCPWM_SELECT_TIMER0_SYNC McpwmSyncSignalT = 1
	MCPWM_SELECT_TIMER1_SYNC McpwmSyncSignalT = 2
	MCPWM_SELECT_TIMER2_SYNC McpwmSyncSignalT = 3
	MCPWM_SELECT_GPIO_SYNC0  McpwmSyncSignalT = 4
	MCPWM_SELECT_GPIO_SYNC1  McpwmSyncSignalT = 5
	MCPWM_SELECT_GPIO_SYNC2  McpwmSyncSignalT = 6
)

type McpwmTimerSyncTriggerT c.Int

const (
	MCPWM_SWSYNC_SOURCE_SYNCIN   McpwmTimerSyncTriggerT = 0
	MCPWM_SWSYNC_SOURCE_TEZ      McpwmTimerSyncTriggerT = 1
	MCPWM_SWSYNC_SOURCE_TEP      McpwmTimerSyncTriggerT = 2
	MCPWM_SWSYNC_SOURCE_DISABLED McpwmTimerSyncTriggerT = 3
)

type McpwmFaultInputLevelT c.Int

const (
	MCPWM_LOW_LEVEL_TGR  McpwmFaultInputLevelT = 0
	MCPWM_HIGH_LEVEL_TGR McpwmFaultInputLevelT = 1
)

type McpwmCaptureOnEdgeT c.Int

const (
	MCPWM_NEG_EDGE  McpwmCaptureOnEdgeT = 1
	MCPWM_POS_EDGE  McpwmCaptureOnEdgeT = 2
	MCPWM_BOTH_EDGE McpwmCaptureOnEdgeT = 3
)

type McpwmCounterTypeT c.Int

const (
	MCPWM_FREEZE_COUNTER  McpwmCounterTypeT = 0
	MCPWM_UP_COUNTER      McpwmCounterTypeT = 1
	MCPWM_DOWN_COUNTER    McpwmCounterTypeT = 2
	MCPWM_UP_DOWN_COUNTER McpwmCounterTypeT = 3
	MCPWM_COUNTER_MAX     McpwmCounterTypeT = 4
)

type McpwmDutyTypeT c.Int

const (
	MCPWM_DUTY_MODE_0          McpwmDutyTypeT = 0
	MCPWM_DUTY_MODE_1          McpwmDutyTypeT = 1
	MCPWM_DUTY_MODE_FORCE_LOW  McpwmDutyTypeT = 2
	MCPWM_DUTY_MODE_FORCE_HIGH McpwmDutyTypeT = 3
	MCPWM_DUTY_MODE_MAX        McpwmDutyTypeT = 4
)

type McpwmDeadtimeTypeT c.Int

const (
	MCPWM_DEADTIME_BYPASS             McpwmDeadtimeTypeT = 0
	MCPWM_BYPASS_RED                  McpwmDeadtimeTypeT = 1
	MCPWM_BYPASS_FED                  McpwmDeadtimeTypeT = 2
	MCPWM_ACTIVE_HIGH_MODE            McpwmDeadtimeTypeT = 3
	MCPWM_ACTIVE_LOW_MODE             McpwmDeadtimeTypeT = 4
	MCPWM_ACTIVE_HIGH_COMPLIMENT_MODE McpwmDeadtimeTypeT = 5
	MCPWM_ACTIVE_LOW_COMPLIMENT_MODE  McpwmDeadtimeTypeT = 6
	MCPWM_ACTIVE_RED_FED_FROM_PWMXA   McpwmDeadtimeTypeT = 7
	MCPWM_ACTIVE_RED_FED_FROM_PWMXB   McpwmDeadtimeTypeT = 8
	MCPWM_DEADTIME_TYPE_MAX           McpwmDeadtimeTypeT = 9
)

type McpwmOutputActionT c.Int

const (
	MCPWM_ACTION_NO_CHANGE  McpwmOutputActionT = 0
	MCPWM_ACTION_FORCE_LOW  McpwmOutputActionT = 1
	MCPWM_ACTION_FORCE_HIGH McpwmOutputActionT = 2
	MCPWM_ACTION_TOGGLE     McpwmOutputActionT = 3
)

type McpwmActionOnPwmxaT McpwmOutputActionT
type McpwmActionOnPwmxbT McpwmOutputActionT
type McpwmCaptureSignalT c.Int

const (
	MCPWM_SELECT_CAP0 McpwmCaptureSignalT = 0
	MCPWM_SELECT_CAP1 McpwmCaptureSignalT = 1
	MCPWM_SELECT_CAP2 McpwmCaptureSignalT = 2
)

type McpwmCaptureChannelIdT McpwmCaptureSignalT

/**
 * @brief event data that will be passed into ISR callback
 */

type CapEventDataT struct {
	CapEdge  McpwmCaptureOnEdgeT
	CapValue c.Uint32T
}

// llgo:type C
type CapIsrCbT func(McpwmUnitT, McpwmCaptureChannelIdT, *CapEventDataT, c.Pointer) bool

/**
 * @brief MCPWM config structure
 */

type McpwmConfigT struct {
	Frequency   c.Uint32T
	CmprA       c.Float
	CmprB       c.Float
	DutyMode    McpwmDutyTypeT
	CounterMode McpwmCounterTypeT
}

/**
 * @brief MCPWM carrier configuration structure
 */

type McpwmCarrierConfigT struct {
	CarrierPeriod  c.Uint8T
	CarrierDuty    c.Uint8T
	PulseWidthInOs c.Uint8T
	CarrierIvtMode McpwmCarrierOutIvtT
}

/**
 * @brief MCPWM config capture structure
 */

type McpwmCaptureConfigT struct {
	CapEdge     McpwmCaptureOnEdgeT
	CapPrescale c.Uint32T
	CaptureCb   CapIsrCbT
	UserData    c.Pointer
}

/**
 * @brief MCPWM config sync structure
 */

type McpwmSyncConfigT struct {
	SyncSig        McpwmSyncSignalT
	TimerVal       c.Uint32T
	CountDirection McpwmTimerDirectionT
}

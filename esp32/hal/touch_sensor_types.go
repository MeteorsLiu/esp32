package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TouchPadT c.Int

const (
	TOUCH_PAD_NUM0 TouchPadT = 0
	TOUCH_PAD_NUM1 TouchPadT = 1
	TOUCH_PAD_NUM2 TouchPadT = 2
	TOUCH_PAD_NUM3 TouchPadT = 3
	TOUCH_PAD_NUM4 TouchPadT = 4
	TOUCH_PAD_NUM5 TouchPadT = 5
	TOUCH_PAD_NUM6 TouchPadT = 6
	TOUCH_PAD_NUM7 TouchPadT = 7
	TOUCH_PAD_NUM8 TouchPadT = 8
	TOUCH_PAD_NUM9 TouchPadT = 9
	TOUCH_PAD_MAX  TouchPadT = 10
)

type TouchHighVoltT c.Int

const (
	TOUCH_HVOLT_KEEP TouchHighVoltT = -1
	TOUCH_HVOLT_2V4  TouchHighVoltT = 0
	TOUCH_HVOLT_2V5  TouchHighVoltT = 1
	TOUCH_HVOLT_2V6  TouchHighVoltT = 2
	TOUCH_HVOLT_2V7  TouchHighVoltT = 3
	TOUCH_HVOLT_MAX  TouchHighVoltT = 4
)

type TouchLowVoltT c.Int

const (
	TOUCH_LVOLT_KEEP TouchLowVoltT = -1
	TOUCH_LVOLT_0V5  TouchLowVoltT = 0
	TOUCH_LVOLT_0V6  TouchLowVoltT = 1
	TOUCH_LVOLT_0V7  TouchLowVoltT = 2
	TOUCH_LVOLT_0V8  TouchLowVoltT = 3
	TOUCH_LVOLT_MAX  TouchLowVoltT = 4
)

type TouchVoltAttenT c.Int

const (
	TOUCH_HVOLT_ATTEN_KEEP TouchVoltAttenT = -1
	TOUCH_HVOLT_ATTEN_1V5  TouchVoltAttenT = 0
	TOUCH_HVOLT_ATTEN_1V   TouchVoltAttenT = 1
	TOUCH_HVOLT_ATTEN_0V5  TouchVoltAttenT = 2
	TOUCH_HVOLT_ATTEN_0V   TouchVoltAttenT = 3
	TOUCH_HVOLT_ATTEN_MAX  TouchVoltAttenT = 4
)

type TouchCntSlopeT c.Int

const (
	TOUCH_PAD_SLOPE_0   TouchCntSlopeT = 0
	TOUCH_PAD_SLOPE_1   TouchCntSlopeT = 1
	TOUCH_PAD_SLOPE_2   TouchCntSlopeT = 2
	TOUCH_PAD_SLOPE_3   TouchCntSlopeT = 3
	TOUCH_PAD_SLOPE_4   TouchCntSlopeT = 4
	TOUCH_PAD_SLOPE_5   TouchCntSlopeT = 5
	TOUCH_PAD_SLOPE_6   TouchCntSlopeT = 6
	TOUCH_PAD_SLOPE_7   TouchCntSlopeT = 7
	TOUCH_PAD_SLOPE_MAX TouchCntSlopeT = 8
)

type TouchTieOptT c.Int

const (
	TOUCH_PAD_TIE_OPT_LOW   TouchTieOptT = 0
	TOUCH_PAD_TIE_OPT_HIGH  TouchTieOptT = 1
	TOUCH_PAD_TIE_OPT_FLOAT TouchTieOptT = 2
	TOUCH_PAD_TIE_OPT_MAX   TouchTieOptT = 3
)

type TouchFsmModeT c.Int

const (
	TOUCH_FSM_MODE_TIMER TouchFsmModeT = 0
	TOUCH_FSM_MODE_SW    TouchFsmModeT = 1
	TOUCH_FSM_MODE_MAX   TouchFsmModeT = 2
)

type TouchTriggerModeT c.Int

const (
	TOUCH_TRIGGER_BELOW TouchTriggerModeT = 0
	TOUCH_TRIGGER_ABOVE TouchTriggerModeT = 1
	TOUCH_TRIGGER_MAX   TouchTriggerModeT = 2
)

type TouchTriggerSrcT c.Int

const (
	TOUCH_TRIGGER_SOURCE_BOTH TouchTriggerSrcT = 0
	TOUCH_TRIGGER_SOURCE_SET1 TouchTriggerSrcT = 1
	TOUCH_TRIGGER_SOURCE_MAX  TouchTriggerSrcT = 2
)

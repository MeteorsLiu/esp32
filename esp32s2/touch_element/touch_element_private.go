package touch_element

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const TE_TAG = "Touch Element"
const TE_DEBUG_TAG = "Touch Element Debug"

type TeStateT c.Int

const (
	TE_STATE_IDLE    TeStateT = 0
	TE_STATE_PRESS   TeStateT = 1
	TE_STATE_RELEASE TeStateT = 2
)

type TeDevStateT TeStateT
type TeDevTypeT TouchElemTypeT

type TeDevT struct {
	Sens               c.Float
	Channel            TouchPadT
	Type               TeDevTypeT
	State              TeDevStateT
	IsUseLastThreshold bool
}
type TeClassTypeT c.Int

const (
	TE_CLS_TYPE_BUTTON TeClassTypeT = 0
	TE_CLS_TYPE_SLIDER TeClassTypeT = 1
	TE_CLS_TYPE_MATRIX TeClassTypeT = 2
	TE_CLS_TYPE_MAX    TeClassTypeT = 3
)

type TeObjectMethodsT struct {
	Handle       TouchElemHandleT
	CheckChannel c.Pointer
	SetThreshold c.Pointer
	ProcessState c.Pointer
	UpdateState  c.Pointer
}

/* -------------------------------------------- Waterproof basic type --------------------------------------------- */

type TeWaterproofS struct {
	GuardDevice      *TeDevT
	MaskHandle       *TouchElemHandleT
	ShieldChannel    TouchPadT
	IsShieldLevelSet bool
}
type TeWaterproofHandleT *TeWaterproofS

/* -------------------------------------------- Sleep basic type --------------------------------------------- */

type TeSleepS struct {
	WakeupHandle         TouchElemHandleT
	NonVolatileThreshold *c.Uint32T
}
type TeSleepHandleT *TeSleepS

/* -------------------------------------------- Button basic type --------------------------------------------- */

type TeButtonHandleConfigT struct {
	DispatchMethod TouchElemDispatchT
	Callback       TouchButtonCallbackT
	EventMask      c.Uint32T
	Arg            c.Pointer
}
type TeButtonStateT TeStateT

type TeButtonS struct {
	Config       *TeButtonHandleConfigT
	Device       *TeDevT
	CurrentState TeButtonStateT
	LastState    TeButtonStateT
	Event        TouchButtonEventT
	TriggerCnt   c.Uint32T
	TriggerThr   c.Uint32T
}
type TeButtonHandleT *TeButtonS

/* -------------------------------------------- Slider basic type --------------------------------------------- */

type TeSliderHandleConfigT struct {
	DispatchMethod TouchElemDispatchT
	Callback       TouchSliderCallbackT
	EventMask      c.Uint32T
	Arg            c.Pointer
}
type TeSliderStateT TeStateT

type TeSliderS struct {
	Config              *TeSliderHandleConfigT
	Device              **TeDevT
	CurrentState        TeSliderStateT
	LastState           TeSliderStateT
	Event               TouchSliderEventT
	PositionScale       c.Float
	QuantifySignalArray *c.Float
	ChannelBcm          *c.Uint32T
	ChannelBcmUpdateCnt c.Uint32T
	FilterResetCnt      c.Uint32T
	LastPosition        c.Uint32T
	Position            TouchSliderPositionT
	PositionRange       c.Uint8T
	ChannelSum          c.Uint8T
	PosFilterWindow     *c.Uint8T
	PosWindowIdx        c.Uint8T
	IsFirstSample       bool
}
type TeSliderHandleT *TeSliderS

/* -------------------------------------------- Matrix basic type --------------------------------------------- */

type TeMatrixHandleConfigT struct {
	DispatchMethod TouchElemDispatchT
	Callback       TouchMatrixCallbackT
	EventMask      c.Uint32T
	Arg            c.Pointer
}
type TeMatrixStateT TeStateT

type TeMatrixS struct {
	Config       *TeMatrixHandleConfigT
	Device       **TeDevT
	CurrentState TeMatrixStateT
	LastState    TeMatrixStateT
	Event        TouchMatrixEventT
	TriggerCnt   c.Uint32T
	TriggerThr   c.Uint32T
	Position     TouchMatrixPositionT
	XChannelNum  c.Uint8T
	YChannelNum  c.Uint8T
}
type TeMatrixHandleT *TeMatrixS

/* --------------------------------------------- Global system methods ---------------------------------------------- */
//go:linkname TeReadSmoothSignal C.te_read_smooth_signal
func TeReadSmoothSignal(channel_num TouchPadT) c.Uint32T

//go:linkname TeSystemCheckState C.te_system_check_state
func TeSystemCheckState() bool

// TODO: Refactor this function with function overload
//
//go:linkname TeDevInit C.te_dev_init
func TeDevInit(device **TeDevT, device_num c.Uint8T, type_ TeDevTypeT, channel *TouchPadT, sens *c.Float, divider c.Float) EspErrT

//go:linkname TeDevDeinit C.te_dev_deinit
func TeDevDeinit(device **TeDevT, device_num c.Uint8T)

// llgo:link (*TeDevT).TeDevSetThreshold C.te_dev_set_threshold
func (recv_ *TeDevT) TeDevSetThreshold() EspErrT {
	return 0
}

// llgo:link TouchElemMessageT.TeEventGive C.te_event_give
func (recv_ TouchElemMessageT) TeEventGive() EspErrT {
	return 0
}

//go:linkname TeGetTimerPeriod C.te_get_timer_period
func TeGetTimerPeriod() c.Uint8T

// llgo:link (*TeObjectMethodsT).TeObjectMethodRegister C.te_object_method_register
func (recv_ *TeObjectMethodsT) TeObjectMethodRegister(object_type TeClassTypeT) {
}

// llgo:link TeClassTypeT.TeObjectMethodUnregister C.te_object_method_unregister
func (recv_ TeClassTypeT) TeObjectMethodUnregister() {
}

//go:linkname TeObjectCheckChannel C.te_object_check_channel
func TeObjectCheckChannel(channel_array *TouchPadT, channel_sum c.Uint8T) bool

//go:linkname WaterproofCheckMaskHandle C.waterproof_check_mask_handle
func WaterproofCheckMaskHandle(te_handle TouchElemHandleT) bool

//go:linkname TeIsTouchDsleepWakeup C.te_is_touch_dsleep_wakeup
func TeIsTouchDsleepWakeup() bool

//go:linkname TeGetSleepChannel C.te_get_sleep_channel
func TeGetSleepChannel() TouchPadT

//go:linkname IsButtonObjectHandle C.is_button_object_handle
func IsButtonObjectHandle(element_handle TouchElemHandleT) bool

//go:linkname IsSliderObjectHandle C.is_slider_object_handle
func IsSliderObjectHandle(element_handle TouchElemHandleT) bool

//go:linkname IsMatrixObjectHandle C.is_matrix_object_handle
func IsMatrixObjectHandle(element_handle TouchElemHandleT) bool

//go:linkname ButtonEnableWakeupCalibration C.button_enable_wakeup_calibration
func ButtonEnableWakeupCalibration(button_handle TeButtonHandleT, en bool)

//go:linkname SliderEnableWakeupCalibration C.slider_enable_wakeup_calibration
func SliderEnableWakeupCalibration(slider_handle TeSliderHandleT, en bool)

//go:linkname MatrixEnableWakeupCalibration C.matrix_enable_wakeup_calibration
func MatrixEnableWakeupCalibration(matrix_handle TeMatrixHandleT, en bool)

package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const RTCIO_LL_PIN_FUNC = 0

type RtcioLlFuncT c.Int

const (
	RTCIO_LL_FUNC_RTC     RtcioLlFuncT = 0
	RTCIO_LL_FUNC_DIGITAL RtcioLlFuncT = 1
)

type RtcioLlWakeTypeT c.Int

const (
	RTCIO_LL_WAKEUP_DISABLE    RtcioLlWakeTypeT = 0
	RTCIO_LL_WAKEUP_LOW_LEVEL  RtcioLlWakeTypeT = 4
	RTCIO_LL_WAKEUP_HIGH_LEVEL RtcioLlWakeTypeT = 5
)

type RtcioLlIntrTypeT c.Int

const (
	RTCIO_INTR_DISABLE    RtcioLlIntrTypeT = 0
	RTCIO_INTR_POSEDGE    RtcioLlIntrTypeT = 1
	RTCIO_INTR_NEGEDGE    RtcioLlIntrTypeT = 2
	RTCIO_INTR_ANYEDGE    RtcioLlIntrTypeT = 3
	RTCIO_INTR_LOW_LEVEL  RtcioLlIntrTypeT = 4
	RTCIO_INTR_HIGH_LEVEL RtcioLlIntrTypeT = 5
)

type RtcioLlOutModeT c.Int

const (
	RTCIO_LL_OUTPUT_NORMAL RtcioLlOutModeT = 0
	RTCIO_LL_OUTPUT_OD     RtcioLlOutModeT = 1
)

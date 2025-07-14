package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const RTCIO_LL_GPIO_NUM_OFFSET = 7

type RtcioLlFuncT c.Int

const (
	RTCIO_LL_FUNC_RTC     RtcioLlFuncT = 0
	RTCIO_LL_FUNC_DIGITAL RtcioLlFuncT = 1
)

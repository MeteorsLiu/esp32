package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GptimerClockSourceT SocPeriphGptimerClkSrcT
type GptimerCountDirectionT c.Int

const (
	GPTIMER_COUNT_DOWN GptimerCountDirectionT = 0
	GPTIMER_COUNT_UP   GptimerCountDirectionT = 1
)

type GptimerEtmTaskTypeT c.Int

const (
	GPTIMER_ETM_TASK_START_COUNT GptimerEtmTaskTypeT = 0
	GPTIMER_ETM_TASK_STOP_COUNT  GptimerEtmTaskTypeT = 1
	GPTIMER_ETM_TASK_EN_ALARM    GptimerEtmTaskTypeT = 2
	GPTIMER_ETM_TASK_RELOAD      GptimerEtmTaskTypeT = 3
	GPTIMER_ETM_TASK_CAPTURE     GptimerEtmTaskTypeT = 4
	GPTIMER_ETM_TASK_MAX         GptimerEtmTaskTypeT = 5
)

type GptimerEtmEventTypeT c.Int

const (
	GPTIMER_ETM_EVENT_ALARM_MATCH GptimerEtmEventTypeT = 0
	GPTIMER_ETM_EVENT_MAX         GptimerEtmEventTypeT = 1
)

package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LpCoreEtmEventTypeT c.Int

const (
	LP_CORE_EVENT_ERR_INTR   LpCoreEtmEventTypeT = 0
	LP_CORE_EVENT_START_INTR LpCoreEtmEventTypeT = 1
	LP_CORE_EVENT_MAX        LpCoreEtmEventTypeT = 2
)

type LpCoreEtmTaskTypeT c.Int

const (
	LP_CORE_TASK_WAKEUP_CPU LpCoreEtmTaskTypeT = 0
	LP_CORE_TASK_MAX        LpCoreEtmTaskTypeT = 1
)

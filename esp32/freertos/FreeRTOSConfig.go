package freertos

import _ "unsafe"

const STACK_OVERHEAD_CHECKER = 0
const STACK_OVERHEAD_OPTIMIZATION = 0
const STACK_OVERHEAD_APPTRACE = 0
const STACK_OVERHEAD_WATCHPOINT = 0
const ConfigUSE_PREEMPTION = 1
const ConfigUSE_TIME_SLICING = 1
const ConfigUSE_16_BIT_TICKS = 0
const ConfigIDLE_SHOULD_YIELD = 0
const ConfigKERNEL_INTERRUPT_PRIORITY = 1
const ConfigUSE_MUTEXES = 1
const ConfigUSE_RECURSIVE_MUTEXES = 1
const ConfigUSE_COUNTING_SEMAPHORES = 1
const ConfigUSE_QUEUE_SETS = 1
const ConfigUSE_TASK_NOTIFICATIONS = 1
const ConfigENABLE_BACKWARD_COMPATIBILITY = 0
const ConfigSUPPORT_STATIC_ALLOCATION = 1
const ConfigSUPPORT_DYNAMIC_ALLOCATION = 1
const ConfigAPPLICATION_ALLOCATED_HEAP = 1
const ConfigUSE_IDLE_HOOK = 0
const ConfigUSE_TICK_HOOK = 0
const ConfigCHECK_FOR_STACK_OVERFLOW = 2
const ConfigRECORD_STACK_HIGH_ADDRESS = 1
const ConfigUSE_CO_ROUTINES = 0
const ConfigMAX_CO_ROUTINE_PRIORITIES = 2
const ConfigUSE_TIMERS = 1
const INCLUDE_vTaskPrioritySet = 1
const INCLUDE_uxTaskPriorityGet = 1
const INCLUDE_vTaskDelete = 1
const INCLUDE_vTaskSuspend = 1
const INCLUDE_vTaskDelay = 1
const INCLUDE_xTaskGetIdleTaskHandle = 1
const INCLUDE_xTaskAbortDelay = 1
const INCLUDE_xSemaphoreGetMutexHolder = 1
const INCLUDE_xTaskGetHandle = 1
const INCLUDE_uxTaskGetStackHighWaterMark = 1
const INCLUDE_eTaskGetState = 1
const INCLUDE_xTaskResumeFromISR = 1
const INCLUDE_xTimerPendFunctionCall = 1
const INCLUDE_xTaskGetSchedulerState = 1
const INCLUDE_xTaskGetCurrentTaskHandle = 1
const ConfigTHREAD_LOCAL_STORAGE_DELETE_CALLBACKS = 1
const ConfigCHECK_MUTEX_GIVEN_BY_OWNER = 1

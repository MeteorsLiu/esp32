package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspSleepEventCbIndexT c.Int

const (
	SLEEP_EVENT_HW_EXIT_SLEEP       EspSleepEventCbIndexT = 0
	SLEEP_EVENT_SW_CLK_READY        EspSleepEventCbIndexT = 1
	SLEEP_EVENT_SW_EXIT_SLEEP       EspSleepEventCbIndexT = 2
	SLEEP_EVENT_SW_GOTO_SLEEP       EspSleepEventCbIndexT = 3
	SLEEP_EVENT_HW_TIME_START       EspSleepEventCbIndexT = 4
	SLEEP_EVENT_HW_GOTO_SLEEP       EspSleepEventCbIndexT = 5
	SLEEP_EVENT_SW_CPU_TO_MEM_START EspSleepEventCbIndexT = 6
	SLEEP_EVENT_SW_CPU_TO_MEM_END   EspSleepEventCbIndexT = 7
	SLEEP_EVENT_HW_PLL_EN_START     EspSleepEventCbIndexT = 8
	SLEEP_EVENT_HW_PLL_EN_STOP      EspSleepEventCbIndexT = 9
	SLEEP_EVENT_CB_INDEX_NUM        EspSleepEventCbIndexT = 10
)

// llgo:type C
type EspSleepEventCbT func(c.Pointer, c.Pointer) EspErrT

/**
 * @brief Function entry parameter types for light sleep event callback functions (if CONFIG_FREERTOS_USE_TICKLESS_IDLE)
 */

type X_espSleepEventCbConfigT struct {
	Cb      EspSleepEventCbT
	UserArg c.Pointer
	Prior   c.Uint32T
	Next    *X_espSleepEventCbConfigT
}
type EspSleepEventCbConfigT X_espSleepEventCbConfigT

type X_espSleepEventCbsConfigT struct {
	SleepEventCbConfig [10]*EspSleepEventCbConfigT
}
type EspSleepEventCbsConfigT X_espSleepEventCbsConfigT

/**
 * @brief Designed to execute functions in the esp_sleep_event_cb_config_t linked list
 *
 * @param event_id   Designed to annotate the corresponding event_cb in g_sleep_event_cbs_config
 * @param ext_arg    Designed to pass external parameters
 * @return None
 */
// llgo:link EspSleepEventCbIndexT.EspSleepExecuteEventCallbacks C.esp_sleep_execute_event_callbacks
func (recv_ EspSleepEventCbIndexT) EspSleepExecuteEventCallbacks(ext_arg c.Pointer) {
}

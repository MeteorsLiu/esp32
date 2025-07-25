package esp_system

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Initialize the crosscore interrupt system for this CPU.
 * This needs to be called once on every CPU that is used
 * by FreeRTOS.
 *
 * If multicore FreeRTOS support is enabled, this will be
 * called automatically by the startup code and should not
 * be called manually.
 */
//go:linkname EspCrosscoreIntInit C.esp_crosscore_int_init
func EspCrosscoreIntInit()

/**
 * Send an interrupt to a CPU indicating it should yield its
 * currently running task in favour of a higher-priority task
 * that presumably just woke up.
 *
 * This is used internally by FreeRTOS in multicore mode
 * and should not be called by the user.
 *
 * @param core_id Core that should do the yielding
 */
//go:linkname EspCrosscoreIntSendYield C.esp_crosscore_int_send_yield
func EspCrosscoreIntSendYield(core_id c.Int)

/**
 * Send an interrupt to a CPU indicating it should update its
 * CCOMPARE1 value due to a frequency switch.
 *
 * This is used internally when dynamic frequency switching is
 * enabled, and should not be called from application code.
 *
 * @param core_id Core that should update its CCOMPARE1 value
 */
//go:linkname EspCrosscoreIntSendFreqSwitch C.esp_crosscore_int_send_freq_switch
func EspCrosscoreIntSendFreqSwitch(core_id c.Int)

//go:linkname EspCrosscoreIntSendGdbCall C.esp_crosscore_int_send_gdb_call
func EspCrosscoreIntSendGdbCall(core_id c.Int)

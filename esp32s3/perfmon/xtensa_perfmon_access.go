package perfmon

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**@{*/
/**
 * @brief      Init Performance Monitoor
 *
 * Initialize performance monitor register with define values
 *
 * @param[in] id: performance counter number
 * @param[in] select: select value from PMCTRLx register
 * @param[in] mask: mask value from PMCTRLx register
 * @param[in] kernelcnt: kernelcnt value from PMCTRLx register
 * @param[in] tracelevel: tracelevel value from PMCTRLx register
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if one of the arguments is not correct
 */
//go:linkname XtensaPerfmonInit C.xtensa_perfmon_init
func XtensaPerfmonInit(id c.Int, select_ c.Uint16T, mask c.Uint16T, kernelcnt c.Int, tracelevel c.Int) EspErrT

/**@{*/
/**
 * @brief      Reset PM counter
 *
 * Reset PM counter. Writes 0 to the PMx register.
 *
 * @param[in] id: performance counter number
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if id out of range
 */
//go:linkname XtensaPerfmonReset C.xtensa_perfmon_reset
func XtensaPerfmonReset(id c.Int) EspErrT

/**@{*/
/**
 * @brief      Start PM counters
 *
 * Start all PM counters synchronously. Write 1 to the PGM register
 */
//go:linkname XtensaPerfmonStart C.xtensa_perfmon_start
func XtensaPerfmonStart()

/**@{*/
/**
 * @brief      Stop PM counters
 *
 * Stop all PM counters synchronously. Write 0 to the PGM register
 */
//go:linkname XtensaPerfmonStop C.xtensa_perfmon_stop
func XtensaPerfmonStop()

/**@{*/
/**
 * @brief      Read PM counter
 *
 * Read value of defined PM counter.
 *
 * @param[in] id: performance counter number
 *
 * @return
 *      - Performance counter value
 */
//go:linkname XtensaPerfmonValue C.xtensa_perfmon_value
func XtensaPerfmonValue(id c.Int) c.Uint32T

/**@{*/
/**
 * @brief      Read PM overflow state
 *
 * Read overflow value of defined PM counter.
 *
 * @param[in] id: performance counter number
 *
 * @return
 *      - ESP_OK if there is no overflow (overflow = 0)
 *      - ESP_FAIL if overflow occure (overflow = 1)
 */
//go:linkname XtensaPerfmonOverflow C.xtensa_perfmon_overflow
func XtensaPerfmonOverflow(id c.Int) EspErrT

/**@{*/
/**
 * @brief      Dump PM values
 *
 * Dump all PM register to the console.
 *
 */
//go:linkname XtensaPerfmonDump C.xtensa_perfmon_dump
func XtensaPerfmonDump()

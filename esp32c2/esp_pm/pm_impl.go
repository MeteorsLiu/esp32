package esp_pm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PmModeT c.Int

const (
	PM_MODE_LIGHT_SLEEP PmModeT = 0
	PM_MODE_APB_MIN     PmModeT = 1
	PM_MODE_APB_MAX     PmModeT = 2
	PM_MODE_CPU_MAX     PmModeT = 3
	PM_MODE_COUNT       PmModeT = 4
)

/**
 * @brief Get the mode corresponding to a certain lock
 * @param type lock type
 * @param arg argument value for this lock (passed to esp_pm_lock_create)
 * @return lowest power consumption mode which meets the constraints of the lock
 */
// llgo:link EspPmLockTypeT.EspPmImplGetMode C.esp_pm_impl_get_mode
func (recv_ EspPmLockTypeT) EspPmImplGetMode(arg c.Int) PmModeT {
	return 0
}

/**
 * @brief Get CPU clock frequency by power mode
 * @param mode power mode
 * @return CPU clock frequency
 */
// llgo:link PmModeT.EspPmImplGetCpuFreq C.esp_pm_impl_get_cpu_freq
func (recv_ PmModeT) EspPmImplGetCpuFreq() c.Int {
	return 0
}

type PmTimeT c.Int64T
type PmModeSwitchT c.Int

const (
	MODE_LOCK   PmModeSwitchT = 0
	MODE_UNLOCK PmModeSwitchT = 1
)

/**
 * @brief Switch between power modes when lock is taken or released
 * @param mode pm_mode_t corresponding to the lock being taken or released,
 *             as returned by \ref esp_pm_impl_get_mode
 * @param lock_or_unlock
 *              - MODE_LOCK: lock was taken. Implementation needs to make sure
 *                that the constraints of the lock are met by switching to the
 *                given 'mode' or any of the higher power ones.
 *              - MODE_UNLOCK: lock was released. If all the locks for given
 *                mode are released, and no locks for higher power modes are
 *                taken, implementation can switch to one of lower power modes.
 * @param now timestamp when the lock was taken or released. Passed as
 *            a minor optimization, so that the implementation does not need to
 *            call pm_get_time again.
 */
// llgo:link PmModeT.EspPmImplSwitchMode C.esp_pm_impl_switch_mode
func (recv_ PmModeT) EspPmImplSwitchMode(lock_or_unlock PmModeSwitchT, now PmTimeT) {
}

/**
 * @brief Call once at startup to initialize pm implementation
 */
//go:linkname EspPmImplInit C.esp_pm_impl_init
func EspPmImplInit()

/**
 * @brief Hook function for the idle task
 * Must be called from the IDLE task on each CPU before entering waiti state.
 */
//go:linkname EspPmImplIdleHook C.esp_pm_impl_idle_hook
func EspPmImplIdleHook()

/**
 * @brief Hook function for the interrupt dispatcher
 * Must be called soon after entering the ISR
 */
//go:linkname EspPmImplIsrHook C.esp_pm_impl_isr_hook
func EspPmImplIsrHook()

/**
 * @brief Hook function implementing `waiti` instruction, should be invoked from idle task context
 */
//go:linkname EspPmImplWaiti C.esp_pm_impl_waiti
func EspPmImplWaiti()

// llgo:type C
type SkipLightSleepCbT func() bool

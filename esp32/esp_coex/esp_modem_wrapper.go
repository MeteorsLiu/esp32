package esp_coex

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

//go:linkname EspCoexCommonEnvIsChipWrapper C.esp_coex_common_env_is_chip_wrapper
func EspCoexCommonEnvIsChipWrapper() bool

//go:linkname EspCoexCommonSpinLockCreateWrapper C.esp_coex_common_spin_lock_create_wrapper
func EspCoexCommonSpinLockCreateWrapper() c.Pointer

//go:linkname EspCoexCommonIntDisableWrapper C.esp_coex_common_int_disable_wrapper
func EspCoexCommonIntDisableWrapper(wifi_int_mux c.Pointer) c.Uint32T

//go:linkname EspCoexCommonIntRestoreWrapper C.esp_coex_common_int_restore_wrapper
func EspCoexCommonIntRestoreWrapper(wifi_int_mux c.Pointer, tmp c.Uint32T)

//go:linkname EspCoexCommonTaskYieldFromIsrWrapper C.esp_coex_common_task_yield_from_isr_wrapper
func EspCoexCommonTaskYieldFromIsrWrapper()

//go:linkname EspCoexCommonSemphrCreateWrapper C.esp_coex_common_semphr_create_wrapper
func EspCoexCommonSemphrCreateWrapper(max c.Uint32T, init c.Uint32T) c.Pointer

//go:linkname EspCoexCommonSemphrDeleteWrapper C.esp_coex_common_semphr_delete_wrapper
func EspCoexCommonSemphrDeleteWrapper(semphr c.Pointer)

//go:linkname EspCoexCommonSemphrTakeWrapper C.esp_coex_common_semphr_take_wrapper
func EspCoexCommonSemphrTakeWrapper(semphr c.Pointer, block_time_tick c.Uint32T) c.Int32T

//go:linkname EspCoexCommonSemphrGiveWrapper C.esp_coex_common_semphr_give_wrapper
func EspCoexCommonSemphrGiveWrapper(semphr c.Pointer) c.Int32T

//go:linkname EspCoexCommonTimerDisarmWrapper C.esp_coex_common_timer_disarm_wrapper
func EspCoexCommonTimerDisarmWrapper(timer c.Pointer)

//go:linkname EspCoexCommonTimerDoneWrapper C.esp_coex_common_timer_done_wrapper
func EspCoexCommonTimerDoneWrapper(ptimer c.Pointer)

//go:linkname EspCoexCommonTimerSetfnWrapper C.esp_coex_common_timer_setfn_wrapper
func EspCoexCommonTimerSetfnWrapper(ptimer c.Pointer, pfunction c.Pointer, parg c.Pointer)

//go:linkname EspCoexCommonTimerArmUsWrapper C.esp_coex_common_timer_arm_us_wrapper
func EspCoexCommonTimerArmUsWrapper(ptimer c.Pointer, us c.Uint32T, repeat bool)

//go:linkname EspCoexCommonMallocInternalWrapper C.esp_coex_common_malloc_internal_wrapper
func EspCoexCommonMallocInternalWrapper(size c.SizeT) c.Pointer

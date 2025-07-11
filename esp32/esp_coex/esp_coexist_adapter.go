package esp_coex

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const COEX_ADAPTER_VERSION = 0x00000002
const COEX_ADAPTER_MAGIC = 0xDEADBEAF
const COEX_ADAPTER_FUNCS_TIME_BLOCKING = 0xffffffff

type CoexAdapterFuncsT struct {
	X_version           c.Int32T
	X_spinLockCreate    c.Pointer
	X_spinLockDelete    c.Pointer
	X_intDisable        c.Pointer
	X_intEnable         c.Pointer
	X_taskYieldFromIsr  c.Pointer
	X_semphrCreate      c.Pointer
	X_semphrDelete      c.Pointer
	X_semphrTakeFromIsr c.Pointer
	X_semphrGiveFromIsr c.Pointer
	X_semphrTake        c.Pointer
	X_semphrGive        c.Pointer
	X_isInIsr           c.Pointer
	X_mallocInternal    c.Pointer
	X_free              c.Pointer
	X_espTimerGetTime   c.Pointer
	X_envIsChip         c.Pointer
	X_timerDisarm       c.Pointer
	X_timerDone         c.Pointer
	X_timerSetfn        c.Pointer
	X_timerArmUs        c.Pointer
	X_debugMatrixInit   c.Pointer
	X_xtalFreqGet       c.Pointer
	X_magic             c.Int32T
}

type CoexVersionT struct {
	Major c.Uint8T
	Minor c.Uint8T
	Patch c.Uint8T
}

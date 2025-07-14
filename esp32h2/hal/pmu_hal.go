package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PmuHalContextT struct {
	Dev *PmuDevT
}

// llgo:link (*PmuHalContextT).PmuHalHpSetDigitalPowerUpWaitCycle C.pmu_hal_hp_set_digital_power_up_wait_cycle
func (recv_ *PmuHalContextT) PmuHalHpSetDigitalPowerUpWaitCycle(power_supply_wait_cycle c.Uint32T, power_up_wait_cycle c.Uint32T) {
}

// llgo:link (*PmuHalContextT).PmuHalHpGetDigitalPowerUpWaitCycle C.pmu_hal_hp_get_digital_power_up_wait_cycle
func (recv_ *PmuHalContextT) PmuHalHpGetDigitalPowerUpWaitCycle() c.Uint32T {
	return 0
}

// llgo:link (*PmuHalContextT).PmuHalLpSetDigitalPowerUpWaitCycle C.pmu_hal_lp_set_digital_power_up_wait_cycle
func (recv_ *PmuHalContextT) PmuHalLpSetDigitalPowerUpWaitCycle(power_supply_wait_cycle c.Uint32T, power_up_wait_cycle c.Uint32T) {
}

// llgo:link (*PmuHalContextT).PmuHalLpGetDigitalPowerUpWaitCycle C.pmu_hal_lp_get_digital_power_up_wait_cycle
func (recv_ *PmuHalContextT) PmuHalLpGetDigitalPowerUpWaitCycle() c.Uint32T {
	return 0
}

// llgo:link (*PmuHalContextT).PmuHalHpSetSleepActiveBackupEnable C.pmu_hal_hp_set_sleep_active_backup_enable
func (recv_ *PmuHalContextT) PmuHalHpSetSleepActiveBackupEnable() {
}

// llgo:link (*PmuHalContextT).PmuHalHpSetSleepActiveBackupDisable C.pmu_hal_hp_set_sleep_active_backup_disable
func (recv_ *PmuHalContextT) PmuHalHpSetSleepActiveBackupDisable() {
}

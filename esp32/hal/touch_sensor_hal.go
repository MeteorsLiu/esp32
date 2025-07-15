package hal

import _ "unsafe"

/**
 * Get the touch pad which caused wakeup from deep sleep.
 *
 * @param pad_num pointer to touch pad which caused wakeup.
 */
// llgo:link (*TouchPadT).TouchHalGetWakeupStatus C.touch_hal_get_wakeup_status
func (recv_ *TouchPadT) TouchHalGetWakeupStatus() {
}

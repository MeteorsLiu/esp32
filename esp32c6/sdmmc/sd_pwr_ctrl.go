package sdmmc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SdPwrCtrlDrvT struct {
	Unused [8]uint8
}
type SdPwrCtrlHandleT *SdPwrCtrlDrvT

/**
 * @brief Set SD IO voltage by a registered SD power control driver handle
 *
 * @param[in] handle      SD power control driver handle
 * @param[in] voltage_mv  Voltage in mV
 */
//go:linkname SdPwrCtrlSetIoVoltage C.sd_pwr_ctrl_set_io_voltage
func SdPwrCtrlSetIoVoltage(handle SdPwrCtrlHandleT, voltage_mv c.Int) EspErrT

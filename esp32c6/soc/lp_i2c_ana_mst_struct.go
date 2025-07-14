package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: configure_register */
/** Type of i2c0_ctrl register
 *  need_des
 */

type LpI2cAnaMstI2c0CtrlRegT struct {
	Val c.Uint32T
}

/** Type of i2c0_conf register
 *  need_des
 */

type LpI2cAnaMstI2c0ConfRegT struct {
	Val c.Uint32T
}

/** Type of i2c0_data register
 *  need_des
 */

type LpI2cAnaMstI2c0DataRegT struct {
	Val c.Uint32T
}

/** Type of ana_conf1 register
 *  need_des
 */

type LpI2cAnaMstAnaConf1RegT struct {
	Val c.Uint32T
}

/** Type of nouse register
 *  need_des
 */

type LpI2cAnaMstNouseRegT struct {
	Val c.Uint32T
}

/** Type of device_en register
 *  need_des
 */

type LpI2cAnaMstDeviceEnRegT struct {
	Val c.Uint32T
}

/** Type of date register
 *  need_des
 */

type LpI2cAnaMstDateRegT struct {
	Val c.Uint32T
}

type LpI2cAnaMstDevT struct {
	I2c0Ctrl    LpI2cAnaMstI2c0CtrlRegT
	I2c0Conf    LpI2cAnaMstI2c0ConfRegT
	I2c0Data    LpI2cAnaMstI2c0DataRegT
	AnaConf1    LpI2cAnaMstAnaConf1RegT
	Nouse       LpI2cAnaMstNouseRegT
	DeviceEn    LpI2cAnaMstDeviceEnRegT
	Reserved018 [249]c.Uint32T
	Date        LpI2cAnaMstDateRegT
}

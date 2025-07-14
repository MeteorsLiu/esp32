package esp_phy

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Structure holding PHY init parameters
 */

type EspPhyInitDataT struct {
	Params [128]c.Uint8T
}
type EspPhyModemT c.Int

const (
	PHY_MODEM_WIFI       EspPhyModemT = 1
	PHY_MODEM_BT         EspPhyModemT = 2
	PHY_MODEM_IEEE802154 EspPhyModemT = 4
	PHY_MODEM_MAX        EspPhyModemT = 5
)

/**
 * @brief Opaque PHY calibration data
 */

type EspPhyCalibrationDataT struct {
	Version [4]c.Uint8T
	Mac     [6]c.Uint8T
	Opaque  [1894]c.Uint8T
}
type EspPhyCalibrationModeT c.Int

const (
	PHY_RF_CAL_PARTIAL EspPhyCalibrationModeT = 0
	PHY_RF_CAL_NONE    EspPhyCalibrationModeT = 1
	PHY_RF_CAL_FULL    EspPhyCalibrationModeT = 2
)

/**
 * @brief Enable PHY and RF module
 *
 * PHY and RF module should be enabled in order to use WiFi or BT.
 * Now PHY and RF enabling job is done automatically when start WiFi or BT. Users should not
 * call this API in their application.
 *
 * @param modem the modem to call the phy enable.
 */
// llgo:link EspPhyModemT.EspPhyEnable C.esp_phy_enable
func (recv_ EspPhyModemT) EspPhyEnable() {
}

/**
 * @brief Disable PHY and RF module
 *
 * PHY module should be disabled in order to shutdown WiFi or BT.
 * Now PHY and RF disabling job is done automatically when stop WiFi or BT. Users should not
 * call this API in their application.
 *
 * @param modem the modem to call the phy disable.
 */
// llgo:link EspPhyModemT.EspPhyDisable C.esp_phy_disable
func (recv_ EspPhyModemT) EspPhyDisable() {
}

/**
 * @brief Enable BTBB module
 *
 * BTBB module should be enabled in order to use IEEE802154 or BT.
 * Now BTBB enabling job is done automatically when start IEEE802154 or BT. Users should not
 * call this API in their application.
 *
 */
//go:linkname EspBtbbEnable C.esp_btbb_enable
func EspBtbbEnable()

/**
 * @brief Disable BTBB module
 *
 * Disable BTBB module, used by IEEE802154 or Bluetooth.
 * Users should not call this API in their application.
 *
 */
//go:linkname EspBtbbDisable C.esp_btbb_disable
func EspBtbbDisable()

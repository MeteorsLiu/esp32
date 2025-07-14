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
 * @brief Get PHY init data
 *
 * If "Use a partition to store PHY init data" option is set in menuconfig,
 * This function will load PHY init data from a partition. Otherwise,
 * PHY init data will be compiled into the application itself, and this function
 * will return a pointer to PHY init data located in read-only memory (DROM).
 *
 * If "Use a partition to store PHY init data" option is enabled, this function
 * may return NULL if the data loaded from flash is not valid.
 *
 * @note Call esp_phy_release_init_data to release the pointer obtained using
 * this function after the call to esp_wifi_init.
 *
 * @return pointer to PHY init data structure
 */
//go:linkname EspPhyGetInitData C.esp_phy_get_init_data
func EspPhyGetInitData() *EspPhyInitDataT

/**
 * @brief Release PHY init data
 * @param data  pointer to PHY init data structure obtained from
 *              esp_phy_get_init_data function
 */
// llgo:link (*EspPhyInitDataT).EspPhyReleaseInitData C.esp_phy_release_init_data
func (recv_ *EspPhyInitDataT) EspPhyReleaseInitData() {
}

/**
 * @brief Function called by esp_phy_load_cal_and_init to load PHY calibration data
 *
 * This is a convenience function which can be used to load PHY calibration
 * data from NVS. Data can be stored to NVS using esp_phy_store_cal_data_to_nvs
 * function.
 *
 * If calibration data is not present in the NVS, or
 * data is not valid (was obtained for a chip with a different MAC address,
 * or obtained for a different version of software), this function will
 * return an error.
 *
 * @param out_cal_data pointer to calibration data structure to be filled with
 *                     loaded data.
 * @return ESP_OK on success
 */
// llgo:link (*EspPhyCalibrationDataT).EspPhyLoadCalDataFromNvs C.esp_phy_load_cal_data_from_nvs
func (recv_ *EspPhyCalibrationDataT) EspPhyLoadCalDataFromNvs() EspErrT {
	return 0
}

/**
 * @brief Function called by esp_phy_load_cal_and_init to store PHY calibration data
 *
 * This is a convenience function which can be used to store PHY calibration
 * data to the NVS. Calibration data is returned by esp_phy_load_cal_and_init function.
 * Data saved using this function to the NVS can later be loaded using
 * esp_phy_store_cal_data_to_nvs function.
 *
 * @param cal_data pointer to calibration data which has to be saved.
 * @return ESP_OK on success
 */
// llgo:link (*EspPhyCalibrationDataT).EspPhyStoreCalDataToNvs C.esp_phy_store_cal_data_to_nvs
func (recv_ *EspPhyCalibrationDataT) EspPhyStoreCalDataToNvs() EspErrT {
	return 0
}

/**
 * @brief Erase PHY calibration data which is stored in the NVS
 *
 * This is a function which can be used to trigger full calibration as a last-resort remedy
 * if partial calibration is used. It can be called in the application based on some conditions
 * (e.g. an option provided in some diagnostic mode).
 *
 * @return ESP_OK on success
 * @return others on fail. Please refer to NVS API return value error number.
 */
//go:linkname EspPhyEraseCalDataInNvs C.esp_phy_erase_cal_data_in_nvs
func EspPhyEraseCalDataInNvs() EspErrT

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

/**
 * @brief Load calibration data from NVS and initialize PHY and RF module
 */
//go:linkname EspPhyLoadCalAndInit C.esp_phy_load_cal_and_init
func EspPhyLoadCalAndInit()

/**
 * @brief Initialize backup memory for Phy power up/down
 */
//go:linkname EspPhyModemInit C.esp_phy_modem_init
func EspPhyModemInit()

/**
 * @brief Deinitialize backup memory for Phy power up/down
 * Set phy_init_flag if all modems deinit on ESP32C3
 */
//go:linkname EspPhyModemDeinit C.esp_phy_modem_deinit
func EspPhyModemDeinit()

/**
 * @brief Enable WiFi/BT common clock
 *
 */
//go:linkname EspPhyCommonClockEnable C.esp_phy_common_clock_enable
func EspPhyCommonClockEnable()

/**
 * @brief Disable WiFi/BT common clock
 *
 */
//go:linkname EspPhyCommonClockDisable C.esp_phy_common_clock_disable
func EspPhyCommonClockDisable()

/**
 * @brief Update the corresponding PHY init type according to the country code of Wi-Fi.
 *
 * @param country country code
 * @return ESP_OK on success.
 * @return esp_err_t code describing the error on fail
 */
//go:linkname EspPhyUpdateCountryInfo C.esp_phy_update_country_info
func EspPhyUpdateCountryInfo(country *c.Char) EspErrT

/**
 * @brief Get PHY lib version
 * @return PHY lib version.
 */
//go:linkname GetPhyVersionStr C.get_phy_version_str
func GetPhyVersionStr() *c.Char

/**
 * @brief Set PHY init parameters
 * @param param is 1 means combo module
 */
//go:linkname PhyInitParamSet C.phy_init_param_set
func PhyInitParamSet(param c.Uint8T)

/**
 * @brief Power on Bluetooth Wi-Fi power domain
 */
//go:linkname EspWifiBtPowerDomainOn C.esp_wifi_bt_power_domain_on
func EspWifiBtPowerDomainOn()

/**
 * @brief Power off Bluetooth Wi-Fi power domain
 */
//go:linkname EspWifiBtPowerDomainOff C.esp_wifi_bt_power_domain_off
func EspWifiBtPowerDomainOff()

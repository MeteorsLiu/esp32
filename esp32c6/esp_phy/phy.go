package esp_phy

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_CAL_DATA_CHECK_FAIL = 1

type PhyI2cMasterCommandAttributeT struct {
	CmdType c.Uint8T
	Config  struct {
		Start  c.Uint8T
		End    c.Uint8T
		HostId c.Uint8T
	}
}

/**
 * @brief Initialize PHY module and do RF calibration
 * @param[in] init_data Initialization parameters to be used by the PHY
 * @param[inout] cal_data As input, calibration data previously obtained. As output, will contain new calibration data.
 * @param[in] cal_mode  RF calibration mode
 * @return ESP_CAL_DATA_CHECK_FAIL if the calibration data checksum fails or if the calibration data is outdated, other values are reserved for future use
 */
// llgo:link (*EspPhyInitDataT).RegisterChipv7Phy C.register_chipv7_phy
func (recv_ *EspPhyInitDataT) RegisterChipv7Phy(cal_data *EspPhyCalibrationDataT, cal_mode EspPhyCalibrationModeT) c.Int {
	return 0
}

/**
 * @brief Get the format version of calibration data used by PHY library.
 * @return Format version number, OR'ed with BIT(16) if PHY is in WIFI only mode.
 */
//go:linkname PhyGetRfCalVersion C.phy_get_rf_cal_version
func PhyGetRfCalVersion() c.Uint32T

/**
 * @brief Open PHY and RF.
 */
//go:linkname PhyWakeupInit C.phy_wakeup_init
func PhyWakeupInit()

/**
 * @brief Shutdown PHY and RF.
 */
//go:linkname PhyCloseRf C.phy_close_rf
func PhyCloseRf()

/**
 * @brief Disable PHY temperature sensor.
 */
//go:linkname PhyXpdTsens C.phy_xpd_tsens
func PhyXpdTsens()

/**
 * @brief Get the configuration info of PHY i2c master command memory.
 *
 * @param[out] attr  the configuration info of PHY i2c master command memory
 * @param[out] size  the count of PHY i2c master command memory configuration
 */
// llgo:link (*PhyI2cMasterCommandAttributeT).PhyI2cMasterCommandMemCfg C.phy_i2c_master_command_mem_cfg
func (recv_ *PhyI2cMasterCommandAttributeT) PhyI2cMasterCommandMemCfg(size *c.Int) {
}

/**
 * @brief Enable phy track pll
 *
 */
//go:linkname PhyTrackPllInit C.phy_track_pll_init
func PhyTrackPllInit()

/**
 * @brief Disable phy track pll
 *
 */
//go:linkname PhyTrackPllDeinit C.phy_track_pll_deinit
func PhyTrackPllDeinit()

/**
 * @brief Set the flag recorded which modem has already enabled phy
 *
 */
// llgo:link EspPhyModemT.PhySetModemFlag C.phy_set_modem_flag
func (recv_ EspPhyModemT) PhySetModemFlag() {
}

/**
 * @brief Clear the flag to record which modem calls phy disenable
 */
// llgo:link EspPhyModemT.PhyClrModemFlag C.phy_clr_modem_flag
func (recv_ EspPhyModemT) PhyClrModemFlag() {
}

/**
 * @brief Get the flag recorded which modem has already enabled phy
 *
 */
//go:linkname PhyGetModemFlag C.phy_get_modem_flag
func PhyGetModemFlag() EspPhyModemT

/**
 * @brief Get the PHY lock, only used in esp_phy, the user should not use this function.
 *
 */
//go:linkname PhyGetLock C.phy_get_lock
func PhyGetLock() X_lockT

/**
 * @brief Call this funnction to track pll immediately.
 *
 */
//go:linkname PhyTrackPll C.phy_track_pll
func PhyTrackPll()

/**
 * @brief PHY antenna default configuration
 *
 */
//go:linkname AntDftCfg C.ant_dft_cfg
func AntDftCfg(default_ant bool)

/**
 * @brief PHY tx antenna config
 *
 */
//go:linkname AntTxCfg C.ant_tx_cfg
func AntTxCfg(ant0 c.Uint8T)

/**
 * @brief PHY rx antenna config
 *
 */
//go:linkname AntRxCfg C.ant_rx_cfg
func AntRxCfg(auto_en bool, ant0 c.Uint8T, ant1 c.Uint8T)

/**
 * @brief PHY antenna need update
 *
 */
//go:linkname PhyAntNeedUpdate C.phy_ant_need_update
func PhyAntNeedUpdate() bool

/**
 * @brief PHY antenna need update
 *
 */
//go:linkname PhyAntClrUpdateFlag C.phy_ant_clr_update_flag
func PhyAntClrUpdateFlag()

/**
 * @brief PHY antenna configuration update
 *
 */
//go:linkname PhyAntUpdate C.phy_ant_update
func PhyAntUpdate()

/**
 * @brief Get the REGDMA config value of the BBPLL in analog i2c master burst mode
 *
 * @return  the BBPLL REGDMA configure value of i2c master burst mode
 */
//go:linkname PhyAnaI2cMasterBurstBbpllConfig C.phy_ana_i2c_master_burst_bbpll_config
func PhyAnaI2cMasterBurstBbpllConfig() c.Uint32T

/**
 * @brief Get the REGDMA config value of the RF PHY on or off in analog i2c master burst mode
 *
 * @param[in] on true for enable RF PHY, false for disable RF PHY.
 *
 * @return  the RF on or off configure value of i2c master burst mode
 */
//go:linkname PhyAnaI2cMasterBurstRfOnoff C.phy_ana_i2c_master_burst_rf_onoff
func PhyAnaI2cMasterBurstRfOnoff(on bool) c.Uint32T

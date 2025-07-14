package esp_driver_i2s

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief LP I2S pin configurations
 */

type LpI2sPdmRxGpioConfigT struct {
	Clk c.Int
	Din c.Int
}

/**
 * @brief I2S slot configuration for PDM RX mode
 */

type LpI2sPdmRxSlotConfigT struct {
	DataBitWidth   I2sDataBitWidthT
	SlotBitWidth   I2sSlotBitWidthT
	SlotMode       I2sSlotModeT
	SlotMask       I2sPdmSlotMaskT
	WsPol          bool
	HpEn           bool
	HpCutOffFreqHz c.Float
	AmplifyNum     c.Uint32T
}

/**
 * @brief LP I2S PDM configuration
 */

type LpI2sPdmRxConfigT struct {
	PinCfg  LpI2sPdmRxGpioConfigT
	SlotCfg LpI2sPdmRxSlotConfigT
}

/**
 * @brief Init LP I2S to PDM mode
 *
 * @param[in] handle   LP I2S channel handle
 * @param[in] pdm_cfg  PDM configuration
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state
 */
//go:linkname LpI2sChannelInitPdmRxMode C.lp_i2s_channel_init_pdm_rx_mode
func LpI2sChannelInitPdmRxMode(handle LpI2sChanHandleT, pdm_cfg *LpI2sPdmRxConfigT) EspErrT

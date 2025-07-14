package esp_driver_i2s

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief LP I2S pin configurations
 */

type LpI2sStdGpioConfigT struct {
	Bck c.Int
	Ws  c.Int
	Din c.Int
}

/**
 * @brief LP I2S slot configuration for standard mode
 */

type LpI2sStdSlotConfigT struct {
	DataBitWidth I2sDataBitWidthT
	SlotBitWidth I2sSlotBitWidthT
	SlotMode     I2sSlotModeT
	SlotMask     I2sStdSlotMaskT
	WsWidth      c.Uint32T
	WsPol        bool
	BitShift     bool
	LeftAlign    bool
	BigEndian    bool
	BitOrderLsb  bool
}

/**
 * @brief LP I2S STD configuration
 */

type LpI2sStdConfigT struct {
	PinCfg  LpI2sStdGpioConfigT
	SlotCfg LpI2sStdSlotConfigT
}

/**
 * @brief Init LP I2S to STD mode
 *
 * @param[in] handle   LP I2S channel handle
 * @param[in] std_cfg  STD configuration
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_ERR_INVALID_ARG:   Invalid argument
 *        - ESP_ERR_INVALID_STATE: Invalid state
 */
//go:linkname LpI2sChannelInitStdMode C.lp_i2s_channel_init_std_mode
func LpI2sChannelInitStdMode(handle LpI2sChanHandleT, std_cfg *LpI2sStdConfigT) EspErrT

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

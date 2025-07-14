package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type CamDevT LcdCamDevT

/**
 * @brief CAM hardware interface object data
 */

type CamHalContext struct {
	Hw *CamDevT
}
type CamHalContextT CamHalContext

/**
 * @brief CAM HAL driver configuration
 */

type CamHalConfig struct {
	Port       c.Int
	ByteSwapEn bool
}
type CamHalConfigT CamHalConfig

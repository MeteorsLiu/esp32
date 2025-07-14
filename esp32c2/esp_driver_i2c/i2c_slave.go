package esp_driver_i2c

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief I2C slave specific configurations
 */

type I2cSlaveConfigT struct {
	I2cPort      I2cPortNumT
	SdaIoNum     GpioNumT
	SclIoNum     GpioNumT
	ClkSource    I2cClockSourceT
	SendBufDepth c.Uint32T
	SlaveAddr    c.Uint16T
	AddrBitLen   I2cAddrBitLenT
	IntrPriority c.Int
	Flags        struct {
		AllowPd c.Uint32T
	}
}

/**
 * @brief Group of I2C slave callbacks (e.g. get i2c slave stretch cause). But take care of potential concurrency issues.
 * @note The callbacks are all running under ISR context
 * @note When CONFIG_I2C_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well.
 */

type I2cSlaveEventCallbacksT struct {
	OnRecvDone I2cSlaveReceivedCallbackT
}

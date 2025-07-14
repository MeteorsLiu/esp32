package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TempSensorDacOffsetT c.Int

const (
	TSENS_DAC_L0      TempSensorDacOffsetT = 0
	TSENS_DAC_L1      TempSensorDacOffsetT = 1
	TSENS_DAC_L2      TempSensorDacOffsetT = 2
	TSENS_DAC_L3      TempSensorDacOffsetT = 3
	TSENS_DAC_L4      TempSensorDacOffsetT = 4
	TSENS_DAC_MAX     TempSensorDacOffsetT = 5
	TSENS_DAC_DEFAULT TempSensorDacOffsetT = 2
)

/**
 * @brief Configuration for temperature sensor reading
 */

type TempSensorConfigT struct {
	DacOffset TempSensorDacOffsetT
	ClkDiv    c.Uint8T
}

package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const CLK_LL_RC_FAST_WAIT_DEFAULT = 20
const CLK_LL_RC_FAST_ENABLE_WAIT_DEFAULT = 5
const CLK_LL_APLL_SDM_STOP_VAL_1 = 0x09
const CLK_LL_APLL_SDM_STOP_VAL_2_REV0 = 0x69
const CLK_LL_APLL_SDM_STOP_VAL_2_REV1 = 0x49
const CLK_LL_APLL_CAL_DELAY_1 = 0x0f
const CLK_LL_APLL_CAL_DELAY_2 = 0x3f
const CLK_LL_APLL_CAL_DELAY_3 = 0x1f

type ClkLlXtal32kEnableModeT c.Int

const (
	CLK_LL_XTAL32K_ENABLE_MODE_CRYSTAL   ClkLlXtal32kEnableModeT = 0
	CLK_LL_XTAL32K_ENABLE_MODE_EXTERNAL  ClkLlXtal32kEnableModeT = 1
	CLK_LL_XTAL32K_ENABLE_MODE_BOOTSTRAP ClkLlXtal32kEnableModeT = 2
)

/**
 * @brief XTAL32K_CLK configuration structure
 */

type ClkLlXtal32kConfigT struct {
	Dac  c.Uint32T
	Dres c.Uint32T
	Dgm  c.Uint32T
	Dbuf c.Uint32T
}

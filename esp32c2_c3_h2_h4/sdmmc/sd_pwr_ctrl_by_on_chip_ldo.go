package sdmmc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief LDO configurations
 */

type SdPwrCtrlLdoConfigT struct {
	LdoChanId c.Int
}

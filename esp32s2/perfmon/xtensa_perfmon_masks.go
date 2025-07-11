package perfmon

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const MAX_PERFMON_EVENTS = 119

/**
 * @brief Description for select parameter
 *
 * Structure defines description for different select values for performance counters
 */

type XtensaPerfmonSelect struct {
	Select      c.Int
	Description *c.Char
}
type XtensaPerfmonSelectT XtensaPerfmonSelect

/**
 * @brief Description for mask parameter
 *
 * Structure defines description for different select and mask values for performance counters
 */

type XtensaPerfmonMasks struct {
	Select      c.Int
	Mask        c.Int
	Description *c.Char
}
type XtensaPerfmonMasksT XtensaPerfmonMasks

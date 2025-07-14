package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Control/Status registers */
/** Type of set_start register
 *  DS start control register
 */

type DsSetStartRegT struct {
	Val c.Uint32T
}

/** Type of set_continue register
 *  DS continue control register
 */

type DsSetContinueRegT struct {
	Val c.Uint32T
}

/** Type of set_finish register
 *  DS finish control register
 */

type DsSetFinishRegT struct {
	Val c.Uint32T
}

/** Type of query_busy register
 *  DS query busy register
 */

type DsQueryBusyRegT struct {
	Val c.Uint32T
}

/** Type of query_key_wrong register
 *  DS query key-wrong counter register
 */

type DsQueryKeyWrongRegT struct {
	Val c.Uint32T
}

/** Type of query_check register
 *  DS query check result register
 */

type DsQueryCheckRegT struct {
	Val c.Uint32T
}

/** Group: version control register */
/** Type of date register
 *  DS version control register
 */

type DsDateRegT struct {
	Val c.Uint32T
}

type DsDevT struct {
	Y             [128]c.Uint32T
	M             [128]c.Uint32T
	Rb            [128]c.Uint32T
	Box           [12]c.Uint32T
	Iv            [4]c.Uint32T
	Reserved640   [112]c.Uint32T
	X             [128]c.Uint32T
	Z             [128]c.Uint32T
	ReservedC00   [128]c.Uint32T
	SetStart      DsSetStartRegT
	SetContinue   DsSetContinueRegT
	SetFinish     DsSetFinishRegT
	QueryBusy     DsQueryBusyRegT
	QueryKeyWrong DsQueryKeyWrongRegT
	QueryCheck    DsQueryCheckRegT
	ReservedE18   [2]c.Uint32T
	Date          DsDateRegT
}

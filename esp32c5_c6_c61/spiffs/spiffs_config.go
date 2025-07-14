package spiffs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SPIFFS_TAG = "SPIFFS"
const X_SPIPRIad = "%08x"
const X_SPIPRIbl = "%04x"
const X_SPIPRIpg = "%04x"
const X_SPIPRIsp = "%04x"
const X_SPIPRIfd = "%d"
const X_SPIPRIid = "%04x"
const X_SPIPRIfl = "%02x"
const SPIFFS_BUFFER_HELP = 0
const SPIFFS_SINGLETON = 0
const SPIFFS_ALIGNED_OBJECT_INDEX_TABLES = 0
const SPIFFS_HAL_CALLBACK_EXTRA = 1
const SPIFFS_FILEHDL_OFFSET = 0
const SPIFFS_READ_ONLY = 0
const SPIFFS_TEMPORAL_FD_CACHE = 1
const SPIFFS_TEMPORAL_CACHE_HIT_SCORE = 4
const SPIFFS_IX_MAP = 1
const SPIFFS_TEST_VISUALISATION = 0

type S32T c.Int32T
type U32T c.Uint32T
type S16T c.Int16T
type U16T c.Uint16T
type S8T c.Int8T
type U8T c.Uint8T

type SpiffsT struct {
	Unused [8]uint8
}

// llgo:link (*SpiffsT).SpiffsApiLock C.spiffs_api_lock
func (recv_ *SpiffsT) SpiffsApiLock() {
}

// llgo:link (*SpiffsT).SpiffsApiUnlock C.spiffs_api_unlock
func (recv_ *SpiffsT) SpiffsApiUnlock() {
}

type SpiffsBlockIx U16T
type SpiffsPageIx U16T
type SpiffsObjId U16T
type SpiffsSpanIx U16T

package esp_hid

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type HiddLeCccValueT struct {
	Value c.Uint16T
}

type HiddReportItemT struct {
	MapIndex     c.Uint8T
	ReportId     c.Uint8T
	ReportType   c.Uint8T
	ProtocolMode c.Uint8T
	Usage        EspHidUsageT
	ValueLen     c.Uint16T
	Index        c.Uint8T
	Handle       c.Uint16T
	CccHandle    c.Uint16T
	Ccc          HiddLeCccValueT
}

//go:linkname EspHiddProcessEventDataHandler C.esp_hidd_process_event_data_handler
func EspHiddProcessEventDataHandler(event_handler_arg c.Pointer, event_base EspEventBaseT, event_id c.Int32T, event_data c.Pointer)

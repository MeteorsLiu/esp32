package esp_pm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspPmTraceEventT c.Int

const (
	ESP_PM_TRACE_IDLE            EspPmTraceEventT = 0
	ESP_PM_TRACE_TICK            EspPmTraceEventT = 1
	ESP_PM_TRACE_FREQ_SWITCH     EspPmTraceEventT = 2
	ESP_PM_TRACE_CCOMPARE_UPDATE EspPmTraceEventT = 3
	ESP_PM_TRACE_ISR_HOOK        EspPmTraceEventT = 4
	ESP_PM_TRACE_SLEEP           EspPmTraceEventT = 5
	ESP_PM_TRACE_TYPE_MAX        EspPmTraceEventT = 6
)

//go:linkname EspPmTraceInit C.esp_pm_trace_init
func EspPmTraceInit()

// llgo:link EspPmTraceEventT.EspPmTraceEnter C.esp_pm_trace_enter
func (recv_ EspPmTraceEventT) EspPmTraceEnter(core_id c.Int) {
}

// llgo:link EspPmTraceEventT.EspPmTraceExit C.esp_pm_trace_exit
func (recv_ EspPmTraceEventT) EspPmTraceExit(core_id c.Int) {
}

package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspEtmEventT struct {
	Unused [8]uint8
}

type EspEtmTaskT struct {
	Unused [8]uint8
}
type EtmTriggerPeripheralT c.Int

const (
	ETM_TRIG_PERIPH_GPIO     EtmTriggerPeripheralT = 0
	ETM_TRIG_PERIPH_GDMA     EtmTriggerPeripheralT = 1
	ETM_TRIG_PERIPH_GPTIMER  EtmTriggerPeripheralT = 2
	ETM_TRIG_PERIPH_SYSTIMER EtmTriggerPeripheralT = 3
	ETM_TRIG_PERIPH_MCPWM    EtmTriggerPeripheralT = 4
	ETM_TRIG_PERIPH_ANA_CMPR EtmTriggerPeripheralT = 5
	ETM_TRIG_PERIPH_TSENS    EtmTriggerPeripheralT = 6
	ETM_TRIG_PERIPH_I2S      EtmTriggerPeripheralT = 7
	ETM_TRIG_PERIPH_LP_CORE  EtmTriggerPeripheralT = 8
)

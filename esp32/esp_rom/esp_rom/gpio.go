package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const GPIO_ID_PIN0 = 0
const GPIO_FUNC_IN_HIGH = 0x38
const GPIO_FUNC_IN_LOW = 0x30

type GPIOINTTYPE c.Int

const (
	GPIO_PIN_INTR_DISABLE GPIOINTTYPE = 0
	GPIO_PIN_INTR_POSEDGE GPIOINTTYPE = 1
	GPIO_PIN_INTR_NEGEDGE GPIOINTTYPE = 2
	GPIO_PIN_INTR_ANYEDGE GPIOINTTYPE = 3
	GPIO_PIN_INTR_LOLEVEL GPIOINTTYPE = 4
	GPIO_PIN_INTR_HILEVEL GPIOINTTYPE = 5
)

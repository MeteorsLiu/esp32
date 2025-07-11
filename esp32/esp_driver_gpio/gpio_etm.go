package esp_driver_gpio

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const GPIO_ETM_EVENT_EDGE_TYPES = 3
const GPIO_ETM_TASK_ACTION_TYPES = 3

type GpioEtmEventEdgeT c.Int

const (
	GPIO_ETM_EVENT_EDGE_POS GpioEtmEventEdgeT = 1
	GPIO_ETM_EVENT_EDGE_NEG GpioEtmEventEdgeT = 2
	GPIO_ETM_EVENT_EDGE_ANY GpioEtmEventEdgeT = 3
)

/**
 * @brief GPIO ETM event configuration
 *
 * If more than one kind of ETM edge event want to be triggered on the same GPIO pin, you can configure them together.
 * It helps to save GPIO ETM event channel resources for other GPIOs.
 */

type GpioEtmEventConfigT struct {
	Unused [8]uint8
}
type GpioEtmTaskActionT c.Int

const (
	GPIO_ETM_TASK_ACTION_SET GpioEtmTaskActionT = 1
	GPIO_ETM_TASK_ACTION_CLR GpioEtmTaskActionT = 2
	GPIO_ETM_TASK_ACTION_TOG GpioEtmTaskActionT = 3
)

/**
 * @brief GPIO ETM task configuration
 *
 * If multiple actions wants to be added to the same GPIO pin, you have to configure all the GPIO ETM tasks together.
 */

type GpioEtmTaskConfigT struct {
	Unused [8]uint8
}

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

/**
 * @brief Create an ETM event object for the GPIO peripheral
 *
 * @note The created ETM event object can be deleted later by calling `esp_etm_del_event`
 * @note The newly created ETM event object is not bind to any GPIO, you need to call `gpio_etm_event_bind_gpio` to bind the wanted GPIO
 * @note Every success call to this function will acquire a free GPIO ETM event channel
 *
 * @param[in] config GPIO ETM event configuration
 * @param[out] ret_event Returned ETM event handle
 * @param[out] ...  Other returned ETM event handles if any (the order of the returned event handles is aligned with the array order in field `edges` in `gpio_etm_event_config_t`)
 * @return
 *      - ESP_OK: Create ETM event successfully
 *      - ESP_ERR_INVALID_ARG: Create ETM event failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create ETM event failed because of out of memory
 *      - ESP_ERR_NOT_FOUND: Create ETM event failed because all events are used up and no more free one
 *      - ESP_FAIL: Create ETM event failed because of other reasons
 */
// llgo:link (*GpioEtmEventConfigT).GpioNewEtmEvent C.gpio_new_etm_event
func (recv_ *GpioEtmEventConfigT) GpioNewEtmEvent(ret_event *EspEtmEventHandleT, __llgo_va_list ...interface{}) EspErrT {
	return 0
}

/**
 * @brief Bind the GPIO with the ETM event
 *
 * @note Calling this function multiple times with different GPIO number can override the previous setting immediately.
 * @note Only GPIO ETM object can call this function
 *
 * @param[in] event ETM event handle that created by `gpio_new_etm_event`
 * @param[in] gpio_num GPIO number that can trigger the ETM event
 * @return
 *      - ESP_OK: Set the GPIO for ETM event successfully
 *      - ESP_ERR_INVALID_ARG: Set the GPIO for ETM event failed because of invalid argument, e.g. GPIO is not input capable, ETM event is not of GPIO type
 *      - ESP_FAIL: Set the GPIO for ETM event failed because of other reasons
 */
//go:linkname GpioEtmEventBindGpio C.gpio_etm_event_bind_gpio
func GpioEtmEventBindGpio(event EspEtmEventHandleT, gpio_num c.Int) EspErrT

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

/**
 * @brief Create an ETM task object for the GPIO peripheral
 *
 * @note The created ETM task object can be deleted later by calling `esp_etm_del_task`
 * @note The GPIO ETM task works like a container, a newly created ETM task object doesn't have GPIO members to be managed.
 *       You need to call `gpio_etm_task_add_gpio` to put one or more GPIOs to the container.
 * @note Every success call to this function will acquire a free GPIO ETM task channel
 *
 * @param[in] config GPIO ETM task configuration
 * @param[out] ret_task Returned ETM task handle
 * @param[out] ...  Other returned ETM task handles if any (the order of the returned task handles is aligned with the array order in field `actions` in `gpio_etm_task_config_t`)
 * @return
 *      - ESP_OK: Create ETM task successfully
 *      - ESP_ERR_INVALID_ARG: Create ETM task failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create ETM task failed because of out of memory
 *      - ESP_ERR_NOT_FOUND: Create ETM task failed because all tasks are used up and no more free one
 *      - ESP_FAIL: Create ETM task failed because of other reasons
 */
// llgo:link (*GpioEtmTaskConfigT).GpioNewEtmTask C.gpio_new_etm_task
func (recv_ *GpioEtmTaskConfigT) GpioNewEtmTask(ret_task *EspEtmTaskHandleT, __llgo_va_list ...interface{}) EspErrT {
	return 0
}

/**
 * @brief Add GPIO to the ETM task.
 *
 * @note You can call this function multiple times to add more GPIOs
 * @note Only GPIO ETM object can call this function
 *
 * @param[in] task ETM task handle that created by `gpio_new_etm_task`
 * @param[in] gpio_num GPIO number that can be controlled by the ETM task
 * @return
 *      - ESP_OK: Add GPIO to the ETM task successfully
 *      - ESP_ERR_INVALID_ARG: Add GPIO to the ETM task failed because of invalid argument, e.g. GPIO is not output capable, ETM task is not of GPIO type
 *      - ESP_ERR_INVALID_STATE: Add GPIO to the ETM task failed because the GPIO is used by other ETM task already
 *      - ESP_FAIL: Add GPIO to the ETM task failed because of other reasons
 */
//go:linkname GpioEtmTaskAddGpio C.gpio_etm_task_add_gpio
func GpioEtmTaskAddGpio(task EspEtmTaskHandleT, gpio_num c.Int) EspErrT

/**
 * @brief Remove the GPIO from the ETM task
 *
 * @note Before deleting the ETM task, you need to remove all the GPIOs from the ETM task by this function
 * @note Only GPIO ETM object can call this function
 *
 * @param[in] task ETM task handle that created by `gpio_new_etm_task`
 * @param[in] gpio_num GPIO number that to be remove from the ETM task
 * @return
 *      - ESP_OK: Remove the GPIO from the ETM task successfully
 *      - ESP_ERR_INVALID_ARG: Remove the GPIO from the ETM task failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Remove the GPIO from the ETM task failed because the GPIO is not controlled by this ETM task
 *      - ESP_FAIL: Remove the GPIO from the ETM task failed because of other reasons
 */
//go:linkname GpioEtmTaskRmGpio C.gpio_etm_task_rm_gpio
func GpioEtmTaskRmGpio(task EspEtmTaskHandleT, gpio_num c.Int) EspErrT

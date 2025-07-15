package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 *
 * Type by which event groups are referenced.  For example, a call to
 * xEventGroupCreate() returns an EventGroupHandle_t variable that can then
 * be used as a parameter to other event group functions.
 *
 * \ingroup EventGroup
 */

type EventGroupDefT struct {
	Unused [8]uint8
}
type EventGroupHandleT *EventGroupDefT
type EventBitsT c.Int

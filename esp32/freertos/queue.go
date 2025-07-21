package freertos

import _ "unsafe"

/**
 * Type by which queues are referenced.  For example, a call to xQueueCreate()
 * returns an QueueHandle_t variable that can then be used as a parameter to
 * xQueueSend(), xQueueReceive(), etc.
 */

type QueueDefinition struct {
	Unused [8]uint8
}
type QueueHandleT *QueueDefinition
type QueueSetHandleT *QueueDefinition
type QueueSetMemberHandleT *QueueDefinition

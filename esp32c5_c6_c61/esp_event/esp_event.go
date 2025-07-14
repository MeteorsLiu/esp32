package esp_event

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// / Configuration for creating event loops
type EspEventLoopArgsT struct {
	QueueSize     c.Int32T
	TaskName      *c.Char
	TaskPriority  UBaseTypeT
	TaskStackSize c.Uint32T
	TaskCoreId    BaseTypeT
}

/**
 * @brief Create a new event loop.
 *
 * @param[in] event_loop_args configuration structure for the event loop to create
 * @param[out] event_loop handle to the created event loop
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_ERR_INVALID_ARG: event_loop_args or event_loop was NULL
 *  - ESP_ERR_NO_MEM: Cannot allocate memory for event loops list
 *  - ESP_FAIL: Failed to create task loop
 *  - Others: Fail
 */
// llgo:link (*EspEventLoopArgsT).EspEventLoopCreate C.esp_event_loop_create
func (recv_ *EspEventLoopArgsT) EspEventLoopCreate(event_loop *EspEventLoopHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete an existing event loop.
 *
 * @param[in] event_loop event loop to delete, must not be NULL
 *
 * @return
 *  - ESP_OK: Success
 *  - Others: Fail
 */
//go:linkname EspEventLoopDelete C.esp_event_loop_delete
func EspEventLoopDelete(event_loop EspEventLoopHandleT) EspErrT

/**
 * @brief Create default event loop
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_ERR_NO_MEM: Cannot allocate memory for event loops list
 *  - ESP_ERR_INVALID_STATE: Default event loop has already been created
 *  - ESP_FAIL: Failed to create task loop
 *  - Others: Fail
 */
//go:linkname EspEventLoopCreateDefault C.esp_event_loop_create_default
func EspEventLoopCreateDefault() EspErrT

/**
 * @brief Delete the default event loop
 *
 * @return
 *  - ESP_OK: Success
 *  - Others: Fail
 */
//go:linkname EspEventLoopDeleteDefault C.esp_event_loop_delete_default
func EspEventLoopDeleteDefault() EspErrT

/**
 * @brief Dispatch events posted to an event loop.
 *
 * This function is used to dispatch events posted to a loop with no dedicated task, i.e. task name was set to NULL
 * in event_loop_args argument during loop creation. This function includes an argument to limit the amount of time
 * it runs, returning control to the caller when that time expires (or some time afterwards). There is no guarantee
 * that a call to this function will exit at exactly the time of expiry. There is also no guarantee that events have
 * been dispatched during the call, as the function might have spent all the allotted time waiting on the event queue.
 * Once an event has been dequeued, however, it is guaranteed to be dispatched. This guarantee contributes to not being
 * able to exit exactly at time of expiry as (1) blocking on internal mutexes is necessary for dispatching the dequeued
 * event, and (2) during  dispatch of the dequeued event there is no way to control the time occupied by handler code
 * execution. The guaranteed time of exit is therefore the allotted time + amount of time required to dispatch
 * the last dequeued event.
 *
 * In cases where waiting on the queue times out, ESP_OK is returned and not ESP_ERR_TIMEOUT, since it is
 * normal behavior.
 *
 * @param[in] event_loop event loop to dispatch posted events from, must not be NULL
 * @param[in] ticks_to_run number of ticks to run the loop
 *
 * @note encountering an unknown event that has been posted to the loop will only generate a warning, not an error.
 *
 * @return
 *  - ESP_OK: Success
 *  - Others: Fail
 */
//go:linkname EspEventLoopRun C.esp_event_loop_run
func EspEventLoopRun(event_loop EspEventLoopHandleT, ticks_to_run TickTypeT) EspErrT

/**
 * @brief Register an event handler to the system event loop (legacy).
 *
 * This function can be used to register a handler for either: (1) specific events,
 * (2) all events of a certain event base, or (3) all events known by the system event loop.
 *
 *  - specific events: specify exact event_base and event_id
 *  - all events of a certain base: specify exact event_base and use ESP_EVENT_ANY_ID as the event_id
 *  - all events known by the loop: use ESP_EVENT_ANY_BASE for event_base and ESP_EVENT_ANY_ID as the event_id
 *
 * Registering multiple handlers to events is possible. Registering a single handler to multiple events is
 * also possible. However, registering the same handler to the same event multiple times would cause the
 * previous registrations to be overwritten.
 *
 * @param[in] event_base the base ID of the event to register the handler for
 * @param[in] event_id the ID of the event to register the handler for
 * @param[in] event_handler the handler function which gets called when the event is dispatched
 * @param[in] event_handler_arg data, aside from event data, that is passed to the handler when it is called
 *
 * @note the event loop library does not maintain a copy of event_handler_arg, therefore the user should
 * ensure that event_handler_arg still points to a valid location by the time the handler gets called
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_ERR_NO_MEM: Cannot allocate memory for the handler
 *  - ESP_ERR_INVALID_ARG: Invalid combination of event base and event ID
 *  - Others: Fail
 */
//go:linkname EspEventHandlerRegister C.esp_event_handler_register
func EspEventHandlerRegister(event_base EspEventBaseT, event_id c.Int32T, event_handler EspEventHandlerT, event_handler_arg c.Pointer) EspErrT

/**
 * @brief Register an event handler to a specific loop (legacy).
 *
 * This function behaves in the same manner as esp_event_handler_register, except the additional
 * specification of the event loop to register the handler to.
 *
 * @param[in] event_loop the event loop to register this handler function to, must not be NULL
 * @param[in] event_base the base ID of the event to register the handler for
 * @param[in] event_id the ID of the event to register the handler for
 * @param[in] event_handler the handler function which gets called when the event is dispatched
 * @param[in] event_handler_arg data, aside from event data, that is passed to the handler when it is called
 *
 * @note the event loop library does not maintain a copy of event_handler_arg, therefore the user should
 * ensure that event_handler_arg still points to a valid location by the time the handler gets called
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_ERR_NO_MEM: Cannot allocate memory for the handler
 *  - ESP_ERR_INVALID_ARG: Invalid combination of event base and event ID
 *  - Others: Fail
 */
//go:linkname EspEventHandlerRegisterWith C.esp_event_handler_register_with
func EspEventHandlerRegisterWith(event_loop EspEventLoopHandleT, event_base EspEventBaseT, event_id c.Int32T, event_handler EspEventHandlerT, event_handler_arg c.Pointer) EspErrT

/**
 * @brief Register an instance of event handler to a specific loop.
 *
 * This function can be used to register a handler for either: (1) specific events,
 * (2) all events of a certain event base, or (3) all events known by the system event loop.
 *
 *  - specific events: specify exact event_base and event_id
 *  - all events of a certain base: specify exact event_base and use ESP_EVENT_ANY_ID as the event_id
 *  - all events known by the loop: use ESP_EVENT_ANY_BASE for event_base and ESP_EVENT_ANY_ID as the event_id
 *
 * Besides the error, the function returns an instance object as output parameter to identify each registration.
 * This is necessary to remove (unregister) the registration before the event loop is deleted.
 *
 * Registering multiple handlers to events, registering a single handler to multiple events as well as registering
 * the same handler to the same event multiple times is possible.
 * Each registration yields a distinct instance object which identifies it over the registration
 * lifetime.
 *
 * @param[in] event_loop the event loop to register this handler function to, must not be NULL
 * @param[in] event_base the base ID of the event to register the handler for
 * @param[in] event_id the ID of the event to register the handler for
 * @param[in] event_handler the handler function which gets called when the event is dispatched
 * @param[in] event_handler_arg data, aside from event data, that is passed to the handler when it is called
 * @param[out] instance An event handler instance object related to the registered event handler and data, can be NULL.
 *             This needs to be kept if the specific callback instance should be unregistered before deleting the whole
 *             event loop. Registering the same event handler multiple times is possible and yields distinct instance
 *             objects. The data can be the same for all registrations.
 *             If no unregistration is needed, but the handler should be deleted when the event loop is deleted,
 *             instance can be NULL.
 *
 * @note the event loop library does not maintain a copy of event_handler_arg, therefore the user should
 * ensure that event_handler_arg still points to a valid location by the time the handler gets called
 *
 * @note Calling this function with instance set to NULL is equivalent to calling esp_event_handler_register_with.
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_ERR_NO_MEM: Cannot allocate memory for the handler
 *  - ESP_ERR_INVALID_ARG: Invalid combination of event base and event ID or instance is NULL
 *  - Others: Fail
 */
//go:linkname EspEventHandlerInstanceRegisterWith C.esp_event_handler_instance_register_with
func EspEventHandlerInstanceRegisterWith(event_loop EspEventLoopHandleT, event_base EspEventBaseT, event_id c.Int32T, event_handler EspEventHandlerT, event_handler_arg c.Pointer, instance *EspEventHandlerInstanceT) EspErrT

/**
 * @brief Register an instance of event handler to the default loop.
 *
 * This function does the same as esp_event_handler_instance_register_with, except that it registers the
 * handler to the default event loop.
 *
 * @param[in] event_base the base ID of the event to register the handler for
 * @param[in] event_id the ID of the event to register the handler for
 * @param[in] event_handler the handler function which gets called when the event is dispatched
 * @param[in] event_handler_arg data, aside from event data, that is passed to the handler when it is called
 * @param[out] instance An event handler instance object related to the registered event handler and data, can be NULL.
 *             This needs to be kept if the specific callback instance should be unregistered before deleting the whole
 *             event loop. Registering the same event handler multiple times is possible and yields distinct instance
 *             objects. The data can be the same for all registrations.
 *             If no unregistration is needed, but the handler should be deleted when the event loop is deleted,
 *             instance can be NULL.
 *
 * @note the event loop library does not maintain a copy of event_handler_arg, therefore the user should
 * ensure that event_handler_arg still points to a valid location by the time the handler gets called
 *
 * @note Calling this function with instance set to NULL is equivalent to calling esp_event_handler_register.
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_ERR_NO_MEM: Cannot allocate memory for the handler
 *  - ESP_ERR_INVALID_ARG: Invalid combination of event base and event ID or instance is NULL
 *  - Others: Fail
 */
//go:linkname EspEventHandlerInstanceRegister C.esp_event_handler_instance_register
func EspEventHandlerInstanceRegister(event_base EspEventBaseT, event_id c.Int32T, event_handler EspEventHandlerT, event_handler_arg c.Pointer, instance *EspEventHandlerInstanceT) EspErrT

/**
 * @brief Unregister a handler with the system event loop (legacy).
 *
 * Unregisters a handler, so it will no longer be called during dispatch.
 * Handlers can be unregistered for any combination of event_base and event_id which were previously registered.
 * To unregister a handler, the event_base and event_id arguments must match exactly the arguments passed to
 * esp_event_handler_register() when that handler was registered. Passing ESP_EVENT_ANY_BASE and/or ESP_EVENT_ANY_ID
 * will only unregister handlers that were registered with the same wildcard arguments.
 *
 * @note When using ESP_EVENT_ANY_ID, handlers registered to specific event IDs using the same base will not be
 *       unregistered. When using ESP_EVENT_ANY_BASE, events registered to specific bases will also not be
 *       unregistered. This avoids accidental unregistration of handlers registered by other users or components.
 *
 * @param[in] event_base the base of the event with which to unregister the handler
 * @param[in] event_id the ID of the event with which to unregister the handler
 * @param[in] event_handler the handler to unregister
 *
 * @return ESP_OK success
 * @return ESP_ERR_INVALID_ARG invalid combination of event base and event ID
 * @return others fail
 */
//go:linkname EspEventHandlerUnregister C.esp_event_handler_unregister
func EspEventHandlerUnregister(event_base EspEventBaseT, event_id c.Int32T, event_handler EspEventHandlerT) EspErrT

/**
 * @brief Unregister a handler from a specific event loop (legacy).
 *
 * This function behaves in the same manner as esp_event_handler_unregister, except the additional specification of
 * the event loop to unregister the handler with.
 *
 * @param[in] event_loop the event loop with which to unregister this handler function, must not be NULL
 * @param[in] event_base the base of the event with which to unregister the handler
 * @param[in] event_id the ID of the event with which to unregister the handler
 * @param[in] event_handler the handler to unregister
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_ERR_INVALID_ARG: Invalid combination of event base and event ID
 *  - Others: Fail
 */
//go:linkname EspEventHandlerUnregisterWith C.esp_event_handler_unregister_with
func EspEventHandlerUnregisterWith(event_loop EspEventLoopHandleT, event_base EspEventBaseT, event_id c.Int32T, event_handler EspEventHandlerT) EspErrT

/**
 * @brief Unregister a handler instance from a specific event loop.
 *
 * Unregisters a handler instance, so it will no longer be called during dispatch.
 * Handler instances can be unregistered for any combination of event_base and event_id which were previously
 * registered. To unregister a handler instance, the event_base and event_id arguments must match exactly the
 * arguments passed to esp_event_handler_instance_register() when that handler instance was registered.
 * Passing ESP_EVENT_ANY_BASE and/or ESP_EVENT_ANY_ID will only unregister handler instances that were registered
 * with the same wildcard arguments.
 *
 * @note When using ESP_EVENT_ANY_ID, handlers registered to specific event IDs using the same base will not be
 *       unregistered. When using ESP_EVENT_ANY_BASE, events registered to specific bases will also not be
 *       unregistered. This avoids accidental unregistration of handlers registered by other users or components.
 *
 * @param[in] event_loop the event loop with which to unregister this handler function, must not be NULL
 * @param[in] event_base the base of the event with which to unregister the handler
 * @param[in] event_id the ID of the event with which to unregister the handler
 * @param[in] instance the instance object of the registration to be unregistered
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_ERR_INVALID_ARG: Invalid combination of event base and event ID
 *  - Others: Fail
 */
//go:linkname EspEventHandlerInstanceUnregisterWith C.esp_event_handler_instance_unregister_with
func EspEventHandlerInstanceUnregisterWith(event_loop EspEventLoopHandleT, event_base EspEventBaseT, event_id c.Int32T, instance EspEventHandlerInstanceT) EspErrT

/**
 * @brief Unregister a handler from the system event loop.
 *
 * This function does the same as esp_event_handler_instance_unregister_with, except that it unregisters the
 * handler instance from the default event loop.
 *
 * @param[in] event_base the base of the event with which to unregister the handler
 * @param[in] event_id the ID of the event with which to unregister the handler
 * @param[in] instance the instance object of the registration to be unregistered
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_ERR_INVALID_ARG: Invalid combination of event base and event ID
 *  - Others: Fail
 */
//go:linkname EspEventHandlerInstanceUnregister C.esp_event_handler_instance_unregister
func EspEventHandlerInstanceUnregister(event_base EspEventBaseT, event_id c.Int32T, instance EspEventHandlerInstanceT) EspErrT

/**
 * @brief Posts an event to the system default event loop. The event loop library keeps a copy of event_data and manages
 * the copy's lifetime automatically (allocation + deletion); this ensures that the data the
 * handler receives is always valid.
 *
 * @param[in] event_base the event base that identifies the event
 * @param[in] event_id the event ID that identifies the event
 * @param[in] event_data the data, specific to the event occurrence, that gets passed to the handler
 * @param[in] event_data_size the size of the event data
 * @param[in] ticks_to_wait number of ticks to block on a full event queue
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_ERR_TIMEOUT: Time to wait for event queue to unblock expired,
 *                      queue full when posting from ISR
 *  - ESP_ERR_INVALID_ARG: Invalid combination of event base and event ID
 *  - Others: Fail
 */
//go:linkname EspEventPost C.esp_event_post
func EspEventPost(event_base EspEventBaseT, event_id c.Int32T, event_data c.Pointer, event_data_size c.SizeT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Posts an event to the specified event loop. The event loop library keeps a copy of event_data and manages
 * the copy's lifetime automatically (allocation + deletion); this ensures that the data the
 * handler receives is always valid.
 *
 * This function behaves in the same manner as esp_event_post, except the additional specification of the event loop
 * to post the event to.
 *
 * @param[in] event_loop the event loop to post to, must not be NULL
 * @param[in] event_base the event base that identifies the event
 * @param[in] event_id the event ID that identifies the event
 * @param[in] event_data the data, specific to the event occurrence, that gets passed to the handler
 * @param[in] event_data_size the size of the event data
 * @param[in] ticks_to_wait number of ticks to block on a full event queue
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_ERR_TIMEOUT: Time to wait for event queue to unblock expired,
 *                      queue full when posting from ISR
 *  - ESP_ERR_INVALID_ARG: Invalid combination of event base and event ID
 *  - Others: Fail
 */
//go:linkname EspEventPostTo C.esp_event_post_to
func EspEventPostTo(event_loop EspEventLoopHandleT, event_base EspEventBaseT, event_id c.Int32T, event_data c.Pointer, event_data_size c.SizeT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Special variant of esp_event_post for posting events from interrupt handlers.
 *
 * @param[in] event_base the event base that identifies the event
 * @param[in] event_id the event ID that identifies the event
 * @param[in] event_data the data, specific to the event occurrence, that gets passed to the handler
 * @param[in] event_data_size the size of the event data; max is 4 bytes
 * @param[out] task_unblocked an optional parameter (can be NULL) which indicates that an event task with
 *                            higher priority than currently running task has been unblocked by the posted event;
 *                            a context switch should be requested before the interrupt is existed.
 *
 * @note this function is only available when CONFIG_ESP_EVENT_POST_FROM_ISR is enabled
 * @note when this function is called from an interrupt handler placed in IRAM, this function should
 *       be placed in IRAM as well by enabling CONFIG_ESP_EVENT_POST_FROM_IRAM_ISR
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_FAIL: Event queue for the default event loop full
 *  - ESP_ERR_INVALID_ARG: Invalid combination of event base and event ID,
 *                          data size of more than 4 bytes
 *  - Others: Fail
 */
//go:linkname EspEventIsrPost C.esp_event_isr_post
func EspEventIsrPost(event_base EspEventBaseT, event_id c.Int32T, event_data c.Pointer, event_data_size c.SizeT, task_unblocked *BaseTypeT) EspErrT

/**
 * @brief Special variant of esp_event_post_to for posting events from interrupt handlers
 *
 * @param[in] event_loop the event loop to post to, must not be NULL
 * @param[in] event_base the event base that identifies the event
 * @param[in] event_id the event ID that identifies the event
 * @param[in] event_data the data, specific to the event occurrence, that gets passed to the handler
 * @param[in] event_data_size the size of the event data
 * @param[out] task_unblocked an optional parameter (can be NULL) which indicates that an event task with
 *                            higher priority than currently running task has been unblocked by the posted event;
 *                            a context switch should be requested before the interrupt is existed.
 *
 * @note this function is only available when CONFIG_ESP_EVENT_POST_FROM_ISR is enabled
 * @note when this function is called from an interrupt handler placed in IRAM, this function should
 *       be placed in IRAM as well by enabling CONFIG_ESP_EVENT_POST_FROM_IRAM_ISR
 *
 * @return
 *  - ESP_OK: Success
 *  - ESP_FAIL: Event queue for the loop full
 *  - ESP_ERR_INVALID_ARG: Invalid combination of event base and event ID,
 *                          data size of more than 4 bytes
 *  - Others: Fail
 */
//go:linkname EspEventIsrPostTo C.esp_event_isr_post_to
func EspEventIsrPostTo(event_loop EspEventLoopHandleT, event_base EspEventBaseT, event_id c.Int32T, event_data c.Pointer, event_data_size c.SizeT, task_unblocked *BaseTypeT) EspErrT

/**
* @brief Dumps statistics of all event loops.
*
* Dumps event loop info in the format:
*
@verbatim
      event loop
          handler
          handler
          ...
      event loop
          handler
          handler
          ...

 where:

  event loop
      format: address,name rx:total_received dr:total_dropped
      where:
          address - memory address of the event loop
          name - name of the event loop, 'none' if no dedicated task
          total_received - number of successfully posted events
          total_dropped - number of events unsuccessfully posted due to queue being full

  handler
      format: address ev:base,id inv:total_invoked run:total_runtime
      where:
          address - address of the handler function
          base,id - the event specified by event base and ID this handler executes
          total_invoked - number of times this handler has been invoked
          total_runtime - total amount of time used for invoking this handler

@endverbatim
*
* @param[in] file the file stream to output to
*
* @note this function is a noop when CONFIG_ESP_EVENT_LOOP_PROFILING is disabled
*
* @return
*  - ESP_OK: Success
*  - ESP_ERR_NO_MEM: Cannot allocate memory for event loops list
*  - Others: Fail
*/
//go:linkname EspEventDump C.esp_event_dump
func EspEventDump(file *c.FILE) EspErrT

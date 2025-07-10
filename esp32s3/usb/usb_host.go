package usb

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const USB_HOST_LIB_EVENT_FLAGS_NO_CLIENTS = 0x01
const USB_HOST_LIB_EVENT_FLAGS_ALL_FREE = 0x02

type UsbHostClientHandleS struct {
	Unused [8]uint8
}
type UsbHostClientHandleT *UsbHostClientHandleS
type UsbHostClientEventT c.Int

const (
	USB_HOST_CLIENT_EVENT_NEW_DEV  UsbHostClientEventT = 0
	USB_HOST_CLIENT_EVENT_DEV_GONE UsbHostClientEventT = 1
)

/**
 * @brief Client event message
 *
 * Client event messages are sent to each client of the USB Host Library in order to notify them of various
 * USB Host Library events such as:
 * - Addition of new devices
 * - Removal of existing devices
 *
 * @note The event message structure has a union with members corresponding to each particular event. Based on the event
 *       type, only the relevant member field should be accessed.
 */

type UsbHostClientEventMsgT struct {
	Event UsbHostClientEventT
}

/**
 * @brief Current information about the USB Host Library obtained via usb_host_lib_info()
 */

type UsbHostLibInfoT struct {
	NumDevices c.Int
	NumClients c.Int
}

// llgo:type C
type UsbHostClientEventCbT func(*UsbHostClientEventMsgT, c.Pointer)

/**
 * @brief USB Host Library configuration
 *
 * Configuration structure of the USB Host Library. Provided in the usb_host_install() function
 */

type UsbHostConfigT struct {
	SkipPhySetup      bool
	RootPortUnpowered bool
	IntrFlags         c.Int
	EnumFilterCb      UsbHostEnumFilterCbT
}

/**
 * @brief USB Host Library Client configuration
 *
 * Configuration structure for a USB Host Library client. Provided in usb_host_client_register()
 */

type UsbHostClientConfigT struct {
	IsSynchronous  bool
	MaxNumEventMsg c.Int
}

/**
 * @brief Install the USB Host Library
 *
 * - This function should be called only once to install the USB Host Library
 * - This function should be called before any other USB Host Library functions are called
 *
 * @note If skip_phy_setup is set in the install configuration, the user is responsible for ensuring that the underlying
 *       Host Controller is enabled and the USB PHY (internal or external) is already setup before this function is
 *       called.
 * @param[in] config USB Host Library configuration
 *
 * @return
 *    - ESP_OK: USB Host installed successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 *    - ESP_ERR_INVALID_STATE: USB Host Library is not in correct state to be installed
 *      (eg. the library itself of one of it's drivers is already installed)
 *    - ESP_ERR_NO_MEM: Insufficient memory
 */
// llgo:link (*UsbHostConfigT).UsbHostInstall C.usb_host_install
func (recv_ *UsbHostConfigT) UsbHostInstall() EspErrT {
	return 0
}

/**
 * @brief Uninstall the USB Host Library
 *
 * - This function should be called to uninstall the USB Host Library, thereby freeing its resources
 * - All clients must have been deregistered before calling this function
 * - All devices must have been freed by calling usb_host_device_free_all() and receiving the
 *   USB_HOST_LIB_EVENT_FLAGS_ALL_FREE event flag
 *
 * @note If skip_phy_setup was set when the Host Library was installed, the user is responsible for disabling the
 *       underlying Host Controller and USB PHY (internal or external).
 *
 * @return
 *    - ESP_OK: USB Host uninstalled successfully
 *    - ESP_ERR_INVALID_STATE: USB Host Library is not installed, or has unfinished actions
 */
//go:linkname UsbHostUninstall C.usb_host_uninstall
func UsbHostUninstall() EspErrT

/**
 * @brief Handle USB Host Library events
 *
 * - This function handles all of the USB Host Library's processing and should be called repeatedly in a loop
 * - Check event_flags_ret to see if an flags are set indicating particular USB Host Library events
 * - This function should never be called by multiple threads simultaneously
 *
 * @note This function can block
 * @param[in] timeout_ticks Timeout in ticks to wait for an event to occur
 * @param[out] event_flags_ret Event flags that indicate what USB Host Library event occurred.
 *
 * @return
 *    - ESP_OK: No events to be handled
 *    - ESP_ERR_INVALID_STATE: USB Host Library is not installed
 *    - ESP_ERR_TIMEOUT: Semaphore waiting for events has timed out
 */
//go:linkname UsbHostLibHandleEvents C.usb_host_lib_handle_events
func UsbHostLibHandleEvents(timeout_ticks TickTypeT, event_flags_ret *c.Uint32T) EspErrT

/**
 * @brief Unblock the USB Host Library handler
 *
 * - This function simply unblocks the USB Host Library event handling function (usb_host_lib_handle_events())
 *
 * @return
 *    - ESP_OK: USB Host library unblocked successfully
 *    - ESP_ERR_INVALID_STATE: USB Host Library is not installed
 */
//go:linkname UsbHostLibUnblock C.usb_host_lib_unblock
func UsbHostLibUnblock() EspErrT

/**
 * @brief Get current information about the USB Host Library
 *
 * @param[out] info_ret USB Host Library Information
 *
 * @return
 *    - ESP_OK: USB Host Library info obtained successfully
 *    - ESP_ERR_INVALID_STATE: USB Host Library is not installed
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 */
// llgo:link (*UsbHostLibInfoT).UsbHostLibInfo C.usb_host_lib_info
func (recv_ *UsbHostLibInfoT) UsbHostLibInfo() EspErrT {
	return 0
}

/**
 * @brief Power the root port ON or OFF
 *
 * - Powering ON the root port will allow device connections to occur
 * - Powering OFF the root port will disconnect all devices downstream off the root port and prevent
 *   any further device connections.
 *
 * @note If 'usb_host_config_t.root_port_unpowered' was set on USB Host Library installation, users must call this
 *       function to power ON the root port before any device connections can occur.
 *
 * @param[in] enable True to power the root port ON, false to power OFF
 * @return
 *    - ESP_OK: Root port power enabled/disabled
 *    - ESP_ERR_INVALID_STATE: Root port already powered or HUB driver not installed
 */
//go:linkname UsbHostLibSetRootPortPower C.usb_host_lib_set_root_port_power
func UsbHostLibSetRootPortPower(enable bool) EspErrT

/**
 * @brief Register a client of the USB Host Library
 *
 * - This function registers a client of the USB Host Library
 * - Once a client is registered, its processing function usb_host_client_handle_events() should be called repeatedly
 *
 * @param[in] client_config Client configuration
 * @param[out] client_hdl_ret Client handle
 *
 * @return
 *    - ESP_OK: USB Host Library client registered successfully
 *    - ESP_ERR_INVALID_STATE: USB Host Library is not installed
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 *    - ESP_ERR_NO_MEM: Insufficient memory
 */
// llgo:link (*UsbHostClientConfigT).UsbHostClientRegister C.usb_host_client_register
func (recv_ *UsbHostClientConfigT) UsbHostClientRegister(client_hdl_ret *UsbHostClientHandleT) EspErrT {
	return 0
}

/**
 * @brief Deregister a USB Host Library client
 *
 * - This function deregisters a client of the USB Host Library
 * - The client must have closed all previously opened devices before attempting to deregister
 *
 * @param[in] client_hdl Client handle
 *
 * @return
 *    - ESP_OK: USB Host Library client deregistered successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 */
//go:linkname UsbHostClientDeregister C.usb_host_client_deregister
func UsbHostClientDeregister(client_hdl UsbHostClientHandleT) EspErrT

/**
 * @brief USB Host Library client processing function
 *
 * - This function handles all of a client's processing and should be called repeatedly in a loop
 * - For a particular client, this function should never be called by multiple threads simultaneously
 *
 * @note This function can block
 * @param[in] client_hdl Client handle
 * @param[in] timeout_ticks Timeout in ticks to wait for an event to occur
 *
 * @return
 *    - ESP_OK: No event to be handled
 *    - ESP_ERR_INVALID_STATE: USB Host Library is not installed
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 *    - ESP_ERR_TIMEOUT: Semaphore waiting for events has timed out
 */
//go:linkname UsbHostClientHandleEvents C.usb_host_client_handle_events
func UsbHostClientHandleEvents(client_hdl UsbHostClientHandleT, timeout_ticks TickTypeT) EspErrT

/**
 * @brief Unblock a client
 *
 * - This function simply unblocks a client if it is blocked on the usb_host_client_handle_events() function.
 * - This function is useful when need to unblock a client in order to deregister it.
 *
 * @param[in] client_hdl Client handle
 *
 * @return
 *    - ESP_OK: Client unblocked successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 */
//go:linkname UsbHostClientUnblock C.usb_host_client_unblock
func UsbHostClientUnblock(client_hdl UsbHostClientHandleT) EspErrT

/**
 * @brief Open a device
 *
 * - This function allows a client to open a device
 * - A client must open a device first before attempting to use it (e.g., sending transfers, device requests etc.)
 *
 * @param[in] client_hdl Client handle
 * @param[in] dev_addr Device's address
 * @param[out] dev_hdl_ret Device's handle
 *
 * @return
 *    - ESP_OK: Device opened successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 */
//go:linkname UsbHostDeviceOpen C.usb_host_device_open
func UsbHostDeviceOpen(client_hdl UsbHostClientHandleT, dev_addr c.Uint8T, dev_hdl_ret *UsbDeviceHandleT) EspErrT

/**
 * @brief Close a device
 *
 * - This function allows a client to close a device
 * - A client must close a device after it has finished using the device (claimed interfaces must also be released)
 * - A client must close all devices it has opened before de-registering
 *
 * @note This function can block
 * @param[in] client_hdl Client handle
 * @param[in] dev_hdl Device handle
 *
 * @return
 *    - ESP_OK: Device closed successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 *    - ESP_ERR_NOT_FOUND: The client never opened the device (the device address not found among opened devices)
 *    - ESP_ERR_INVALID_STATE: The client has not released all interfaces from the device
 */
//go:linkname UsbHostDeviceClose C.usb_host_device_close
func UsbHostDeviceClose(client_hdl UsbHostClientHandleT, dev_hdl UsbDeviceHandleT) EspErrT

/**
 * @brief Indicate that all devices can be freed when possible
 *
 * - This function marks all devices as waiting to be freed
 * - If a device is not opened by any clients, it will be freed immediately
 * - If a device is opened by at least one client, the device will be free when the last client closes that device.
 * - Wait for the USB_HOST_LIB_EVENT_FLAGS_ALL_FREE flag to be set by usb_host_lib_handle_events() in order to know
 *   when all devices have been freed
 * - This function is useful when cleaning up devices before uninstalling the USB Host Library
 *
 * @return
 *    - ESP_OK: All devices already freed (i.e., there were no devices)
 *    - ESP_ERR_INVALID_STATE: Client must be deregistered
 *    - ESP_ERR_NOT_FINISHED: There are one or more devices that still need to be freed,
 *      wait for USB_HOST_LIB_EVENT_FLAGS_ALL_FREE event
 */
//go:linkname UsbHostDeviceFreeAll C.usb_host_device_free_all
func UsbHostDeviceFreeAll() EspErrT

/**
 * @brief Fill a list of device address
 *
 * - This function fills an empty list with the address of connected devices
 * - The Device addresses can then used in usb_host_device_open()
 * - If there are more devices than the list_len, this function will only fill up to list_len number of devices.
 *
 * @param[in] list_len Length of the empty list
 * @param[inout] dev_addr_list Empty list to be filled
 * @param[out] num_dev_ret Number of devices
 *
 * @return
 *    - ESP_OK: Device list filled successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 */
//go:linkname UsbHostDeviceAddrListFill C.usb_host_device_addr_list_fill
func UsbHostDeviceAddrListFill(list_len c.Int, dev_addr_list *c.Uint8T, num_dev_ret *c.Int) EspErrT

/**
 * @brief Get device's information
 *
 * - This function gets some basic information of a device
 * - The device must be opened first before attempting to get its information
 *
 * @note This function can block
 * @param[in] dev_hdl Device handle
 * @param[out] dev_info Device information
 *
 * @return
 *    - ESP_OK: Device information obtained successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 */
//go:linkname UsbHostDeviceInfo C.usb_host_device_info
func UsbHostDeviceInfo(dev_hdl UsbDeviceHandleT, dev_info *UsbDeviceInfoT) EspErrT

/**
 * @brief Get device's device descriptor
 *
 * - A client must call usb_host_device_open() first
 * - No control transfer is sent. The device's descriptor is cached on enumeration
 * - This function simple returns a pointer to the cached descriptor
 *
 * @note No control transfer is sent. The device's descriptor is cached on enumeration
 * @param[in] dev_hdl Device handle
 * @param[out] device_desc Device descriptor
 *
 * @return
 *    - ESP_OK: Device descriptor obtained successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 */
//go:linkname UsbHostGetDeviceDescriptor C.usb_host_get_device_descriptor
func UsbHostGetDeviceDescriptor(dev_hdl UsbDeviceHandleT, device_desc **UsbDeviceDescT) EspErrT

/**
 * @brief Get device's active configuration descriptor
 *
 * - A client must call usb_host_device_open() first
 * - No control transfer is sent. The device's active configuration descriptor is cached on enumeration
 * - This function simple returns a pointer to the cached descriptor
 *
 * @note This function can block
 * @note No control transfer is sent. A device's active configuration descriptor is cached on enumeration
 * @param[in] dev_hdl Device handle
 * @param[out] config_desc Configuration descriptor
 *
 * @return
 *    - ESP_OK: Active configuration descriptor obtained successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 */
//go:linkname UsbHostGetActiveConfigDescriptor C.usb_host_get_active_config_descriptor
func UsbHostGetActiveConfigDescriptor(dev_hdl UsbDeviceHandleT, config_desc **UsbConfigDescT) EspErrT

/**
 * @brief Get get device's configuration descriptor
 *
 * - The USB Host library only caches a device's active configuration descriptor.
 * - This function reads any configuration descriptor of a particular device (specified by bConfigurationValue).
 * - This function will read the specified configuration descriptor via control transfers, and allocate memory
 *   to store that descriptor.
 * - Users can call usb_host_free_config_desc() to free the descriptor's memory afterwards.
 *
 * @note This function can block
 * @note A client must call usb_host_device_open() on the device first
 * @param[in] client_hdl Client handle - usb_host_client_handle_events() should be called repeatedly in a separate task
 *            to handle client events
 * @param[in] dev_hdl Device handle
 * @param[in] bConfigurationValue Index of device's configuration descriptor to be read
 * @param[out] config_desc_ret Returned configuration descriptor
 * @note bConfigurationValue starts from index 1
 *
 * @return
 *    - ESP_OK: Configuration descriptor obtained successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 *    - ESP_ERR_NOT_SUPPORTED: Invalid bConfigurationValue value (the device does not have this configuration value)
 *    - ESP_ERR_NO_MEM: Insufficient memory
 */
//go:linkname UsbHostGetConfigDesc C.usb_host_get_config_desc
func UsbHostGetConfigDesc(client_hdl UsbHostClientHandleT, dev_hdl UsbDeviceHandleT, bConfigurationValue c.Uint8T, config_desc_ret **UsbConfigDescT) EspErrT

/**
 * @brief Free a configuration descriptor
 *
 * This function frees a configuration descriptor that was returned by usb_host_get_config_desc()
 *
 * @param[out] config_desc Configuration descriptor
 *
 * @return
 *    - ESP_OK: Configuration descriptor freed successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 */
// llgo:link (*UsbConfigDescT).UsbHostFreeConfigDesc C.usb_host_free_config_desc
func (recv_ *UsbConfigDescT) UsbHostFreeConfigDesc() EspErrT {
	return 0
}

/**
 * @brief Function for a client to claim a device's interface
 *
 * - A client must claim a device's interface before attempting to communicate with any of its endpoints
 * - Once an interface is claimed by a client, it cannot be claimed by any other client.
 *
 * @note This function can block
 * @param[in] client_hdl Client handle
 * @param[in] dev_hdl Device handle
 * @param[in] bInterfaceNumber Interface number
 * @param[in] bAlternateSetting Interface alternate setting number
 *
 * @return
 *    - ESP_OK: Interface claimed successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 *    - ESP_ERR_INVALID_STATE: USB Host is not a correct state to claim an interface
 *    - ESP_ERR_NOT_FOUND: Interface or EP not found
 *    - ESP_ERR_NO_MEM: Insufficient memory
 */
//go:linkname UsbHostInterfaceClaim C.usb_host_interface_claim
func UsbHostInterfaceClaim(client_hdl UsbHostClientHandleT, dev_hdl UsbDeviceHandleT, bInterfaceNumber c.Uint8T, bAlternateSetting c.Uint8T) EspErrT

/**
 * @brief Function for a client to release a previously claimed interface
 *
 * - A client should release a device's interface after it no longer needs to communicate with the interface
 * - A client must release all of its interfaces of a device it has claimed before being able to close the device
 *
 * @note This function can block
 * @param[in] client_hdl Client handle
 * @param[in] dev_hdl Device handle
 * @param[in] bInterfaceNumber Interface number
 *
 * @return
 *    - ESP_OK: Interface released successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 *    - ESP_ERR_INVALID_STATE: Client never opened the USB Device, or interface currently can not be freed
 *    - ESP_ERR_NOT_FOUND: Interface number not found in the list of interfaces
 */
//go:linkname UsbHostInterfaceRelease C.usb_host_interface_release
func UsbHostInterfaceRelease(client_hdl UsbHostClientHandleT, dev_hdl UsbDeviceHandleT, bInterfaceNumber c.Uint8T) EspErrT

/**
 * @brief Halt a particular endpoint
 *
 * - The device must have been opened by a client
 * - The endpoint must be part of an interface claimed by a client
 * - Once halted, the endpoint must be cleared using usb_host_endpoint_clear() before it can communicate again
 *
 * @note This function can block
 * @param[in] dev_hdl Device handle
 * @param[in] bEndpointAddress Endpoint address
 *
 * @return
 *    - ESP_OK: Endpoint halted successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 *    - ESP_ERR_NOT_FOUND: Endpoint address not found
 *    - ESP_ERR_INVALID_STATE: Endpoint pipe is not in the correct state/condition to execute a command
 */
//go:linkname UsbHostEndpointHalt C.usb_host_endpoint_halt
func UsbHostEndpointHalt(dev_hdl UsbDeviceHandleT, bEndpointAddress c.Uint8T) EspErrT

/**
 * @brief Flush a particular endpoint
 *
 * - The device must have been opened by a client
 * - The endpoint must be part of an interface claimed by a client
 * - The endpoint must have been halted (either through a transfer error, or usb_host_endpoint_halt())
 * - Flushing an endpoint will caused an queued up transfers to be canceled
 *
 * @note This function can block
 * @param[in] dev_hdl Device handle
 * @param[in] bEndpointAddress Endpoint address
 *
 * @return
 *    - ESP_OK: Endpoint flushed successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 *    - ESP_ERR_NOT_FOUND: Endpoint address not found
 *    - ESP_ERR_INVALID_STATE: Endpoint pipe is not in the correct state/condition to execute a command
 */
//go:linkname UsbHostEndpointFlush C.usb_host_endpoint_flush
func UsbHostEndpointFlush(dev_hdl UsbDeviceHandleT, bEndpointAddress c.Uint8T) EspErrT

/**
 * @brief Clear a halt on a particular endpoint
 *
 * - The device must have been opened by a client
 * - The endpoint must be part of an interface claimed by a client
 * - The endpoint must have been halted (either through a transfer error, or usb_host_endpoint_halt())
 * - If the endpoint has any queued up transfers, clearing a halt will resume their execution
 *
 * @note This function can block
 * @param[in] dev_hdl Device handle
 * @param[in] bEndpointAddress Endpoint address
 *
 * @return
 *    - ESP_OK: Endpoint cleared successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 *    - ESP_ERR_NOT_FOUND: Endpoint address not found
 *    - ESP_ERR_INVALID_STATE: Endpoint pipe is not in the correct state/condition to execute a command
 */
//go:linkname UsbHostEndpointClear C.usb_host_endpoint_clear
func UsbHostEndpointClear(dev_hdl UsbDeviceHandleT, bEndpointAddress c.Uint8T) EspErrT

/**
 * @brief Allocate a transfer object
 *
 * - This function allocates a transfer object
 * - Each transfer object has a fixed sized buffer specified on allocation
 * - The resulting data_buffer_size can be bigger that the requested size. This is to ensure that the data buffer is cache aligned
 * - A transfer object can be re-used indefinitely
 * - A transfer can be submitted using usb_host_transfer_submit() or usb_host_transfer_submit_control()
 *
 * @param[in] data_buffer_size Size of the transfer's data buffer
 * @param[in] num_isoc_packets Number of isochronous packets in transfer (set to 0 for non-isochronous transfers)
 * @param[out] transfer Transfer object
 *
 * @return
 *    - ESP_OK: Transfer object allocated successfully
 *    - ESP_ERR_NO_MEM: Insufficient memory
 */
//go:linkname UsbHostTransferAlloc C.usb_host_transfer_alloc
func UsbHostTransferAlloc(data_buffer_size c.SizeT, num_isoc_packets c.Int, transfer **UsbTransferT) EspErrT

/**
 * @brief Free a transfer object
 *
 * - Free a transfer object previously allocated using usb_host_transfer_alloc()
 * - The transfer must not be in-flight when attempting to free it
 * - If a NULL pointer is passed, this function will simply return ESP_OK
 *
 * @param[in] transfer Transfer object
 *
 * @return
 *    - ESP_OK: Transfer object freed successfully
 */
// llgo:link (*UsbTransferT).UsbHostTransferFree C.usb_host_transfer_free
func (recv_ *UsbTransferT) UsbHostTransferFree() EspErrT {
	return 0
}

/**
 * @brief Submit a non-control transfer
 *
 * - Submit a transfer to a particular endpoint. The device and endpoint number is specified inside the transfer
 * - The transfer must be properly initialized before submitting
 * - On completion, the transfer's callback will be called from the client's usb_host_client_handle_events() function.
 *
 * @param[in] transfer Initialized transfer object
 *
 * @return
 *    - ESP_OK: Transfer submitted successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 *    - ESP_ERR_NOT_FINISHED: Transfer already in-flight
 *    - ESP_ERR_NOT_FOUND: Endpoint address not found
 *    - ESP_ERR_INVALID_STATE: Endpoint pipe is not in a correct state to submit transfer
 */
// llgo:link (*UsbTransferT).UsbHostTransferSubmit C.usb_host_transfer_submit
func (recv_ *UsbTransferT) UsbHostTransferSubmit() EspErrT {
	return 0
}

/**
 * @brief Submit a control transfer
 *
 * - Submit a control transfer to a particular device. The client must have opened the device first
 * - The transfer must be properly initialized before submitting. The first 8 bytes of the transfer's data buffer should
 *   contain the control transfer setup packet
 * - On completion, the transfer's callback will be called from the client's usb_host_client_handle_events() function.
 *
 * @param[in] client_hdl Client handle
 * @param[in] transfer Initialized transfer object
 *
 * @return
 *    - ESP_OK: Control transfer submitted successfully
 *    - ESP_ERR_INVALID_ARG: Invalid argument
 *    - ESP_ERR_NOT_FINISHED: Transfer already in-flight
 *    - ESP_ERR_NOT_FOUND: Endpoint address not found
 *    - ESP_ERR_INVALID_STATE: Endpoint pipe is not in a correct state to submit transfer
 */
//go:linkname UsbHostTransferSubmitControl C.usb_host_transfer_submit_control
func UsbHostTransferSubmitControl(client_hdl UsbHostClientHandleT, transfer *UsbTransferT) EspErrT

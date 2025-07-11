package usb

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Get the next descriptor
 *
 * Given a particular descriptor within a full configuration descriptor, get the next descriptor within the
 * configuration descriptor. This is a convenience function that can be used to walk each individual descriptor within
 * a full configuration descriptor.
 *
 * @param[in] cur_desc Current descriptor
 * @param[in] wTotalLength Total length of the configuration descriptor
 * @param[inout] offset Byte offset relative to the start of the configuration descriptor. On input, it is the offset of
 *                      the current descriptor. On output, it is the offset of the returned descriptor.
 * @return usb_standard_desc_t* Next descriptor, NULL if end of configuration descriptor reached
 */
// llgo:link (*UsbStandardDescT).UsbParseNextDescriptor C.usb_parse_next_descriptor
func (recv_ *UsbStandardDescT) UsbParseNextDescriptor(wTotalLength c.Uint16T, offset *c.Int) *UsbStandardDescT {
	return nil
}

/**
 * @brief Get the next descriptor of a particular type
 *
 * Given a particular descriptor within a full configuration descriptor, get the next descriptor of a particular type
 * (i.e., using the bDescriptorType value) within the configuration descriptor.
 *
 * @param[in] cur_desc Current descriptor
 * @param[in] wTotalLength Total length of the configuration descriptor
 * @param[in] bDescriptorType Type of the next descriptor to get
 * @param[inout] offset Byte offset relative to the start of the configuration descriptor. On input, it is the offset of
 *                      the current descriptor. On output, it is the offset of the returned descriptor.
 * @return usb_standard_desc_t* Next descriptor, NULL if end descriptor is not found or configuration descriptor reached
 */
// llgo:link (*UsbStandardDescT).UsbParseNextDescriptorOfType C.usb_parse_next_descriptor_of_type
func (recv_ *UsbStandardDescT) UsbParseNextDescriptorOfType(wTotalLength c.Uint16T, bDescriptorType c.Uint8T, offset *c.Int) *UsbStandardDescT {
	return nil
}

/**
 * @brief Get the number of alternate settings for a bInterfaceNumber
 *
 * Given a particular configuration descriptor, for a particular bInterfaceNumber, get the number of alternate settings
 * available for that interface (i.e., the max possible value of bAlternateSetting for that bInterfaceNumber).
 *
 * @param[in] config_desc Pointer to the start of a full configuration descriptor
 * @param[in] bInterfaceNumber Interface number
 * @return int The number of alternate settings that the interface has, -1 if bInterfaceNumber not found
 */
// llgo:link (*UsbConfigDescT).UsbParseInterfaceNumberOfAlternate C.usb_parse_interface_number_of_alternate
func (recv_ *UsbConfigDescT) UsbParseInterfaceNumberOfAlternate(bInterfaceNumber c.Uint8T) c.Int {
	return 0
}

/**
 * @brief Get a particular interface descriptor (using bInterfaceNumber and bAlternateSetting)
 *
 * Given a full configuration descriptor, get a particular interface descriptor.
 *
 * @note To get the number of alternate settings for a particular bInterfaceNumber, call
 *       usb_parse_interface_number_of_alternate()
 *
 * @param[in] config_desc Pointer to the start of a full configuration descriptor
 * @param[in] bInterfaceNumber Interface number
 * @param[in] bAlternateSetting Alternate setting number
 * @param[out] offset Byte offset of the interface descriptor relative to the start of the configuration descriptor. Can be NULL.
 * @return const usb_intf_desc_t* Pointer to interface descriptor, NULL if not found.
 */
// llgo:link (*UsbConfigDescT).UsbParseInterfaceDescriptor C.usb_parse_interface_descriptor
func (recv_ *UsbConfigDescT) UsbParseInterfaceDescriptor(bInterfaceNumber c.Uint8T, bAlternateSetting c.Uint8T, offset *c.Int) *UsbIntfDescT {
	return nil
}

/**
 * @brief Get an endpoint descriptor within an interface descriptor
 *
 * Given an interface descriptor, get the Nth endpoint descriptor of the interface. The number of endpoints in an
 * interface is indicated by the bNumEndpoints field of the interface descriptor.
 *
 * @note If bNumEndpoints is 0, it means the interface uses the default endpoint only
 *
 * @param[in] intf_desc Pointer to thee start of an interface descriptor
 * @param[in] index Endpoint index
 * @param[in] wTotalLength Total length of the containing configuration descriptor
 * @param[inout] offset Byte offset relative to the start of the configuration descriptor. On input, it is the offset
 *                      of the interface descriptor. On output, it is the offset of the endpoint descriptor.
 * @return const usb_ep_desc_t* Pointer to endpoint descriptor, NULL if not found.
 */
// llgo:link (*UsbIntfDescT).UsbParseEndpointDescriptorByIndex C.usb_parse_endpoint_descriptor_by_index
func (recv_ *UsbIntfDescT) UsbParseEndpointDescriptorByIndex(index c.Int, wTotalLength c.Uint16T, offset *c.Int) *UsbEpDescT {
	return nil
}

/**
 * @brief Get an endpoint descriptor based on an endpoint's address
 *
 * Given a configuration descriptor, get an endpoint descriptor based on it's bEndpointAddress, bAlternateSetting, and
 * bInterfaceNumber.
 *
 * @param[in] config_desc Pointer to the start of a full configuration descriptor
 * @param[in] bInterfaceNumber Interface number
 * @param[in] bAlternateSetting Alternate setting number
 * @param[in] bEndpointAddress Endpoint address
 * @param[out] offset Byte offset of the endpoint descriptor relative to the start of the configuration descriptor. Can be NULL
 * @return const usb_ep_desc_t* Pointer to endpoint descriptor, NULL if not found.
 */
// llgo:link (*UsbConfigDescT).UsbParseEndpointDescriptorByAddress C.usb_parse_endpoint_descriptor_by_address
func (recv_ *UsbConfigDescT) UsbParseEndpointDescriptorByAddress(bInterfaceNumber c.Uint8T, bAlternateSetting c.Uint8T, bEndpointAddress c.Uint8T, offset *c.Int) *UsbEpDescT {
	return nil
}

// llgo:type C
type PrintClassDescriptorCb func(*UsbStandardDescT)

/**
 * @brief Print device descriptor
 *
 * @param[in] devc_desc Device descriptor
 */
// llgo:link (*UsbDeviceDescT).UsbPrintDeviceDescriptor C.usb_print_device_descriptor
func (recv_ *UsbDeviceDescT) UsbPrintDeviceDescriptor() {
}

/**
 * @brief Print configuration descriptor
 *
 * - This function prints the full contents of a configuration descriptor (including interface and endpoint descriptors)
 * - When a non-standard descriptor is encountered, this function will call the class_specific_cb if it is provided
 *
 * @param[in] cfg_desc Configuration descriptor
 * @param[in] class_specific_cb Class specific descriptor callback. Can be NULL
 */
// llgo:link (*UsbConfigDescT).UsbPrintConfigDescriptor C.usb_print_config_descriptor
func (recv_ *UsbConfigDescT) UsbPrintConfigDescriptor(class_specific_cb PrintClassDescriptorCb) {
}

/**
 * @brief Print a string descriptor
 *
 * This function will only print ASCII characters of the UTF-16 encoded string
 *
 * @param[in] str_desc String descriptor
 */
// llgo:link (*UsbStrDescT).UsbPrintStringDescriptor C.usb_print_string_descriptor
func (recv_ *UsbStrDescT) UsbPrintStringDescriptor() {
}

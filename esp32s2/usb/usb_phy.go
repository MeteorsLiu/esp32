package usb

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const USB_PHY_SUPPORTS_P4_OTG11 = 1

type UsbPhyStatusT c.Int

const (
	USB_PHY_STATUS_FREE   UsbPhyStatusT = 0
	USB_PHY_STATUS_IN_USE UsbPhyStatusT = 1
)

type UsbPhyActionT c.Int

const (
	USB_PHY_ACTION_HOST_ALLOW_CONN    UsbPhyActionT = 0
	USB_PHY_ACTION_HOST_FORCE_DISCONN UsbPhyActionT = 1
	USB_PHY_ACTION_MAX                UsbPhyActionT = 2
)

/**
 * @brief USB external PHY IO pins configuration structure
 */

type UsbPhyExtIoConfT struct {
	VpIoNum        c.Int
	VmIoNum        c.Int
	RcvIoNum       c.Int
	SuspendNIoNum  c.Int
	OenIoNum       c.Int
	VpoIoNum       c.Int
	VmoIoNum       c.Int
	FsEdgeSelIoNum c.Int
}

/**
 * @brief USB OTG IO pins configuration structure
 */

type UsbPhyOtgIoConfT struct {
	IddigIoNum       c.Int
	AvalidIoNum      c.Int
	VbusvalidIoNum   c.Int
	IdpullupIoNum    c.Int
	DppulldownIoNum  c.Int
	DmpulldownIoNum  c.Int
	DrvvbusIoNum     c.Int
	BvalidIoNum      c.Int
	SessendIoNum     c.Int
	ChrgvbusIoNum    c.Int
	DischrgvbusIoNum c.Int
}

/**
 * @brief USB PHY configure struct
 *
 * At minimum the PHY controller and PHY target must be initialized.
 */

type UsbPhyConfigT struct {
	Controller UsbPhyControllerT
	Target     UsbPhyTargetT
	OtgMode    UsbOtgModeT
	OtgSpeed   UsbPhySpeedT
	ExtIoConf  *UsbPhyExtIoConfT
	OtgIoConf  *UsbPhyOtgIoConfT
}

type PhyContextT struct {
	Unused [8]uint8
}
type UsbPhyHandleT *PhyContextT

/**
 * @brief Initialize a new USB PHY
 *        Configure at least PHY source.
 *
 * This function will enable the OTG Controller
 *
 * @param[in]  config     USB PHY configuration struct
 * @param[out] handle_ret USB PHY context handle
 *
 * @return
 *     - ESP_OK                 Success
 *     - ESP_ERR_INVALID_STATE  USB PHY already initialized.
 *     - ESP_ERR_NO_MEM USB_OTG Installation failed due to no mem.
 *     - ESP_ERR_NOT_SUPPORTED  Selected PHY is not supported on this target.
 *     - ESP_ERR_INVALID_ARG    Invalid input argument.
 */
// llgo:link (*UsbPhyConfigT).UsbNewPhy C.usb_new_phy
func (recv_ *UsbPhyConfigT) UsbNewPhy(handle_ret *UsbPhyHandleT) EspErrT {
	return 0
}

/**
 * @brief Configure OTG mode for a USB PHY
 *
 * @param handle Pointer of USB PHY context handle
 * @param mode USB OTG mode
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error.
 *     - ESP_FAIL OTG set mode fail.
 */
//go:linkname UsbPhyOtgSetMode C.usb_phy_otg_set_mode
func UsbPhyOtgSetMode(handle UsbPhyHandleT, mode UsbOtgModeT) EspErrT

/**
 * @brief Configure USB speed for a USB PHY that is operating as an OTG Device
 *
 * @param handle Pointer of USB PHY context handle
 * @param mode USB speed
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error.
 *     - ESP_FAIL OTG set speed fail.
 */
//go:linkname UsbPhyOtgDevSetSpeed C.usb_phy_otg_dev_set_speed
func UsbPhyOtgDevSetSpeed(handle UsbPhyHandleT, speed UsbPhySpeedT) EspErrT

/**
 * @brief Take a action for a USB PHY
 *
 * @param handle Pointer of USB PHY context handle
 * @param action USB PHY action
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error.
 *     - ESP_FAIL Action cannot be performed.
 */
//go:linkname UsbPhyAction C.usb_phy_action
func UsbPhyAction(handle UsbPhyHandleT, action UsbPhyActionT) EspErrT

/**
 * @brief Delete a USB PHY
 *
 * @param handle Pointer of USB PHY context handle
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error.
 */
//go:linkname UsbDelPhy C.usb_del_phy
func UsbDelPhy(handle UsbPhyHandleT) EspErrT

/**
 * @brief Get status of a USB PHY
 *
 * @param[in] target The specific PHY target to check
 * @param[out] status Status of the PHY
 *
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG Parameter error.
 *     - ESP_ERR_INVALID_STATE USB PHY not installed.
 */
//go:linkname UsbPhyGetPhyStatus C.usb_phy_get_phy_status
func UsbPhyGetPhyStatus(target UsbPhyTargetT, status *UsbPhyStatusT) EspErrT

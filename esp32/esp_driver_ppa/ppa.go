package esp_driver_ppa

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PpaOperationT c.Int

const (
	PPA_OPERATION_SRM     PpaOperationT = 0
	PPA_OPERATION_BLEND   PpaOperationT = 1
	PPA_OPERATION_FILL    PpaOperationT = 2
	PPA_OPERATION_INVALID PpaOperationT = 3
)

type PpaClientT struct {
	Unused [8]uint8
}
type PpaClientHandleT *PpaClientT

/**
 * @brief A collection of configuration items that used for registering a PPA client
 */

type PpaClientConfigT struct {
	OperType           PpaOperationT
	MaxPendingTransNum c.Uint32T
	DataBurstLength    PpaDataBurstLengthT
}

/**
 * @brief Register a PPA client to do a specific PPA operation
 *
 * @param[in] config Pointer to a collection of configurations for the client
 * @param[out] ret_client Returned client handle
 *
 * @return
 *      - ESP_OK: Register the PPA client successfully
 *      - ESP_ERR_INVALID_ARG: Register the PPA client failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Register the PPA client failed because out of memory
 *      - ESP_FAIL: Register the PPA client failed because of other error
 */
// llgo:link (*PpaClientConfigT).PpaRegisterClient C.ppa_register_client
func (recv_ *PpaClientConfigT) PpaRegisterClient(ret_client *PpaClientHandleT) EspErrT {
	return 0
}

/**
 * @brief Unregister a PPA client
 *
 * @note This will also free the resources occupied by the client
 *
 * @param[in] ppa_client PPA client handle, allocated by `ppa_register_client`
 *
 * @return
 *      - ESP_OK: Unregister the PPA client successfully
 *      - ESP_ERR_INVALID_ARG: Unregister the PPA client failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Unregister the PPA client failed because there are unfinished transactions
 */
//go:linkname PpaUnregisterClient C.ppa_unregister_client
func PpaUnregisterClient(ppa_client PpaClientHandleT) EspErrT

/**
 * @brief Type of PPA event data
 */

type PpaEventDataT struct {
	Unused [8]uint8
}

// llgo:type C
type PpaEventCallbackT func(PpaClientHandleT, *PpaEventDataT, c.Pointer) bool

/**
 * @brief Group of supported PPA callbacks
 */

type PpaEventCallbacksT struct {
	OnTransDone PpaEventCallbackT
}

/**
 * @brief Register event callbacks for a PPA client
 *
 * @param[in] ppa_client PPA client handle
 * @param[in] cbs Structure with all PPA callbacks
 *
 * @note Any user private data that wants to be passed directly to callback's user_data is provided per PPA transaction.
 *       Please check the `user_data` field in `ppa_xxx_oper_config_t` structure.
 *
 * @return
 *      - ESP_OK: Register event callbacks for the PPA client successfully
 *      - ESP_ERR_INVALID_ARG: Register event callbacks for the PPA client failed because of invalid argument
 */
//go:linkname PpaClientRegisterEventCallbacks C.ppa_client_register_event_callbacks
func PpaClientRegisterEventCallbacks(ppa_client PpaClientHandleT, cbs *PpaEventCallbacksT) EspErrT

/**
 * @brief A collection of configuration items for an input picture and the target block inside the picture
 */

type PpaInPicBlkConfigT struct {
	Buffer       c.Pointer
	PicW         c.Uint32T
	PicH         c.Uint32T
	BlockW       c.Uint32T
	BlockH       c.Uint32T
	BlockOffsetX c.Uint32T
	BlockOffsetY c.Uint32T
	YuvRange     PpaColorRangeT
	YuvStd       PpaColorConvStdRgbYuvT
}

/**
 * @brief A collection of configuration items for an output picture and the target block inside the picture
 */

type PpaOutPicBlkConfigT struct {
	Buffer       c.Pointer
	BufferSize   c.Uint32T
	PicW         c.Uint32T
	PicH         c.Uint32T
	BlockOffsetX c.Uint32T
	BlockOffsetY c.Uint32T
	YuvRange     PpaColorRangeT
	YuvStd       PpaColorConvStdRgbYuvT
}
type PpaTransModeT c.Int

const (
	PPA_TRANS_MODE_BLOCKING     PpaTransModeT = 0
	PPA_TRANS_MODE_NON_BLOCKING PpaTransModeT = 1
)

/**
 * @brief A collection of configuration items to do a PPA SRM operation transaction
 */

type PpaSrmOperConfigT struct {
	In              PpaInPicBlkConfigT
	Out             PpaOutPicBlkConfigT
	RotationAngle   PpaSrmRotationAngleT
	ScaleX          c.Float
	ScaleY          c.Float
	MirrorX         bool
	MirrorY         bool
	RgbSwap         bool
	ByteSwap        bool
	AlphaUpdateMode PpaAlphaUpdateModeT
	Mode            PpaTransModeT
	UserData        c.Pointer
}

/**
 * @brief Perform a scaling-rotating-mirroring (SRM) operation to a picture
 *
 * @param[in] ppa_client PPA client handle that has been registered to do SRM operations
 * @param[in] config Pointer to a collection of configurations for the SRM operation transaction, ppa_srm_oper_config_t
 *
 * @return
 *      - ESP_OK: Perform a SRM operation successfully
 *      - ESP_ERR_INVALID_ARG: Perform a SRM operation failed because of invalid argument
 *      - ESP_FAIL: Perform a SRM operation failed because the client's pending transactions has reached its maximum capacity
 */
//go:linkname PpaDoScaleRotateMirror C.ppa_do_scale_rotate_mirror
func PpaDoScaleRotateMirror(ppa_client PpaClientHandleT, config *PpaSrmOperConfigT) EspErrT

/**
 * @brief A collection of configuration items to do a PPA blend operation transaction
 */

type PpaBlendOperConfigT struct {
	InBg              PpaInPicBlkConfigT
	InFg              PpaInPicBlkConfigT
	Out               PpaOutPicBlkConfigT
	BgRgbSwap         bool
	BgByteSwap        bool
	BgAlphaUpdateMode PpaAlphaUpdateModeT
	FgRgbSwap         bool
	FgByteSwap        bool
	FgAlphaUpdateMode PpaAlphaUpdateModeT
	FgFixRgbVal       ColorPixelRgb888DataT
	BgCkEn            bool
	BgCkRgbLowThres   ColorPixelRgb888DataT
	BgCkRgbHighThres  ColorPixelRgb888DataT
	FgCkEn            bool
	FgCkRgbLowThres   ColorPixelRgb888DataT
	FgCkRgbHighThres  ColorPixelRgb888DataT
	CkRgbDefaultVal   ColorPixelRgb888DataT
	CkReverseBg2fg    bool
	Mode              PpaTransModeT
	UserData          c.Pointer
}

/**
 * @brief Perform a blending operation to a picture
 *
 * @param[in] ppa_client PPA client handle that has been registered to do blend operations
 * @param[in] config Pointer to a collection of configurations for the blend operation transaction, ppa_blend_oper_config_t
 *
 * @return
 *      - ESP_OK: Perform a blend operation successfully
 *      - ESP_ERR_INVALID_ARG: Perform a blend operation failed because of invalid argument
 *      - ESP_FAIL: Perform a blend operation failed because the client's pending transactions has reached its maximum capacity
 */
//go:linkname PpaDoBlend C.ppa_do_blend
func PpaDoBlend(ppa_client PpaClientHandleT, config *PpaBlendOperConfigT) EspErrT

/**
 * @brief A collection of configuration items to do a PPA fill operation transaction
 */

type PpaFillOperConfigT struct {
	Out           PpaOutPicBlkConfigT
	FillBlockW    c.Uint32T
	FillBlockH    c.Uint32T
	FillArgbColor ColorPixelArgb8888DataT
	Mode          PpaTransModeT
	UserData      c.Pointer
}

/**
 * @brief Perform a filling operation to a picture
 *
 * @param[in] ppa_client PPA client handle that has been registered to do fill operations
 * @param[in] config Pointer to a collection of configurations for the fill operation transaction, ppa_fill_oper_config_t
 *
 * @return
 *      - ESP_OK: Perform a fill operation successfully
 *      - ESP_ERR_INVALID_ARG: Perform a fill operation failed because of invalid argument
 *      - ESP_FAIL: Perform a fill operation failed because the client's pending transactions has reached its maximum capacity
 */
//go:linkname PpaDoFill C.ppa_do_fill
func PpaDoFill(ppa_client PpaClientHandleT, config *PpaFillOperConfigT) EspErrT

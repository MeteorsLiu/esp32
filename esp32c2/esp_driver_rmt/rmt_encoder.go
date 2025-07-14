package esp_driver_rmt

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RmtEncodeStateT c.Int

const (
	RMT_ENCODING_RESET    RmtEncodeStateT = 0
	RMT_ENCODING_COMPLETE RmtEncodeStateT = 1
	RMT_ENCODING_MEM_FULL RmtEncodeStateT = 2
)

// llgo:type C
type RmtEncodeSimpleCbT func(c.Pointer, c.SizeT, c.SizeT, c.SizeT, *RmtSymbolWordT, *bool, c.Pointer) c.SizeT

/**
 * @brief Bytes encoder configuration
 */

type RmtBytesEncoderConfigT struct {
	Bit0  RmtSymbolWordT
	Bit1  RmtSymbolWordT
	Flags struct {
		MsbFirst c.Uint32T
	}
}

/**
 * @brief Copy encoder configuration
 */

type RmtCopyEncoderConfigT struct {
	Unused [8]uint8
}

/**
 * @brief Simple callback encoder configuration
 */

type RmtSimpleEncoderConfigT struct {
	Callback     RmtEncodeSimpleCbT
	Arg          c.Pointer
	MinChunkSize c.SizeT
}

/**
 * @brief Create RMT bytes encoder, which can encode byte stream into RMT symbols
 *
 * @param[in] config Bytes encoder configuration
 * @param[out] ret_encoder Returned encoder handle
 * @return
 *      - ESP_OK: Create RMT bytes encoder successfully
 *      - ESP_ERR_INVALID_ARG: Create RMT bytes encoder failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create RMT bytes encoder failed because out of memory
 *      - ESP_FAIL: Create RMT bytes encoder failed because of other error
 */
// llgo:link (*RmtBytesEncoderConfigT).RmtNewBytesEncoder C.rmt_new_bytes_encoder
func (recv_ *RmtBytesEncoderConfigT) RmtNewBytesEncoder(ret_encoder *RmtEncoderHandleT) EspErrT {
	return 0
}

/**
 * @brief Update the configuration of the bytes encoder
 *
 * @note The configurations of the bytes encoder is also set up by `rmt_new_bytes_encoder()`.
 *       This function is used to update the configuration of the bytes encoder at runtime.
 *
 * @param[in] bytes_encoder Bytes encoder handle, created by e.g `rmt_new_bytes_encoder()`
 * @param[in] config Bytes encoder configuration
 * @return
 *      - ESP_OK: Update RMT bytes encoder successfully
 *      - ESP_ERR_INVALID_ARG: Update RMT bytes encoder failed because of invalid argument
 *      - ESP_FAIL: Update RMT bytes encoder failed because of other error
 */
//go:linkname RmtBytesEncoderUpdateConfig C.rmt_bytes_encoder_update_config
func RmtBytesEncoderUpdateConfig(bytes_encoder RmtEncoderHandleT, config *RmtBytesEncoderConfigT) EspErrT

/**
 * @brief Create RMT copy encoder, which copies the given RMT symbols into RMT memory
 *
 * @param[in] config Copy encoder configuration
 * @param[out] ret_encoder Returned encoder handle
 * @return
 *      - ESP_OK: Create RMT copy encoder successfully
 *      - ESP_ERR_INVALID_ARG: Create RMT copy encoder failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create RMT copy encoder failed because out of memory
 *      - ESP_FAIL: Create RMT copy encoder failed because of other error
 */
// llgo:link (*RmtCopyEncoderConfigT).RmtNewCopyEncoder C.rmt_new_copy_encoder
func (recv_ *RmtCopyEncoderConfigT) RmtNewCopyEncoder(ret_encoder *RmtEncoderHandleT) EspErrT {
	return 0
}

/**
 * @brief Create RMT simple callback encoder, which uses a callback to convert incoming
 *        data into RMT symbols.
 *
 * @param[in] config Simple callback encoder configuration
 * @param[out] ret_encoder Returned encoder handle
 * @return
 *      - ESP_OK: Create RMT simple callback encoder successfully
 *      - ESP_ERR_INVALID_ARG: Create RMT simple callback encoder failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create RMT simple callback encoder failed because out of memory
 *      - ESP_FAIL: Create RMT simple callback encoder failed because of other error
 */
// llgo:link (*RmtSimpleEncoderConfigT).RmtNewSimpleEncoder C.rmt_new_simple_encoder
func (recv_ *RmtSimpleEncoderConfigT) RmtNewSimpleEncoder(ret_encoder *RmtEncoderHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete RMT encoder
 *
 * @param[in] encoder RMT encoder handle, created by e.g `rmt_new_bytes_encoder()`
 * @return
 *      - ESP_OK: Delete RMT encoder successfully
 *      - ESP_ERR_INVALID_ARG: Delete RMT encoder failed because of invalid argument
 *      - ESP_FAIL: Delete RMT encoder failed because of other error
 */
//go:linkname RmtDelEncoder C.rmt_del_encoder
func RmtDelEncoder(encoder RmtEncoderHandleT) EspErrT

/**
 * @brief Reset RMT encoder
 *
 * @param[in] encoder RMT encoder handle, created by e.g `rmt_new_bytes_encoder()`
 * @return
 *      - ESP_OK: Reset RMT encoder successfully
 *      - ESP_ERR_INVALID_ARG: Reset RMT encoder failed because of invalid argument
 *      - ESP_FAIL: Reset RMT encoder failed because of other error
 */
//go:linkname RmtEncoderReset C.rmt_encoder_reset
func RmtEncoderReset(encoder RmtEncoderHandleT) EspErrT

/**
 * @brief A helper function to allocate a proper memory for RMT encoder
 *
 * @param size Size of memory to be allocated
 * @return Pointer to the allocated memory if the allocation is successful, NULL otherwise
 */
//go:linkname RmtAllocEncoderMem C.rmt_alloc_encoder_mem
func RmtAllocEncoderMem(size c.SizeT) c.Pointer

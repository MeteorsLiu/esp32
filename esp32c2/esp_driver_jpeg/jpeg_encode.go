package esp_driver_jpeg

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief JPEG encoder configure structure
 */

type JpegEncodeCfgT struct {
	Height       c.Uint32T
	Width        c.Uint32T
	SrcType      JpegEncInputFormatT
	SubSample    JpegDownSamplingTypeT
	ImageQuality c.Uint32T
}

/**
 * @brief Configuration parameters for the JPEG encode engine.
 */

type JpegEncodeEngineCfgT struct {
	IntrPriority c.Int
	TimeoutMs    c.Int
}

/**
 * @brief JPEG encoder memory allocation config
 */

type JpegEncodeMemoryAllocCfgT struct {
	BufferDirection JpegEncBufferAllocDirectionT
}

/**
 * @brief Allocate JPEG encoder
 *
 * @param[in] enc_eng_cfg config for jpeg encoder
 * @param[out] ret_encoder handle for jpeg encoder
 * @return
 *      - ESP_OK: JPEG encoder initialized successfully.
 *      - ESP_ERR_INVALID_ARG: JPEG encoder initialization failed because of invalid argument.
 *      - ESP_ERR_NO_MEM: Create JPEG encoder failed because of out of memory.
 */
// llgo:link (*JpegEncodeEngineCfgT).JpegNewEncoderEngine C.jpeg_new_encoder_engine
func (recv_ *JpegEncodeEngineCfgT) JpegNewEncoderEngine(ret_encoder *JpegEncoderHandleT) EspErrT {
	return 0
}

/**
 * @brief Process encoding of JPEG data using the specified encoder engine.
 *
 * This function processes the encoding of JPEG data using the provided encoder engine
 * and configuration. It takes an input buffer containing the raw image data, performs
 * encoding based on the configuration settings, and outputs the compressed bitstream.
 *
 * @param[in] encoder_engine Handle to the JPEG encoder engine to be used for encoding.
 * @param[in] encode_cfg Pointer to the configuration structure for the JPEG encoding process.
 * @param[in] encode_inbuf Pointer to the input buffer containing the raw image data.
 * @param[in] inbuf_size Size of the input buffer in bytes.
 * @param[in] encode_outbuf Pointer to the output buffer where the compressed bitstream will be stored.
 * @param[in] outbuf_size The size of output buffer.
 * @param[out] out_size Pointer to a variable where the size of the output bitstream will be stored.
 *
 * @return
 *      - ESP_OK: JPEG encoder process successfully.
 *      - ESP_ERR_INVALID_ARG: JPEG encoder process failed because of invalid argument.
 *      - ESP_ERR_TIMEOUT: JPEG encoder process timeout.
 */
//go:linkname JpegEncoderProcess C.jpeg_encoder_process
func JpegEncoderProcess(encoder_engine JpegEncoderHandleT, encode_cfg *JpegEncodeCfgT, encode_inbuf *c.Uint8T, inbuf_size c.Uint32T, encode_outbuf *c.Uint8T, outbuf_size c.Uint32T, out_size *c.Uint32T) EspErrT

/**
 * @brief Release resources used by a JPEG encoder instance.
 *
 * This function releases the resources used by the specified JPEG encoder instance. The encoder instance is
 * specified by the `encoder_engine` parameter.
 *
 * @param[in] encoder_engine Handle of the JPEG encoder instance to release resources for.
 * @return
 *      - ESP_OK: Delete JPEG encoder successfully.
 *      - ESP_ERR_INVALID_ARG: Delete JPEG encoder failed because of invalid argument.
 */
//go:linkname JpegDelEncoderEngine C.jpeg_del_encoder_engine
func JpegDelEncoderEngine(encoder_engine JpegEncoderHandleT) EspErrT

/**
 * @brief A helper function to allocate memory space for JPEG encoder.
 *
 * @param[in] size The size of memory to allocate.
 * @param[in] mem_cfg Memory configuration for memory allocation
 * @param[out] allocated_size Actual allocated buffer size.
 * @return Pointer to the allocated memory space, or NULL if allocation fails.
 */
//go:linkname JpegAllocEncoderMem C.jpeg_alloc_encoder_mem
func JpegAllocEncoderMem(size c.SizeT, mem_cfg *JpegEncodeMemoryAllocCfgT, allocated_size *c.SizeT) c.Pointer

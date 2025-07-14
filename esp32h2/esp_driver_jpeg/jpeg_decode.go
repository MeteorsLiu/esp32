package esp_driver_jpeg

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Configuration parameters for a JPEG decoder image process.
 */

type JpegDecodeCfgT struct {
	OutputFormat JpegDecOutputFormatT
	RgbOrder     JpegDecRgbElementOrderT
	ConvStd      JpegYuvRgbConvStdT
}

/**
 * @brief Configuration parameters for the JPEG decoder engine.
 */

type JpegDecodeEngineCfgT struct {
	IntrPriority c.Int
	TimeoutMs    c.Int
}

/**
 * @brief Structure for jpeg decode header
 */

type JpegDecodePictureInfoT struct {
	Width        c.Uint32T
	Height       c.Uint32T
	SampleMethod JpegDownSamplingTypeT
}

/**
 * @brief JPEG decoder memory allocation config
 */

type JpegDecodeMemoryAllocCfgT struct {
	BufferDirection JpegDecBufferAllocDirectionT
}

/**
 * @brief Acquire a JPEG decode engine with the specified configuration.
 *
 * This function acquires a JPEG decode engine with the specified configuration. The configuration
 * parameters are provided through the `dec_eng_cfg` structure, and the resulting JPEG decoder handle
 * is returned through the `ret_decoder` pointer.
 *
 * @param[in] dec_eng_cfg Pointer to the JPEG decode engine configuration.
 * @param[out] ret_decoder Pointer to a variable that will receive the JPEG decoder handle.
 * @return
 *      - ESP_OK: JPEG decoder initialized successfully.
 *      - ESP_ERR_INVALID_ARG: JPEG decoder initialization failed because of invalid argument.
 *      - ESP_ERR_NO_MEM: Create JPEG decoder failed because of out of memory.
 */
// llgo:link (*JpegDecodeEngineCfgT).JpegNewDecoderEngine C.jpeg_new_decoder_engine
func (recv_ *JpegDecodeEngineCfgT) JpegNewDecoderEngine(ret_decoder *JpegDecoderHandleT) EspErrT {
	return 0
}

/**
 * @brief Helper function for getting information about a JPEG image.
 *
 * This function analyzes the provided JPEG image data and retrieves information about the image,
 * such as its width, height. The image data is specified by the `bit_stream` pointer and the `stream_size` parameter.
 * The resulting image information is returned through the `picture_info` structure.
 *
 * @note This function doesn't depend on any jpeg hardware, it helps user to know jpeg information from jpeg header. For example,
 * user can get picture width and height via this function and malloc a reasonable size buffer for jpeg engine process.
 *
 * @param[in] bit_stream Pointer to the buffer containing the JPEG image data.
 * @param[in] stream_size Size of the JPEG image data in bytes. Note that parse beginning partial of picture also works, but the beginning partial should be enough given.
 * @param[out] picture_info Pointer to the structure that will receive the image information.
 * @return
 *      - ESP_OK: JPEG decoder get jpg image header successfully.
 *      - ESP_ERR_INVALID_ARG: JPEG decoder get header info failed because of invalid argument.
 */
//go:linkname JpegDecoderGetInfo C.jpeg_decoder_get_info
func JpegDecoderGetInfo(bit_stream *c.Uint8T, stream_size c.Uint32T, picture_info *JpegDecodePictureInfoT) EspErrT

/**
 * @brief Process a JPEG image with the specified decoder instance.
 *
 * This function processes the provided JPEG image data using the specified JPEG decoder instance. The input
 * JPEG image data is specified by the `bit_stream` pointer and the `stream_size` parameter. The resulting
 * decoded image data is written to the `decode_outbuf` buffer, and the length of the output image data is
 * returned through the `out_size` pointer.
 *
 * @note 1.Please make sure that the content of `bit_stream` pointer cannot be modified until this function returns.
 *       2.Please note that the output size of image is always the multiple of 16 depends on protocol of JPEG.
 *
 * @param[in] decoder_engine Handle of the JPEG decoder instance to use for processing.
 * @param[in] decode_cfg Config structure of decoder.
 * @param[in] bit_stream Pointer to the buffer containing the input JPEG image data.
 * @param[in] stream_size Size of the input JPEG image data in bytes.
 * @param[in] decode_outbuf Pointer to the buffer that will receive the decoded image data.
 * @param[in] outbuf_size The size of `decode_outbuf`
 * @param[out] out_size Pointer to a variable that will receive the length of the output image data.
 * @return
 *      - ESP_OK: JPEG decoder process successfully.
 *      - ESP_ERR_INVALID_ARG: JPEG decoder process failed because of invalid argument.
 */
//go:linkname JpegDecoderProcess C.jpeg_decoder_process
func JpegDecoderProcess(decoder_engine JpegDecoderHandleT, decode_cfg *JpegDecodeCfgT, bit_stream *c.Uint8T, stream_size c.Uint32T, decode_outbuf *c.Uint8T, outbuf_size c.Uint32T, out_size *c.Uint32T) EspErrT

/**
 * @brief Release resources used by a JPEG decoder instance.
 *
 * This function releases the resources used by the specified JPEG decoder instance. The decoder instance is
 * specified by the `decoder_engine` parameter.
 *
 * @param decoder_engine Handle of the JPEG decoder instance to release resources for.
 * @return
 *      - ESP_OK: Delete JPEG decoder successfully.
 *      - ESP_ERR_INVALID_ARG: Delete JPEG decoder failed because of invalid argument.
 */
//go:linkname JpegDelDecoderEngine C.jpeg_del_decoder_engine
func JpegDelDecoderEngine(decoder_engine JpegDecoderHandleT) EspErrT

/**
 * @brief A helper function to allocate memory space for JPEG decoder.
 *
 * @param[in] size The size of memory to allocate.
 * @param[in] mem_cfg Memory configuration for memory allocation
 * @param[out] allocated_size Actual allocated buffer size.
 * @return Pointer to the allocated memory space, or NULL if allocation fails.
 */
//go:linkname JpegAllocDecoderMem C.jpeg_alloc_decoder_mem
func JpegAllocDecoderMem(size c.SizeT, mem_cfg *JpegDecodeMemoryAllocCfgT, allocated_size *c.SizeT) c.Pointer

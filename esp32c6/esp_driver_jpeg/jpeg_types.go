package esp_driver_jpeg

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type JpegDecOutputFormatT c.Int

const (
	JPEG_DECODE_OUT_FORMAT_RGB888 JpegDecOutputFormatT = 33554432
	JPEG_DECODE_OUT_FORMAT_RGB565 JpegDecOutputFormatT = 33554434
	JPEG_DECODE_OUT_FORMAT_GRAY   JpegDecOutputFormatT = 67108865
	JPEG_DECODE_OUT_FORMAT_YUV444 JpegDecOutputFormatT = 50331648
	JPEG_DECODE_OUT_FORMAT_YUV422 JpegDecOutputFormatT = 50331649
	JPEG_DECODE_OUT_FORMAT_YUV420 JpegDecOutputFormatT = 50331650
)

type JpegYuvRgbConvStdT c.Int

const (
	JPEG_YUV_RGB_CONV_STD_BT601 JpegYuvRgbConvStdT = 0
	JPEG_YUV_RGB_CONV_STD_BT709 JpegYuvRgbConvStdT = 1
)

type JpegDecRgbElementOrderT c.Int

const (
	JPEG_DEC_RGB_ELEMENT_ORDER_BGR JpegDecRgbElementOrderT = 1
	JPEG_DEC_RGB_ELEMENT_ORDER_RGB JpegDecRgbElementOrderT = 0
)

type JpegDecBufferAllocDirectionT c.Int

const (
	JPEG_DEC_ALLOC_INPUT_BUFFER  JpegDecBufferAllocDirectionT = 0
	JPEG_DEC_ALLOC_OUTPUT_BUFFER JpegDecBufferAllocDirectionT = 1
)

type JpegEncInputFormatT c.Int

const (
	JPEG_ENCODE_IN_FORMAT_RGB888 JpegEncInputFormatT = 33554432
	JPEG_ENCODE_IN_FORMAT_RGB565 JpegEncInputFormatT = 33554434
	JPEG_ENCODE_IN_FORMAT_GRAY   JpegEncInputFormatT = 67108865
	JPEG_ENCODE_IN_FORMAT_YUV422 JpegEncInputFormatT = 50331649
)

type JpegEncBufferAllocDirectionT c.Int

const (
	JPEG_ENC_ALLOC_INPUT_BUFFER  JpegEncBufferAllocDirectionT = 0
	JPEG_ENC_ALLOC_OUTPUT_BUFFER JpegEncBufferAllocDirectionT = 1
)

type JpegDecoderT struct {
	Unused [8]uint8
}
type JpegDecoderHandleT *JpegDecoderT

type JpegEncoderT struct {
	Unused [8]uint8
}
type JpegEncoderHandleT *JpegEncoderT

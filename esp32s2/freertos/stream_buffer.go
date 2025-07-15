package freertos

import _ "unsafe"

/**
 * Type by which stream buffers are referenced.  For example, a call to
 * xStreamBufferCreate() returns an StreamBufferHandle_t variable that can
 * then be used as a parameter to xStreamBufferSend(), xStreamBufferReceive(),
 * etc.
 */

type StreamBufferDefT struct {
	Unused [8]uint8
}
type StreamBufferHandleT *StreamBufferDefT

// llgo:type C
type StreamBufferCallbackFunctionT func(StreamBufferHandleT, BaseTypeT, BaseTypeT)

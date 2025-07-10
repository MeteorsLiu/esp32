package esp_eth

import _ "unsafe"

const EMAC_DMA_BUF_SIZE_AUTO = 0

type EmacEspDmaT struct {
	Unused [8]uint8
}
type EmacEspDmaHandleT *EmacEspDmaT

type EmacEspDmaConfigT struct {
	Unused [8]uint8
}

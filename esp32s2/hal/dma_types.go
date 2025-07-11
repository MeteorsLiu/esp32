package hal

import _ "unsafe"

type DmaDescriptorS struct {
	Unused [8]uint8
}
type DmaDescriptorT DmaDescriptorS
type DmaDescriptorAlign4T DmaDescriptorT

type DmaDescriptorAlign8S struct {
	Unused [8]uint8
}
type DmaDescriptorAlign8T DmaDescriptorAlign8S

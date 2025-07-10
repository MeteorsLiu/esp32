package xtensa

import _ "unsafe"

const XTOS_KEEPON_MEM = 0x00000100
const XTOS_KEEPON_MEM_SHIFT = 8
const XTOS_KEEPON_DEBUG = 0x00001000
const XTOS_KEEPON_DEBUG_SHIFT = 12
const XTOS_IDMA_NO_WAIT = 0x00010000
const XTOS_IDMA_WAIT_STANDBY = 0x00020000
const XTOS_COREF_PSO = 0x00000001
const XTOS_COREF_PSO_SHIFT = 0

// llgo:type C
type X_xtosHandlerFunc func()
type X_xtosHandler *X_xtosHandlerFunc

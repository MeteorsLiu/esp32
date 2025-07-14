package riscv

import _ "unsafe"

const CSR_PMACFG0 = 0xBC0
const CSR_PMAADDR0 = 0xBD0
const PMA_SHIFT = 2
const PMA_TOR = 0x40000000
const PMA_NA4 = 0x80000000
const PMA_NAPOT = 0xC0000000
const CSR_PMPCFG0 = 0x3A0
const CSR_PMPADDR0 = 0x3B0
const PMPADDR_ALL = 0xFFFFFFFF
const MEXSTATUS = 0x7E1
const MHINT = 0x7C5
const LDPC0 = 0xBE0
const LDPC1 = 0xBE1
const STPC0 = 0xBF0
const STPC1 = 0xBF1
const STPC2 = 0xBF2

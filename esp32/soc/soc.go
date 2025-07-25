package soc

import _ "unsafe"

const SOC_MAX_CONTIGUOUS_RAM_SIZE = 0x400000
const SPI_CLK_DIV = 4
const TICKS_PER_US_ROM = 26
const SOC_DROM_LOW = 0x3F400000
const SOC_DROM_HIGH = 0x3F800000
const SOC_DRAM_LOW = 0x3FFAE000
const SOC_DRAM_HIGH = 0x40000000
const SOC_IROM_LOW = 0x400D0000
const SOC_IROM_HIGH = 0x40400000
const SOC_IROM_MASK_LOW = 0x40000000
const SOC_IROM_MASK_HIGH = 0x40070000
const SOC_CACHE_PRO_LOW = 0x40070000
const SOC_CACHE_PRO_HIGH = 0x40078000
const SOC_CACHE_APP_LOW = 0x40078000
const SOC_CACHE_APP_HIGH = 0x40080000
const SOC_IRAM_LOW = 0x40080000
const SOC_IRAM_HIGH = 0x400AA000
const SOC_RTC_IRAM_LOW = 0x400C0000
const SOC_RTC_IRAM_HIGH = 0x400C2000
const SOC_RTC_DRAM_LOW = 0x3FF80000
const SOC_RTC_DRAM_HIGH = 0x3FF82000
const SOC_RTC_DATA_LOW = 0x50000000
const SOC_RTC_DATA_HIGH = 0x50002000
const SOC_EXTRAM_DATA_LOW = 0x3F800000
const SOC_EXTRAM_DATA_HIGH = 0x3FC00000
const SOC_DIRAM_IRAM_LOW = 0x400A0000
const SOC_DIRAM_IRAM_HIGH = 0x400C0000
const SOC_DIRAM_DRAM_LOW = 0x3FFE0000
const SOC_DIRAM_DRAM_HIGH = 0x40000000
const SOC_DIRAM_INVERTED = 1
const SOC_DMA_LOW = 0x3FFAE000
const SOC_DMA_HIGH = 0x40000000
const SOC_BYTE_ACCESSIBLE_LOW = 0x3FF90000
const SOC_BYTE_ACCESSIBLE_HIGH = 0x40000000
const SOC_MEM_INTERNAL_LOW = 0x3FF90000
const SOC_MEM_INTERNAL_HIGH = 0x400C2000
const SOC_ROM_STACK_START = 0x3ffe3f20
const ETS_WMAC_INUM = 0
const ETS_BT_HOST_INUM = 1
const ETS_WBB_INUM = 4
const ETS_T1_WDT_INUM = 24
const ETS_MEMACCESS_ERR_INUM = 25
const ETS_IPC_ISR_INUM = 28
const ETS_SLC_INUM = 1
const ETS_UART0_INUM = 5
const ETS_UART1_INUM = 5
const ETS_INVALID_INUM = 6

package soc

import _ "unsafe"

const SOC_IROM_LOW = 0x42000000
const SOC_IROM_MASK_LOW = 0x40000000
const SOC_IROM_MASK_HIGH = 0x40020000
const SOC_DROM_MASK_LOW = 0x40000000
const SOC_DROM_MASK_HIGH = 0x40020000
const SOC_IRAM_LOW = 0x40800000
const SOC_IRAM_HIGH = 0x40850000
const SOC_DRAM_LOW = 0x40800000
const SOC_DRAM_HIGH = 0x40850000
const SOC_RTC_IRAM_LOW = 0x50000000
const SOC_RTC_IRAM_HIGH = 0x50001000
const SOC_RTC_DRAM_LOW = 0x50000000
const SOC_RTC_DRAM_HIGH = 0x50001000
const SOC_RTC_DATA_LOW = 0x50000000
const SOC_RTC_DATA_HIGH = 0x50001000
const SOC_DIRAM_IRAM_LOW = 0x40800000
const SOC_DIRAM_IRAM_HIGH = 0x40850000
const SOC_DIRAM_DRAM_LOW = 0x40800000
const SOC_DIRAM_DRAM_HIGH = 0x40850000
const SOC_DMA_LOW = 0x40800000
const SOC_DMA_HIGH = 0x40850000
const SOC_BYTE_ACCESSIBLE_LOW = 0x40800000
const SOC_BYTE_ACCESSIBLE_HIGH = 0x40850000
const SOC_MEM_INTERNAL_LOW = 0x40800000
const SOC_MEM_INTERNAL_HIGH = 0x40850000
const SOC_MEM_INTERNAL_LOW1 = 0x40800000
const SOC_MEM_INTERNAL_HIGH1 = 0x40850000
const SOC_PERIPHERAL_LOW = 0x60000000
const SOC_PERIPHERAL_HIGH = 0x60100000
const SOC_CPU_SUBSYSTEM_LOW = 0x20000000
const SOC_CPU_SUBSYSTEM_HIGH = 0x30000000
const SOC_ROM_STACK_START = 0x4084f380
const SOC_ROM_STACK_SIZE = 0x2000
const ETS_T1_WDT_INUM = 24
const ETS_CACHEERR_INUM = 25
const ETS_MEMPROT_ERR_INUM = 26
const ETS_ASSIST_DEBUG_INUM = 27
const ETS_MAX_INUM = 31
const ETS_SLC_INUM = 1
const ETS_UART0_INUM = 5
const ETS_UART1_INUM = 5
const ETS_SPI2_INUM = 1
const ETS_GPIO_INUM = 4
const ETS_INVALID_INUM = 0
const SOC_INTERRUPT_LEVEL_MEDIUM = 4

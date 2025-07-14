package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PeriphInterrputT c.Int

const (
	ETS_WIFI_MAC_INTR_SOURCE              PeriphInterrputT = 0
	ETS_WIFI_MAC_NMI_SOURCE               PeriphInterrputT = 1
	ETS_WIFI_PWR_INTR_SOURCE              PeriphInterrputT = 2
	ETS_WIFI_BB_INTR_SOURCE               PeriphInterrputT = 3
	ETS_BT_MAC_INTR_SOURCE                PeriphInterrputT = 4
	ETS_BT_BB_INTR_SOURCE                 PeriphInterrputT = 5
	ETS_BT_BB_NMI_SOURCE                  PeriphInterrputT = 6
	ETS_RWBT_INTR_SOURCE                  PeriphInterrputT = 7
	ETS_RWBLE_INTR_SOURCE                 PeriphInterrputT = 8
	ETS_RWBT_NMI_SOURCE                   PeriphInterrputT = 9
	ETS_RWBLE_NMI_SOURCE                  PeriphInterrputT = 10
	ETS_I2C_MASTER_SOURCE                 PeriphInterrputT = 11
	ETS_SLC0_INTR_SOURCE                  PeriphInterrputT = 12
	ETS_SLC1_INTR_SOURCE                  PeriphInterrputT = 13
	ETS_APB_CTRL_INTR_SOURCE              PeriphInterrputT = 14
	ETS_UHCI0_INTR_SOURCE                 PeriphInterrputT = 15
	ETS_GPIO_INTR_SOURCE                  PeriphInterrputT = 16
	ETS_GPIO_NMI_SOURCE                   PeriphInterrputT = 17
	ETS_SPI1_INTR_SOURCE                  PeriphInterrputT = 18
	ETS_SPI2_INTR_SOURCE                  PeriphInterrputT = 19
	ETS_I2S0_INTR_SOURCE                  PeriphInterrputT = 20
	ETS_UART0_INTR_SOURCE                 PeriphInterrputT = 21
	ETS_UART1_INTR_SOURCE                 PeriphInterrputT = 22
	ETS_LEDC_INTR_SOURCE                  PeriphInterrputT = 23
	ETS_EFUSE_INTR_SOURCE                 PeriphInterrputT = 24
	ETS_TWAI_INTR_SOURCE                  PeriphInterrputT = 25
	ETS_USB_SERIAL_JTAG_INTR_SOURCE       PeriphInterrputT = 26
	ETS_RTC_CORE_INTR_SOURCE              PeriphInterrputT = 27
	ETS_RMT_INTR_SOURCE                   PeriphInterrputT = 28
	ETS_I2C_EXT0_INTR_SOURCE              PeriphInterrputT = 29
	ETS_TIMER1_INTR_SOURCE                PeriphInterrputT = 30
	ETS_TIMER2_INTR_SOURCE                PeriphInterrputT = 31
	ETS_TG0_T0_LEVEL_INTR_SOURCE          PeriphInterrputT = 32
	ETS_TG0_WDT_LEVEL_INTR_SOURCE         PeriphInterrputT = 33
	ETS_TG1_T0_LEVEL_INTR_SOURCE          PeriphInterrputT = 34
	ETS_TG1_WDT_LEVEL_INTR_SOURCE         PeriphInterrputT = 35
	ETS_CACHE_IA_INTR_SOURCE              PeriphInterrputT = 36
	ETS_SYSTIMER_TARGET0_INTR_SOURCE      PeriphInterrputT = 37
	ETS_SYSTIMER_TARGET1_INTR_SOURCE      PeriphInterrputT = 38
	ETS_SYSTIMER_TARGET2_INTR_SOURCE      PeriphInterrputT = 39
	ETS_SYSTIMER_TARGET0_EDGE_INTR_SOURCE PeriphInterrputT = 37
	ETS_SYSTIMER_TARGET1_EDGE_INTR_SOURCE PeriphInterrputT = 38
	ETS_SYSTIMER_TARGET2_EDGE_INTR_SOURCE PeriphInterrputT = 39
	ETS_SPI_MEM_REJECT_CACHE_INTR_SOURCE  PeriphInterrputT = 40
	ETS_ICACHE_PRELOAD0_INTR_SOURCE       PeriphInterrputT = 41
	ETS_ICACHE_SYNC0_INTR_SOURCE          PeriphInterrputT = 42
	ETS_APB_ADC_INTR_SOURCE               PeriphInterrputT = 43
	ETS_DMA_CH0_INTR_SOURCE               PeriphInterrputT = 44
	ETS_DMA_CH1_INTR_SOURCE               PeriphInterrputT = 45
	ETS_DMA_CH2_INTR_SOURCE               PeriphInterrputT = 46
	ETS_RSA_INTR_SOURCE                   PeriphInterrputT = 47
	ETS_AES_INTR_SOURCE                   PeriphInterrputT = 48
	ETS_SHA_INTR_SOURCE                   PeriphInterrputT = 49
	ETS_FROM_CPU_INTR0_SOURCE             PeriphInterrputT = 50
	ETS_FROM_CPU_INTR1_SOURCE             PeriphInterrputT = 51
	ETS_FROM_CPU_INTR2_SOURCE             PeriphInterrputT = 52
	ETS_FROM_CPU_INTR3_SOURCE             PeriphInterrputT = 53
	ETS_ASSIST_DEBUG_INTR_SOURCE          PeriphInterrputT = 54
	ETS_DMA_APBPERI_PMS_INTR_SOURCE       PeriphInterrputT = 55
	ETS_CORE0_IRAM0_PMS_INTR_SOURCE       PeriphInterrputT = 56
	ETS_CORE0_DRAM0_PMS_INTR_SOURCE       PeriphInterrputT = 57
	ETS_CORE0_PIF_PMS_INTR_SOURCE         PeriphInterrputT = 58
	ETS_CORE0_PIF_PMS_SIZE_INTR_SOURCE    PeriphInterrputT = 59
	ETS_BAK_PMS_VIOLATE_INTR_SOURCE       PeriphInterrputT = 60
	ETS_CACHE_CORE0_ACS_INTR_SOURCE       PeriphInterrputT = 61
	ETS_MAX_INTR_SOURCE                   PeriphInterrputT = 62
)

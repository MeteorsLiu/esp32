package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PeriphInterrputT c.Int

const (
	ETS_WIFI_MAC_INTR_SOURCE       PeriphInterrputT = 0
	ETS_WIFI_MAC_NMI_SOURCE        PeriphInterrputT = 1
	ETS_WIFI_BB_INTR_SOURCE        PeriphInterrputT = 2
	ETS_BT_MAC_INTR_SOURCE         PeriphInterrputT = 3
	ETS_BT_BB_INTR_SOURCE          PeriphInterrputT = 4
	ETS_BT_BB_NMI_SOURCE           PeriphInterrputT = 5
	ETS_RWBT_INTR_SOURCE           PeriphInterrputT = 6
	ETS_RWBLE_INTR_SOURCE          PeriphInterrputT = 7
	ETS_RWBT_NMI_SOURCE            PeriphInterrputT = 8
	ETS_RWBLE_NMI_SOURCE           PeriphInterrputT = 9
	ETS_SLC0_INTR_SOURCE           PeriphInterrputT = 10
	ETS_SLC1_INTR_SOURCE           PeriphInterrputT = 11
	ETS_UHCI0_INTR_SOURCE          PeriphInterrputT = 12
	ETS_UHCI1_INTR_SOURCE          PeriphInterrputT = 13
	ETS_TG0_T0_LEVEL_INTR_SOURCE   PeriphInterrputT = 14
	ETS_TG0_T1_LEVEL_INTR_SOURCE   PeriphInterrputT = 15
	ETS_TG0_WDT_LEVEL_INTR_SOURCE  PeriphInterrputT = 16
	ETS_TG0_LACT_LEVEL_INTR_SOURCE PeriphInterrputT = 17
	ETS_TG1_T0_LEVEL_INTR_SOURCE   PeriphInterrputT = 18
	ETS_TG1_T1_LEVEL_INTR_SOURCE   PeriphInterrputT = 19
	ETS_TG1_WDT_LEVEL_INTR_SOURCE  PeriphInterrputT = 20
	ETS_TG1_LACT_LEVEL_INTR_SOURCE PeriphInterrputT = 21
	ETS_GPIO_INTR_SOURCE           PeriphInterrputT = 22
	ETS_GPIO_NMI_SOURCE            PeriphInterrputT = 23
	ETS_FROM_CPU_INTR0_SOURCE      PeriphInterrputT = 24
	ETS_FROM_CPU_INTR1_SOURCE      PeriphInterrputT = 25
	ETS_FROM_CPU_INTR2_SOURCE      PeriphInterrputT = 26
	ETS_FROM_CPU_INTR3_SOURCE      PeriphInterrputT = 27
	ETS_SPI0_INTR_SOURCE           PeriphInterrputT = 28
	ETS_SPI1_INTR_SOURCE           PeriphInterrputT = 29
	ETS_SPI2_INTR_SOURCE           PeriphInterrputT = 30
	ETS_SPI3_INTR_SOURCE           PeriphInterrputT = 31
	ETS_I2S0_INTR_SOURCE           PeriphInterrputT = 32
	ETS_I2S1_INTR_SOURCE           PeriphInterrputT = 33
	ETS_UART0_INTR_SOURCE          PeriphInterrputT = 34
	ETS_UART1_INTR_SOURCE          PeriphInterrputT = 35
	ETS_UART2_INTR_SOURCE          PeriphInterrputT = 36
	ETS_SDIO_HOST_INTR_SOURCE      PeriphInterrputT = 37
	ETS_ETH_MAC_INTR_SOURCE        PeriphInterrputT = 38
	ETS_PWM0_INTR_SOURCE           PeriphInterrputT = 39
	ETS_PWM1_INTR_SOURCE           PeriphInterrputT = 40
	ETS_LEDC_INTR_SOURCE           PeriphInterrputT = 43
	ETS_EFUSE_INTR_SOURCE          PeriphInterrputT = 44
	ETS_TWAI_INTR_SOURCE           PeriphInterrputT = 45
	ETS_RTC_CORE_INTR_SOURCE       PeriphInterrputT = 46
	ETS_RMT_INTR_SOURCE            PeriphInterrputT = 47
	ETS_PCNT_INTR_SOURCE           PeriphInterrputT = 48
	ETS_I2C_EXT0_INTR_SOURCE       PeriphInterrputT = 49
	ETS_I2C_EXT1_INTR_SOURCE       PeriphInterrputT = 50
	ETS_RSA_INTR_SOURCE            PeriphInterrputT = 51
	ETS_SPI1_DMA_INTR_SOURCE       PeriphInterrputT = 52
	ETS_SPI2_DMA_INTR_SOURCE       PeriphInterrputT = 53
	ETS_SPI3_DMA_INTR_SOURCE       PeriphInterrputT = 54
	ETS_WDT_INTR_SOURCE            PeriphInterrputT = 55
	ETS_TIMER1_INTR_SOURCE         PeriphInterrputT = 56
	ETS_TIMER2_INTR_SOURCE         PeriphInterrputT = 57
	ETS_TG0_T0_EDGE_INTR_SOURCE    PeriphInterrputT = 58
	ETS_TG0_T1_EDGE_INTR_SOURCE    PeriphInterrputT = 59
	ETS_TG0_WDT_EDGE_INTR_SOURCE   PeriphInterrputT = 60
	ETS_TG0_LACT_EDGE_INTR_SOURCE  PeriphInterrputT = 61
	ETS_TG1_T0_EDGE_INTR_SOURCE    PeriphInterrputT = 62
	ETS_TG1_T1_EDGE_INTR_SOURCE    PeriphInterrputT = 63
	ETS_TG1_WDT_EDGE_INTR_SOURCE   PeriphInterrputT = 64
	ETS_TG1_LACT_EDGE_INTR_SOURCE  PeriphInterrputT = 65
	ETS_MMU_IA_INTR_SOURCE         PeriphInterrputT = 66
	ETS_MPU_IA_INTR_SOURCE         PeriphInterrputT = 67
	ETS_CACHE_IA_INTR_SOURCE       PeriphInterrputT = 68
	ETS_MAX_INTR_SOURCE            PeriphInterrputT = 69
)

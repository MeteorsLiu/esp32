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
	ETS_LP_TIMER_INTR_SOURCE              PeriphInterrputT = 7
	ETS_COEX_INTR_SOURCE                  PeriphInterrputT = 8
	ETS_BLE_TIMER_INTR_SOURCE             PeriphInterrputT = 9
	ETS_BLE_SEC_INTR_SOURCE               PeriphInterrputT = 10
	ETS_I2C_MASTER_SOURCE                 PeriphInterrputT = 11
	ETS_ZB_MAC_SOURCE                     PeriphInterrputT = 12
	ETS_PMU_INTR_SOURCE                   PeriphInterrputT = 13
	ETS_EFUSE_INTR_SOURCE                 PeriphInterrputT = 14
	ETS_LP_RTC_TIMER_INTR_SOURCE          PeriphInterrputT = 15
	ETS_LP_UART_INTR_SOURCE               PeriphInterrputT = 16
	ETS_LP_I2C_INTR_SOURCE                PeriphInterrputT = 17
	ETS_LP_WDT_INTR_SOURCE                PeriphInterrputT = 18
	ETS_LP_PERI_TIMEOUT_INTR_SOURCE       PeriphInterrputT = 19
	ETS_LP_APM_M0_INTR_SOURCE             PeriphInterrputT = 20
	ETS_LP_APM_M1_INTR_SOURCE             PeriphInterrputT = 21
	ETS_FROM_CPU_INTR0_SOURCE             PeriphInterrputT = 22
	ETS_FROM_CPU_INTR1_SOURCE             PeriphInterrputT = 23
	ETS_FROM_CPU_INTR2_SOURCE             PeriphInterrputT = 24
	ETS_FROM_CPU_INTR3_SOURCE             PeriphInterrputT = 25
	ETS_ASSIST_DEBUG_INTR_SOURCE          PeriphInterrputT = 26
	ETS_TRACE_INTR_SOURCE                 PeriphInterrputT = 27
	ETS_CACHE_INTR_SOURCE                 PeriphInterrputT = 28
	ETS_CPU_PERI_TIMEOUT_INTR_SOURCE      PeriphInterrputT = 29
	ETS_GPIO_INTR_SOURCE                  PeriphInterrputT = 30
	ETS_GPIO_NMI_SOURCE                   PeriphInterrputT = 31
	ETS_PAU_INTR_SOURCE                   PeriphInterrputT = 32
	ETS_HP_PERI_TIMEOUT_INTR_SOURCE       PeriphInterrputT = 33
	ETS_MODEM_PERI_TIMEOUT_INTR_SOURCE    PeriphInterrputT = 34
	ETS_HP_APM_M0_INTR_SOURCE             PeriphInterrputT = 35
	ETS_HP_APM_M1_INTR_SOURCE             PeriphInterrputT = 36
	ETS_HP_APM_M2_INTR_SOURCE             PeriphInterrputT = 37
	ETS_HP_APM_M3_INTR_SOURCE             PeriphInterrputT = 38
	ETS_LP_APM0_INTR_SOURCE               PeriphInterrputT = 39
	ETS_MSPI_INTR_SOURCE                  PeriphInterrputT = 40
	ETS_I2S0_INTR_SOURCE                  PeriphInterrputT = 41
	ETS_UHCI0_INTR_SOURCE                 PeriphInterrputT = 42
	ETS_UART0_INTR_SOURCE                 PeriphInterrputT = 43
	ETS_UART1_INTR_SOURCE                 PeriphInterrputT = 44
	ETS_LEDC_INTR_SOURCE                  PeriphInterrputT = 45
	ETS_TWAI0_INTR_SOURCE                 PeriphInterrputT = 46
	ETS_TWAI1_INTR_SOURCE                 PeriphInterrputT = 47
	ETS_USB_SERIAL_JTAG_INTR_SOURCE       PeriphInterrputT = 48
	ETS_RMT_INTR_SOURCE                   PeriphInterrputT = 49
	ETS_I2C_EXT0_INTR_SOURCE              PeriphInterrputT = 50
	ETS_TG0_T0_LEVEL_INTR_SOURCE          PeriphInterrputT = 51
	ETS_TG0_T1_LEVEL_INTR_SOURCE          PeriphInterrputT = 52
	ETS_TG0_WDT_LEVEL_INTR_SOURCE         PeriphInterrputT = 53
	ETS_TG1_T0_LEVEL_INTR_SOURCE          PeriphInterrputT = 54
	ETS_TG1_T1_LEVEL_INTR_SOURCE          PeriphInterrputT = 55
	ETS_TG1_WDT_LEVEL_INTR_SOURCE         PeriphInterrputT = 56
	ETS_SYSTIMER_TARGET0_INTR_SOURCE      PeriphInterrputT = 57
	ETS_SYSTIMER_TARGET1_INTR_SOURCE      PeriphInterrputT = 58
	ETS_SYSTIMER_TARGET2_INTR_SOURCE      PeriphInterrputT = 59
	ETS_SYSTIMER_TARGET0_EDGE_INTR_SOURCE PeriphInterrputT = 57
	ETS_SYSTIMER_TARGET1_EDGE_INTR_SOURCE PeriphInterrputT = 58
	ETS_SYSTIMER_TARGET2_EDGE_INTR_SOURCE PeriphInterrputT = 59
	ETS_APB_ADC_INTR_SOURCE               PeriphInterrputT = 60
	ETS_TEMPERATURE_SENSOR_INTR_SOURCE    PeriphInterrputT = 60
	ETS_MCPWM0_INTR_SOURCE                PeriphInterrputT = 61
	ETS_PCNT_INTR_SOURCE                  PeriphInterrputT = 62
	ETS_PARL_IO_INTR_SOURCE               PeriphInterrputT = 63
	ETS_SLC0_INTR_SOURCE                  PeriphInterrputT = 64
	ETS_SLC_INTR_SOURCE                   PeriphInterrputT = 65
	ETS_DMA_IN_CH0_INTR_SOURCE            PeriphInterrputT = 66
	ETS_DMA_IN_CH1_INTR_SOURCE            PeriphInterrputT = 67
	ETS_DMA_IN_CH2_INTR_SOURCE            PeriphInterrputT = 68
	ETS_DMA_OUT_CH0_INTR_SOURCE           PeriphInterrputT = 69
	ETS_DMA_OUT_CH1_INTR_SOURCE           PeriphInterrputT = 70
	ETS_DMA_OUT_CH2_INTR_SOURCE           PeriphInterrputT = 71
	ETS_GSPI2_INTR_SOURCE                 PeriphInterrputT = 72
	ETS_AES_INTR_SOURCE                   PeriphInterrputT = 73
	ETS_SHA_INTR_SOURCE                   PeriphInterrputT = 74
	ETS_RSA_INTR_SOURCE                   PeriphInterrputT = 75
	ETS_ECC_INTR_SOURCE                   PeriphInterrputT = 76
	ETS_MAX_INTR_SOURCE                   PeriphInterrputT = 77
)

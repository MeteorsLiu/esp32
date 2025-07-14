package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PeriphRetentionModule c.Int

const (
	SLEEP_RETENTION_MODULE_MIN          PeriphRetentionModule = 0
	SLEEP_RETENTION_MODULE_NULL         PeriphRetentionModule = 0
	SLEEP_RETENTION_MODULE_CLOCK_SYSTEM PeriphRetentionModule = 1
	SLEEP_RETENTION_MODULE_CLOCK_MODEM  PeriphRetentionModule = 2
	SLEEP_RETENTION_MODULE_SYS_PERIPH   PeriphRetentionModule = 3
	SLEEP_RETENTION_MODULE_TG0_WDT      PeriphRetentionModule = 4
	SLEEP_RETENTION_MODULE_TG1_WDT      PeriphRetentionModule = 5
	SLEEP_RETENTION_MODULE_TG0_TIMER0   PeriphRetentionModule = 6
	SLEEP_RETENTION_MODULE_TG1_TIMER0   PeriphRetentionModule = 7
	SLEEP_RETENTION_MODULE_GDMA_CH0     PeriphRetentionModule = 8
	SLEEP_RETENTION_MODULE_GDMA_CH1     PeriphRetentionModule = 9
	SLEEP_RETENTION_MODULE_GDMA_CH2     PeriphRetentionModule = 10
	SLEEP_RETENTION_MODULE_ADC          PeriphRetentionModule = 11
	SLEEP_RETENTION_MODULE_I2C0         PeriphRetentionModule = 12
	SLEEP_RETENTION_MODULE_I2C1         PeriphRetentionModule = 13
	SLEEP_RETENTION_MODULE_RMT0         PeriphRetentionModule = 14
	SLEEP_RETENTION_MODULE_UART0        PeriphRetentionModule = 15
	SLEEP_RETENTION_MODULE_UART1        PeriphRetentionModule = 16
	SLEEP_RETENTION_MODULE_I2S0         PeriphRetentionModule = 17
	SLEEP_RETENTION_MODULE_ETM0         PeriphRetentionModule = 18
	SLEEP_RETENTION_MODULE_TEMP_SENSOR  PeriphRetentionModule = 19
	SLEEP_RETENTION_MODULE_TWAI0        PeriphRetentionModule = 20
	SLEEP_RETENTION_MODULE_PARLIO0      PeriphRetentionModule = 21
	SLEEP_RETENTION_MODULE_GPSPI2       PeriphRetentionModule = 22
	SLEEP_RETENTION_MODULE_LEDC         PeriphRetentionModule = 23
	SLEEP_RETENTION_MODULE_PCNT0        PeriphRetentionModule = 24
	SLEEP_RETENTION_MODULE_MCPWM0       PeriphRetentionModule = 25
	SLEEP_RETENTION_MODULE_BLE_MAC      PeriphRetentionModule = 28
	SLEEP_RETENTION_MODULE_BT_BB        PeriphRetentionModule = 29
	SLEEP_RETENTION_MODULE_802154_MAC   PeriphRetentionModule = 30
	SLEEP_RETENTION_MODULE_MAX          PeriphRetentionModule = 31
)

type PeriphRetentionModuleT PeriphRetentionModule

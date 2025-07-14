package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PeriphModuleT c.Int

const (
	PERIPH_LEDC_MODULE           PeriphModuleT = 0
	PERIPH_UART0_MODULE          PeriphModuleT = 1
	PERIPH_UART1_MODULE          PeriphModuleT = 2
	PERIPH_I2C0_MODULE           PeriphModuleT = 3
	PERIPH_TIMG0_MODULE          PeriphModuleT = 4
	PERIPH_TIMG1_MODULE          PeriphModuleT = 5
	PERIPH_UHCI0_MODULE          PeriphModuleT = 6
	PERIPH_SPI_MODULE            PeriphModuleT = 7
	PERIPH_SPI2_MODULE           PeriphModuleT = 8
	PERIPH_RNG_MODULE            PeriphModuleT = 9
	PERIPH_WIFI_MODULE           PeriphModuleT = 10
	PERIPH_BT_MODULE             PeriphModuleT = 11
	PERIPH_WIFI_BT_COMMON_MODULE PeriphModuleT = 12
	PERIPH_BT_BASEBAND_MODULE    PeriphModuleT = 13
	PERIPH_BT_LC_MODULE          PeriphModuleT = 14
	PERIPH_AES_MODULE            PeriphModuleT = 15
	PERIPH_SHA_MODULE            PeriphModuleT = 16
	PERIPH_ECC_MODULE            PeriphModuleT = 17
	PERIPH_GDMA_MODULE           PeriphModuleT = 18
	PERIPH_SYSTIMER_MODULE       PeriphModuleT = 19
	PERIPH_SARADC_MODULE         PeriphModuleT = 20
	PERIPH_TEMPSENSOR_MODULE     PeriphModuleT = 21
	PERIPH_MODEM_RPA_MODULE      PeriphModuleT = 22
	PERIPH_ASSIST_DEBUG_MODULE   PeriphModuleT = 23
	PERIPH_MODULE_MAX            PeriphModuleT = 24
)

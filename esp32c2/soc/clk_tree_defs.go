package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SOC_CLK_RC_FAST_FREQ_APPROX = 17500000
const SOC_CLK_RC_SLOW_FREQ_APPROX = 136000
const SOC_CLK_OSC_SLOW_FREQ_APPROX = 32768

type SocRootClkT c.Int

const (
	SOC_ROOT_CLK_INT_RC_FAST  SocRootClkT = 0
	SOC_ROOT_CLK_INT_RC_SLOW  SocRootClkT = 1
	SOC_ROOT_CLK_EXT_XTAL     SocRootClkT = 2
	SOC_ROOT_CLK_EXT_OSC_SLOW SocRootClkT = 3
)

type SocCpuClkSrcT c.Int

const (
	SOC_CPU_CLK_SRC_XTAL    SocCpuClkSrcT = 0
	SOC_CPU_CLK_SRC_PLL     SocCpuClkSrcT = 1
	SOC_CPU_CLK_SRC_RC_FAST SocCpuClkSrcT = 2
	SOC_CPU_CLK_SRC_INVALID SocCpuClkSrcT = 3
)

type SocRtcSlowClkSrcT c.Int

const (
	SOC_RTC_SLOW_CLK_SRC_RC_SLOW      SocRtcSlowClkSrcT = 0
	SOC_RTC_SLOW_CLK_SRC_OSC_SLOW     SocRtcSlowClkSrcT = 1
	SOC_RTC_SLOW_CLK_SRC_RC_FAST_D256 SocRtcSlowClkSrcT = 2
	SOC_RTC_SLOW_CLK_SRC_INVALID      SocRtcSlowClkSrcT = 3
)

type SocRtcFastClkSrcT c.Int

const (
	SOC_RTC_FAST_CLK_SRC_XTAL_D2  SocRtcFastClkSrcT = 0
	SOC_RTC_FAST_CLK_SRC_XTAL_DIV SocRtcFastClkSrcT = 0
	SOC_RTC_FAST_CLK_SRC_RC_FAST  SocRtcFastClkSrcT = 1
	SOC_RTC_FAST_CLK_SRC_INVALID  SocRtcFastClkSrcT = 2
)

type SocXtalFreqT c.Int

const (
	SOC_XTAL_FREQ_26M SocXtalFreqT = 26
	SOC_XTAL_FREQ_32M SocXtalFreqT = 32
	SOC_XTAL_FREQ_40M SocXtalFreqT = 40
)

type SocModuleClkT c.Int

const (
	SOC_MOD_CLK_CPU          SocModuleClkT = 1
	SOC_MOD_CLK_RTC_FAST     SocModuleClkT = 2
	SOC_MOD_CLK_RTC_SLOW     SocModuleClkT = 3
	SOC_MOD_CLK_APB          SocModuleClkT = 4
	SOC_MOD_CLK_PLL_F40M     SocModuleClkT = 5
	SOC_MOD_CLK_PLL_F60M     SocModuleClkT = 6
	SOC_MOD_CLK_PLL_F80M     SocModuleClkT = 7
	SOC_MOD_CLK_OSC_SLOW     SocModuleClkT = 8
	SOC_MOD_CLK_RC_FAST      SocModuleClkT = 9
	SOC_MOD_CLK_RC_FAST_D256 SocModuleClkT = 10
	SOC_MOD_CLK_XTAL         SocModuleClkT = 11
	SOC_MOD_CLK_INVALID      SocModuleClkT = 12
)

type SocPeriphSystimerClkSrcT c.Int

const (
	SYSTIMER_CLK_SRC_XTAL    SocPeriphSystimerClkSrcT = 11
	SYSTIMER_CLK_SRC_DEFAULT SocPeriphSystimerClkSrcT = 11
)

type SocPeriphGptimerClkSrcT c.Int

const (
	GPTIMER_CLK_SRC_PLL_F40M SocPeriphGptimerClkSrcT = 5
	GPTIMER_CLK_SRC_XTAL     SocPeriphGptimerClkSrcT = 11
	GPTIMER_CLK_SRC_DEFAULT  SocPeriphGptimerClkSrcT = 5
)

type SocPeriphTgClkSrcLegacyT c.Int

const (
	TIMER_SRC_CLK_PLL_F40M SocPeriphTgClkSrcLegacyT = 5
	TIMER_SRC_CLK_XTAL     SocPeriphTgClkSrcLegacyT = 11
	TIMER_SRC_CLK_DEFAULT  SocPeriphTgClkSrcLegacyT = 5
)

type SocPeriphTemperatureSensorClkSrcT c.Int

const (
	TEMPERATURE_SENSOR_CLK_SRC_XTAL    SocPeriphTemperatureSensorClkSrcT = 11
	TEMPERATURE_SENSOR_CLK_SRC_RC_FAST SocPeriphTemperatureSensorClkSrcT = 9
	TEMPERATURE_SENSOR_CLK_SRC_DEFAULT SocPeriphTemperatureSensorClkSrcT = 11
)

type SocPeriphUartClkSrcLegacyT c.Int

const (
	UART_SCLK_PLL_F40M SocPeriphUartClkSrcLegacyT = 5
	UART_SCLK_RTC      SocPeriphUartClkSrcLegacyT = 9
	UART_SCLK_XTAL     SocPeriphUartClkSrcLegacyT = 11
	UART_SCLK_DEFAULT  SocPeriphUartClkSrcLegacyT = 5
)

type SocPeriphSpiClkSrcT c.Int

const (
	SPI_CLK_SRC_DEFAULT  SocPeriphSpiClkSrcT = 5
	SPI_CLK_SRC_PLL_F40M SocPeriphSpiClkSrcT = 5
	SPI_CLK_SRC_XTAL     SocPeriphSpiClkSrcT = 11
)

type SocPeriphI2cClkSrcT c.Int

const (
	I2C_CLK_SRC_XTAL    SocPeriphI2cClkSrcT = 11
	I2C_CLK_SRC_RC_FAST SocPeriphI2cClkSrcT = 9
	I2C_CLK_SRC_DEFAULT SocPeriphI2cClkSrcT = 11
)

type SocPeriphAdcDigiClkSrcT c.Int

const (
	ADC_DIGI_CLK_SRC_XTAL     SocPeriphAdcDigiClkSrcT = 11
	ADC_DIGI_CLK_SRC_PLL_F80M SocPeriphAdcDigiClkSrcT = 7
	ADC_DIGI_CLK_SRC_DEFAULT  SocPeriphAdcDigiClkSrcT = 7
)

type SocPeriphGlitchFilterClkSrcT c.Int

const (
	GLITCH_FILTER_CLK_SRC_APB     SocPeriphGlitchFilterClkSrcT = 4
	GLITCH_FILTER_CLK_SRC_DEFAULT SocPeriphGlitchFilterClkSrcT = 4
)

type SocPeriphMwdtClkSrcT c.Int

const (
	MWDT_CLK_SRC_XTAL     SocPeriphMwdtClkSrcT = 11
	MWDT_CLK_SRC_PLL_F40M SocPeriphMwdtClkSrcT = 5
	MWDT_CLK_SRC_DEFAULT  SocPeriphMwdtClkSrcT = 5
)

type SocPeriphLedcClkSrcLegacyT c.Int

const (
	LEDC_AUTO_CLK        SocPeriphLedcClkSrcLegacyT = 0
	LEDC_USE_PLL_DIV_CLK SocPeriphLedcClkSrcLegacyT = 6
	LEDC_USE_RC_FAST_CLK SocPeriphLedcClkSrcLegacyT = 9
	LEDC_USE_XTAL_CLK    SocPeriphLedcClkSrcLegacyT = 11
	LEDC_USE_RTC8M_CLK   SocPeriphLedcClkSrcLegacyT = 9
)

type SocClkoutSigIdT c.Int

const (
	CLKOUT_SIG_PLL      SocClkoutSigIdT = 1
	CLKOUT_SIG_RC_SLOW  SocClkoutSigIdT = 4
	CLKOUT_SIG_XTAL     SocClkoutSigIdT = 5
	CLKOUT_SIG_PLL_F80M SocClkoutSigIdT = 13
	CLKOUT_SIG_RC_FAST  SocClkoutSigIdT = 14
	CLKOUT_SIG_INVALID  SocClkoutSigIdT = 255
)

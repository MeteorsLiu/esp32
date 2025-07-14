package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SOC_CLK_RC_FAST_FREQ_APPROX = 17500000
const SOC_CLK_RC_SLOW_FREQ_APPROX = 136000
const SOC_CLK_RC32K_FREQ_APPROX = 32768
const SOC_CLK_XTAL32K_FREQ_APPROX = 32768
const SOC_CLK_OSC_SLOW_FREQ_APPROX = 32768

type SocRootClkT c.Int

const (
	SOC_ROOT_CLK_INT_RC_FAST  SocRootClkT = 0
	SOC_ROOT_CLK_INT_RC_SLOW  SocRootClkT = 1
	SOC_ROOT_CLK_EXT_XTAL     SocRootClkT = 2
	SOC_ROOT_CLK_EXT_XTAL32K  SocRootClkT = 3
	SOC_ROOT_CLK_INT_RC32K    SocRootClkT = 4
	SOC_ROOT_CLK_EXT_OSC_SLOW SocRootClkT = 5
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
	SOC_RTC_SLOW_CLK_SRC_RC_SLOW  SocRtcSlowClkSrcT = 0
	SOC_RTC_SLOW_CLK_SRC_XTAL32K  SocRtcSlowClkSrcT = 1
	SOC_RTC_SLOW_CLK_SRC_RC32K    SocRtcSlowClkSrcT = 2
	SOC_RTC_SLOW_CLK_SRC_OSC_SLOW SocRtcSlowClkSrcT = 3
	SOC_RTC_SLOW_CLK_SRC_INVALID  SocRtcSlowClkSrcT = 4
)

type SocRtcFastClkSrcT c.Int

const (
	SOC_RTC_FAST_CLK_SRC_RC_FAST  SocRtcFastClkSrcT = 0
	SOC_RTC_FAST_CLK_SRC_XTAL_D2  SocRtcFastClkSrcT = 1
	SOC_RTC_FAST_CLK_SRC_XTAL_DIV SocRtcFastClkSrcT = 1
	SOC_RTC_FAST_CLK_SRC_INVALID  SocRtcFastClkSrcT = 2
)

type SocXtalFreqT c.Int

const SOC_XTAL_FREQ_40M SocXtalFreqT = 40

type SocModuleClkT c.Int

const (
	SOC_MOD_CLK_CPU       SocModuleClkT = 1
	SOC_MOD_CLK_RTC_FAST  SocModuleClkT = 2
	SOC_MOD_CLK_RTC_SLOW  SocModuleClkT = 3
	SOC_MOD_CLK_PLL_F80M  SocModuleClkT = 4
	SOC_MOD_CLK_PLL_F160M SocModuleClkT = 5
	SOC_MOD_CLK_PLL_F240M SocModuleClkT = 6
	SOC_MOD_CLK_XTAL32K   SocModuleClkT = 7
	SOC_MOD_CLK_RC_FAST   SocModuleClkT = 8
	SOC_MOD_CLK_XTAL      SocModuleClkT = 9
	SOC_MOD_CLK_XTAL_D2   SocModuleClkT = 10
	SOC_MOD_CLK_INVALID   SocModuleClkT = 11
)

type SocPeriphSystimerClkSrcT c.Int

const (
	SYSTIMER_CLK_SRC_XTAL    SocPeriphSystimerClkSrcT = 9
	SYSTIMER_CLK_SRC_RC_FAST SocPeriphSystimerClkSrcT = 8
	SYSTIMER_CLK_SRC_DEFAULT SocPeriphSystimerClkSrcT = 9
)

type SocPeriphGptimerClkSrcT c.Int

const (
	GPTIMER_CLK_SRC_PLL_F80M SocPeriphGptimerClkSrcT = 4
	GPTIMER_CLK_SRC_RC_FAST  SocPeriphGptimerClkSrcT = 8
	GPTIMER_CLK_SRC_XTAL     SocPeriphGptimerClkSrcT = 9
	GPTIMER_CLK_SRC_DEFAULT  SocPeriphGptimerClkSrcT = 4
)

type SocPeriphTgClkSrcLegacyT c.Int

const (
	TIMER_SRC_CLK_PLL_F80M SocPeriphTgClkSrcLegacyT = 4
	TIMER_SRC_CLK_XTAL     SocPeriphTgClkSrcLegacyT = 9
	TIMER_SRC_CLK_DEFAULT  SocPeriphTgClkSrcLegacyT = 4
)

type SocPeriphRmtClkSrcT c.Int

const (
	RMT_CLK_SRC_PLL_F80M SocPeriphRmtClkSrcT = 4
	RMT_CLK_SRC_RC_FAST  SocPeriphRmtClkSrcT = 8
	RMT_CLK_SRC_XTAL     SocPeriphRmtClkSrcT = 9
	RMT_CLK_SRC_DEFAULT  SocPeriphRmtClkSrcT = 4
)

type SocPeriphRmtClkSrcLegacyT c.Int

const (
	RMT_BASECLK_PLL_F80M SocPeriphRmtClkSrcLegacyT = 4
	RMT_BASECLK_XTAL     SocPeriphRmtClkSrcLegacyT = 9
	RMT_BASECLK_DEFAULT  SocPeriphRmtClkSrcLegacyT = 4
)

type SocPeriphTemperatureSensorClkSrcT c.Int

const (
	TEMPERATURE_SENSOR_CLK_SRC_XTAL    SocPeriphTemperatureSensorClkSrcT = 9
	TEMPERATURE_SENSOR_CLK_SRC_RC_FAST SocPeriphTemperatureSensorClkSrcT = 8
	TEMPERATURE_SENSOR_CLK_SRC_DEFAULT SocPeriphTemperatureSensorClkSrcT = 9
)

type SocPeriphUartClkSrcLegacyT c.Int

const (
	UART_SCLK_PLL_F80M SocPeriphUartClkSrcLegacyT = 4
	UART_SCLK_RTC      SocPeriphUartClkSrcLegacyT = 8
	UART_SCLK_XTAL     SocPeriphUartClkSrcLegacyT = 9
	UART_SCLK_DEFAULT  SocPeriphUartClkSrcLegacyT = 4
)

type SocPeriphLpUartClkSrcT c.Int

const (
	LP_UART_SCLK_LP_FAST SocPeriphLpUartClkSrcT = 2
	LP_UART_SCLK_XTAL_D2 SocPeriphLpUartClkSrcT = 10
	LP_UART_SCLK_DEFAULT SocPeriphLpUartClkSrcT = 2
)

type SocPeriphMcpwmTimerClkSrcT c.Int

const (
	MCPWM_TIMER_CLK_SRC_PLL160M SocPeriphMcpwmTimerClkSrcT = 5
	MCPWM_TIMER_CLK_SRC_XTAL    SocPeriphMcpwmTimerClkSrcT = 9
	MCPWM_TIMER_CLK_SRC_DEFAULT SocPeriphMcpwmTimerClkSrcT = 5
)

type SocPeriphMcpwmCaptureClkSrcT c.Int

const (
	MCPWM_CAPTURE_CLK_SRC_PLL160M SocPeriphMcpwmCaptureClkSrcT = 5
	MCPWM_CAPTURE_CLK_SRC_XTAL    SocPeriphMcpwmCaptureClkSrcT = 9
	MCPWM_CAPTURE_CLK_SRC_DEFAULT SocPeriphMcpwmCaptureClkSrcT = 5
)

type SocPeriphMcpwmCarrierClkSrcT c.Int

const (
	MCPWM_CARRIER_CLK_SRC_PLL160M SocPeriphMcpwmCarrierClkSrcT = 5
	MCPWM_CARRIER_CLK_SRC_XTAL    SocPeriphMcpwmCarrierClkSrcT = 9
	MCPWM_CARRIER_CLK_SRC_DEFAULT SocPeriphMcpwmCarrierClkSrcT = 5
)

type SocPeriphI2sClkSrcT c.Int

const (
	I2S_CLK_SRC_DEFAULT  SocPeriphI2sClkSrcT = 5
	I2S_CLK_SRC_PLL_160M SocPeriphI2sClkSrcT = 5
	I2S_CLK_SRC_XTAL     SocPeriphI2sClkSrcT = 9
	I2S_CLK_SRC_EXTERNAL SocPeriphI2sClkSrcT = -1
)

type SocPeriphI2cClkSrcT c.Int

const (
	I2C_CLK_SRC_XTAL    SocPeriphI2cClkSrcT = 9
	I2C_CLK_SRC_RC_FAST SocPeriphI2cClkSrcT = 8
	I2C_CLK_SRC_DEFAULT SocPeriphI2cClkSrcT = 9
)

type SocPeriphLpI2cClkSrcT c.Int

const (
	LP_I2C_SCLK_LP_FAST SocPeriphLpI2cClkSrcT = 2
	LP_I2C_SCLK_XTAL_D2 SocPeriphLpI2cClkSrcT = 10
	LP_I2C_SCLK_DEFAULT SocPeriphLpI2cClkSrcT = 2
)

type SocPeriphSpiClkSrcT c.Int

const (
	SPI_CLK_SRC_DEFAULT  SocPeriphSpiClkSrcT = 4
	SPI_CLK_SRC_PLL_F80M SocPeriphSpiClkSrcT = 4
	SPI_CLK_SRC_XTAL     SocPeriphSpiClkSrcT = 9
	SPI_CLK_SRC_RC_FAST  SocPeriphSpiClkSrcT = 8
)

type SocPeriphSdmClkSrcT c.Int

const (
	SDM_CLK_SRC_XTAL     SocPeriphSdmClkSrcT = 9
	SDM_CLK_SRC_PLL_F80M SocPeriphSdmClkSrcT = 4
	SDM_CLK_SRC_DEFAULT  SocPeriphSdmClkSrcT = 4
)

type SocPeriphGlitchFilterClkSrcT c.Int

const (
	GLITCH_FILTER_CLK_SRC_XTAL     SocPeriphGlitchFilterClkSrcT = 9
	GLITCH_FILTER_CLK_SRC_PLL_F80M SocPeriphGlitchFilterClkSrcT = 4
	GLITCH_FILTER_CLK_SRC_DEFAULT  SocPeriphGlitchFilterClkSrcT = 4
)

type SocPeriphTwaiClkSrcT c.Int

const (
	TWAI_CLK_SRC_XTAL    SocPeriphTwaiClkSrcT = 9
	TWAI_CLK_SRC_DEFAULT SocPeriphTwaiClkSrcT = 9
)

type SocPeriphAdcDigiClkSrcT c.Int

const (
	ADC_DIGI_CLK_SRC_XTAL     SocPeriphAdcDigiClkSrcT = 9
	ADC_DIGI_CLK_SRC_PLL_F80M SocPeriphAdcDigiClkSrcT = 4
	ADC_DIGI_CLK_SRC_RC_FAST  SocPeriphAdcDigiClkSrcT = 8
	ADC_DIGI_CLK_SRC_DEFAULT  SocPeriphAdcDigiClkSrcT = 4
)

type SocPeriphMwdtClkSrcT c.Int

const (
	MWDT_CLK_SRC_XTAL     SocPeriphMwdtClkSrcT = 9
	MWDT_CLK_SRC_PLL_F80M SocPeriphMwdtClkSrcT = 4
	MWDT_CLK_SRC_RC_FAST  SocPeriphMwdtClkSrcT = 8
	MWDT_CLK_SRC_DEFAULT  SocPeriphMwdtClkSrcT = 9
)

type SocPeriphLedcClkSrcLegacyT c.Int

const (
	LEDC_AUTO_CLK        SocPeriphLedcClkSrcLegacyT = 0
	LEDC_USE_PLL_DIV_CLK SocPeriphLedcClkSrcLegacyT = 4
	LEDC_USE_RC_FAST_CLK SocPeriphLedcClkSrcLegacyT = 8
	LEDC_USE_XTAL_CLK    SocPeriphLedcClkSrcLegacyT = 9
	LEDC_USE_RTC8M_CLK   SocPeriphLedcClkSrcLegacyT = 8
)

type SocPeriphParlioClkSrcT c.Int

const (
	PARLIO_CLK_SRC_XTAL      SocPeriphParlioClkSrcT = 9
	PARLIO_CLK_SRC_PLL_F240M SocPeriphParlioClkSrcT = 6
	PARLIO_CLK_SRC_RC_FAST   SocPeriphParlioClkSrcT = 8
	PARLIO_CLK_SRC_EXTERNAL  SocPeriphParlioClkSrcT = -1
	PARLIO_CLK_SRC_DEFAULT   SocPeriphParlioClkSrcT = 6
)

type SocClkoutSigIdT c.Int

const (
	CLKOUT_SIG_PLL      SocClkoutSigIdT = 1
	CLKOUT_SIG_XTAL     SocClkoutSigIdT = 5
	CLKOUT_SIG_PLL_F80M SocClkoutSigIdT = 13
	CLKOUT_SIG_CPU      SocClkoutSigIdT = 16
	CLKOUT_SIG_AHB      SocClkoutSigIdT = 17
	CLKOUT_SIG_APB      SocClkoutSigIdT = 18
	CLKOUT_SIG_XTAL32K  SocClkoutSigIdT = 21
	CLKOUT_SIG_EXT32K   SocClkoutSigIdT = 22
	CLKOUT_SIG_RC_FAST  SocClkoutSigIdT = 23
	CLKOUT_SIG_RC_32K   SocClkoutSigIdT = 24
	CLKOUT_SIG_RC_SLOW  SocClkoutSigIdT = 25
	CLKOUT_SIG_INVALID  SocClkoutSigIdT = 255
)

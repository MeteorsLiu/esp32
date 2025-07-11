package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SOC_CLK_RC_FAST_FREQ_APPROX = 8500000
const SOC_CLK_RC_SLOW_FREQ_APPROX = 90000
const SOC_CLK_XTAL32K_FREQ_APPROX = 32768

type SocRootClkT c.Int

const (
	SOC_ROOT_CLK_INT_RC_FAST SocRootClkT = 0
	SOC_ROOT_CLK_INT_RC_SLOW SocRootClkT = 1
	SOC_ROOT_CLK_EXT_XTAL    SocRootClkT = 2
	SOC_ROOT_CLK_EXT_XTAL32K SocRootClkT = 3
)

type SocCpuClkSrcT c.Int

const (
	SOC_CPU_CLK_SRC_XTAL    SocCpuClkSrcT = 0
	SOC_CPU_CLK_SRC_PLL     SocCpuClkSrcT = 1
	SOC_CPU_CLK_SRC_RC_FAST SocCpuClkSrcT = 2
	SOC_CPU_CLK_SRC_APLL    SocCpuClkSrcT = 3
	SOC_CPU_CLK_SRC_INVALID SocCpuClkSrcT = 4
)

type SocRtcSlowClkSrcT c.Int

const (
	SOC_RTC_SLOW_CLK_SRC_RC_SLOW      SocRtcSlowClkSrcT = 0
	SOC_RTC_SLOW_CLK_SRC_XTAL32K      SocRtcSlowClkSrcT = 1
	SOC_RTC_SLOW_CLK_SRC_RC_FAST_D256 SocRtcSlowClkSrcT = 2
	SOC_RTC_SLOW_CLK_SRC_INVALID      SocRtcSlowClkSrcT = 3
)

type SocRtcFastClkSrcT c.Int

const (
	SOC_RTC_FAST_CLK_SRC_XTAL_D4  SocRtcFastClkSrcT = 0
	SOC_RTC_FAST_CLK_SRC_XTAL_DIV SocRtcFastClkSrcT = 0
	SOC_RTC_FAST_CLK_SRC_RC_FAST  SocRtcFastClkSrcT = 1
	SOC_RTC_FAST_CLK_SRC_INVALID  SocRtcFastClkSrcT = 2
)

type SocXtalFreqT c.Int

const SOC_XTAL_FREQ_40M SocXtalFreqT = 40

type SocModuleClkT c.Int

const (
	SOC_MOD_CLK_CPU          SocModuleClkT = 1
	SOC_MOD_CLK_RTC_FAST     SocModuleClkT = 2
	SOC_MOD_CLK_RTC_SLOW     SocModuleClkT = 3
	SOC_MOD_CLK_APB          SocModuleClkT = 4
	SOC_MOD_CLK_PLL_F160M    SocModuleClkT = 5
	SOC_MOD_CLK_XTAL32K      SocModuleClkT = 6
	SOC_MOD_CLK_RC_FAST      SocModuleClkT = 7
	SOC_MOD_CLK_RC_FAST_D256 SocModuleClkT = 8
	SOC_MOD_CLK_XTAL         SocModuleClkT = 9
	SOC_MOD_CLK_REF_TICK     SocModuleClkT = 10
	SOC_MOD_CLK_APLL         SocModuleClkT = 11
	SOC_MOD_CLK_TEMP_SENSOR  SocModuleClkT = 12
	SOC_MOD_CLK_INVALID      SocModuleClkT = 13
)

type SocPeriphSystimerClkSrcT c.Int

const (
	SYSTIMER_CLK_SRC_XTAL    SocPeriphSystimerClkSrcT = 9
	SYSTIMER_CLK_SRC_DEFAULT SocPeriphSystimerClkSrcT = 9
)

type SocPeriphGptimerClkSrcT c.Int

const (
	GPTIMER_CLK_SRC_APB     SocPeriphGptimerClkSrcT = 4
	GPTIMER_CLK_SRC_XTAL    SocPeriphGptimerClkSrcT = 9
	GPTIMER_CLK_SRC_DEFAULT SocPeriphGptimerClkSrcT = 4
)

type SocPeriphTgClkSrcLegacyT c.Int

const (
	TIMER_SRC_CLK_APB     SocPeriphTgClkSrcLegacyT = 4
	TIMER_SRC_CLK_XTAL    SocPeriphTgClkSrcLegacyT = 9
	TIMER_SRC_CLK_DEFAULT SocPeriphTgClkSrcLegacyT = 4
)

type SocPeriphLcdClkSrcT c.Int

const (
	LCD_CLK_SRC_PLL160M SocPeriphLcdClkSrcT = 5
	LCD_CLK_SRC_DEFAULT SocPeriphLcdClkSrcT = 5
)

type SocPeriphRmtClkSrcT c.Int

const (
	RMT_CLK_SRC_APB      SocPeriphRmtClkSrcT = 4
	RMT_CLK_SRC_REF_TICK SocPeriphRmtClkSrcT = 10
	RMT_CLK_SRC_DEFAULT  SocPeriphRmtClkSrcT = 4
)

type SocPeriphRmtClkSrcLegacyT c.Int

const (
	RMT_BASECLK_APB     SocPeriphRmtClkSrcLegacyT = 4
	RMT_BASECLK_REF     SocPeriphRmtClkSrcLegacyT = 10
	RMT_BASECLK_DEFAULT SocPeriphRmtClkSrcLegacyT = 4
)

type SocPeriphTemperatureSensorClkSrcT c.Int

const (
	TEMPERATURE_SENSOR_CLK_SRC_RC_FAST SocPeriphTemperatureSensorClkSrcT = 12
	TEMPERATURE_SENSOR_CLK_SRC_DEFAULT SocPeriphTemperatureSensorClkSrcT = 12
)

type SocPeriphUartClkSrcLegacyT c.Int

const (
	UART_SCLK_APB      SocPeriphUartClkSrcLegacyT = 4
	UART_SCLK_REF_TICK SocPeriphUartClkSrcLegacyT = 10
	UART_SCLK_DEFAULT  SocPeriphUartClkSrcLegacyT = 4
)

type SocPeriphI2sClkSrcT c.Int

const (
	I2S_CLK_SRC_DEFAULT  SocPeriphI2sClkSrcT = 5
	I2S_CLK_SRC_PLL_160M SocPeriphI2sClkSrcT = 5
	I2S_CLK_SRC_APLL     SocPeriphI2sClkSrcT = 11
)

type SocPeriphI2cClkSrcT c.Int

const (
	I2C_CLK_SRC_APB      SocPeriphI2cClkSrcT = 4
	I2C_CLK_SRC_REF_TICK SocPeriphI2cClkSrcT = 10
	I2C_CLK_SRC_DEFAULT  SocPeriphI2cClkSrcT = 4
)

type SocPeriphSpiClkSrcT c.Int

const (
	SPI_CLK_SRC_DEFAULT SocPeriphSpiClkSrcT = 4
	SPI_CLK_SRC_APB     SocPeriphSpiClkSrcT = 4
)

type SocPeriphSdmClkSrcT c.Int

const (
	SDM_CLK_SRC_APB     SocPeriphSdmClkSrcT = 4
	SDM_CLK_SRC_DEFAULT SocPeriphSdmClkSrcT = 4
)

type SocPeriphGlitchFilterClkSrcT c.Int

const (
	GLITCH_FILTER_CLK_SRC_APB     SocPeriphGlitchFilterClkSrcT = 4
	GLITCH_FILTER_CLK_SRC_DEFAULT SocPeriphGlitchFilterClkSrcT = 4
)

type SocPeriphDacDigiClkSrcT c.Int

const (
	DAC_DIGI_CLK_SRC_APB     SocPeriphDacDigiClkSrcT = 4
	DAC_DIGI_CLK_SRC_APLL    SocPeriphDacDigiClkSrcT = 11
	DAC_DIGI_CLK_SRC_DEFAULT SocPeriphDacDigiClkSrcT = 4
)

type SocPeriphDacCosineClkSrcT c.Int

const (
	DAC_COSINE_CLK_SRC_RTC_FAST SocPeriphDacCosineClkSrcT = 2
	DAC_COSINE_CLK_SRC_DEFAULT  SocPeriphDacCosineClkSrcT = 2
)

type SocPeriphTwaiClkSrcT c.Int

const (
	TWAI_CLK_SRC_APB     SocPeriphTwaiClkSrcT = 4
	TWAI_CLK_SRC_DEFAULT SocPeriphTwaiClkSrcT = 4
)

type SocPeriphAdcDigiClkSrcT c.Int

const (
	ADC_DIGI_CLK_SRC_APB     SocPeriphAdcDigiClkSrcT = 4
	ADC_DIGI_CLK_SRC_APLL    SocPeriphAdcDigiClkSrcT = 11
	ADC_DIGI_CLK_SRC_DEFAULT SocPeriphAdcDigiClkSrcT = 4
)

type SocPeriphAdcRtcClkSrcT c.Int

const (
	ADC_RTC_CLK_SRC_RC_FAST SocPeriphAdcRtcClkSrcT = 7
	ADC_RTC_CLK_SRC_DEFAULT SocPeriphAdcRtcClkSrcT = 7
)

type SocPeriphMwdtClkSrcT c.Int

const (
	MWDT_CLK_SRC_APB     SocPeriphMwdtClkSrcT = 4
	MWDT_CLK_SRC_DEFAULT SocPeriphMwdtClkSrcT = 4
)

type SocPeriphLedcClkSrcLegacyT c.Int

const (
	LEDC_AUTO_CLK        SocPeriphLedcClkSrcLegacyT = 0
	LEDC_USE_APB_CLK     SocPeriphLedcClkSrcLegacyT = 4
	LEDC_USE_RC_FAST_CLK SocPeriphLedcClkSrcLegacyT = 7
	LEDC_USE_REF_TICK    SocPeriphLedcClkSrcLegacyT = 10
	LEDC_USE_XTAL_CLK    SocPeriphLedcClkSrcLegacyT = 9
	LEDC_USE_RTC8M_CLK   SocPeriphLedcClkSrcLegacyT = 7
)

type SocClkoutSigIdT c.Int

const (
	CLKOUT_SIG_PLL      SocClkoutSigIdT = 1
	CLKOUT_SIG_RC_SLOW  SocClkoutSigIdT = 4
	CLKOUT_SIG_XTAL     SocClkoutSigIdT = 5
	CLKOUT_SIG_APLL     SocClkoutSigIdT = 6
	CLKOUT_SIG_REF_TICK SocClkoutSigIdT = 12
	CLKOUT_SIG_PLL_F80M SocClkoutSigIdT = 13
	CLKOUT_SIG_RC_FAST  SocClkoutSigIdT = 14
	CLKOUT_SIG_INVALID  SocClkoutSigIdT = 255
)

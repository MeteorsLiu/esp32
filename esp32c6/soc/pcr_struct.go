package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Register */
/** Type of uart0_conf register
 *  UART0 configuration register
 */

type PcrUart0ConfRegT struct {
	Val c.Uint32T
}

/** Type of uart0_sclk_conf register
 *  UART0_SCLK configuration register
 */

type PcrUart0SclkConfRegT struct {
	Val c.Uint32T
}

/** Type of uart0_pd_ctrl register
 *  UART0 power control register
 */

type PcrUart0PdCtrlRegT struct {
	Val c.Uint32T
}

/** Type of uart1_conf register
 *  UART1 configuration register
 */

type PcrUart1ConfRegT struct {
	Val c.Uint32T
}

/** Type of uart1_sclk_conf register
 *  UART1_SCLK configuration register
 */

type PcrUart1SclkConfRegT struct {
	Val c.Uint32T
}

/** Type of uart1_pd_ctrl register
 *  UART1 power control register
 */

type PcrUart1PdCtrlRegT struct {
	Val c.Uint32T
}

/** Type of mspi_conf register
 *  MSPI configuration register
 */

type PcrMspiConfRegT struct {
	Val c.Uint32T
}

/** Type of mspi_clk_conf register
 *  MSPI_CLK configuration register
 */

type PcrMspiClkConfRegT struct {
	Val c.Uint32T
}

/** Type of i2c_conf register
 *  I2C configuration register
 */

type PcrI2cConfRegT struct {
	Val c.Uint32T
}

/** Type of i2c_sclk_conf register
 *  I2C_SCLK configuration register
 */

type PcrI2cSclkConfRegT struct {
	Val c.Uint32T
}

/** Type of uhci_conf register
 *  UHCI configuration register
 */

type PcrUhciConfRegT struct {
	Val c.Uint32T
}

/** Type of rmt_conf register
 *  RMT configuration register
 */

type PcrRmtConfRegT struct {
	Val c.Uint32T
}

/** Type of rmt_sclk_conf register
 *  RMT_SCLK configuration register
 */

type PcrRmtSclkConfRegT struct {
	Val c.Uint32T
}

/** Type of ledc_conf register
 *  LEDC configuration register
 */

type PcrLedcConfRegT struct {
	Val c.Uint32T
}

/** Type of ledc_sclk_conf register
 *  LEDC_SCLK configuration register
 */

type PcrLedcSclkConfRegT struct {
	Val c.Uint32T
}

/** Type of timergroup0_conf register
 *  TIMERGROUP0 configuration register
 */

type PcrTimergroup0ConfRegT struct {
	Val c.Uint32T
}

/** Type of timergroup0_timer_clk_conf register
 *  TIMERGROUP0_TIMER_CLK configuration register
 */

type PcrTimergroup0TimerClkConfRegT struct {
	Val c.Uint32T
}

/** Type of timergroup0_wdt_clk_conf register
 *  TIMERGROUP0_WDT_CLK configuration register
 */

type PcrTimergroup0WdtClkConfRegT struct {
	Val c.Uint32T
}

/** Type of timergroup1_conf register
 *  TIMERGROUP1 configuration register
 */

type PcrTimergroup1ConfRegT struct {
	Val c.Uint32T
}

/** Type of timergroup1_timer_clk_conf register
 *  TIMERGROUP1_TIMER_CLK configuration register
 */

type PcrTimergroup1TimerClkConfRegT struct {
	Val c.Uint32T
}

/** Type of timergroup1_wdt_clk_conf register
 *  TIMERGROUP1_WDT_CLK configuration register
 */

type PcrTimergroup1WdtClkConfRegT struct {
	Val c.Uint32T
}

/** Type of systimer_conf register
 *  SYSTIMER configuration register
 */

type PcrSystimerConfRegT struct {
	Val c.Uint32T
}

/** Type of systimer_func_clk_conf register
 *  SYSTIMER_FUNC_CLK configuration register
 */

type PcrSystimerFuncClkConfRegT struct {
	Val c.Uint32T
}

/** Type of twai0_conf register
 *  TWAI0 configuration register
 */

type PcrTwai0ConfRegT struct {
	Val c.Uint32T
}

/** Type of twai0_func_clk_conf register
 *  TWAI0_FUNC_CLK configuration register
 */

type PcrTwai0FuncClkConfRegT struct {
	Val c.Uint32T
}

/** Type of twai1_conf register
 *  TWAI1 configuration register
 */

type PcrTwai1ConfRegT struct {
	Val c.Uint32T
}

/** Type of twai1_func_clk_conf register
 *  TWAI1_FUNC_CLK configuration register
 */

type PcrTwai1FuncClkConfRegT struct {
	Val c.Uint32T
}

/** Type of i2s_conf register
 *  I2S configuration register
 */

type PcrI2sConfRegT struct {
	Val c.Uint32T
}

/** Type of i2s_tx_clkm_conf register
 *  I2S_TX_CLKM configuration register
 */

type PcrI2sTxClkmConfRegT struct {
	Val c.Uint32T
}

/** Type of i2s_tx_clkm_div_conf register
 *  I2S_TX_CLKM_DIV configuration register
 */

type PcrI2sTxClkmDivConfRegT struct {
	Val c.Uint32T
}

/** Type of i2s_rx_clkm_conf register
 *  I2S_RX_CLKM configuration register
 */

type PcrI2sRxClkmConfRegT struct {
	Val c.Uint32T
}

/** Type of i2s_rx_clkm_div_conf register
 *  I2S_RX_CLKM_DIV configuration register
 */

type PcrI2sRxClkmDivConfRegT struct {
	Val c.Uint32T
}

/** Type of saradc_conf register
 *  SARADC configuration register
 */

type PcrSaradcConfRegT struct {
	Val c.Uint32T
}

/** Type of saradc_clkm_conf register
 *  SARADC_CLKM configuration register
 */

type PcrSaradcClkmConfRegT struct {
	Val c.Uint32T
}

/** Type of tsens_clk_conf register
 *  TSENS_CLK configuration register
 */

type PcrTsensClkConfRegT struct {
	Val c.Uint32T
}

/** Type of usb_device_conf register
 *  USB_DEVICE configuration register
 */

type PcrUsbDeviceConfRegT struct {
	Val c.Uint32T
}

/** Type of intmtx_conf register
 *  INTMTX configuration register
 */

type PcrIntmtxConfRegT struct {
	Val c.Uint32T
}

/** Type of pcnt_conf register
 *  PCNT configuration register
 */

type PcrPcntConfRegT struct {
	Val c.Uint32T
}

/** Type of etm_conf register
 *  ETM configuration register
 */

type PcrEtmConfRegT struct {
	Val c.Uint32T
}

/** Type of pwm_conf register
 *  PWM configuration register
 */

type PcrPwmConfRegT struct {
	Val c.Uint32T
}

/** Type of pwm_clk_conf register
 *  PWM_CLK configuration register
 */

type PcrPwmClkConfRegT struct {
	Val c.Uint32T
}

/** Type of parl_io_conf register
 *  PARL_IO configuration register
 */

type PcrParlIoConfRegT struct {
	Val c.Uint32T
}

/** Type of parl_clk_rx_conf register
 *  PARL_CLK_RX configuration register
 */

type PcrParlClkRxConfRegT struct {
	Val c.Uint32T
}

/** Type of parl_clk_tx_conf register
 *  PARL_CLK_TX configuration register
 */

type PcrParlClkTxConfRegT struct {
	Val c.Uint32T
}

/** Type of sdio_slave_conf register
 *  SDIO_SLAVE configuration register
 */

type PcrSdioSlaveConfRegT struct {
	Val c.Uint32T
}

/** Type of pvt_monitor_conf register
 *  PVT_MONITOR configuration register
 */

type PcrPvtMonitorConfRegT struct {
	Val c.Uint32T
}

/** Type of pvt_monitor_func_clk_conf register
 *  PVT_MONITOR function clock configuration register
 */

type PcrPvtMonitorFuncClkConfRegT struct {
	Val c.Uint32T
}

/** Type of gdma_conf register
 *  GDMA configuration register
 */

type PcrGdmaConfRegT struct {
	Val c.Uint32T
}

/** Type of spi2_conf register
 *  SPI2 configuration register
 */

type PcrSpi2ConfRegT struct {
	Val c.Uint32T
}

/** Type of spi2_clkm_conf register
 *  SPI2_CLKM configuration register
 */

type PcrSpi2ClkmConfRegT struct {
	Val c.Uint32T
}

/** Type of aes_conf register
 *  AES configuration register
 */

type PcrAesConfRegT struct {
	Val c.Uint32T
}

/** Type of sha_conf register
 *  SHA configuration register
 */

type PcrShaConfRegT struct {
	Val c.Uint32T
}

/** Type of rsa_conf register
 *  RSA configuration register
 */

type PcrRsaConfRegT struct {
	Val c.Uint32T
}

/** Type of rsa_pd_ctrl register
 *  RSA power control register
 */

type PcrRsaPdCtrlRegT struct {
	Val c.Uint32T
}

/** Type of ecc_conf register
 *  ECC configuration register
 */

type PcrEccConfRegT struct {
	Val c.Uint32T
}

/** Type of ecc_pd_ctrl register
 *  ECC power control register
 */

type PcrEccPdCtrlRegT struct {
	Val c.Uint32T
}

/** Type of ds_conf register
 *  DS configuration register
 */

type PcrDsConfRegT struct {
	Val c.Uint32T
}

/** Type of hmac_conf register
 *  HMAC configuration register
 */

type PcrHmacConfRegT struct {
	Val c.Uint32T
}

/** Type of iomux_conf register
 *  IOMUX configuration register
 */

type PcrIomuxConfRegT struct {
	Val c.Uint32T
}

/** Type of iomux_clk_conf register
 *  IOMUX_CLK configuration register
 */

type PcrIomuxClkConfRegT struct {
	Val c.Uint32T
}

/** Type of mem_monitor_conf register
 *  MEM_MONITOR configuration register
 */

type PcrMemMonitorConfRegT struct {
	Val c.Uint32T
}

/** Type of regdma_conf register
 *  REGDMA configuration register
 */

type PcrRegdmaConfRegT struct {
	Val c.Uint32T
}

/** Type of retention_conf register
 *  retention configuration register
 */

type PcrRetentionConfRegT struct {
	Val c.Uint32T
}

/** Type of trace_conf register
 *  TRACE configuration register
 */

type PcrTraceConfRegT struct {
	Val c.Uint32T
}

/** Type of assist_conf register
 *  ASSIST configuration register
 */

type PcrAssistConfRegT struct {
	Val c.Uint32T
}

/** Type of cache_conf register
 *  CACHE configuration register
 */

type PcrCacheConfRegT struct {
	Val c.Uint32T
}

/** Type of modem_apb_conf register
 *  MODEM_APB configuration register
 */

type PcrModemApbConfRegT struct {
	Val c.Uint32T
}

/** Type of timeout_conf register
 *  TIMEOUT configuration register
 */

type PcrTimeoutConfRegT struct {
	Val c.Uint32T
}

/** Type of sysclk_conf register
 *  SYSCLK configuration register
 */

type PcrSysclkConfRegT struct {
	Val c.Uint32T
}

/** Type of cpu_waiti_conf register
 *  CPU_WAITI configuration register
 */

type PcrCpuWaitiConfRegT struct {
	Val c.Uint32T
}

/** Type of cpu_freq_conf register
 *  CPU_FREQ configuration register
 */

type PcrCpuFreqConfRegT struct {
	Val c.Uint32T
}

/** Type of ahb_freq_conf register
 *  AHB_FREQ configuration register
 */

type PcrAhbFreqConfRegT struct {
	Val c.Uint32T
}

/** Type of apb_freq_conf register
 *  APB_FREQ configuration register
 */

type PcrApbFreqConfRegT struct {
	Val c.Uint32T
}

/** Type of pll_div_clk_en register
 *  SPLL DIV clock-gating configuration register
 */

type PcrPllDivClkEnRegT struct {
	Val c.Uint32T
}

/** Type of ctrl_clk_out_en register
 *  CLK_OUT_EN configuration register
 */

type PcrCtrlClkOutEnRegT struct {
	Val c.Uint32T
}

/** Type of ctrl_tick_conf register
 *  TICK configuration register
 */

type PcrCtrlTickConfRegT struct {
	Val c.Uint32T
}

/** Type of ctrl_32k_conf register
 *  32KHz clock configuration register
 */

type PcrCtrl32kConfRegT struct {
	Val c.Uint32T
}

/** Type of sram_power_conf register
 *  HP SRAM/ROM configuration register
 */

type PcrSramPowerConfRegT struct {
	Val c.Uint32T
}

/** Type of reset_event_bypass register
 *  reset event bypass backdoor configuration register
 */

type PcrResetEventBypassRegT struct {
	Val c.Uint32T
}

/** Type of clock_gate register
 *  PCR clock gating configure register
 */

type PcrClockGateRegT struct {
	Val c.Uint32T
}

/** Group: Frequency Statistics Register */
/** Type of sysclk_freq_query_0 register
 *  SYSCLK frequency query 0 register
 */

type PcrSysclkFreqQuery0RegT struct {
	Val c.Uint32T
}

/** Group: FPGA Debug Register */
/** Type of fpga_debug register
 *  fpga debug register
 */

type PcrFpgaDebugRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  Date register.
 */

type PcrDateRegT struct {
	Val c.Uint32T
}

type PcrDevT struct {
	Uart0Conf               PcrUart0ConfRegT
	Uart0SclkConf           PcrUart0SclkConfRegT
	Uart0PdCtrl             PcrUart0PdCtrlRegT
	Uart1Conf               PcrUart1ConfRegT
	Uart1SclkConf           PcrUart1SclkConfRegT
	Uart1PdCtrl             PcrUart1PdCtrlRegT
	MspiConf                PcrMspiConfRegT
	MspiClkConf             PcrMspiClkConfRegT
	I2cConf                 PcrI2cConfRegT
	I2cSclkConf             PcrI2cSclkConfRegT
	UhciConf                PcrUhciConfRegT
	RmtConf                 PcrRmtConfRegT
	RmtSclkConf             PcrRmtSclkConfRegT
	LedcConf                PcrLedcConfRegT
	LedcSclkConf            PcrLedcSclkConfRegT
	Timergroup0Conf         PcrTimergroup0ConfRegT
	Timergroup0TimerClkConf PcrTimergroup0TimerClkConfRegT
	Timergroup0WdtClkConf   PcrTimergroup0WdtClkConfRegT
	Timergroup1Conf         PcrTimergroup1ConfRegT
	Timergroup1TimerClkConf PcrTimergroup1TimerClkConfRegT
	Timergroup1WdtClkConf   PcrTimergroup1WdtClkConfRegT
	SystimerConf            PcrSystimerConfRegT
	SystimerFuncClkConf     PcrSystimerFuncClkConfRegT
	Twai0Conf               PcrTwai0ConfRegT
	Twai0FuncClkConf        PcrTwai0FuncClkConfRegT
	Twai1Conf               PcrTwai1ConfRegT
	Twai1FuncClkConf        PcrTwai1FuncClkConfRegT
	I2sConf                 PcrI2sConfRegT
	I2sTxClkmConf           PcrI2sTxClkmConfRegT
	I2sTxClkmDivConf        PcrI2sTxClkmDivConfRegT
	I2sRxClkmConf           PcrI2sRxClkmConfRegT
	I2sRxClkmDivConf        PcrI2sRxClkmDivConfRegT
	SaradcConf              PcrSaradcConfRegT
	SaradcClkmConf          PcrSaradcClkmConfRegT
	TsensClkConf            PcrTsensClkConfRegT
	UsbDeviceConf           PcrUsbDeviceConfRegT
	IntmtxConf              PcrIntmtxConfRegT
	PcntConf                PcrPcntConfRegT
	EtmConf                 PcrEtmConfRegT
	PwmConf                 PcrPwmConfRegT
	PwmClkConf              PcrPwmClkConfRegT
	ParlIoConf              PcrParlIoConfRegT
	ParlClkRxConf           PcrParlClkRxConfRegT
	ParlClkTxConf           PcrParlClkTxConfRegT
	SdioSlaveConf           PcrSdioSlaveConfRegT
	PvtMonitorConf          PcrPvtMonitorConfRegT
	PvtMonitorFuncClkConf   PcrPvtMonitorFuncClkConfRegT
	GdmaConf                PcrGdmaConfRegT
	Spi2Conf                PcrSpi2ConfRegT
	Spi2ClkmConf            PcrSpi2ClkmConfRegT
	AesConf                 PcrAesConfRegT
	ShaConf                 PcrShaConfRegT
	RsaConf                 PcrRsaConfRegT
	RsaPdCtrl               PcrRsaPdCtrlRegT
	EccConf                 PcrEccConfRegT
	EccPdCtrl               PcrEccPdCtrlRegT
	DsConf                  PcrDsConfRegT
	HmacConf                PcrHmacConfRegT
	IomuxConf               PcrIomuxConfRegT
	IomuxClkConf            PcrIomuxClkConfRegT
	MemMonitorConf          PcrMemMonitorConfRegT
	RegdmaConf              PcrRegdmaConfRegT
	RetentionConf           PcrRetentionConfRegT
	TraceConf               PcrTraceConfRegT
	AssistConf              PcrAssistConfRegT
	CacheConf               PcrCacheConfRegT
	ModemApbConf            PcrModemApbConfRegT
	TimeoutConf             PcrTimeoutConfRegT
	SysclkConf              PcrSysclkConfRegT
	CpuWaitiConf            PcrCpuWaitiConfRegT
	CpuFreqConf             PcrCpuFreqConfRegT
	AhbFreqConf             PcrAhbFreqConfRegT
	ApbFreqConf             PcrApbFreqConfRegT
	SysclkFreqQuery0        PcrSysclkFreqQuery0RegT
	PllDivClkEn             PcrPllDivClkEnRegT
	CtrlClkOutEn            PcrCtrlClkOutEnRegT
	CtrlTickConf            PcrCtrlTickConfRegT
	Ctrl32kConf             PcrCtrl32kConfRegT
	SramPowerConf           PcrSramPowerConfRegT
	Reserved13c             [941]c.Uint32T
	ResetEventBypass        PcrResetEventBypassRegT
	FpgaDebug               PcrFpgaDebugRegT
	ClockGate               PcrClockGateRegT
	Date                    PcrDateRegT
}

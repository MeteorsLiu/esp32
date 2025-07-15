package esp_hw_support

import _ "unsafe"

/**
 * @brief Whether to allow the cpu power domain to be powered off.
 *
 * In light sleep mode, only when the system can provide enough memory
 * for cpu retention, the cpu power domain can be powered off.
 */
//go:linkname CpuDomainPdAllowed C.cpu_domain_pd_allowed
func CpuDomainPdAllowed() bool

/**
 * @brief Configure the parameters of the CPU domain during the sleep process
 *
 * @param light_sleep_enable true for enable light sleep mode, false for disable light sleep mode
 *
 * @return
 *  - ESP_OK on success
 */
//go:linkname SleepCpuConfigure C.sleep_cpu_configure
func SleepCpuConfigure(light_sleep_enable bool) EspErrT

/**
 * @brief Enable cpu retention of some modules.
 *
 * In light sleep mode, before the system goes to sleep, enable the cpu
 * retention of modules such as CPU and I/D-cache tag memory.
 */
//go:linkname SleepEnableCpuRetention C.sleep_enable_cpu_retention
func SleepEnableCpuRetention()

/**
 * @brief Disable cpu retention of some modules.
 *
 * In light sleep mode, after the system exits sleep, disable the cpu
 * retention of modules such as CPU and I/D-cache tag memory.
 */
//go:linkname SleepDisableCpuRetention C.sleep_disable_cpu_retention
func SleepDisableCpuRetention()

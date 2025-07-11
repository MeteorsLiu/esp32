package efuse

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_EFUSE_ADC_CALIB_VER = 2
const RTCCALIB_ESP32S2_ADCCOUNT = 2
const RTCCALIB_ESP32S2_ATTENCOUNT = 4
const RTCCALIB_V1_PARAM_VLOW = 0
const RTCCALIB_V1_PARAM_VHIGH = 1
const RTCCALIB_V2_PARAM_VHIGH = 0
const RTCCALIB_V2_PARAM_VINIT = 1
const RTCCALIB_V1IDX_A10L = 1
const RTCCALIB_V1IDX_A11L = 2
const RTCCALIB_V1IDX_A12L = 3
const RTCCALIB_V1IDX_A13L = 4
const RTCCALIB_V1IDX_A20L = 5
const RTCCALIB_V1IDX_A21L = 6
const RTCCALIB_V1IDX_A22L = 7
const RTCCALIB_V1IDX_A23L = 8
const RTCCALIB_V1IDX_A10H = 9
const RTCCALIB_V1IDX_A11H = 10
const RTCCALIB_V1IDX_A12H = 11
const RTCCALIB_V1IDX_A13H = 12
const RTCCALIB_V1IDX_A20H = 13
const RTCCALIB_V1IDX_A21H = 14
const RTCCALIB_V1IDX_A22H = 15
const RTCCALIB_V1IDX_A23H = 16
const RTCCALIB_V2IDX_A10H = 17
const RTCCALIB_V2IDX_A11H = 18
const RTCCALIB_V2IDX_A12H = 19
const RTCCALIB_V2IDX_A13H = 20
const RTCCALIB_V2IDX_A20H = 21
const RTCCALIB_V2IDX_A21H = 22
const RTCCALIB_V2IDX_A22H = 23
const RTCCALIB_V2IDX_A23H = 24
const RTCCALIB_V2IDX_A10I = 25
const RTCCALIB_V2IDX_A11I = 26
const RTCCALIB_V2IDX_A12I = 27
const RTCCALIB_V2IDX_A13I = 28
const RTCCALIB_V2IDX_A20I = 29
const RTCCALIB_V2IDX_A21I = 30
const RTCCALIB_V2IDX_A22I = 31
const RTCCALIB_V2IDX_A23I = 32
const RTCCALIB_IDX_TMPSENSOR = 33

/**
 * @brief Get rtc calibration version.
 */
//go:linkname EspEfuseRtcTableReadCalibVersion C.esp_efuse_rtc_table_read_calib_version
func EspEfuseRtcTableReadCalibVersion() c.Int

/**
 * @brief Helper function to calculate a tag from human-readable parameters.
 * Tag is used to index the desired data from the efuse.
 * For example, (1, 1, 3, 1) yields the tag RTCCALIB_V1IDX_A13H
 * extra params are used for identification when a adc_num-atten combination has
 * multiple efuse values.
 * @param adc_channel_num verbatim numbering of the ADC channel. For channel 1, use 1 and not 0.
 * @param atten attenuation. use the enum value.
 * @param version the version of the scheme to index for.
 * @param extra_params defined differently for each version.
 * */
//go:linkname EspEfuseRtcTableGetTag C.esp_efuse_rtc_table_get_tag
func EspEfuseRtcTableGetTag(version c.Int, adc_channel_num c.Int, atten c.Int, extra_params c.Int) c.Int

/**
 * @brief Fetches a raw value from efuse and does signed bit parsing
 * @param tag tag obtained with esp_efuse_rtc_table_get_tag
 *
 * */
//go:linkname EspEfuseRtcTableGetRawEfuseValue C.esp_efuse_rtc_table_get_raw_efuse_value
func EspEfuseRtcTableGetRawEfuseValue(tag c.Int) c.Int

/**
 * @brief Fetches a raw value from efuse and resolve it to get
 * the original number that it meant to represent.
 *
 * @param tag tag obtained with esp_efuse_rtc_table_get_tag
 * @param use_zero_inputs Does not perform the raw value fetching before resolving the number,
 * but proceed as if all zeros were read from efuse.
 *
 * */
//go:linkname EspEfuseRtcTableGetParsedEfuseValue C.esp_efuse_rtc_table_get_parsed_efuse_value
func EspEfuseRtcTableGetParsedEfuseValue(tag c.Int, skip_efuse_reading bool) c.Int

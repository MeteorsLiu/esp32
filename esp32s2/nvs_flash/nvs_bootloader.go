package nvs_flash

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Placeholders for buffer pointer and length of string type
 */

type NvsBootloaderStrValuePlaceholderT struct {
	BuffPtr *c.Char
	BuffLen c.SizeT
}

/**
 * @brief Union of value placeholders for all nvs_type_t supported by bootloader code
 */

type NvsBootloaderValuePlaceholderT struct {
	StrVal NvsBootloaderStrValuePlaceholderT
}

/**
 * @brief Structure representing one NVS bootloader entry.
 *
 * This structure serves as read operation input parameters and result value and status placeholder.
 * Before calling the `nvs_bootloader_read` function, populate the namespace_name, key_name and value_type members.
 * If string value has to be read, provide also buffer and its length in the `value.str_val` member.
 *
 * The result_code member will be populated by the function with the result of the read operation.
 * There are 2 possible situations and interpretations of the result_code:
 * If the return value of the `nvs_bootloader_read` was ESP_OK, the result_code will be one of the following:
 * - `ESP_OK`: Entry found, value member contains the data. This is the only case when the value member is populated.
 * - `ESP_ERR_NVS_TYPE_MISMATCH`: Entry was found, but requested datatype doesn't match datatype found in NVS
 * - `ESP_ERR_NVS_NOT_FOUND`: Data was not found.
 * - `ESP_ERR_INVALID_SIZE`: the value found for string is longer than the space provided in placeholder (str_val.buff_len)
 * If the return value of the function was ESP_ERR_INVALID_ARG, the result_code will be one of the following:
 * - `ESP_ERR_NVS_NOT_FOUND`: Check of this parameters was successful.
 * - `ESP_ERR_NVS_INVALID_NAME`: namespace_name is NULL or too long
 * - `ESP_ERR_NVS_KEY_TOO_LONG`: key_name NULL or too long
 * - `ESP_ERR_INVALID_SIZE`: the size of the buffer provided for NVS_TYPE_STR in placeholder (str_val.buff_len) is zero or exceeds maximum value NVS_CONST_STR_LEN_MAX_SIZE
 * - `ESP_ERR_INVALID_ARG`: Invalid datatype requested
 *
 */

type NvsBootloaderReadListT struct {
	NamespaceName  *c.Char
	KeyName        *c.Char
	ValueType      NvsTypeT
	ResultCode     EspErrT
	Value          NvsBootloaderValuePlaceholderT
	NamespaceIndex c.Uint8T
}

/**
 * @brief Reads data specified from the specified NVS partition.
 *
 * This function reads data from the NVS partition specified by `partition_name`.
 * Multiple NVS entries can be read in a single call. The list of entries to read is specified in the `read_list` array.
 * Function indicates overall success or failure by its return value. In case it is ESP_OK or ESP_ERR_INVALID_ARG,
 * result of validation / reading of individual entry is returned in the `result_code` member of each element of the `read_list` array.
 *
 * @param partition_name The name of the NVS partition to read from.
 * @param read_list_count The number of elements in the `read_list` array.
 * @param read_list An array of `nvs_bootloader_read_list_t` structures specifying the keys and buffers for reading data.
 *
 * @return
 * The return value of the function in this file can be one of the following:
 * - `ESP_OK`: The function successfully checked all input parameters and executed successfully.
 *   The individual `result_code` in `read_list` indicates the result of the lookup for a particular requested key.
 * - `ESP_ERR_INVALID_ARG`: The validity of all `read_list` input parameters was
 *   checked and failed for at least one of the parameters. The individual `result_code`
 *   in `read_list` provides the detailed reason. This error code is also returned when read_list is null or read_list_count is 0.
 * - `ESP_ERR_NVS_INVALID_NAME`: The partition name specified is too long or is null.
 * - `ESP_ERR_NVS_PART_NOT_FOUND`: The partition was not found in the partition table.
 * - `ESP_ERR_NVS_CORRUPT_KEY_PART`: Encryption-related problems.
 * - `ESP_ERR_NVS_WRONG_ENCRYPTION`: Encryption-related problems.
 * - `ESP_ERR_INVALID_STATE`: NVS partition or pages related errors - wrong size of partition, header inconsistent / entries
 *   inconsistencies, multiple active pages, page in state INVALID.
 * - `ESP_ERR_NO_MEM`: Cannot allocate memory required to perform the function.
 * - Technical errors in underlying storage.
 */
//go:linkname NvsBootloaderRead C.nvs_bootloader_read
func NvsBootloaderRead(partition_name *c.Char, read_list_count c.SizeT, read_list *NvsBootloaderReadListT) EspErrT

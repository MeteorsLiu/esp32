package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Set the work mode of the operation
 *
 * @param mode Mode of operation
 */
// llgo:link EccModeT.EccHalSetMode C.ecc_hal_set_mode
func (recv_ EccModeT) EccHalSetMode() {
}

/**
 * @brief Set the ECC curve of operation
 *
 * @param curve Curve to use for operation
 */
// llgo:link EccCurveT.EccHalSetCurve C.ecc_hal_set_curve
func (recv_ EccCurveT) EccHalSetCurve() {
}

/**
 * @brief Start calculation
 *
 */
//go:linkname EccHalStartCalc C.ecc_hal_start_calc
func EccHalStartCalc()

/**
 * @brief Check whether the calculation has finished
 *
 * @return - 1 if the hardware has finished calculating
 *         - 0 otherwise
 */
//go:linkname EccHalIsCalcFinished C.ecc_hal_is_calc_finished
func EccHalIsCalcFinished() c.Int

/**
 * @brief Write parameters for point multiplication (K * (Px, Py))
 *
 * @param k Scalar value
 * @param px X coordinate of the ECC point
 * @param py Y coordinate of the ECC point
 * @param len Length (in bytes) of the ECC point
 *            - 32 bytes for SECP256R1
 *            - 24 bytes for SECP192R1
 */
//go:linkname EccHalWriteMulParam C.ecc_hal_write_mul_param
func EccHalWriteMulParam(k *c.Uint8T, px *c.Uint8T, py *c.Uint8T, len c.Uint16T)

/**
 * @brief Write parameters for point verification,
 *        i.e to check if the point lies on the curve
 *
 * @param px X coordinate of the ECC point
 * @param py Y coordinate of the ECC point
 * @param len Length (in bytes) of the ECC point
 *            - 32 for SECP256R1
 *            - 24 for SECP192R1
 */
//go:linkname EccHalWriteVerifyParam C.ecc_hal_write_verify_param
func EccHalWriteVerifyParam(px *c.Uint8T, py *c.Uint8T, len c.Uint16T)

/**
 * @brief Read point multiplication result
 *
 * @param rx X coordinate of the multiplication result
 * @param ry Y coordinate of the multiplication result
 * @param len Length (in bytes) of the ECC point
 *            - 32 for SECP256R1
 *            - 24 for SECP192R1
 *
 * @return - 0 if the operation was successful
 *         - -1 if the operation was not successful
 *
 * In case the operation is not successful, rx and ry will contain
 * all zeros
 */
//go:linkname EccHalReadMulResult C.ecc_hal_read_mul_result
func EccHalReadMulResult(rx *c.Uint8T, ry *c.Uint8T, len c.Uint16T) c.Int

/**
 * @brief Read point verification result
 *
 * @return - 1 if point lies on curve
 *         - 0 otherwise
 */
//go:linkname EccHalReadVerifyResult C.ecc_hal_read_verify_result
func EccHalReadVerifyResult() c.Int

/**
 * @brief Set the mod base value used in MOD operation
 *
 * @param base Identifier of the base to use
 */
// llgo:link EccModBaseT.EccHalSetModBase C.ecc_hal_set_mod_base
func (recv_ EccModBaseT) EccHalSetModBase() {
}

/**
 * @brief Write parameters for Jacobian verification
 *        i.e Check whether (Qx, Qy, Qz) is a point on selected curve
 *
 * @param qx X coordinate of the ECC point in jacobian form
 * @param qy Y coordinate of the ECC point in jacobian form
 * @param qz Z coordinate of the ECC point in jacobian form
 * @param len Length (in bytes) of the ECC point
 *            - 32 bytes for SECP256R1
 *            - 24 bytes for SECP192R1
 */
//go:linkname EccHalWriteJacobVerifyParam C.ecc_hal_write_jacob_verify_param
func EccHalWriteJacobVerifyParam(qx *c.Uint8T, qy *c.Uint8T, qz *c.Uint8T, len c.Uint16T)

/**
 * @brief Read ECC point multiplication result in jacobian form
 *
 * @param rx X coordinate of the multiplication result
 * @param ry Y coordinate of the multiplication result
 * @param rz Z coordinate of the multiplication result
 * @param len Length (in bytes) of the ECC point
 *            - 32 for SECP256R1
 *            - 24 for SECP192R1
 *
 * @return - 0 if the operation was successful
 *         - -1 if the operation was not successful
 *
 * In case the operation is not successful, rx, ry, and rz will contain
 * all zeros
 */
//go:linkname EccHalReadJacobMulResult C.ecc_hal_read_jacob_mul_result
func EccHalReadJacobMulResult(rx *c.Uint8T, ry *c.Uint8T, rz *c.Uint8T, len c.Uint16T) c.Int

/**
 * @brief Write parameters for ECC point addition ((Px, Py, 1) + (Qx, Qy, Qz))
 *
 * @param px X coordinate of the 1st addend ECC point
 * @param py Y coordinate of the 1st addend ECC point
 * @param qx X coordinate of the 2nd addend ECC point in jacobian form
 * @param qy Y coordinate of the 2nd addend ECC point in jacobian form
 * @param qz Z coordinate of the 2nd addend ECC point in jacobian form
 * @param len Length (in bytes) of the ECC point
 *            - 32 bytes for SECP256R1
 *            - 24 bytes for SECP192R1
 */
//go:linkname EccHalWritePointAddParam C.ecc_hal_write_point_add_param
func EccHalWritePointAddParam(px *c.Uint8T, py *c.Uint8T, qx *c.Uint8T, qy *c.Uint8T, qz *c.Uint8T, len c.Uint16T)

/**
 * @brief Read ECC point addition result
 *
 * @param rx X coordinate of the addition result
 * @param ry Y coordinate of the addition result
 * @param rz Z coordinate of the addition result
 * @param len Length (in bytes) of the ECC point
 *            - 32 for SECP256R1
 *            - 24 for SECP192R1
 * @param read_jacob Read the result in Jacobian form
 *
 * @return - 0 if the operation was successful
 *         - -1 otherwise
 */
//go:linkname EccHalReadPointAddResult C.ecc_hal_read_point_add_result
func EccHalReadPointAddResult(rx *c.Uint8T, ry *c.Uint8T, rz *c.Uint8T, len c.Uint16T, read_jacob bool) c.Int

/**
 * @brief Write parameters for mod operations
 *        i.e mod add, mod sub, mod mul, mod inverse mul (or mod division)
 *
 * @param a Value of operand 1
 * @param b Value of operand 2
 * @param len Length (in bytes) of the ECC point
 *            - 32 bytes for SECP256R1
 *            - 24 bytes for SECP192R1
 */
//go:linkname EccHalWriteModOpParam C.ecc_hal_write_mod_op_param
func EccHalWriteModOpParam(a *c.Uint8T, b *c.Uint8T, len c.Uint16T)

/**
 * @brief Read result of mod operations
 *        i.e mod add, mod sub, mod mul, mod inverse mul (or mod division)
 *
 * @param r Result of the mod operation
 * @param len Length (in bytes) of the ECC point
 *            - 32 bytes for SECP256R1
 *            - 24 bytes for SECP192R1
 *
 * @return - 0 if operation successful
 *         - -1 otherwise
 */
//go:linkname EccHalReadModOpResult C.ecc_hal_read_mod_op_result
func EccHalReadModOpResult(r *c.Uint8T, len c.Uint16T) c.Int

/**
 * @brief Enable constant time multiplication operations
 *
 * @param true: enable; false: disable
 */
//go:linkname EccHalEnableConstantTimePointMul C.ecc_hal_enable_constant_time_point_mul
func EccHalEnableConstantTimePointMul(enable bool)

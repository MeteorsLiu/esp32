package unity

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const UNITY_DETAIL1_NAME = "Function"
const UNITY_DETAIL2_NAME = "Argument"

type UNITYUINT8 c.Char
type UNITYUINT16 uint16
type UNITYUINT32 c.Uint
type UNITYINT8 c.Char
type UNITYINT16 int16
type UNITYINT32 c.Int
type UNITYUINT UNITYUINT32
type UNITYINT UNITYINT32
type UNITYFLOAT c.Float
type UNITYDOUBLE c.Double

// llgo:type C
type UnityTestFunction func()
type UNITYDISPLAYSTYLET c.Int

const (
	UNITY_DISPLAY_STYLE_INT     UNITYDISPLAYSTYLET = 20
	UNITY_DISPLAY_STYLE_INT8    UNITYDISPLAYSTYLET = 17
	UNITY_DISPLAY_STYLE_INT16   UNITYDISPLAYSTYLET = 18
	UNITY_DISPLAY_STYLE_INT32   UNITYDISPLAYSTYLET = 20
	UNITY_DISPLAY_STYLE_UINT    UNITYDISPLAYSTYLET = 36
	UNITY_DISPLAY_STYLE_UINT8   UNITYDISPLAYSTYLET = 33
	UNITY_DISPLAY_STYLE_UINT16  UNITYDISPLAYSTYLET = 34
	UNITY_DISPLAY_STYLE_UINT32  UNITYDISPLAYSTYLET = 36
	UNITY_DISPLAY_STYLE_HEX8    UNITYDISPLAYSTYLET = 65
	UNITY_DISPLAY_STYLE_HEX16   UNITYDISPLAYSTYLET = 66
	UNITY_DISPLAY_STYLE_HEX32   UNITYDISPLAYSTYLET = 68
	UNITY_DISPLAY_STYLE_CHAR    UNITYDISPLAYSTYLET = 145
	UNITY_DISPLAY_STYLE_UNKNOWN UNITYDISPLAYSTYLET = 146
)

type UNITYCOMPARISONT c.Int

const (
	UNITY_WITHIN           UNITYCOMPARISONT = 0
	UNITY_EQUAL_TO         UNITYCOMPARISONT = 1
	UNITY_GREATER_THAN     UNITYCOMPARISONT = 2
	UNITY_GREATER_OR_EQUAL UNITYCOMPARISONT = 3
	UNITY_SMALLER_THAN     UNITYCOMPARISONT = 4
	UNITY_SMALLER_OR_EQUAL UNITYCOMPARISONT = 5
	UNITY_NOT_EQUAL        UNITYCOMPARISONT = 0
	UNITY_UNKNOWN          UNITYCOMPARISONT = 1
)

type UNITYFLOATTRAIT c.Int

const (
	UNITY_FLOAT_IS_NOT_INF     UNITYFLOATTRAIT = 0
	UNITY_FLOAT_IS_INF         UNITYFLOATTRAIT = 1
	UNITY_FLOAT_IS_NOT_NEG_INF UNITYFLOATTRAIT = 2
	UNITY_FLOAT_IS_NEG_INF     UNITYFLOATTRAIT = 3
	UNITY_FLOAT_IS_NOT_NAN     UNITYFLOATTRAIT = 4
	UNITY_FLOAT_IS_NAN         UNITYFLOATTRAIT = 5
	UNITY_FLOAT_IS_NOT_DET     UNITYFLOATTRAIT = 6
	UNITY_FLOAT_IS_DET         UNITYFLOATTRAIT = 7
	UNITY_FLOAT_INVALID_TRAIT  UNITYFLOATTRAIT = 8
)

type UNITYFLOATTRAITT UNITYFLOATTRAIT
type UNITYFLAGST c.Int

const (
	UNITY_ARRAY_TO_VAL   UNITYFLAGST = 0
	UNITY_ARRAY_TO_ARRAY UNITYFLAGST = 1
	UNITY_ARRAY_UNKNOWN  UNITYFLAGST = 2
)

type UNITYSTORAGET struct {
	TestFile              *c.Char
	CurrentTestName       *c.Char
	CurrentDetail1        *c.Char
	CurrentDetail2        *c.Char
	CurrentTestLineNumber UNITYUINT
	NumberOfTests         UNITYUINT
	TestFailures          UNITYUINT
	TestIgnores           UNITYUINT
	CurrentTestFailed     UNITYUINT
	CurrentTestIgnored    UNITYUINT
	AbortFrame            JmpBuf
}

/*-------------------------------------------------------
 * Test Suite Management
 *-------------------------------------------------------*/
//go:linkname UnityBegin C.UnityBegin
func UnityBegin(filename *c.Char)

//go:linkname UnityEnd C.UnityEnd
func UnityEnd() c.Int

//go:linkname UnitySetTestFile C.UnitySetTestFile
func UnitySetTestFile(filename *c.Char)

//go:linkname UnityConcludeTest C.UnityConcludeTest
func UnityConcludeTest()

//go:linkname UnityDefaultTestRun C.UnityDefaultTestRun
func UnityDefaultTestRun(Func UnityTestFunction, FuncName *c.Char, FuncLineNum c.Int)

/*-------------------------------------------------------
 * Test Output
 *-------------------------------------------------------*/
//go:linkname UnityPrint C.UnityPrint
func UnityPrint(string *c.Char)

//go:linkname UnityPrintLen C.UnityPrintLen
func UnityPrintLen(string *c.Char, length UNITYUINT32)

// llgo:link UNITYUINT.UnityPrintMask C.UnityPrintMask
func (recv_ UNITYUINT) UnityPrintMask(number UNITYUINT) {
}

// llgo:link UNITYINT.UnityPrintNumberByStyle C.UnityPrintNumberByStyle
func (recv_ UNITYINT) UnityPrintNumberByStyle(style UNITYDISPLAYSTYLET) {
}

// llgo:link UNITYINT.UnityPrintNumber C.UnityPrintNumber
func (recv_ UNITYINT) UnityPrintNumber() {
}

// llgo:link UNITYUINT.UnityPrintNumberUnsigned C.UnityPrintNumberUnsigned
func (recv_ UNITYUINT) UnityPrintNumberUnsigned() {
}

// llgo:link UNITYUINT.UnityPrintNumberHex C.UnityPrintNumberHex
func (recv_ UNITYUINT) UnityPrintNumberHex(nibbles_to_print c.Char) {
}

// llgo:link UNITYDOUBLE.UnityPrintFloat C.UnityPrintFloat
func (recv_ UNITYDOUBLE) UnityPrintFloat() {
}

/*-------------------------------------------------------
 * Test Assertion Functions
 *-------------------------------------------------------
 *  Use the macros below this section instead of calling
 *  these directly. The macros have a consistent naming
 *  convention and will pull in file and line information
 *  for you. */
// llgo:link UNITYINT.UnityAssertEqualNumber C.UnityAssertEqualNumber
func (recv_ UNITYINT) UnityAssertEqualNumber(actual UNITYINT, msg *c.Char, lineNumber UNITYUINT, style UNITYDISPLAYSTYLET) {
}

// llgo:link UNITYINT.UnityAssertGreaterOrLessOrEqualNumber C.UnityAssertGreaterOrLessOrEqualNumber
func (recv_ UNITYINT) UnityAssertGreaterOrLessOrEqualNumber(actual UNITYINT, compare UNITYCOMPARISONT, msg *c.Char, lineNumber UNITYUINT, style UNITYDISPLAYSTYLET) {
}

//go:linkname UnityAssertEqualIntArray C.UnityAssertEqualIntArray
func UnityAssertEqualIntArray(expected c.Pointer, actual c.Pointer, num_elements UNITYUINT32, msg *c.Char, lineNumber UNITYUINT, style UNITYDISPLAYSTYLET, flags UNITYFLAGST)

// llgo:link UNITYINT.UnityAssertBits C.UnityAssertBits
func (recv_ UNITYINT) UnityAssertBits(expected UNITYINT, actual UNITYINT, msg *c.Char, lineNumber UNITYUINT) {
}

//go:linkname UnityAssertEqualString C.UnityAssertEqualString
func UnityAssertEqualString(expected *c.Char, actual *c.Char, msg *c.Char, lineNumber UNITYUINT)

//go:linkname UnityAssertEqualStringLen C.UnityAssertEqualStringLen
func UnityAssertEqualStringLen(expected *c.Char, actual *c.Char, length UNITYUINT32, msg *c.Char, lineNumber UNITYUINT)

//go:linkname UnityAssertEqualStringArray C.UnityAssertEqualStringArray
func UnityAssertEqualStringArray(expected c.Pointer, actual **c.Char, num_elements UNITYUINT32, msg *c.Char, lineNumber UNITYUINT, flags UNITYFLAGST)

//go:linkname UnityAssertEqualMemory C.UnityAssertEqualMemory
func UnityAssertEqualMemory(expected c.Pointer, actual c.Pointer, length UNITYUINT32, num_elements UNITYUINT32, msg *c.Char, lineNumber UNITYUINT, flags UNITYFLAGST)

// llgo:link UNITYUINT.UnityAssertNumbersWithin C.UnityAssertNumbersWithin
func (recv_ UNITYUINT) UnityAssertNumbersWithin(expected UNITYINT, actual UNITYINT, msg *c.Char, lineNumber UNITYUINT, style UNITYDISPLAYSTYLET) {
}

// llgo:link UNITYUINT.UnityAssertNumbersArrayWithin C.UnityAssertNumbersArrayWithin
func (recv_ UNITYUINT) UnityAssertNumbersArrayWithin(expected c.Pointer, actual c.Pointer, num_elements UNITYUINT32, msg *c.Char, lineNumber UNITYUINT, style UNITYDISPLAYSTYLET, flags UNITYFLAGST) {
}

//go:linkname UnityFail C.UnityFail
func UnityFail(message *c.Char, line UNITYUINT)

//go:linkname UnityIgnore C.UnityIgnore
func UnityIgnore(message *c.Char, line UNITYUINT)

//go:linkname UnityMessage C.UnityMessage
func UnityMessage(message *c.Char, line UNITYUINT)

// llgo:link UNITYFLOAT.UnityAssertFloatsWithin C.UnityAssertFloatsWithin
func (recv_ UNITYFLOAT) UnityAssertFloatsWithin(expected UNITYFLOAT, actual UNITYFLOAT, msg *c.Char, lineNumber UNITYUINT) {
}

// llgo:link UNITYFLOAT.UnityAssertFloatsNotWithin C.UnityAssertFloatsNotWithin
func (recv_ UNITYFLOAT) UnityAssertFloatsNotWithin(expected UNITYFLOAT, actual UNITYFLOAT, msg *c.Char, lineNumber UNITYUINT) {
}

// llgo:link UNITYFLOAT.UnityAssertGreaterOrLessFloat C.UnityAssertGreaterOrLessFloat
func (recv_ UNITYFLOAT) UnityAssertGreaterOrLessFloat(actual UNITYFLOAT, compare UNITYCOMPARISONT, msg *c.Char, linenumber UNITYUINT) {
}

// llgo:link UNITYFLOAT.UnityAssertWithinFloatArray C.UnityAssertWithinFloatArray
func (recv_ UNITYFLOAT) UnityAssertWithinFloatArray(expected *UNITYFLOAT, actual *UNITYFLOAT, num_elements UNITYUINT32, msg *c.Char, lineNumber UNITYUINT, flags UNITYFLAGST) {
}

// llgo:link UNITYFLOAT.UnityAssertFloatSpecial C.UnityAssertFloatSpecial
func (recv_ UNITYFLOAT) UnityAssertFloatSpecial(msg *c.Char, lineNumber UNITYUINT, style UNITYFLOATTRAITT) {
}

// llgo:link UNITYDOUBLE.UnityAssertDoublesWithin C.UnityAssertDoublesWithin
func (recv_ UNITYDOUBLE) UnityAssertDoublesWithin(expected UNITYDOUBLE, actual UNITYDOUBLE, msg *c.Char, lineNumber UNITYUINT) {
}

// llgo:link UNITYDOUBLE.UnityAssertDoublesNotWithin C.UnityAssertDoublesNotWithin
func (recv_ UNITYDOUBLE) UnityAssertDoublesNotWithin(expected UNITYDOUBLE, actual UNITYDOUBLE, msg *c.Char, lineNumber UNITYUINT) {
}

// llgo:link UNITYDOUBLE.UnityAssertGreaterOrLessDouble C.UnityAssertGreaterOrLessDouble
func (recv_ UNITYDOUBLE) UnityAssertGreaterOrLessDouble(actual UNITYDOUBLE, compare UNITYCOMPARISONT, msg *c.Char, linenumber UNITYUINT) {
}

// llgo:link UNITYDOUBLE.UnityAssertWithinDoubleArray C.UnityAssertWithinDoubleArray
func (recv_ UNITYDOUBLE) UnityAssertWithinDoubleArray(expected *UNITYDOUBLE, actual *UNITYDOUBLE, num_elements UNITYUINT32, msg *c.Char, lineNumber UNITYUINT, flags UNITYFLAGST) {
}

// llgo:link UNITYDOUBLE.UnityAssertDoubleSpecial C.UnityAssertDoubleSpecial
func (recv_ UNITYDOUBLE) UnityAssertDoubleSpecial(msg *c.Char, lineNumber UNITYUINT, style UNITYFLOATTRAITT) {
}

/*-------------------------------------------------------
 * Helpers
 *-------------------------------------------------------*/
// llgo:link UNITYINT.UnityNumToPtr C.UnityNumToPtr
func (recv_ UNITYINT) UnityNumToPtr(size UNITYUINT8) c.Pointer {
	return nil
}

//go:linkname UnityFloatToPtr C.UnityFloatToPtr
func UnityFloatToPtr(num c.Float) c.Pointer

//go:linkname UnityDoubleToPtr C.UnityDoubleToPtr
func UnityDoubleToPtr(num c.Double) c.Pointer

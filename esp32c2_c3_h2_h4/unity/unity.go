package unity

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const UNITY_VERSION_MAJOR = 2
const UNITY_VERSION_MINOR = 6
const UNITY_VERSION_BUILD = 0

/* These functions are intended to be called before and after each test.
 * If using unity directly, these will need to be provided for each test
 * executable built. If you are using the test runner generator and/or
 * Ceedling, these are optional. */
//go:linkname SetUp C.setUp
func SetUp()

//go:linkname TearDown C.tearDown
func TearDown()

/* These functions are intended to be called at the beginning and end of an
 * entire test suite.  suiteTearDown() is passed the number of tests that
 * failed, and its return value becomes the exit code of main(). If using
 * Unity directly, you're in charge of calling these if they are desired.
 * If using Ceedling or the test runner generator, these will be called
 * automatically if they exist. */
//go:linkname SuiteSetUp C.suiteSetUp
func SuiteSetUp()

//go:linkname SuiteTearDown C.suiteTearDown
func SuiteTearDown(num_failures c.Int) c.Int

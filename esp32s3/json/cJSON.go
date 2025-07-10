package json

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const CJSON_VERSION_MAJOR = 1
const CJSON_VERSION_MINOR = 7
const CJSON_VERSION_PATCH = 18
const CJSON_IsReference = 256
const CJSON_StringIsConst = 512
const CJSON_NESTING_LIMIT = 1000
const CJSON_CIRCULAR_LIMIT = 10000

/* The cJSON structure: */

type CJSON struct {
	Next        *CJSON
	Prev        *CJSON
	Child       *CJSON
	Type        c.Int
	Valuestring *c.Char
	Valueint    c.Int
	Valuedouble c.Double
	String      *c.Char
}

type CJSONHooks struct {
	MallocFn c.Pointer
	FreeFn   c.Pointer
}
type CJSONBool c.Int

/* returns the version of cJSON as a string */
//go:linkname CJSONVersion C.cJSON_Version
func CJSONVersion() *c.Char

/* Supply malloc, realloc and free functions to cJSON */
// llgo:link (*CJSONHooks).CJSONInitHooks C.cJSON_InitHooks
func (recv_ *CJSONHooks) CJSONInitHooks() {
}

/* Memory Management: the caller is always responsible to free the results from all variants of cJSON_Parse (with cJSON_Delete) and cJSON_Print (with stdlib free, cJSON_Hooks.free_fn, or cJSON_free as appropriate). The exception is cJSON_PrintPreallocated, where the caller has full responsibility of the buffer. */
/* Supply a block of JSON, and this returns a cJSON object you can interrogate. */
//go:linkname CJSONParse C.cJSON_Parse
func CJSONParse(value *c.Char) *CJSON

//go:linkname CJSONParseWithLength C.cJSON_ParseWithLength
func CJSONParseWithLength(value *c.Char, buffer_length c.SizeT) *CJSON

/* ParseWithOpts allows you to require (and check) that the JSON is null terminated, and to retrieve the pointer to the final byte parsed. */
/* If you supply a ptr in return_parse_end and parsing fails, then return_parse_end will contain a pointer to the error so will match cJSON_GetErrorPtr(). */
//go:linkname CJSONParseWithOpts C.cJSON_ParseWithOpts
func CJSONParseWithOpts(value *c.Char, return_parse_end **c.Char, require_null_terminated CJSONBool) *CJSON

//go:linkname CJSONParseWithLengthOpts C.cJSON_ParseWithLengthOpts
func CJSONParseWithLengthOpts(value *c.Char, buffer_length c.SizeT, return_parse_end **c.Char, require_null_terminated CJSONBool) *CJSON

/* Render a cJSON entity to text for transfer/storage. */
// llgo:link (*CJSON).CJSONPrint C.cJSON_Print
func (recv_ *CJSON) CJSONPrint() *c.Char {
	return nil
}

/* Render a cJSON entity to text for transfer/storage without any formatting. */
// llgo:link (*CJSON).CJSONPrintUnformatted C.cJSON_PrintUnformatted
func (recv_ *CJSON) CJSONPrintUnformatted() *c.Char {
	return nil
}

/* Render a cJSON entity to text using a buffered strategy. prebuffer is a guess at the final size. guessing well reduces reallocation. fmt=0 gives unformatted, =1 gives formatted */
// llgo:link (*CJSON).CJSONPrintBuffered C.cJSON_PrintBuffered
func (recv_ *CJSON) CJSONPrintBuffered(prebuffer c.Int, fmt CJSONBool) *c.Char {
	return nil
}

/* Render a cJSON entity to text using a buffer already allocated in memory with given length. Returns 1 on success and 0 on failure. */
/* NOTE: cJSON is not always 100% accurate in estimating how much memory it will use, so to be safe allocate 5 bytes more than you actually need */
// llgo:link (*CJSON).CJSONPrintPreallocated C.cJSON_PrintPreallocated
func (recv_ *CJSON) CJSONPrintPreallocated(buffer *c.Char, length c.Int, format CJSONBool) CJSONBool {
	return 0
}

/* Delete a cJSON entity and all subentities. */
// llgo:link (*CJSON).CJSONDelete C.cJSON_Delete
func (recv_ *CJSON) CJSONDelete() {
}

/* Returns the number of items in an array (or object). */
// llgo:link (*CJSON).CJSONGetArraySize C.cJSON_GetArraySize
func (recv_ *CJSON) CJSONGetArraySize() c.Int {
	return 0
}

/* Retrieve item number "index" from array "array". Returns NULL if unsuccessful. */
// llgo:link (*CJSON).CJSONGetArrayItem C.cJSON_GetArrayItem
func (recv_ *CJSON) CJSONGetArrayItem(index c.Int) *CJSON {
	return nil
}

/* Get item "string" from object. Case insensitive. */
// llgo:link (*CJSON).CJSONGetObjectItem C.cJSON_GetObjectItem
func (recv_ *CJSON) CJSONGetObjectItem(string *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONGetObjectItemCaseSensitive C.cJSON_GetObjectItemCaseSensitive
func (recv_ *CJSON) CJSONGetObjectItemCaseSensitive(string *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONHasObjectItem C.cJSON_HasObjectItem
func (recv_ *CJSON) CJSONHasObjectItem(string *c.Char) CJSONBool {
	return 0
}

/* For analysing failed parses. This returns a pointer to the parse error. You'll probably need to look a few chars back to make sense of it. Defined when cJSON_Parse() returns 0. 0 when cJSON_Parse() succeeds. */
//go:linkname CJSONGetErrorPtr C.cJSON_GetErrorPtr
func CJSONGetErrorPtr() *c.Char

/* Check item type and return its value */
// llgo:link (*CJSON).CJSONGetStringValue C.cJSON_GetStringValue
func (recv_ *CJSON) CJSONGetStringValue() *c.Char {
	return nil
}

// llgo:link (*CJSON).CJSONGetNumberValue C.cJSON_GetNumberValue
func (recv_ *CJSON) CJSONGetNumberValue() c.Double {
	return 0
}

/* These functions check the type of an item */
// llgo:link (*CJSON).CJSONIsInvalid C.cJSON_IsInvalid
func (recv_ *CJSON) CJSONIsInvalid() CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONIsFalse C.cJSON_IsFalse
func (recv_ *CJSON) CJSONIsFalse() CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONIsTrue C.cJSON_IsTrue
func (recv_ *CJSON) CJSONIsTrue() CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONIsBool C.cJSON_IsBool
func (recv_ *CJSON) CJSONIsBool() CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONIsNull C.cJSON_IsNull
func (recv_ *CJSON) CJSONIsNull() CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONIsNumber C.cJSON_IsNumber
func (recv_ *CJSON) CJSONIsNumber() CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONIsString C.cJSON_IsString
func (recv_ *CJSON) CJSONIsString() CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONIsArray C.cJSON_IsArray
func (recv_ *CJSON) CJSONIsArray() CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONIsObject C.cJSON_IsObject
func (recv_ *CJSON) CJSONIsObject() CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONIsRaw C.cJSON_IsRaw
func (recv_ *CJSON) CJSONIsRaw() CJSONBool {
	return 0
}

/* These calls create a cJSON item of the appropriate type. */
//go:linkname CJSONCreateNull C.cJSON_CreateNull
func CJSONCreateNull() *CJSON

//go:linkname CJSONCreateTrue C.cJSON_CreateTrue
func CJSONCreateTrue() *CJSON

//go:linkname CJSONCreateFalse C.cJSON_CreateFalse
func CJSONCreateFalse() *CJSON

// llgo:link CJSONBool.CJSONCreateBool C.cJSON_CreateBool
func (recv_ CJSONBool) CJSONCreateBool() *CJSON {
	return nil
}

//go:linkname CJSONCreateNumber C.cJSON_CreateNumber
func CJSONCreateNumber(num c.Double) *CJSON

//go:linkname CJSONCreateString C.cJSON_CreateString
func CJSONCreateString(string *c.Char) *CJSON

/* raw json */
//go:linkname CJSONCreateRaw C.cJSON_CreateRaw
func CJSONCreateRaw(raw *c.Char) *CJSON

//go:linkname CJSONCreateArray C.cJSON_CreateArray
func CJSONCreateArray() *CJSON

//go:linkname CJSONCreateObject C.cJSON_CreateObject
func CJSONCreateObject() *CJSON

/* Create a string where valuestring references a string so
 * it will not be freed by cJSON_Delete */
//go:linkname CJSONCreateStringReference C.cJSON_CreateStringReference
func CJSONCreateStringReference(string *c.Char) *CJSON

/* Create an object/array that only references it's elements so
 * they will not be freed by cJSON_Delete */
// llgo:link (*CJSON).CJSONCreateObjectReference C.cJSON_CreateObjectReference
func (recv_ *CJSON) CJSONCreateObjectReference() *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONCreateArrayReference C.cJSON_CreateArrayReference
func (recv_ *CJSON) CJSONCreateArrayReference() *CJSON {
	return nil
}

/* These utilities create an Array of count items.
 * The parameter count cannot be greater than the number of elements in the number array, otherwise array access will be out of bounds.*/
//go:linkname CJSONCreateIntArray C.cJSON_CreateIntArray
func CJSONCreateIntArray(numbers *c.Int, count c.Int) *CJSON

//go:linkname CJSONCreateFloatArray C.cJSON_CreateFloatArray
func CJSONCreateFloatArray(numbers *c.Float, count c.Int) *CJSON

//go:linkname CJSONCreateDoubleArray C.cJSON_CreateDoubleArray
func CJSONCreateDoubleArray(numbers *c.Double, count c.Int) *CJSON

//go:linkname CJSONCreateStringArray C.cJSON_CreateStringArray
func CJSONCreateStringArray(strings **c.Char, count c.Int) *CJSON

/* Append item to the specified array/object. */
// llgo:link (*CJSON).CJSONAddItemToArray C.cJSON_AddItemToArray
func (recv_ *CJSON) CJSONAddItemToArray(item *CJSON) CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONAddItemToObject C.cJSON_AddItemToObject
func (recv_ *CJSON) CJSONAddItemToObject(string *c.Char, item *CJSON) CJSONBool {
	return 0
}

/* Use this when string is definitely const (i.e. a literal, or as good as), and will definitely survive the cJSON object.
 * WARNING: When this function was used, make sure to always check that (item->type & cJSON_StringIsConst) is zero before
 * writing to `item->string` */
// llgo:link (*CJSON).CJSONAddItemToObjectCS C.cJSON_AddItemToObjectCS
func (recv_ *CJSON) CJSONAddItemToObjectCS(string *c.Char, item *CJSON) CJSONBool {
	return 0
}

/* Append reference to item to the specified array/object. Use this when you want to add an existing cJSON to a new cJSON, but don't want to corrupt your existing cJSON. */
// llgo:link (*CJSON).CJSONAddItemReferenceToArray C.cJSON_AddItemReferenceToArray
func (recv_ *CJSON) CJSONAddItemReferenceToArray(item *CJSON) CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONAddItemReferenceToObject C.cJSON_AddItemReferenceToObject
func (recv_ *CJSON) CJSONAddItemReferenceToObject(string *c.Char, item *CJSON) CJSONBool {
	return 0
}

/* Remove/Detach items from Arrays/Objects. */
// llgo:link (*CJSON).CJSONDetachItemViaPointer C.cJSON_DetachItemViaPointer
func (recv_ *CJSON) CJSONDetachItemViaPointer(item *CJSON) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONDetachItemFromArray C.cJSON_DetachItemFromArray
func (recv_ *CJSON) CJSONDetachItemFromArray(which c.Int) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONDeleteItemFromArray C.cJSON_DeleteItemFromArray
func (recv_ *CJSON) CJSONDeleteItemFromArray(which c.Int) {
}

// llgo:link (*CJSON).CJSONDetachItemFromObject C.cJSON_DetachItemFromObject
func (recv_ *CJSON) CJSONDetachItemFromObject(string *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONDetachItemFromObjectCaseSensitive C.cJSON_DetachItemFromObjectCaseSensitive
func (recv_ *CJSON) CJSONDetachItemFromObjectCaseSensitive(string *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONDeleteItemFromObject C.cJSON_DeleteItemFromObject
func (recv_ *CJSON) CJSONDeleteItemFromObject(string *c.Char) {
}

// llgo:link (*CJSON).CJSONDeleteItemFromObjectCaseSensitive C.cJSON_DeleteItemFromObjectCaseSensitive
func (recv_ *CJSON) CJSONDeleteItemFromObjectCaseSensitive(string *c.Char) {
}

/* Update array items. */
// llgo:link (*CJSON).CJSONInsertItemInArray C.cJSON_InsertItemInArray
func (recv_ *CJSON) CJSONInsertItemInArray(which c.Int, newitem *CJSON) CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONReplaceItemViaPointer C.cJSON_ReplaceItemViaPointer
func (recv_ *CJSON) CJSONReplaceItemViaPointer(item *CJSON, replacement *CJSON) CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONReplaceItemInArray C.cJSON_ReplaceItemInArray
func (recv_ *CJSON) CJSONReplaceItemInArray(which c.Int, newitem *CJSON) CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONReplaceItemInObject C.cJSON_ReplaceItemInObject
func (recv_ *CJSON) CJSONReplaceItemInObject(string *c.Char, newitem *CJSON) CJSONBool {
	return 0
}

// llgo:link (*CJSON).CJSONReplaceItemInObjectCaseSensitive C.cJSON_ReplaceItemInObjectCaseSensitive
func (recv_ *CJSON) CJSONReplaceItemInObjectCaseSensitive(string *c.Char, newitem *CJSON) CJSONBool {
	return 0
}

/* Duplicate a cJSON item */
// llgo:link (*CJSON).CJSONDuplicate C.cJSON_Duplicate
func (recv_ *CJSON) CJSONDuplicate(recurse CJSONBool) *CJSON {
	return nil
}

/* Duplicate will create a new, identical cJSON item to the one you pass, in new memory that will
 * need to be released. With recurse!=0, it will duplicate any children connected to the item.
 * The item->next and ->prev pointers are always zero on return from Duplicate. */
/* Recursively compare two cJSON items for equality. If either a or b is NULL or invalid, they will be considered unequal.
 * case_sensitive determines if object keys are treated case sensitive (1) or case insensitive (0) */
// llgo:link (*CJSON).CJSONCompare C.cJSON_Compare
func (recv_ *CJSON) CJSONCompare(b *CJSON, case_sensitive CJSONBool) CJSONBool {
	return 0
}

/* Minify a strings, remove blank characters(such as ' ', '\t', '\r', '\n') from strings.
 * The input pointer json cannot point to a read-only address area, such as a string constant,
 * but should point to a readable and writable address area. */
//go:linkname CJSONMinify C.cJSON_Minify
func CJSONMinify(json *c.Char)

/* Helper functions for creating and adding items to an object at the same time.
 * They return the added item or NULL on failure. */
// llgo:link (*CJSON).CJSONAddNullToObject C.cJSON_AddNullToObject
func (recv_ *CJSON) CJSONAddNullToObject(name *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONAddTrueToObject C.cJSON_AddTrueToObject
func (recv_ *CJSON) CJSONAddTrueToObject(name *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONAddFalseToObject C.cJSON_AddFalseToObject
func (recv_ *CJSON) CJSONAddFalseToObject(name *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONAddBoolToObject C.cJSON_AddBoolToObject
func (recv_ *CJSON) CJSONAddBoolToObject(name *c.Char, boolean CJSONBool) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONAddNumberToObject C.cJSON_AddNumberToObject
func (recv_ *CJSON) CJSONAddNumberToObject(name *c.Char, number c.Double) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONAddStringToObject C.cJSON_AddStringToObject
func (recv_ *CJSON) CJSONAddStringToObject(name *c.Char, string *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONAddRawToObject C.cJSON_AddRawToObject
func (recv_ *CJSON) CJSONAddRawToObject(name *c.Char, raw *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONAddObjectToObject C.cJSON_AddObjectToObject
func (recv_ *CJSON) CJSONAddObjectToObject(name *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONAddArrayToObject C.cJSON_AddArrayToObject
func (recv_ *CJSON) CJSONAddArrayToObject(name *c.Char) *CJSON {
	return nil
}

/* helper for the cJSON_SetNumberValue macro */
// llgo:link (*CJSON).CJSONSetNumberHelper C.cJSON_SetNumberHelper
func (recv_ *CJSON) CJSONSetNumberHelper(number c.Double) c.Double {
	return 0
}

/* Change the valuestring of a cJSON_String object, only takes effect when type of object is cJSON_String */
// llgo:link (*CJSON).CJSONSetValuestring C.cJSON_SetValuestring
func (recv_ *CJSON) CJSONSetValuestring(valuestring *c.Char) *c.Char {
	return nil
}

/* malloc/free objects using the malloc/free functions that have been set with cJSON_InitHooks */
//go:linkname CJSONMalloc C.cJSON_malloc
func CJSONMalloc(size c.SizeT) c.Pointer

//go:linkname CJSONFree C.cJSON_free
func CJSONFree(object c.Pointer)

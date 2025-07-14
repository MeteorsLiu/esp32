package json

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/* Implement RFC6901 (https://tools.ietf.org/html/rfc6901) JSON Pointer spec. */
// llgo:link (*CJSON).CJSONUtilsGetPointer C.cJSONUtils_GetPointer
func (recv_ *CJSON) CJSONUtilsGetPointer(pointer *c.Char) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONUtilsGetPointerCaseSensitive C.cJSONUtils_GetPointerCaseSensitive
func (recv_ *CJSON) CJSONUtilsGetPointerCaseSensitive(pointer *c.Char) *CJSON {
	return nil
}

/* Implement RFC6902 (https://tools.ietf.org/html/rfc6902) JSON Patch spec. */
/* NOTE: This modifies objects in 'from' and 'to' by sorting the elements by their key */
// llgo:link (*CJSON).CJSONUtilsGeneratePatches C.cJSONUtils_GeneratePatches
func (recv_ *CJSON) CJSONUtilsGeneratePatches(to *CJSON) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONUtilsGeneratePatchesCaseSensitive C.cJSONUtils_GeneratePatchesCaseSensitive
func (recv_ *CJSON) CJSONUtilsGeneratePatchesCaseSensitive(to *CJSON) *CJSON {
	return nil
}

/* Utility for generating patch array entries. */
// llgo:link (*CJSON).CJSONUtilsAddPatchToArray C.cJSONUtils_AddPatchToArray
func (recv_ *CJSON) CJSONUtilsAddPatchToArray(operation *c.Char, path *c.Char, value *CJSON) {
}

/* Returns 0 for success. */
// llgo:link (*CJSON).CJSONUtilsApplyPatches C.cJSONUtils_ApplyPatches
func (recv_ *CJSON) CJSONUtilsApplyPatches(patches *CJSON) c.Int {
	return 0
}

// llgo:link (*CJSON).CJSONUtilsApplyPatchesCaseSensitive C.cJSONUtils_ApplyPatchesCaseSensitive
func (recv_ *CJSON) CJSONUtilsApplyPatchesCaseSensitive(patches *CJSON) c.Int {
	return 0
}

/* Implement RFC7386 (https://tools.ietf.org/html/rfc7396) JSON Merge Patch spec. */
/* target will be modified by patch. return value is new ptr for target. */
// llgo:link (*CJSON).CJSONUtilsMergePatch C.cJSONUtils_MergePatch
func (recv_ *CJSON) CJSONUtilsMergePatch(patch *CJSON) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONUtilsMergePatchCaseSensitive C.cJSONUtils_MergePatchCaseSensitive
func (recv_ *CJSON) CJSONUtilsMergePatchCaseSensitive(patch *CJSON) *CJSON {
	return nil
}

/* generates a patch to move from -> to */
/* NOTE: This modifies objects in 'from' and 'to' by sorting the elements by their key */
// llgo:link (*CJSON).CJSONUtilsGenerateMergePatch C.cJSONUtils_GenerateMergePatch
func (recv_ *CJSON) CJSONUtilsGenerateMergePatch(to *CJSON) *CJSON {
	return nil
}

// llgo:link (*CJSON).CJSONUtilsGenerateMergePatchCaseSensitive C.cJSONUtils_GenerateMergePatchCaseSensitive
func (recv_ *CJSON) CJSONUtilsGenerateMergePatchCaseSensitive(to *CJSON) *CJSON {
	return nil
}

/* Given a root object and a target object, construct a pointer from one to the other. */
// llgo:link (*CJSON).CJSONUtilsFindPointerFromObjectTo C.cJSONUtils_FindPointerFromObjectTo
func (recv_ *CJSON) CJSONUtilsFindPointerFromObjectTo(target *CJSON) *c.Char {
	return nil
}

/* Sorts the members of the object into alphabetical order. */
// llgo:link (*CJSON).CJSONUtilsSortObject C.cJSONUtils_SortObject
func (recv_ *CJSON) CJSONUtilsSortObject() {
}

// llgo:link (*CJSON).CJSONUtilsSortObjectCaseSensitive C.cJSONUtils_SortObjectCaseSensitive
func (recv_ *CJSON) CJSONUtilsSortObjectCaseSensitive() {
}

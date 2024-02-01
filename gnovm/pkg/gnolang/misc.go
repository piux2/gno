package gnolang

import (
	"fmt"
	"strings"
	"unicode"
	"unsafe"

	"github.com/gnolang/gno/tm2/pkg/crypto"
)

//----------------------------------------
// Misc.

func cp(bz []byte) (ret []byte) {
	ret = make([]byte, len(bz))
	copy(ret, bz)
	return ret
}

// Returns the associated machine operation for binary AST operations.  TODO:
// to make this faster and inlineable, remove the switch statement and create a
// mathematical mapping between them.
func word2BinaryOp(w Word) Op {
	switch w {
	case ADD:
		return OpAdd
	case SUB:
		return OpSub
	case MUL:
		return OpMul
	case QUO:
		return OpQuo
	case REM:
		return OpRem
	case BAND:
		return OpBand
	case BOR:
		return OpBor
	case XOR:
		return OpXor
	case SHL:
		return OpShl
	case SHR:
		return OpShr
	case BAND_NOT:
		return OpBandn
	case LAND:
		return OpLand
	case LOR:
		return OpLor
	case EQL:
		return OpEql
	case LSS:
		return OpLss
	case GTR:
		return OpGtr
	case NEQ:
		return OpNeq
	case LEQ:
		return OpLeq
	case GEQ:
		return OpGeq
	default:
		panic(fmt.Sprintf("unexpected binary operation word %v", w.String()))
	}
}

func word2UnaryOp(w Word) Op {
	switch w {
	case ADD:
		return OpUpos
	case SUB:
		return OpUneg
	case NOT:
		return OpUnot
	case XOR:
		return OpUxor
	case MUL:
		panic("unexpected unary operation * - use StarExpr instead")
	case BAND:
		panic("unexpected unary operation & - use RefExpr instead")
	case ARROW:
		return OpUrecv
	default:
		panic("unexpected unary operation")
	}
}

func toString(n Node) string {
	if n == nil {
		return "<nil>"
	}
	return n.String()
}

// true if the first rune is uppercase.
func isUpper(s string) bool {
	var first rune
	for _, c := range s {
		first = c
		break
	}
	return unicode.IsUpper(first)
}

//----------------------------------------
// converting uintptr to bytes.

const sizeOfUintPtr = unsafe.Sizeof(uintptr(0))

func uintptrToBytes(u *uintptr) []byte {
	return (*[sizeOfUintPtr]byte)(unsafe.Pointer(u))[:]
}

func defaultPkgName(gopkgPath string) Name {
	parts := strings.Split(gopkgPath, "/")
	last := parts[len(parts)-1]
	parts = strings.Split(last, "-")
	name := parts[len(parts)-1]
	name = strings.ToLower(name)
	return Name(name)
}

//----------------------------------------
// value convenience

func toTypeValue(t Type) TypeValue {
	return TypeValue{
		Type: t,
	}
}

//----------------------------------------
// reserved & uverse names

var reservedNames = map[Name]struct{}{
	"break": {}, "default": {}, "func": {}, "interface": {}, "select": {},
	"case": {}, "defer": {}, "go": {}, "map": {}, "struct": {},
	"chan": {}, "else": {}, "goto": {}, "package": {}, "switch": {},
	"const": {}, "fallthrough": {}, "if": {}, "range": {}, "type": {},
	"continue": {}, "for": {}, "import": {}, "return": {}, "var": {},
}

// if true, caller should generally panic.
func isReservedName(n Name) bool {
	_, ok := reservedNames[n]
	return ok
}

// scans uverse static node for blocknames. (slow)
func isUverseName(n Name) bool {
	uverseNames := UverseNode().GetBlockNames()
	for _, name := range uverseNames {
		if name == n {
			return true
		}
	}
	return false
}

//----------------------------------------
// other

// For keeping record of package & realm coins.
func DerivePkgAddr(pkgPath string) crypto.Address {
	// NOTE: must not collide with pubkey addrs.
	return crypto.AddressFromPreimage([]byte("pkgPath:" + pkgPath))
}

var opToStringMap = map[Op]string{
	/* Control operators */
	OpInvalid:             "OpInvalid",
	OpHalt:                "OpHalt",
	OpNoop:                "OpNoop",
	OpExec:                "OpExec",
	OpPrecall:             "OpPrecall",
	OpCall:                "OpCall",
	OpCallNativeBody:      "OpCallNativeBody",
	OpReturn:              "OpReturn",
	OpReturnFromBlock:     "OpReturnFromBlock",
	OpReturnToBlock:       "OpReturnToBlock",
	OpDefer:               "OpDefer",
	OpCallDeferNativeBody: "OpCallDeferNativeBody",
	OpGo:                  "OpGo",
	OpSelect:              "OpSelect",
	OpSwitchClause:        "OpSwitchClause",
	OpSwitchClauseCase:    "OpSwitchClauseCase",
	OpTypeSwitch:          "OpTypeSwitch",
	OpIfCond:              "OpIfCond",
	OpPopValue:            "OpPopValue",
	OpPopResults:          "OpPopResults",
	OpPopBlock:            "OpPopBlock",
	OpPopFrameAndReset:    "OpPopFrameAndReset",
	OpPanic1:              "OpPanic1",
	OpPanic2:              "OpPanic2",

	/* Unary & binary operators */
	OpUpos:  "OpUpos",
	OpUneg:  "OpUneg",
	OpUnot:  "OpUnot",
	OpUxor:  "OpUxor",
	OpUrecv: "OpUrecv",
	OpLor:   "OpLor",
	OpLand:  "OpLand",
	OpEql:   "OpEql",
	OpNeq:   "OpNeq",
	OpLss:   "OpLss",
	OpLeq:   "OpLeq",
	OpGtr:   "OpGtr",
	OpGeq:   "OpGeq",
	OpAdd:   "OpAdd",
	OpSub:   "OpSub",
	OpBor:   "OpBor",
	OpXor:   "OpXor",
	OpMul:   "OpMul",
	OpQuo:   "OpQuo",
	OpRem:   "OpRem",
	OpShl:   "OpShl",
	OpShr:   "OpShr",
	OpBand:  "OpBand",
	OpBandn: "OpBandn",

	/* Other expression operators */
	OpEval:         "OpEval",
	OpBinary1:      "OpBinary1",
	OpIndex1:       "OpIndex1",
	OpIndex2:       "OpIndex2",
	OpSelector:     "OpSelector",
	OpSlice:        "OpSlice",
	OpStar:         "OpStar",
	OpRef:          "OpRef",
	OpTypeAssert1:  "OpTypeAssert1",
	OpTypeAssert2:  "OpTypeAssert2",
	OpStaticTypeOf: "OpStaticTypeOf",
	OpCompositeLit: "OpCompositeLit",
	OpArrayLit:     "OpArrayLit",
	OpSliceLit:     "OpSliceLit",
	OpSliceLit2:    "OpSliceLit2",
	OpMapLit:       "OpMapLit",
	OpStructLit:    "OpStructLit",
	OpFuncLit:      "OpFuncLit",
	OpConvert:      "OpConvert",

	/* Native operators */
	OpArrayLitGoNative:  "OpArrayLitGoNative",
	OpSliceLitGoNative:  "OpSliceLitGoNative",
	OpStructLitGoNative: "OpStructLitGoNative",
	OpCallGoNative:      "OpCallGoNative",

	/* Type operators */
	OpFieldType:       "OpFieldType",
	OpArrayType:       "OpArrayType",
	OpSliceType:       "OpSliceType",
	OpPointerType:     "OpPointerType",
	OpInterfaceType:   "OpInterfaceType",
	OpChanType:        "OpChanType",
	OpFuncType:        "OpFuncType",
	OpMapType:         "OpMapType",
	OpStructType:      "OpStructType",
	OpMaybeNativeType: "OpMaybeNativeType",

	/* Statement operators */
	OpAssign:      "OpAssign",
	OpAddAssign:   "OpAddAssign",
	OpSubAssign:   "OpSubAssign",
	OpMulAssign:   "OpMulAssign",
	OpQuoAssign:   "OpQuoAssign",
	OpRemAssign:   "OpRemAssign",
	OpBandAssign:  "OpBandAssign",
	OpBandnAssign: "OpBandnAssign",
	OpBorAssign:   "OpBorAssign",
	OpXorAssign:   "OpXorAssign",
	OpShlAssign:   "OpShlAssign",
	OpShrAssign:   "OpShrAssign",
	OpDefine:      "OpDefine",
	OpInc:         "OpInc",
	OpDec:         "OpDec",

	/* Decl operators */
	OpValueDecl: "OpValueDecl",
	OpTypeDecl:  "OpTypeDecl",

	/* Loop (sticky) operators (>= 0xD0) */
	OpSticky:            "OpSticky",
	OpBody:              "OpBody",
	OpForLoop:           "OpForLoop",
	OpRangeIter:         "OpRangeIter",
	OpRangeIterString:   "OpRangeIterString",
	OpRangeIterMap:      "OpRangeIterMap",
	OpRangeIterArrayPtr: "OpRangeIterArrayPtr",
	OpReturnCallDefers:  "OpReturnCallDefers",
}

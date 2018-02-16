// Package MapAST is an abstract syntax tree for the go language.
package mapast

// MaxSubnodes - The current limit of how many certain type subnodes can some
// nodes have.
const MaxSubnodes = 1000000

var storage = make([]byte, MaxSubnodes)

// RootMatter node is equivalent to a package block or an universe block
// containing all Go source text contained within the map. There is only one
// RootMatter located at key zero (0).
var RootMatter = []byte("RootMatter")

// FileMatter contains the file block elements. It represents a single go file.
// FileMatter node is a child of RootMatter node.
var FileMatter = []byte("FileMatter")

// PackageDef is a child of FileMatter. Contains a string holding the file
// package name, and an optional CommentRow containing the import comment.
var PackageDef = []byte("PackageDef")

// ImportStmt holds a single import declaration. It is a child of ImportsDef or
// of a FileMatter. If the parent is FileMatter, it is a standalone import
// declaration. ImportStmt contains one or two strings.
var ImportStmt = []byte("ImportStmt")

// ImportsDef is a bracketed container for multiple ImportStmt nodes.
// It is a child of FileMatter.
var ImportsDef = []byte("ImportsDef")

// TypedIdent field contains several string identifiers known as names, followed
// by a RootOfType node representing the type shared by the named identifiers.
// If contained within a StructType, it can be optionally tagged using the last
// string child (the tag).
var TypedIdent = []byte("TypedIdent")

// RootOfType marks the root of the type expression tree. It's only child is
// usually an Expression.
var RootOfType = []byte("RootOfType")

// TypDefStmt type declaration allows to declare an alias or a type. Bracketed
// form is currently not available.
var TypDefStmt = []byte("TypDefStmt")

// StructType is a sequence of named elements, called fields. Some fields can
// share their type, using a TypedIdent node.
var StructType = []byte("StructType")

// BranchStmt is a sole statement. One of semicolon, break, continue,
// fallthrough, goto. It has no children.
var BranchStmt = []byte("BranchStmt")

// GoDferStmt node is a statement node used to invoke a call Expression using
// the go or defer keyword. It's only child is the Expression to be invoked.
var GoDferStmt = []byte("GoDferStmt")

// ReturnStmt node is a return statement followed by it's children. Strings are
// not allowed, a string child must be wrapped by ExpressionIdentifier.
var ReturnStmt = []byte("ReturnStmt")

// IncDecStmt is an increment++ or decrement-- statement. It's only child is
// a string identifier or a more complex Expression.
var IncDecStmt = []byte("IncDecStmt")

// VarDefStmt is a standalone or a multiple variable or constant declaration.
// It's children are single or multiple AssignStmt nodes, one for each row.
var VarDefStmt = []byte("VarDefStmt")

// LblGotoCnt is a labeled statement, or a statement that uses label: break,
// continue or goto statement. It has one child, the string known as label name.
var LblGotoCnt = []byte("LblGotoCnt")

// IfceTypExp node contains one or several IfceMethod or RootOfType nodes.
var IfceTypExp = []byte("IfceTypExp")

// CommentRow contains exactly one string, holding the comment verbatim, with
// optional leading or trailing newlines.
var CommentRow = []byte("CommentRow")

// Expression node is one of the 38 differend kinds of Expression. It contains
// strings, IfceTypExps, ClosureExps, StructTypes or more Expressions.
var Expression = storage[0:]

// BlocOfCode node holds statements or other BlocOfCode nodes. Some BlocOfCode
// kinds have a header followed by the opening brace. Other BlocOfCode kinds lack
// braces altogether and use colon instead.
var BlocOfCode = storage[16:]

// ToplevFunc is a child function of FileMatter. The first child is a string.
// The rest of children can be TypedIdent nodes. The trailing child is
// an optional BlocOfCode.
var ToplevFunc = storage[32:]

// AssignStmt contains left hand side entries followed by an optional RootOfType
// and implicit equality kind operator, followed by a right hand side entries.
var AssignStmt = storage[48:]

// ClosureExp is a function literal. The children are TypedIdent nodes.
var ClosureExp = storage[64:]

// IfceMethod node is a child of IfceTypExp. It's children are TypedIdent nodes.
// The name of interface method is stored in the first TypedIdent child, the one
// that would otherwise work as a receiver field.
var IfceMethod = storage[80:]

// Constructor for ToplevFunc node. Argc is the count of proper arguments
// excluding results and receiver.
func ToplevFuncNode(receiver bool, argc uint64) []byte {
	var recv = 0
	if receiver {
		recv = 1
	}
	return ToplevFunc[:recv+1 : int(argc)+recv+1]
}

// Constructor for BlocOfCode node. Headelemscount is the count of header
// elements including semicolons.
func BlocOfCodeNode(kind byte, headelemscount uint64) []byte {
	return BlocOfCode[:int(kind)+1 : int(headelemscount)+int(BlocOfCodeTotalCount)]
}

// Constructor for Expression node. Elemscount is the number of children elements.
func ExpressionNode(kind byte, elemscount uint64) []byte {
	return Expression[:int(kind)+1 : int(elemscount+uint64(ExpressionTotalCount))]
}

// Constructor for BranchStmt node.
func BranchStmtNode(kind byte) []byte {
	return BranchStmt[:int(kind)+1]
}

// Constructor for IncDecStmt node.
func IncDecStmtNode(kind byte) []byte {
	return IncDecStmt[:int(kind)+1]
}

// Constructor for AssignStmt node. Elemscount is the number of elements
// including both side elements and the central type node (if any).
func AssignStmtNode(kind byte, elemscount uint64) []byte {
	return AssignStmt[:int(kind)+1 : int(elemscount+uint64(AssignStmtTotalCount))]
}

// Constructor for ClosureExp node. Paramscount is the number of parameters.
func ClosureExpNode(paramscount uint64) []byte {
	return ClosureExp[:int(paramscount)+1]
}

// Constructor for GoDferStmt node.
func GoDferStmtNode(kind byte) []byte {
	return GoDferStmt[:int(kind)+1]
}

// Constructor for LblGotoCnt node.
func LblGotoCntNode(kind byte) []byte {
	return LblGotoCnt[:int(kind)+1]
}

// Constructor for VarDefStmt node.
func VarDefStmtNode(kind byte) []byte {
	return VarDefStmt[:int(kind)+1]
}

// Constructor for TypDefStmt node.
func TypDefStmtNode(kind byte) []byte {
	return TypDefStmt[:int(kind)+1]
}

// BlocOfCodePlain is a plain code block. Child of BlocOfCode, ToplevFunc or
// ClosureExp. Has zero header entries.
const BlocOfCodePlain byte = 0

// BlocOfCodeIf is an if code block. It may occur right after BlocOfCodeIfElse,
// forming an if else if sequence of blocks. Has 1 or 3 header entries.
const BlocOfCodeIf byte = 1

// BlocOfCodeIfElse. Must be followed by BlocOfCodePlain or BlocOfCodeIf node.
const BlocOfCodeIfElse byte = 2

// BlocOfCodeSwitch is a non-type switch. Can contain only BlocOfCodeCase and
// BlocOfCodeDefault nodes.
const BlocOfCodeSwitch byte = 3

// BlocOfCodeFor is "for" or "for something range something" loop. The range
// (if any) is part of AssignStmt child.
const BlocOfCodeFor byte = 4

// BlocOfCodeForRange is a for range loop. This loop code begins with "for
// range".
const BlocOfCodeForRange byte = 5

// BlocOfCodeTypeSwitch switches using a .(type) header Expression.
const BlocOfCodeTypeSwitch byte = 6

// BlocOfCodeSelect contains no block header elements. The children are
// BlocOfCodeCommunicate or BlocOfCodeCommunicateDefault nodes.
const BlocOfCodeSelect byte = 7

// BlocOfCodeCase contains Expression header elements that will appear comma
// separated.
const BlocOfCodeCase byte = 8

// BlocOfCodeDefault is a default clause, child of BlocOfCodeSwitch or
// a BlocOfCodeTypeSwitch. It has zero header elements.
const BlocOfCodeDefault byte = 9

// BlocOfCodeNone is unused, but contains elements but no brackets.
const BlocOfCodeNone byte = 10

// BlocOfCodeCommunicate is a communicate clause. It has exactly one Expression
// or AssignStmt node in header. Child of BlocOfCodeSelect.
const BlocOfCodeCommunicate byte = 11

// BlocOfCodeCommunicateDefault is a communicate default clause. It has no
// header entries. Child of BlocOfCodeSelect.
const BlocOfCodeCommunicateDefault byte = 12

// BlocOfCodeTotalCount is a total count sentinel. Do not use.
const BlocOfCodeTotalCount byte = 13

// BranchStmtSemi is an explicit semicolon. Child of a BlocOfCode node.
const BranchStmtSemi byte = 0

// BranchStmtSemi is a break statement. Child of a BlocOfCode node. Does not
// refer to a label.
const BranchStmtBreak byte = 1

// BranchStmtContinue is a continue statement. Child of a BlocOfCode node. Does
// not refer to a label.
const BranchStmtContinue byte = 2

// BranchStmtFallthrough is a fallthrough statement. Child of a BlocOfCode node.
const BranchStmtFallthrough byte = 3

// BranchStmtGoto is a goto statement. Child of a BlocOfCode node. Does not
// refer to a label.
const BranchStmtGoto byte = 4

// GoDferStmtGo is a go statement used to invoke a call Expression using
// the go or defer keyword. It's only child is the Expression to be invoked.
const GoDferStmtGo byte = 0

// GoDferStmtDefer is a defer statement used to invoke a call Expression using
// the go or defer keyword. It's only child is the Expression to be invoked.
const GoDferStmtDefer byte = 1

// ExpressionBrackets has exactly one child that appears in a round brackets.
const ExpressionBrackets byte = 0

// ExpressionOrOr is a boolean or operator.
const ExpressionOrOr byte = 1

// ExpressionAndAnd is a boolean and operator.
const ExpressionAndAnd byte = 2

// ExpressionEqual is an equality operator.
const ExpressionEqual byte = 3

// ExpressionNotEq is an inequality operator.
const ExpressionNotEq byte = 4

// ExpressionLessThan is a numeric less than operator.
const ExpressionLessThan byte = 5

// ExpressionLessEq is a numeric less than or equal operator.
const ExpressionLessEq byte = 6

// ExpressionGrtEq is a numeric greater than or equal operator.
const ExpressionGrtEq byte = 7

// ExpressionGrtThan is a numeric greater than operator.
const ExpressionGrtThan byte = 8

// ExpressionPlus is a numeric plus operator.
const ExpressionPlus byte = 9

// ExpressionMinus is a numeric minus operator.
const ExpressionMinus byte = 10

// ExpressionOr is a bitwise or operator.
const ExpressionOr byte = 11

// ExpressionXor is a bitwise xor operator.
const ExpressionXor byte = 12

// ExpressionMul is a numeric multiply operator.
const ExpressionMul byte = 13

// ExpressionDiv is a numeric divide operator.
const ExpressionDiv byte = 14

// ExpressionMod is a numeric remainder operator.
const ExpressionMod byte = 15

// ExpressionAnd is a bitwise and operator.
const ExpressionAnd byte = 16

// ExpressionAndNot is a bitwise and not operator.
const ExpressionAndNot byte = 17

// ExpressionLSh is a numeric left shift operator.
const ExpressionLSh byte = 18

// ExpressionRSh is a numeric right shift operator.
const ExpressionRSh byte = 19

// ExpressionAndAnd is a boolean not operator.
const ExpressionNot byte = 20

// ExpressionDot is a dot separated selector expression or a qualified
// identifier.
const ExpressionDot byte = 21

// ExpressionSlice is a slice expression.
const ExpressionSlice byte = 22

// ExpressionComposite is a composite literal.
const ExpressionComposite byte = 23

// ExpressionCall is a call or method call expression, it is not variadic.
const ExpressionCall byte = 24

// ExpressionArrow is a send statement or an unary receive operation.
const ExpressionArrow byte = 25

// ExpressionArrayType is an array type.
const ExpressionArrayType byte = 26

// ExpressionSliceType is a slice type.
const ExpressionSliceType byte = 27

// ExpressionKeyVal is a key value expression.
const ExpressionKeyVal byte = 28

// ExpressionType is a binary type assertion. If unary, it asserts the reserved
// word type.
const ExpressionType byte = 29

// ExpressionCall is a variadic call or method call expression.
const ExpressionCallDotDotDot byte = 30

// ExpressionComposed is a Literal Value nested in a composite literal.
const ExpressionComposed byte = 31

// ExpressionIndex is an index expression.
const ExpressionIndex byte = 32

// ExpressionMap is a map type
const ExpressionMap byte = 33

// ExpressionIdentifier is a string wrapper. It has one string child node.
const ExpressionIdentifier byte = 34

// ExpressionChan is a channel type without a direction arrow operator.
const ExpressionChan byte = 35

// ExpressionInChan is a channel type with a send direction arrow operator.
const ExpressionInChan byte = 36

// ExpressionOutChan is a channel type with a receive direction arrow operator.
const ExpressionOutChan byte = 37

// ExpressionTotalCount is a total count sentinel. Do not use.
const ExpressionTotalCount byte = 38

// IncDecStmtPlusPlus is an IncDec statement. It's the increment ++ statement.
const IncDecStmtPlusPlus byte = 0

// IncDecStmtMinusMinus is an IncDec statement. It's the decrement -- statement.
const IncDecStmtMinusMinus byte = 1

// AssignStmtEqual is an assignment without an assignment operation.
const AssignStmtEqual byte = 0

// AssignStmtColonEq is a short variable declaration
const AssignStmtColonEq byte = 1

// AssignStmtAndNot is an assignment with an and not assignment operation.
const AssignStmtAndNot byte = 2

// AssignStmtAdd is an assignment with an add assignment operation.
const AssignStmtAdd byte = 3

// AssignStmtSub is an assignment with a subtract assignment operation.
const AssignStmtSub byte = 4

// AssignStmtMul is an assignment with a multiply assignment operation.
const AssignStmtMul byte = 5

// AssignStmtQuo is an assignment with a division assignment operation.
const AssignStmtQuo byte = 6

// AssignStmtRem is an assignment with a remainder assignment operation.
const AssignStmtRem byte = 7

// AssignStmtAnd is an assignment with a bitwise and assignment operation.
const AssignStmtAnd byte = 8

// AssignStmtOr is an assignment with a bitwise or assignment operation.
const AssignStmtOr byte = 9

// AssignStmtXor is an assignment with a bitwise xor assignment operation.
const AssignStmtXor byte = 10

// AssignStmtShl is an assignment with a left shift assignment operation.
const AssignStmtShl byte = 11

// AssignStmtShr is an assignment with a right shift assignment operation.
const AssignStmtShr byte = 12

// AssignStmtIotaIsLast is used for constant declaration for lines that follow
// the iota row.
const AssignStmtIotaIsLast byte = 13

// AssignStmtTypeIsLast is used in variable declaration giving the variables
// a type but no initial values are assigned.
const AssignStmtTypeIsLast byte = 14

// AssignStmtMoreEqual is an assignment without an assignment operation. It has
// one expression on the right hand side.
const AssignStmtMoreEqual byte = 15

// AssignStmtMoreColonEq is a short variable declaration with one expression on
// the right hand side.
const AssignStmtMoreColonEq byte = 16

// AssignStmtMoreEqualRange is an assignment without an assignment operation. It
//has one expression on the right hand side. Used within a range loop.
const AssignStmtMoreEqualRange byte = 17

// AssignStmtMoreColonEqRange is a short variable declaration with one
// expression on the right hand side. Used within a range loop.
const AssignStmtMoreColonEqRange byte = 18

// AssignStmtTotalCount is a total count sentinel. Do not use.
const AssignStmtTotalCount byte = 19

// VarDefStmtVar is a standalone or a multiple variable declaration.
const VarDefStmtVar byte = 0

// VarDefStmtConst is a standalone or a multiple constant declaration.
const VarDefStmtConst byte = 1

// VarDefStmtTotalCount is a total count sentinel. Do not use.
const VarDefStmtTotalCount byte = 2

// LblGotoCntLabel is a labeled statement followed by a colon.
const LblGotoCntLabel byte = 0

// LblGotoCntGoto is a goto followed by a label.
const LblGotoCntGoto byte = 1

// LblGotoCntContinue is a continue followed by a label.
const LblGotoCntContinue byte = 2

// LblGotoCntBreak is a break followed by a label.
const LblGotoCntBreak byte = 3

// TypedIdentNormal is the names followed by type.
const TypedIdentNormal byte = 0

// TypedIdentEquals is the names followed by an equal sign followed by a type.
const TypedIdentEquals byte = 1

// TypedIdentEllipsis is the names followed by an ellipsis followed by a type.
const TypedIdentEllipsis byte = 2

// TypedIdentTagged is the names followed by type followed by a string tag.
const TypedIdentTagged byte = 3

// TypDefStmtNormal contains type definition.
const TypDefStmtNormal byte = 0

// TypDefStmtAlias contains alias declaration.
const TypDefStmtAlias byte = 1

// CommentRowEnder is a comment starting at the end of a previous element line.
const CommentRowEnder byte = 0

// CommentRowNormal is a comment occupying it's own row.
const CommentRowNormal byte = 1

// CommentRowSeparate is a comment following an empty line(s).
const CommentRowSeparate byte = 2

// PackageDefNormal is a package statement not following an empty line(s).
const PackageDefNormal byte = 0

// PackageDefSeparate is a package statement separated by an empty line(s).
const PackageDefSeparate byte = 1

// O is an one way function. Given a node key it calculates the key of its first
// child node. The other keys of child nodes follow by adding 1, 2, 3... to
// the result.
func O(n uint64) uint64 {
	var v = n*3935559000370003845 + 2691343689449507681
	v ^= v >> 21
	v ^= v << 37
	v ^= v >> 4
	v *= 4768777513237032717
	v ^= v << 20
	v ^= v >> 41
	v ^= v << 5
	return v
}

// The Poke function tests whether a given iterator key in an ast is occupied.
func Poke(ast map[uint64][]byte, iterator uint64) bool {
	_, ok := ast[iterator]
	return ok
}

// Which determines which node a given byte slice represents. It resets capacity
// and the length to the maximum capacity and length. If node is a string, Which
// returns nil.
func Which(node []byte) []byte {
	if node == nil {
		return nil
	} else if &node[0] == &RootMatter[0] {
		return RootMatter
	} else if &node[0] == &FileMatter[0] {
		return FileMatter
	} else if &node[0] == &PackageDef[0] {
		return PackageDef
	} else if &node[0] == &ImportStmt[0] {
		return ImportStmt
	} else if &node[0] == &ImportsDef[0] {
		return ImportsDef
	} else if &node[0] == &ToplevFunc[0] {
		return ToplevFunc
	} else if &node[0] == &TypedIdent[0] {
		return TypedIdent
	} else if &node[0] == &RootOfType[0] {
		return RootOfType
	} else if &node[0] == &BlocOfCode[0] {
		return BlocOfCode
	} else if &node[0] == &TypDefStmt[0] {
		return TypDefStmt
	} else if &node[0] == &StructType[0] {
		return StructType
	} else if &node[0] == &BranchStmt[0] {
		return BranchStmt
	} else if &node[0] == &GoDferStmt[0] {
		return GoDferStmt
	} else if &node[0] == &Expression[0] {
		return Expression
	} else if &node[0] == &ReturnStmt[0] {
		return ReturnStmt
	} else if &node[0] == &IncDecStmt[0] {
		return IncDecStmt
	} else if &node[0] == &AssignStmt[0] {
		return AssignStmt
	} else if &node[0] == &VarDefStmt[0] {
		return VarDefStmt
	} else if &node[0] == &LblGotoCnt[0] {
		return LblGotoCnt
	} else if &node[0] == &ClosureExp[0] {
		return ClosureExp
	} else if &node[0] == &IfceTypExp[0] {
		return IfceTypExp
	} else if &node[0] == &IfceMethod[0] {
		return IfceMethod
	} else if &node[0] == &CommentRow[0] {
		return CommentRow
	}
	return nil
}

// Printer is used to print strings to standard error. Empty strings are printed
// as if they were newlines.
func Printer(s string) {
	if len(s) == 0 {
		println("")
	} else {
		print(s)
	}
}

func itoA(n int) string {
	if n > 9999 {
		return "many"
	}
	var a = [4]byte{'0', '0', '0', '0'}
	a[3] += byte(n % 10)
	a[2] += byte((n / 10) % 10)
	a[1] += byte((n / 100) % 10)
	a[0] += byte((n / 1000) % 10)
	if a[0] == '0' {
		if a[1] == '0' {
			if a[2] == '0' {
				return string(a[3:])
			}
			return string(a[2:])
		}
		return string(a[1:])
	}
	return string(a[:])
}

// PrintDump prints an ast dump. This a is xml like output that can be used to
// debug to see if the ast is correct. Only nested nodes under iterator position
// are printed.
func PrintDump(ast map[uint64][]byte, iterator uint64, pad int) bool {
	return Dump(Printer, ast, iterator, pad)
}

// Dump prints an ast dump. This is a xml like output that can be used to
// debug to see if the ast is correct. Only nested nodes under iterator position
// are printed.
func Dump(print func(string), ast map[uint64][]byte, iterator uint64, pad int) bool {
	if pad > 50 {
		pad = 50
	}
	print("                                                      "[:pad])
	nod, ok := ast[iterator]
	if !ok {
		print("[-]")
		print("")
		return false
	} else if nod == nil {
		print(" [nil]")
		print("")
	} else if Which(ast[iterator]) != nil {
		print(" [")
		switch len(Which(ast[iterator])) {
		case 10:
			print(string(Which(ast[iterator])))

		case MaxSubnodes:
			print("Expression")

		case MaxSubnodes - 16:
			print("BlocOfCode")

		case MaxSubnodes - 32:
			print("ToplevFunc")

		case MaxSubnodes - 48:
			print("AssignStmt")

		case MaxSubnodes - 64:
			print("ClosureExp")

		case MaxSubnodes - 80:
			print("IfceMethod")

		}
		print(" ")
		print(itoA(len(ast[iterator]) - 1))
		print(" ")
		print(itoA(cap(ast[iterator]) - 1))
		print("]")
		print("")
	} else {
		print(" [string " + string(nod) + "]")
		print("")
	}
	for i := uint64(0); true; i++ {
		if !Dump(print, ast, O(iterator)+i, pad+1) {
			return true
		}
	}
	return true
}

// PrintCode prints go source code from an abstract syntax tree to a standard
// error.
func PrintCode(ast map[uint64][]byte, iterator uint64, parent uint64) {
	Code(Printer, ast, iterator, parent)
}

// Code generates go source code from an abstract syntax tree.
func Code(print func(string), ast map[uint64][]byte, iterator uint64, parent uint64) {
	const uint64big = ^uint64(0) - 1
	var ast_o_iterator = string(ast[O(iterator)])
	if ast[iterator] != nil {
		switch &(ast[iterator])[0] {
		case &CommentRow[0]:
			if len(ast[(iterator)])-1 == int(CommentRowSeparate) {
				print("")
			}
			print(ast_o_iterator)

		case &PackageDef[0]:
			if len(ast[(iterator)])-1 == int(PackageDefSeparate) {
				print("")
			}
			print("package ")
			print(ast_o_iterator)

		case &ImportStmt[0]:
			var defparent = ast[(parent)] != nil && &ast[(parent)][0] == &ImportsDef[0]
			if defparent {
			} else {
				print("import ")
			}
			print(ast_o_iterator)
			another := string(ast[O(iterator)+1])
			if len(another) > 0 {
				print(" ")
				print(another)
			}
			if defparent {
				print("")
			}

		case &ImportsDef[0]:
			print("import (")
			print("")

		case &ToplevFunc[0]:
			print("func ")
			if len(ast[(iterator)]) == 2 {
				print("(")
			}

		case &BlocOfCode[0]:
			switch byte(len(ast[(iterator)]) - 1) {
			case BlocOfCodePlain:

			case BlocOfCodeIf:
				fallthrough

			case BlocOfCodeIfElse:
				print("if ")

			case BlocOfCodeFor:
				print("for ")

			case BlocOfCodeForRange:
				print("for range ")

			case BlocOfCodeSwitch:
				fallthrough

			case BlocOfCodeTypeSwitch:
				print("switch ")

			case BlocOfCodeSelect:
				print("select ")

			case BlocOfCodeCase:
				print("case ")

			case BlocOfCodeCommunicateDefault:
				fallthrough

			case BlocOfCodeDefault:
				print("default:")
				print("")

			case BlocOfCodeCommunicate:
				print("case ")

			case BlocOfCodeNone:

			}
			if cap(ast[(iterator)]) == int(BlocOfCodeTotalCount) {
				if len(ast[(iterator)])-1 < int(BlocOfCodeCase) {
					print("{")
					print("")
				}
			}

		case &RootOfType[0]:

		case &TypDefStmt[0]:
			print("type ")
			print(ast_o_iterator)
			print(" ")
			var op = byte(len(ast[(iterator)]) - 1)
			if op == TypDefStmtAlias {
				print("= ")
			}

		case &StructType[0]:
			print("struct{")
			if ast[O(iterator)] != nil {
				print("")
			}

		case &IfceTypExp[0]:
			print("interface{")
			if ast[O(iterator)] != nil {
				print("")
			}

		case &GoDferStmt[0]:
			switch byte(len(ast[(iterator)]) - 1) {
			case GoDferStmtGo:
				print("go ")

			case GoDferStmtDefer:
				print("defer ")

			}

		case &BranchStmt[0]:
			switch byte(len(ast[(iterator)]) - 1) {
			case BranchStmtSemi:
				print(";")

			case BranchStmtBreak:
				print("break")

			case BranchStmtContinue:
				print("continue")

			case BranchStmtFallthrough:
				print("fallthrough")

			case BranchStmtGoto:
				print("goto")

			}

		case &Expression[0]:
			var op = byte(len(ast[(iterator)]) - 1)
			var l = byte(cap(ast[(iterator)]) - int(ExpressionTotalCount))
			if l == 1 {
				switch op {
				case ExpressionBrackets:
					print("(")

				case ExpressionPlus:
					print("+")

				case ExpressionMinus:
					print("-")

				case ExpressionXor:
					print("^")

				case ExpressionMul:
					print("*")

				case ExpressionAnd:
					print("&")

				case ExpressionNot:
					print("!")

				case ExpressionArrow:
					print("<-")

				case ExpressionChan:
					print("chan ")

				case ExpressionInChan:
					print("chan<- ")

				case ExpressionOutChan:
					print("<-chan ")

				case ExpressionComposed:
					print("{")

				}
			} else {
				switch op {
				case ExpressionArrayType:
					print("[]")

				case ExpressionComposed:
					print("{")
					if l == 0 {
						print("}")
					}

				case ExpressionSliceType:
					print("[")

				case ExpressionMap:
					print("map[")

				}
			}

		case &ReturnStmt[0]:
			print("return ")

		case &VarDefStmt[0]:
			var op = byte(len(ast[(iterator)]) - 1)
			var multi = len(ast[O(iterator)+1]) > 0
			var none = len(ast[O(iterator)]) == 0
			switch op {
			case VarDefStmtVar:
				print("var ")

			case VarDefStmtConst:
				print("const ")

			}
			if multi {
				print("(")
				print("")
			} else if none {
				print("()")
			}

		case &LblGotoCnt[0]:
			var op = byte(len(ast[(iterator)]) - 1)
			switch op {
			case LblGotoCntGoto:
				print("goto ")

			case LblGotoCntContinue:
				print("continue ")

			case LblGotoCntBreak:
				print("break ")

			}

		case &ClosureExp[0]:
			print("func(")
			var end = ast[O(iterator)] == nil || &ast[O(iterator)][0] == &BlocOfCode[0]
			var separ = uint64(len(ast[(iterator)]) - 1)
			if separ == 0 {
				print(")(")
			}
			if end {
				print(")")
			}

		case &TypedIdent[0]:
			if ast[O(iterator)+1] == nil || &ast[O(iterator)+1][0] != &RootOfType[0] {
				var op = byte(len(ast[(iterator)]) - 1)
				switch op {
				case TypedIdentEllipsis:
					print("...")

				}
			}

		}
	}
	for i := uint64(0); i < uint64big; i++ {
		if Poke(ast, O(iterator)+i) {
			Code(print, ast, O(iterator)+i, iterator)
		} else {
			i = uint64big
		}
		if ast[iterator] != nil {
			switch &(ast[iterator])[0] {
			case &ImportsDef[0]:
				if i == uint64big {
					print(")")
				}

			case &ToplevFunc[0]:
				var alpha = ast[O(iterator)+i] == nil || &ast[O(iterator)+i][0] == &BlocOfCode[0]
				var beta = ast[O(iterator)+i+1] == nil || &ast[O(iterator)+i+1][0] == &BlocOfCode[0]
				var gamma = cap(ast[(iterator)]) == len(ast[(iterator)])
				var epsil = len(ast[(iterator)])-1 != 0
				var omega = i != uint64big
				var theta = i == 0
				var phi = i == 1
				var rho = i+1 == uint64(cap(ast[(iterator)]))
				if alpha {
				} else if beta {
					if gamma && omega && epsil == phi && theta != phi {
						if phi {
							print(") ")
						}
						print(ast_o_iterator)
						print("()")
					} else {
						print(")")
					}
				} else {
					if !theta {
						if epsil && phi {
							print(") ")
							print(ast_o_iterator)
							print("(")
						} else if !rho {
							print(", ")
						}
					} else if !epsil {
						print(ast_o_iterator)
						print("(")
					}
					if rho {
						print(")(")
					}
				}
				if i == uint64big {
					print("")
				}

			case &TypedIdent[0]:
				if ast[O(iterator)+i] != nil && &ast[O(iterator)+i][0] != &RootOfType[0] {
					if ast[O(iterator)+i-1] != nil && &ast[O(iterator)+i-1][0] == &RootOfType[0] {
						print(" ")
					}
					print(string(ast[O(iterator)+i]))
					if ast[O(iterator)+i+1] != nil && &ast[O(iterator)+i+1][0] != &RootOfType[0] {
						print(", ")
					} else {
						var op = byte(len(ast[(iterator)]) - 1)
						switch op {
						case TypedIdentTagged:
							fallthrough

						case TypedIdentNormal:
							print(" ")

						case TypedIdentEquals:
							print(" = ")

						case TypedIdentEllipsis:
							print(" ...")

						}
					}
				}

			case &BlocOfCode[0]:
				if i != uint64big {
					if i+uint64(BlocOfCodeTotalCount)+1 == uint64(cap(ast[(iterator)])) {
						if len(ast[(iterator)])-1 == int(BlocOfCodeCase) {
							print(":")
							print("")
						} else if len(ast[(iterator)])-1 == int(BlocOfCodeCommunicate) {
							print(":")
							print("")
						} else if len(ast[(iterator)])-1 < int(BlocOfCodeCase) {
							print("{")
							print("")
						}
					} else if i+uint64(BlocOfCodeTotalCount)+1 > uint64(cap(ast[(iterator)])) {
						print("")
					} else if i+uint64(BlocOfCodeTotalCount)+1 < uint64(cap(ast[(iterator)])) {
						if len(ast[(iterator)])-1 == int(BlocOfCodeCase) {
							print(", ")
						}
					}
				} else if i == uint64big {
					if len(ast[(iterator)])-1 < int(BlocOfCodeCase) {
						print("}")
					}
					if len(ast[(iterator)])-1 == int(BlocOfCodeIfElse) {
						print(" else")
					}
				}

			case &TypDefStmt[0]:
				if i == uint64big {
					print("")
				}

			case &IfceTypExp[0]:
				fallthrough

			case &StructType[0]:
				if i == uint64big {
					print("}")
				} else {
					print("")
				}

			case &GoDferStmt[0]:

			case &Expression[0]:
				var caseheader = true
				var blockheader = true
				var op = byte(len(ast[(iterator)]) - 1)
				var l = uint64(cap(ast[(iterator)]) - int(ExpressionTotalCount))
				if Which(ast[O(iterator)+i]) == nil {
					if len(ast[O(iterator)+i]) > 0 {
						print(string(ast[O(iterator)+i]))
					}
				}
				if i != uint64big {
					i := i + 1
					if ((i >= 1) == (l > 1)) && (i != l) {
						if i == 0 {
						} else {
							switch op {
							case ExpressionDot:
								print(".")

							case ExpressionCallDotDotDot:
								fallthrough

							case ExpressionCall:
								if i == 1 {
									print("(")
								} else {
									print(",")
								}

							case ExpressionOrOr:
								print(" || ")

							case ExpressionAndAnd:
								print(" && ")

							case ExpressionEqual:
								print(" == ")

							case ExpressionNotEq:
								print(" != ")

							case ExpressionLessThan:
								print(" < ")

							case ExpressionLessEq:
								print(" <= ")

							case ExpressionGrtEq:
								print(" >= ")

							case ExpressionGrtThan:
								print(" > ")

							case ExpressionPlus:
								print(" + ")

							case ExpressionMinus:
								print(" - ")

							case ExpressionOr:
								print(" | ")

							case ExpressionXor:
								print(" ^ ")

							case ExpressionMul:
								print(" * ")

							case ExpressionDiv:
								print(" / ")

							case ExpressionMod:
								print(" % ")

							case ExpressionAnd:
								print(" & ")

							case ExpressionAndNot:
								print(" &^ ")

							case ExpressionLSh:
								print(" << ")

							case ExpressionRSh:
								print(" >> ")

							case ExpressionArrow:
								print(" <- ")

							case ExpressionKeyVal:
								print(":")

							case ExpressionIndex:
								print("[")

							case ExpressionSlice:
								if i == 1 {
									print("[")
								} else {
									print(":")
								}

							case ExpressionComposite:
								if i == 1 {
									print("{")
								} else {
									print(", ")
								}

							case ExpressionComposed:
								print(", ")

							case ExpressionMap:
								fallthrough

							case ExpressionSliceType:
								print("]")

							case ExpressionType:
								print(".(")

							}
						}
					} else if (i == 1) && (l == 1) {
						switch op {
						case ExpressionBrackets:
							print(")")

						case ExpressionCallDotDotDot:
							fallthrough

						case ExpressionCall:
							print("()")

						case ExpressionComposed:
							print("}")

						case ExpressionComposite:
							print("{}")

						case ExpressionType:
							print(".(type)")

						}
					} else if (i == l) && (i > 0) {
						switch op {
						case ExpressionIndex:
							print("]")

						case ExpressionSlice:
							if l == 2 {
								print(":]")
							} else {
								print("]")
							}

						case ExpressionComposed:
							fallthrough

						case ExpressionComposite:
							print("}")

						case ExpressionCallDotDotDot:
							print("...)")

						case ExpressionCall:
							print(")")

						case ExpressionType:
							print(")")

						}
					}
					if i == l {
						if caseheader {
						} else if blockheader {
							print(" ")
						}
					}
				}

			case &ReturnStmt[0]:
				if i != uint64big && ast[O(iterator)+i+1] != nil {
					print(", ")
				}

			case &IncDecStmt[0]:
				if i == 0 {
					if Which(ast[O(iterator)]) == nil {
						if len(ast_o_iterator) > 0 {
							print(ast_o_iterator)
						}
					}
				}
				if i == uint64big {
					var op = byte(len(ast[(iterator)]) - 1)
					switch op {
					case IncDecStmtPlusPlus:
						print("++")

					case IncDecStmtMinusMinus:
						print("--")

					}
				}

			case &AssignStmt[0]:
				var blockheader = true
				var op = byte(len(ast[(iterator)]) - 1)
				var l = uint64(cap(ast[(iterator)]) - int(AssignStmtTotalCount))
				if Which(ast[O(iterator)+i]) == nil {
					if len(string(ast[O(iterator)+i])) > 0 {
						print(string(ast[O(iterator)+i]))
					}
				}
				if i != uint64big {
					i := i + 1
					if i == l {
						if blockheader {
							print(" ")
						} else {
							print("")
						}
					} else if (i+1 == l) && (op > AssignStmtTypeIsLast) {
						switch op {
						case AssignStmtMoreEqual:
							print(" = ")

						case AssignStmtMoreColonEq:
							print(" := ")

						case AssignStmtMoreEqualRange:
							print(" = range ")

						case AssignStmtMoreColonEqRange:
							print(" := range ")

						case AssignStmtIotaIsLast:
							print(", ")

						}
					} else if (i == (l+1)>>1) && (op != AssignStmtTypeIsLast) {
						switch op {
						case AssignStmtEqual:
							print(" = ")

						case AssignStmtColonEq:
							print(" := ")

						case AssignStmtAdd:
							print(" += ")

						case AssignStmtSub:
							print(" -= ")

						case AssignStmtMul:
							print(" *= ")

						case AssignStmtQuo:
							print(" /= ")

						case AssignStmtRem:
							print(" %= ")

						case AssignStmtAnd:
							print(" &= ")

						case AssignStmtOr:
							print(" |= ")

						case AssignStmtXor:
							print(" ^= ")

						case AssignStmtShl:
							print(" <<= ")

						case AssignStmtShr:
							print(" >>= ")

						case AssignStmtAndNot:
							print(" &^= ")

						case AssignStmtMoreEqual:
							fallthrough

						case AssignStmtMoreColonEq:
							fallthrough

						case AssignStmtMoreEqualRange:
							fallthrough

						case AssignStmtMoreColonEqRange:
							print(", ")

						case AssignStmtIotaIsLast:
							print(", ")

						}
					} else if (i == l>>1) && (op < AssignStmtTypeIsLast) {
						print(" ")
					} else if i > 0 && ((i+1 != l) || (op != AssignStmtTypeIsLast)) {
						print(", ")
					} else if i > 0 {
						print(" ")
					}
				}

			case &RootOfType[0]:
				if i == uint64big && Which(ast[O(iterator)]) == nil {
					if len(ast_o_iterator) > 0 {
						print(ast_o_iterator)
					}
				}

			case &LblGotoCnt[0]:
				if i == uint64big {
					var op = byte(len(ast[(iterator)]) - 1)
					switch op {
					case LblGotoCntBreak:
						fallthrough

					case LblGotoCntContinue:
						fallthrough

					case LblGotoCntGoto:
						print(ast_o_iterator)

					case LblGotoCntLabel:
						print(ast_o_iterator)
						print(": ")

					}
				}

			case &VarDefStmt[0]:
				if i != uint64big {
					if len(ast[O(iterator)+i]) != 0 {
						if len(ast[O(iterator)+i+1]) != 0 {
							print("")
						}
					}
					if len(ast[O(iterator)+i-1]) != 0 {
						if len(ast[O(iterator)+i+1]) == 0 {
							print("")
							print(")")
							print("")
						}
					}
				}

			case &FileMatter[0]:
				if i != uint64big {
					var xyz = ast[O(iterator)+i+1] == nil || &ast[O(iterator)+i+1][0] == &CommentRow[0]
					var abc = ast[O(iterator)+i+0] != nil && &ast[O(iterator)+i+0][0] == &CommentRow[0]
					var def = ast[O(iterator)+i+1] != nil && &ast[O(iterator)+i+1][0] == &CommentRow[0]
					var end = len(ast[O(iterator)+i+1])-1 == int(CommentRowEnder)
					var ene = len(ast[O(iterator)+i+0])-1 == int(CommentRowEnder)
					if !xyz || !end {
						print("")
					}
					if abc && def && end && ene {
						print("")
					}
				}

			case &ClosureExp[0]:
				if i != uint64big {
					var end = ast[O(iterator)+i+1] == nil || &ast[O(iterator)+i+1][0] == &BlocOfCode[0]
					var xyz = ast[O(iterator)+i+0] == nil || &ast[O(iterator)+i+0][0] == &BlocOfCode[0]
					if end {
						if !xyz {
							print(")")
						}
					} else if !xyz {
						var separ = uint64(len(ast[(iterator)]) - 1)
						if i+1 == separ {
							print(")(")
						} else {
							print(", ")
						}
					}
				}

			case &IfceMethod[0]:
				if i == 0 {
					var end = ast[O(iterator)+i+1] == nil
					if end {
						print("()")
					} else {
						var op = uint64(len(ast[(iterator)]) - 1)
						if 0 == op {
							print("()")
						}
						print("(")
					}
				} else if i != uint64big {
					var end = ast[O(iterator)+i+1] == nil
					if end {
						print(")")
					} else {
						var op = uint64(len(ast[(iterator)]) - 1)
						if i == op {
							print(")(")
						} else {
							print(", ")
						}
					}
				}

			}
		}
	}
}

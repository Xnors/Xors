package tokens2bytecodes

type Opcode int
type ArgType interface{}

const (
	OP_IMPORT_MODULE Opcode = iota
	OP_LOAD_MODULE

	OP_LOAD_CONST

	OP_LOAD_NAME
	OP_STORE_NAME

	OP_LOAD_FAST
	OP_STORE_FAST

	OP_CALL_FUNCTION
	OP_RETURN_VALUE

	OP_POP_TOP
)

const NO_ARG int = 114514

type Bytecode struct {
	Op  Opcode
	Arg ArgType
}

func NewBytecode(op Opcode, arg ArgType) Bytecode {
	return Bytecode{Op: op, Arg: arg}
}

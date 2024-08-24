package XVM

import (
	st "Xors/stack"
	t2bc "Xors/tokens2bytecodes"
	"fmt"
)

func RunBytecodes(bytecodes []t2bc.Bytecode) {
	var stack st.Stack = *st.NewStack()

	for _, bc := range bytecodes {
		// fmt.Print(bc, "\n")

		switch bc.Op {
		case t2bc.OP_LOAD_CONST:
			stack.Push(bc.Arg)
			// fmt.Println(stack)

		case t2bc.OP_CALL_FUNCTION:
			switch bc.Arg {
			case t2bc.F_OUT:
				str, err := stack.Peek()
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}

				fmt.Print(str.(string), "\n")

			}
		case t2bc.OP_POP_TOP:
			stack.Pop()
		}
	}
}

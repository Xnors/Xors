package tokens2bytecodes

import (
	"strconv"
	"strings"
)

func Token2Bytecodes(tokens []string) []Bytecode {
	var bytecodes []Bytecode

	for index, token := range tokens {
		switch token {
		case "out":
			index++
			if strings.Contains(tokens[index], "\"") || strings.Contains(tokens[index], "'") {
				bytecodes = append(bytecodes, NewBytecode(OP_LOAD_CONST, tokens[index]))
			} else if isNumber(tokens[index]) {
				bytecodes = append(bytecodes, NewBytecode(OP_LOAD_CONST, tokens[index]))
			}

			index--
			bytecodes = append(bytecodes, NewBytecode(OP_CALL_FUNCTION, ArgType(F_OUT)))
			bytecodes = append(bytecodes, NewBytecode(OP_POP_TOP, ArgType(NO_ARG)))
		}

	}

	return bytecodes
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

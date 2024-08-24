package main

import (
	vm "Xors/XVM"
	"Xors/tokenizer"
	t2bc "Xors/tokens2bytecodes"
	"fmt"
)

func main() {
	// code := `
	// use std;

	// std::out("hello world");

	// for 5{
	// 	std::out("hello world");
	// }

	// for i <- 1...5{
	// 	std::out(i);
	// }
	// `
	code := `
	use strings;

	out "第一句";
	out 114+514;
	out strings::merge("hello" , "world");

	fun int_add | a int b int | int{
		ret a+b;
	}
	
	out int_add(10, 20);
	`

	tokens := tokenizer.Tokenize(code)
	for _, token := range tokens {
		fmt.Printf("%v\n", token)
	}

	fmt.Print("\n")

	bytecodes := t2bc.Token2Bytecodes(tokens)

	fmt.Print(bytecodes)

	fmt.Print("\n\n以下是程序运行结果：\n\n")

	vm.RunBytecodes(bytecodes)

}

package main

import (
	t2bc "Xors/tokens2bytecodes"
	runner "Xors/runBytecodes"
	"Xors/tokenizer"
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
	code:=`
	out "hello world";
	out "hello worl";
	out 114514;
	out a;
	`

	tokens := tokenizer.Tokenize(code)
	for _, token := range tokens {
		fmt.Printf("%v\n", token)
	}

	fmt.Print("\n")

	bytecodes := t2bc.Token2Bytecodes(tokens)

	fmt.Print(bytecodes)

	fmt.Print("\n\n以下是程序运行结果：\n\n")

	runner.RunBytecodes(bytecodes)

}

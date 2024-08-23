package main

import (
	"Xors/tokenizer"
	"fmt"
)

func main() {
	code := `
	use std;

	std::out("hello world");

	for 5{
		std::out("hello world");
	}

	for i <- 1...5{
		std::out(i);
	}
	`
	tokens := tokenizer.Tokenize(code)
	for _, token := range tokens {
		fmt.Printf("%v\n", token)
	}
}

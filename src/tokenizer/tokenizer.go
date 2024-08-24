package tokenizer

import (
	"strconv"
	"strings"
)

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// filterEmptyStrings 移除切片中的所有多余的换行符和空格
func filterEmptyStrings(slice []string) []string {
	var result []string
	for _, item := range slice {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func Tokenize(code string) []string {
	var tokens []string
	var isInsideString bool // 在字符串中(引号中)
	var stringBuidler = strings.Builder{}
	var alreadySeenColon bool = false // 已经遇到冒号了
	var seenNumber bool = false       // 数字的个数

	var addToken = func(token string) {
		tokens = append(tokens, token)
	}
	var write_reset = func(char rune) {
		if isInsideString {
			stringBuidler.WriteRune(char)
			return
		}else if seenNumber {
			stringBuidler.WriteRune(char)
			return
		}
		

		addToken(stringBuidler.String())

		stringBuidler.Reset()
	}

	var write_reset_addchar = func(char rune) {
		write_reset(char)
		addToken(string(char))
	}

	for _, char := range code {
		switch char {
		case '"', '\'':
			if isInsideString {
				isInsideString = false
				stringBuidler.WriteRune(char)
				write_reset(char)
			} else {
				write_reset(char)
				isInsideString = true
				stringBuidler.WriteRune(char)
			}

		case '(', ')':
			write_reset_addchar(char)
		case '{', '}':
			write_reset_addchar(char)
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			if seenNumber == false {
				seenNumber = true
				stringBuidler.WriteRune(char)
			} else {
				seenNumber = false
				stringBuidler.WriteRune(char)
				// write_reset_addchar(char)
			}
		case ':':
			if alreadySeenColon {
				addToken("::")
				stringBuidler.Reset()
				alreadySeenColon = false
			} else {
				alreadySeenColon = true
			}

		case ' ':
			write_reset(char)

		case ';':
			write_reset_addchar(char)

		default:
			if !(char == '\t' || char == '\n') {
				stringBuidler.WriteRune(char)
			}
		}
	}
	tokens = filterEmptyStrings(tokens)
	return tokens
}

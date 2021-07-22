package arraysandstrings

import (
	"fmt"
	"strings"
)

/*
String Compression
Cracking the coding interview, 6th edition, page 91
Assumption: the given string only includes uppercase and lowercase letters
If the compressed string is not shorter than the original string, return the original
*/

type stringCompression struct {
	string string
}

func NewStringCompression(string string) *stringCompression {
	return &stringCompression{
		string: string,
	}
}

func (sc *stringCompression) compressString() string {
	var sb strings.Builder
	length := sc.compressed(&sb)
	if len(sc.string) <= 2 || len(sc.string) <= length {
		return sc.string
	} else {
		return sb.String()
	}

}

func (sc *stringCompression) compressed(sb *strings.Builder) int {
	length := 0
	repeat := 0
	for index := range sc.string {
		repeat += 1
		if index+1 >= len(sc.string) || sc.string[index] != sc.string[index+1] {
			length += len(string(rune(repeat)))
			sb.WriteByte(sc.string[index])
			// sb.WriteByte(byte(48+repeat)) => works only if repeat < 10
			sb.Write([]byte(fmt.Sprintf("%v", repeat)))
			repeat = 0
		}
	}
	return length
}

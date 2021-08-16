package utils

import "fmt"

type StructIsNilException struct {
	DataStructure string
	FuncName      string
}

func (s *StructIsNilException) Error() string {
	return fmt.Sprintf("error in function %s: %s cannot be nil", s.FuncName, s.DataStructure)
}

type InvalidAlphabeticWord struct {
	Word string
}

func (a *InvalidAlphabeticWord) Error() string {
	return fmt.Sprintf("error: word %v contaons a non-alphabetic character", a.Word)
}

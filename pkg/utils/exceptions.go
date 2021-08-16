package utils

import "fmt"

type StructIsNilException struct {
	DataStructure string
	FuncName      string
}

func (s *StructIsNilException) Error() string {
	return fmt.Sprintf("Error in function %s: %s cannot be nil", s.FuncName, s.DataStructure)
}

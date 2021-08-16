package Stacks

import "fmt"

type StackTopIsNilException struct {
	funcName string
}

func (s *StackTopIsNilException) Error() string {
	return fmt.Sprintf("function %v cannot operate with empty stack", s.funcName)
}

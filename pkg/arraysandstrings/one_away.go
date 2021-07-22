package arraysandstrings

import "strings"

/*
Cracking the coding interview, 6th edition, page 91
One Away
*/

type oneAway struct {
	firstString  string
	secondString string
}

func NewOnwAway(firstString, secondString string) *oneAway {
	return &oneAway{
		firstString:  firstString,
		secondString: secondString,
	}
}

func (oa *oneAway) isOneAway() bool {
	if len(oa.firstString) == len(oa.secondString) {
		countDifference := 0
		for index := range oa.firstString {
			if oa.firstString[index] != oa.secondString[index] {
				countDifference += 1
			}
			if countDifference > 1 {
				return false
			}
		}
		return true
	} else if len(oa.firstString) == len(oa.secondString)+1 && strings.Contains(oa.firstString, oa.secondString) {
		return true
	} else if len(oa.secondString) == len(oa.firstString)+1 && strings.Contains(oa.secondString, oa.firstString) {
		return true
	}
	return false
}

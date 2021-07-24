package arraysandstrings

import "strings"

/*
String Rotation
Cracking he coding interview, 6th edition, page 91
*/

type stringRotation struct {
	firstString  string
	secondString string
}

func NewStringRotation(firstString string, secondString string) *stringRotation {
	return &stringRotation{
		firstString:  firstString,
		secondString: secondString,
	}
}

func (sr *stringRotation) areRotations() bool {
	if len(sr.firstString) == len(sr.secondString) {
		return strings.Contains(sr.firstString+sr.firstString, sr.secondString)
	} else {
		return false
	}

}

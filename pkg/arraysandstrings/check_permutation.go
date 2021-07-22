package arraysandstrings

/*
Cracking the coding interview, 6th edition page 90
Check permutation: Given two strings, write a method to decide if one is a permutation of the other one
assumption: it's case sensitive and whitespace is significant
*/

type checkPermutation struct {
	firstString  string
	secondString string
}

func NewCheckPermutation(firstString string, secondString string) *checkPermutation {
	return &checkPermutation{
		firstString:  firstString,
		secondString: secondString,
	}
}

func (cp *checkPermutation) arePermute() bool {
	hmap := make(map[int32]int)
	if len(cp.firstString) != len(cp.secondString) {
		return false
	}
	for _, char := range cp.firstString {
		if _, ok := hmap[char]; ok {
			hmap[char] += 1
		} else {
			hmap[char] = 1
		}
	}
	for _, char := range cp.secondString {
		if _, ok := hmap[char]; ok {
			hmap[char] -= 1
		} else {
			return false
		}
		if hmap[char] < 0 {
			return false
		}
	}
	return true
}

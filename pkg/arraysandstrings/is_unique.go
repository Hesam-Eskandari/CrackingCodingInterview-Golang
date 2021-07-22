package arraysandstrings

/*
Cracking the coding interview, 6th edition, page 90
Is Unique: Implement an algorithm to determine if a string has all unique characters.
What if you cannot use additional data structures.
*/

type uniqueString struct {
	str string
}

func NewUniqueString(str string) *uniqueString {
	return &uniqueString{
		str: str,
	}
}

func (s *uniqueString) IsUnique() bool {
	hmap := make(map[byte]int)
	for index := range s.str {
		if _, ok := hmap[s.str[index]]; ok {
			hmap[s.str[index]] += 1
		} else {
			hmap[s.str[index]] = 1
		}
	}
	for _, value := range hmap {
		if value != 1 {
			return false
		}
	}
	return true
}

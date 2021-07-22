package arraysandstrings

/*
Cracking the coding interview, 6th edition, page 90
URLify: write a method to replace all spaces in a string with '%20'.
You may assume that the string has sufficient space at the end to hold the additional characters.
And that you are given the true length of the string.
*/

type url struct {
	string string
	length int
}

func NewURLify(string string, length int) *url {
	return &url{string: string, length: length}
}

func (u *url) Urlify() {
	byteArr := []byte(u.string)
	end := len(byteArr)
	j := len(byteArr) - 1
	for index := range byteArr {
		indexBackward := end - index - 1
		if byteArr[indexBackward] != 32 {
			temp := byteArr[indexBackward]
			byteArr[indexBackward] = byteArr[j]
			byteArr[j] = temp
			j -= 1
		} else if byteArr[indexBackward] == 32 && indexBackward < u.length {
			j -= 3
		}
	}
	for index := range byteArr {

		if byteArr[index] == 32 && index <= len(byteArr)-3 {
			byteArr[index] = 37
			byteArr[index+1] = 50
			byteArr[index+2] = 48
		}
	}
	u.string = string(byteArr)
}

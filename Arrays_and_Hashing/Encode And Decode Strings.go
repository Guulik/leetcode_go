package Arrays_and_Hashing

import "strconv"

func Encode(strs []string) string {
	result := ""
	for _, str := range strs {
		result += "$" + strconv.Itoa(len(str)) + str
	}
	return result
}

func Decode(str string) []string {
	strs := make([]string, 0, 5)
	var s string
	var length uint8
	for i := 0; i < len(str); {
		if str[i] == 36 {
			length = str[i+1] - 48
			i += 2
		}
		j := 0
		for ; j < int(length); j++ {
			s += string(str[i+j])
		}
		i += j
		strs = append(strs, s)
		s = ""
	}
	return strs
}

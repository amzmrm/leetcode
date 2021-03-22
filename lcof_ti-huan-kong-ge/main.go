package leetcode

func replaceSpace(s string) string {
	var numsOfBlank int
	for i := range s {
		if s[i] == ' ' {
			numsOfBlank++
		}
	}

	res := make([]byte, len(s)+numsOfBlank*2)
	left, right := len(s)-1, len(s)+numsOfBlank*2-1
	for left >= 0 {
		if s[left] == ' ' {
			res[right] = '0'
			res[right-1] = '2'
			res[right-2] = '%'
			right -= 3
		} else {
			res[right] = s[left]
			right--
		}
		left--
	}
	return string(res)
}

// func replaceSpace(s string) string {
// 	var res []byte
// 	for i := range s {
// 		if s[i] == ' ' {
// 			res = append(res, '%', '2', '0')
// 		} else {
// 			res = append(res, s[i])
// 		}
// 	}
// 	return string(res)
// }

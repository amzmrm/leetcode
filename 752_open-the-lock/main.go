package leetcode

import "strings"

const (
	ascii0 uint8 = 48
	ascii9 uint8 = 57
)

func openLock(deadends []string, target string) int {
	ends := make(map[string]bool, len(deadends))
	for i := range deadends {
		ends[deadends[i]] = true
	}
	if _, ok := ends["0000"]; ok {
		return -1
	}
	if _, ok := ends[target]; ok {
		return -1
	}

	queue := []string{"0000"}
	visited := make(map[string]bool)
	visited["0000"] = true
	step := 0
	var curr, up, down string
	var size int
	for len(queue) != 0 {
		size = len(queue)
		for j := 0; j < size; j++ {
			curr, queue = queue[0], queue[1:]
			if _, ok := ends[curr]; ok {
				continue
			}
			if curr == target {
				return step
			}

			for i := 0; i < 4; i++ {
				up = plusOne(curr, i)
				if _, ok := visited[up]; !ok {
					queue = append(queue, up)
					visited[up] = true
				}
				down = minusOne(curr, i)
				if _, ok := visited[down]; !ok {
					queue = append(queue, down)
					visited[down] = true
				}
			}
		}
		step++
	}

	return -1
}

func plusOne(s string, index int) string {
	digit := s[index]
	if digit == ascii9 {
		digit = ascii0
	} else {
		digit++
	}
	sb := strings.Builder{}
	for i := range s {
		if i == index {
			sb.WriteByte(digit)
			continue
		}
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func minusOne(s string, index int) string {
	digit := s[index]
	if digit == ascii0 {
		digit = ascii9
	} else {
		digit--
	}
	sb := strings.Builder{}
	for i := range s {
		if i == index {
			sb.WriteByte(digit)
			continue
		}
		sb.WriteByte(s[i])
	}
	return sb.String()
}

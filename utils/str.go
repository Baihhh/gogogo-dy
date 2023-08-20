package utils

import "strconv"

func StringToInt(str string) int64 {
	intValue, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return intValue
}

package util

import (
	"strconv"
)

func IsFloat(number string) bool {
	if number != "" {
		var _ , error = strconv.ParseFloat(number, 32);
		if error == nil {
			return true
		}
	}
	return false
}

func IsInt(number string) bool {
	if number != "" {
		var _ , error = strconv.ParseInt(number,0,16);
		if error == nil {
			return true
		}
	}
	return false
}
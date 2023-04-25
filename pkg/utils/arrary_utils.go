package utils

import "strconv"

func Contains[T comparable](arr *[]T, e T) bool {
	for _, a := range *arr {
		if a == e {
			return true
		}
	}
	return false
}

func StringToInt(str string) uint64 {
	i, _ := strconv.ParseUint(str, 10, 64)
	return i
}

func StringArrayToInt64Array(arr []string) []uint64 {
	var intArr []uint64
	for _, v := range arr {
		intArr = append(intArr, StringToInt(v))
	}
	return intArr
}

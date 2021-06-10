package utils

import "strconv"

// StrToInt32 function to convert string to int32
func StrToInt32(str string) int32 {
	result, err := strconv.Atoi(str)
	if err != nil {
		Log.Error("Failed to convert string to int")
	}
	return int32(result)
}

// StrToInt64 function to convert string to int64
func StrToInt64(str string) int64 {
	result, err := strconv.Atoi(str)
	if err != nil {
		Log.Error("Failed to convert string to int")
	}
	return int64(result)
}

package utils

import (
	"strconv"
)

func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Int64ToString(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}

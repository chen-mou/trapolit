package utils

import "time"

func NowFormat(format string) string {
	return time.Now().Format(format)
}

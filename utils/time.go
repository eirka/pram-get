package utils

import (
	"time"
)

var StartTime = time.Now()

func GetTime() string {
	return fmt.Sprintf("%s", time.Since(StartTime)*time.Second)
}

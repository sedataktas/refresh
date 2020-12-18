package runner

import (
	"fmt"
	"time"
)

func GetTime() string {
	now := time.Now()
	return fmt.Sprintf("%d:%d:%02d", now.Hour(), now.Minute(), now.Second())
}

package helper

import (
	"fmt"
	"go-hub/config"
	"time"
)

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

func NowTime() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.Cfg.Application.Timezone)
	return time.Now().In(chinaTimezone)
}

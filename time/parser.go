package time

import (
	"regexp"
	"strings"
	"time"
)

// 处理时间格式化

var (
	reg *regexp.Regexp
)

// ParseTime 解析时间字符串
// 支持的 format:
//	"2020-05-04 17:22:00"
// 	"2020/05/04 17:22:00"
func ParseTime(v string, offset int) (*time.Time, error) {
	timeStr := strings.Replace(v, "/", "-", -1)
	if reg.Match([]byte(timeStr)) {
		timeStr = strings.TrimSpace(timeStr) // 去除首位空格
		timeStr = strings.Replace(timeStr, " ", "T", -1)
		timeStr = timeStr + "Z"

		t, err := time.Parse(time.RFC3339, timeStr)
		if err != nil {
			return nil, err
		}

		t = t.Add(time.Duration(-1 * offset * int(time.Hour)))
		t = t.UTC()
		return &t, nil
	}

	return nil, nil
}

func init() {
	reg = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})`)
}

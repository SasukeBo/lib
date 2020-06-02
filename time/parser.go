package time

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// 处理时间格式化

var regs map[string]*regexp.Regexp

const (
	ParseTimePattern1 = `^(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})$`
	ParseTimePattern2 = `^(\d{4})/(\d{2})/(\d{2}) (\d{2}):(\d{2}):(\d{2})$`
	ParseTimePattern3 = `^(\d{14})$`
)

// ParseTime 解析时间字符串
// 时间串 和 时区offset，例如东八区 offset 为 8
// 返回解析后的UTC时间
func ParseTime(v string, offset int) *time.Time {
	timeStr := strings.TrimSpace(v) // 去除首位空格
	var pattern string

	for p, r := range regs {
		if r.Match([]byte(timeStr)) {
			pattern = p
			break
		}
	}

	switch pattern {
	case ParseTimePattern1:
		return parseRFC3339(timeStr, offset)
	case ParseTimePattern2:
		timeStr = strings.Replace(timeStr, "/", "-", -1)
		return parseRFC3339(timeStr, offset)
	case ParseTimePattern3:
		timeStr = fmt.Sprintf("%s-%s-%s %s:%s:%s", timeStr[:4], timeStr[4:6], timeStr[6:8], timeStr[8:10], timeStr[10:12], timeStr[12:])
		return parseRFC3339(timeStr, offset)
	}

	return nil
}

func parseRFC3339(source string, offset int) *time.Time {
	source = strings.Replace(source, " ", "T", -1)
	source = source + "Z"

	t, err := time.Parse(time.RFC3339, source)
	if err != nil {
		return nil
	}

	t = t.Add(time.Duration(-1 * offset * int(time.Hour)))
	t = t.UTC()
	return &t
}

func init() {
	regs = make(map[string]*regexp.Regexp)
	patterns := []string{
		ParseTimePattern1,
		ParseTimePattern2,
		ParseTimePattern3,
	}

	for _, p := range patterns {
		regs[p] = regexp.MustCompile(p)
	}
}

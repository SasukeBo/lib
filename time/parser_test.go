package time

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	std := time.Date(2020, time.June, 2, 5, 45, 25, 0, time.UTC)
	t.Run("Parse 2020-06-02 13:45:25", func(t *testing.T) {
		ret := ParseTime("2020-06-02 13:45:25", 8)
		if !std.Equal(*ret) {
			t.Errorf("want %v \n got %v\n", std, ret)
		}
	})

	t.Run("Parse 2020/06/02 13:45:25", func(t *testing.T) {
		ret := ParseTime("2020/06/02 13:45:25", 8)
		if !std.Equal(*ret) {
			t.Errorf("want %v \n got %v\n", std, ret)
		}
	})

	t.Run("Parse 20200602134525 type", func(t *testing.T) {
		ret := ParseTime("20200602134525", 8)
		if !std.Equal(*ret) {
			t.Errorf("want %v \n got %v\n", std, ret)
		}
	})
}

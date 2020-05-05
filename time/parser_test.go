package time

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	tStr1 := "2020-03-30 10:32:31"
	tStr2 := "2020/03/30 10:32:31"
	t1, err := ParseTime(tStr1, 8)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("t1 is %v\n", t1)
	t2, err := ParseTime(tStr2, 8)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("t2 is %v\n", t2)
}

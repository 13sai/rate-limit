package xrate

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	limit := New(time.Minute, 30, 2, "abc")
	for i := 0; i < 10; i++ {
		t.Logf("aa=%v", limit.Allow("1"))
		time.Sleep(1 * time.Second)
	}
}

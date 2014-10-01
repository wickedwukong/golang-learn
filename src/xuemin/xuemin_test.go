package xuemin

import "testing"

func TestName(t *testing.T) {
	name := Name()
	if (name != "xuemin") {
		t.Error("expected xuemin, got ", name)
	}
}
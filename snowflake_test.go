package snowflake

import (
	"testing"
)

func TestGet(t *testing.T) {
	n, e := Get()
	if e != nil {
		t.Errorf(e.Error())
		return
	}
	t.Log(n)
}

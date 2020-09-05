package bitset

import (
	"testing"
)

func Test(t *testing.T) {
	bs := New(128)

	if bs.Test(123) {
		t.Error(`test 123, expected false`)
	}

	bs.Set(5)
	bs.Set(9)

	if !bs.Test(5) {
		t.Error(`test 5, expected true`)
	}

	dmp := bs.Marshal()
	bs.Unmarshal(dmp)

	bs.Unset(5)

	if !bs.Test(9) {
		t.Error(`test 9, expected true`)
	}
}

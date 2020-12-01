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
	bs.Set(98)

	if !bs.Test(5) {
		t.Error(`test 5, expected true`)
	}

	_ = bs.SaveFile(`./bitset.csv`)
	_ = bs.LoadFile(`./bitset.csv`)

	bs.Unset(5)

	if !bs.Test(9) {
		t.Error(`test 9, expected true`)
	}
}

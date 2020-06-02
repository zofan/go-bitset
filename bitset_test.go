package bitset

import (
	"testing"
)

func Test(t *testing.T) {
	bs := New(128)

	err := bs.Set(1234)
	if err != ErrOutOfRange {
		t.Error(`set 1234, expected ErrOutOfRange`)
	}

	err = bs.Unset(1234)
	if err != ErrOutOfRange {
		t.Error(`unset 1234, expected ErrOutOfRange`)
	}

	if bs.Test(1234) {
		t.Error(`test 1234, expected false`)
	}

	if bs.Test(123) {
		t.Error(`test 123, expected false`)
	}

	err = bs.Set(5)
	if err != nil {
		t.Error(`set 5, expected empty error`)
	}

	err = bs.Set(9)
	if err != nil {
		t.Error(`set 9, expected empty error`)
	}

	if !bs.Test(5) {
		t.Error(`test 5, expected true`)
	}

	dmp := bs.Bytes()
	bs.Load(dmp)

	err = bs.Unset(5)
	if err != nil {
		t.Error(`unset 5, expected empty error`)
	}

	if !bs.Test(9) {
		t.Error(`test 9, expected true`)
	}
}

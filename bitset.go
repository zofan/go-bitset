package bitset

import (
	"encoding/binary"
	"errors"
)

var ErrOutOfRange = errors.New(`bit number out of range`)

type BitSet struct {
	set  []uint64
	size uint64
}

func New(size uint64) *BitSet {
	if size%64 != 0 {
		panic(`invalid size`)
	}

	return &BitSet{
		set:  make([]uint64, size/64),
		size: size,
	}
}

func (bs *BitSet) Reset() {
	bs.set = make([]uint64, bs.size/64)
}

func (bs *BitSet) Load(raw []byte) {
	bs.Reset()

	for n := 0; uint64(n) < bs.size/64; n++ {
		bs.set[n] = binary.BigEndian.Uint64(raw[n*8:])
	}
}

func (bs *BitSet) Bytes() []byte {
	raw := make([]byte, len(bs.set)*8)
	for n, bv := range bs.set {
		binary.BigEndian.PutUint64(raw[n*8:], bv)
	}

	return raw
}

func (bs *BitSet) Test(bitNum uint64) bool {
	if bitNum > bs.size {
		return false
	}

	i := bitNum / 64
	b := uint64(bitNum % 64)
	bv := uint64(1 << (b - 1))

	return bs.set[i]&bv == bv
}

func (bs *BitSet) Set(bitNum uint64) error {
	if bitNum > bs.size {
		return ErrOutOfRange
	}

	i := bitNum / 64
	b := uint64(bitNum % 64)

	bs.set[i] = bs.set[i] | (1 << (b - 1))

	return nil
}

func (bs *BitSet) Unset(bitNum uint64) error {
	if bitNum > bs.size {
		return ErrOutOfRange
	}

	i := bitNum / 64
	b := uint64(bitNum % 64)

	bs.set[i] = bs.set[i] &^ (1 << (b - 1))

	return nil
}

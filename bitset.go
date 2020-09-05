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

func (bs *BitSet) Size() uint64 {
	return bs.size
}

func (bs *BitSet) Reset() {
	bs.set = make([]uint64, bs.size/64)
}

func (bs *BitSet) Unmarshal(raw []byte) {
	bs.size = binary.BigEndian.Uint64(raw[0:])
	bs.Reset()

	for n := 0; uint64(n) < bs.size/64; n++ {
		bs.set[n] = binary.BigEndian.Uint64(raw[(n*8)+8:])
	}
}

func (bs *BitSet) Marshal() []byte {
	raw := make([]byte, (len(bs.set)*8)+8)

	binary.BigEndian.PutUint64(raw[0:], bs.size)

	for n, bv := range bs.set {
		binary.BigEndian.PutUint64(raw[(n*8)+8:], bv)
	}

	return raw
}

func (bs *BitSet) Test(bitNum uint64) bool {
	i := bitNum / 64
	b := bitNum % 64
	bv := uint64(1 << (b - 1))

	return bs.set[i]&bv == bv
}

func (bs *BitSet) Set(bitNum uint64) {
	i := bitNum / 64
	b := bitNum % 64

	bs.set[i] |= 1 << (b - 1)
}

func (bs *BitSet) Unset(bitNum uint64) {
	i := bitNum / 64
	b := bitNum % 64

	bs.set[i] &= ^(1 << (b - 1))
}

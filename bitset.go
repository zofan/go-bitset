package bitset

import (
	"encoding/binary"
	"sync"
)

type BitSet struct {
	set  []bits
	size uint64
}

type bits struct {
	bits uint64
	sync.Mutex
}

func New(size uint64) *BitSet {
	if size%64 != 0 {
		panic(`bitset: invalid size`)
	}

	return &BitSet{
		set:  make([]bits, size/64),
		size: size,
	}
}

func (bs *BitSet) Size() uint64 {
	return bs.size
}

func (bs *BitSet) Reset() {
	bs.set = make([]bits, bs.size/64)
}

func (bs *BitSet) UnmarshalBinary(raw []byte) error {
	bs.size = binary.BigEndian.Uint64(raw)
	bs.Reset()

	for n := 0; uint64(n) < bs.size/64; n++ {
		bs.set[n] = bits{bits: binary.BigEndian.Uint64(raw[(n*8)+8:])}
	}
}

func (bs *BitSet) MarshalBinary() ([]byte, error) {
	raw := make([]byte, (len(bs.set)*8)+8)

	binary.BigEndian.PutUint64(raw, bs.size)

	for n, s := range bs.set {
		binary.BigEndian.PutUint64(raw[(n*8)+8:], s.bits)
	}

	return raw, nil
}

func (bs *BitSet) Test(bitNum uint64) bool {
	v := uint64(1 << ((bitNum % 64) - 1))

	bs.set[bitNum/64].Lock()
	r := bs.set[bitNum/64].bits&v == v
	bs.set[bitNum/64].Unlock()

	return r
}

func (bs *BitSet) Set(bitNum uint64) {
	bs.set[bitNum/64].Lock()
	bs.set[bitNum/64].bits |= 1 << ((bitNum % 64) - 1)
	bs.set[bitNum/64].Unlock()
}

func (bs *BitSet) Unset(bitNum uint64) {
	bs.set[bitNum/64].Lock()
	bs.set[bitNum/64].bits &= ^(1 << ((bitNum % 64) - 1))
	bs.set[bitNum/64].Unlock()
}

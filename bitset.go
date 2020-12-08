package bitset

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
)

type BitSet struct {
	set []bits
}

type bits struct {
	bits uint64
	mu   sync.RWMutex
}

func New(size uint64) *BitSet {
	if size%64 != 0 {
		panic(`bitset: invalid size`)
	}

	return &BitSet{
		set: make([]bits, size/64),
	}
}

func (bs *BitSet) Size() uint64 {
	return uint64(len(bs.set) * 64)
}

func (bs *BitSet) Reset() {
	bs.set = make([]bits, len(bs.set))
}

func (bs *BitSet) Test(bitNum uint64) bool {
	v := uint64(1 << ((bitNum % 64) - 1))

	bs.set[bitNum/64].mu.RLock()
	r := bs.set[bitNum/64].bits&v == v
	bs.set[bitNum/64].mu.RUnlock()

	return r
}

func (bs *BitSet) Set(bitNum uint64) {
	bs.set[bitNum/64].mu.Lock()
	bs.set[bitNum/64].bits |= 1 << ((bitNum % 64) - 1)
	bs.set[bitNum/64].mu.Unlock()
}

func (bs *BitSet) Unset(bitNum uint64) {
	bs.set[bitNum/64].mu.Lock()
	bs.set[bitNum/64].bits &= ^(1 << ((bitNum % 64) - 1))
	bs.set[bitNum/64].mu.Unlock()
}

func (bs *BitSet) LoadFile(file string) error {
	fh, err := os.OpenFile(file, os.O_CREATE|os.O_RDONLY, 0664)
	if err != nil {
		return err
	}
	defer fh.Close()

	bs.Reset()

	s := bufio.NewScanner(fh)

	var i int
	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		if len(line) == 0 {
			continue
		}

		n, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			continue
		}

		bs.set[i] = bits{bits: uint64(n)}
		i++
	}

	return nil
}

func (bs *BitSet) SaveFile(file string) error {
	fh, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return err
	}
	defer fh.Close()

	for _, bits := range bs.set {
		bits.mu.RLock()
		if _, err = fh.WriteString(strconv.Itoa(int(bits.bits)) + "\n"); err != nil {
			return err
		}
		bits.mu.RUnlock()
	}

	return nil
}

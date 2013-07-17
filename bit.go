// © 2013 the Bits Authors under the MIT license. See AUTHORS for the list of authors.

// Package bit implements functionality for reading streams of bits from an io.Reader.
package bit

import (
	"io"
)

var mask = [...]uint8{
	0x0,
	0x1,
	0x3,
	0x7,
	0xF,
	0x1F,
	0x3F,
	0x7F,
	0xFF,
}

// Reader provides methods for reading bits.
type Reader struct {
	in io.Reader
	b  uint8
	n  uint
}

// NewReader returns a new Reader that reads bits the given io.Reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{in: r}
}

// Uint8 reads and return n bits, up to 8.  It panicks if n is greater than 8.
func (r *Reader) Uint8(n uint) (uint8, error) {
	if n > 8 {
		panic("Too many bits for Uint8")
	}

	var vl uint8
	for n > 0 {
		if r.n == 0 {
			if err := r.nextByte(); err != nil {
				return 0, err
			}
		}

		m := r.n
		if r.n >= n {
			m = n
		}

		shift := r.n - m
		b := (r.b >> shift) & mask[m]
		vl = (vl << m) | b

		n -= m
		r.n -= m
	}

	return vl, nil
}

func (r *Reader) nextByte() error {
	var b [1]uint8
	if _, err := io.ReadFull(r.in, b[:]); err != nil {
		return err
	}
	r.b = b[0]
	r.n = 8
	return nil
}

// Uint64 reads and return n bits, up to 64.  It panicks if n is greater than 64.
func (r *Reader) Uint64(n uint) (uint64, error) {
	if n > 64 {
		panic("Too many bits for Uint64")
	}
	var vl uint64
	for n > 0 {
		m := n
		if n > 8 {
			m = 8
		}
		b, err := r.Uint8(m)
		if err != nil {
			return 0, err
		}
		vl = (vl << m) | uint64(b)
		n -= m
	}
	return vl, nil
}

// © 2013 the Bits Authors under the MIT license. See AUTHORS for the list of authors.

// Package bit implements functionality for reading streams of bits from an io.Reader.
package bit

import (
	"io"
)

// Reader provides methods for reading bits.
// Reader buffers bits up to the next byte boundary.
type Reader struct {
	in io.Reader
	b  uint64
	n  uint
}

// NewReader returns a new Reader that reads bits the given io.Reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{in: r}
}

// Read returns the next n bits, up to 64.  It panicks if n is greater than 64.
func (r *Reader) Read(n uint) (uint64, error) {
	if n > 64 {
		panic("Attempt to read too many bits")
	}

	var vl uint64
	for n > 0 {
		if r.n == 0 {
			if err := r.fillBuffer(n); err != nil {
				return 0, err
			}
		}

		m := r.n
		if r.n >= n {
			m = n
		}

		shift := r.n - m
		b := (r.b >> shift) & mask(m)
		vl = (vl << m) | b

		n -= m
		r.n -= m
	}

	return vl, nil
}

func (r *Reader) fillBuffer(n uint) error {
	bytes := n / 8
	if bytes*8 < n {
		bytes++
	}

	if bytes > 8 {
		panic("Too many bytes in fillBuffer")
	}

	var buf [8]uint8
	if _, err := io.ReadFull(r.in, buf[:bytes]); err != nil {
		return err
	}

	r.b = uint64(buf[0])
	for _, b := range buf[1:bytes] {
		r.b <<= 8
		r.b |= uint64(b)
	}
	r.n = bytes * 8

	return nil
}

func mask(i uint) uint64 {
	return ^uint64(0) >> (64 - i)
}

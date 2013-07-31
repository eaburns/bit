// © 2013 the Bits Authors under the MIT license. See AUTHORS for the list of authors.

// Package bit implements functionality for reading streams of bits from an io.Reader.
package bit

import (
	"io"
)

// Reader provides methods for reading bits.
// Reader buffers bits up to the next byte boundary.
type Reader struct {
	in  io.Reader
	b   uint64
	n   uint
	err error
}

// NewReader returns a new Reader that reads bits the given io.Reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{in: r}
}

// Err returns the first error encountered.
func (r *Reader) Err() error {
	return r.err
}

// Read returns the next n bits, up to 64.  It panicks if n is greater than 64.
// If an error is encountered, or if an error was encountered during a previous call
// to Read, then 0 is returned and the first error can be retrieved by Err().
func (r *Reader) Read(n uint) uint64 {
	if r.err != nil {
		return 0
	}

	if n > 64 {
		panic("Attempt to read too many bits")
	}

	var vl uint64
	for n > 0 {
		if r.n == 0 {
			if r.b, r.n, r.err = buffer(r.in, n); r.err != nil {
				return 0
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

	return vl
}

func buffer(r io.Reader, n uint) (uint64, uint, error) {
	bytes := n / 8
	if bytes*8 < n {
		bytes++
	}

	if bytes > 8 {
		panic("Too many bytes in fillBuffer")
	}

	var buf [8]uint8
	if _, err := io.ReadFull(r, buf[:bytes]); err != nil {
		return 0, 0, err
	}

	v := uint64(buf[0])
	for _, b := range buf[1:bytes] {
		v <<= 8
		v |= uint64(b)
	}

	return v, bytes * 8, nil
}

func mask(i uint) uint64 {
	return ^uint64(0) >> (64 - i)
}

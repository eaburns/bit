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
// If an EOF is encountered and some bits have been read then io.ErrUnexpectedEOF
// is returned, otherwise if an EOF is encountered and nothing was read then io.EOF
// is returned.
func (r *Reader) Read(n uint) (uint64, error) {
	if n > 64 {
		panic("Attempt to read too many bits")
	}

	var vl uint64
	left := n
	for left > 0 {
		if r.n == 0 {
			var err error
			switch r.b, r.n, err = buffer(r.in, left); {
			case err == io.EOF && left != n:
				return 0, io.ErrUnexpectedEOF
			case err != nil:
				return 0, err
			}
		}

		m := r.n
		if r.n >= left {
			m = left
		}

		shift := r.n - m
		b := (r.b >> shift) & mask(m)
		vl = (vl << m) | b

		left -= m
		r.n -= m
	}

	return vl, nil
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

// ReadFields performs a series of reads of differing numbers of bits, specified by the argument
// list.  If no error occurs, then the results of the reads are returned in a slice, otherwise the
// first encountered error is returned.   If the error is EOF and any number of bits has been
// beer read then io.ErrUnexpectedEOF is returned, if the error is EOF and nothing was read
// then io.EOF is returned.  If any of the read sizes specified by the argument are greater than
// 64 then ReadFields panicks.
func (r *Reader) ReadFields(ns ...uint) ([]uint64, error) {
	fs := make([]uint64, len(ns))
	var err error
	for i, n := range ns {
		switch fs[i], err = r.Read(n); {
		case err == io.EOF && i != 0:
			return nil, io.ErrUnexpectedEOF
		case err != nil:
			return nil, err
		}
	}
	return fs, nil
}

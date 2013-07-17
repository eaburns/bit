// © 2013 the Bits Authors under the MIT license. See AUTHORS for the list of authors.

package bit

import (
	"bytes"
	"testing"
)

func TestUint8(t *testing.T) {
	tests := []struct {
		data []byte
		ns   []uint
		vals []uint8
	}{
		{[]byte{0xFF}, []uint{1, 1, 1, 1, 1, 1, 1, 1}, []uint8{1, 1, 1, 1, 1, 1, 1, 1}},
		{[]byte{0xFF}, []uint{2, 2, 2, 2}, []uint8{0x3, 0x3, 0x3, 0x3}},
		{[]byte{0xFF}, []uint{3, 3, 2}, []uint8{0x7, 0x7, 0x3}},
		{[]byte{0xFF}, []uint{4, 4}, []uint8{0xF, 0xF}},
		{[]byte{0xFF}, []uint{5, 3}, []uint8{0x1F, 0x7}},
		{[]byte{0xFF}, []uint{6, 2}, []uint8{0x3F, 0x3}},
		{[]byte{0xFF}, []uint{7, 1}, []uint8{0x7F, 0x1}},
		{[]byte{0xFF}, []uint{8}, []uint8{0xFF}},

		{[]byte{0xAA}, []uint{1, 1, 1, 1, 1, 1, 1, 1}, []uint8{1, 0, 1, 0, 1, 0, 1, 0}},
		{[]byte{0xAA}, []uint{2, 2, 2, 2}, []uint8{0x2, 0x2, 0x2, 0x2}},
		{[]byte{0xAA}, []uint{3, 3, 2}, []uint8{0x5, 0x2, 0x2}},
		{[]byte{0xAA}, []uint{4, 4}, []uint8{0xA, 0xA}},
		{[]byte{0xAA}, []uint{5, 3}, []uint8{0x15, 0x2}},
		{[]byte{0xAA}, []uint{6, 2}, []uint8{0x2A, 0x2}},
		{[]byte{0xAA}, []uint{7, 1}, []uint8{0x55, 0x0}},
		{[]byte{0xAA}, []uint{8}, []uint8{0xAA}},

		{
			[]byte{0xAA, 0x55},
			[]uint{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[]uint8{1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1},
		},

		{
			[]byte{0xAA, 0x55},
			[]uint{7, 8, 1},
			[]uint8{0x55, 0x2A, 0x1},
		},

		{
			[]byte{0xAA, 0x55},
			[]uint{3, 3, 3, 3, 3, 1},
			[]uint8{0x5, 0x2, 0x4, 0x5, 0x2, 0x1},
		},
	}

	for _, test := range tests {
		r := NewReader(bytes.NewReader(test.data))
		if len(test.ns) != len(test.vals) {
			panic("Number of reads does not match number of results")
		}
		for i, n := range test.ns {
			m, err := r.Uint8(n)
			if err != nil {
				panic("Unexpected error: " + err.Error())
			}
			if m != test.vals[i] {
				t.Errorf("%v with reads %v: read %d gave %x, expected %x\n", test.data, test.ns, i, m, test.vals[i])
				break
			}
		}
	}
}

func TestUint64(t *testing.T) {
	tests := []struct {
		data []byte
		ns   []uint
		vals []uint64
	}{
		{[]byte{0xFF}, []uint{1, 1, 1, 1, 1, 1, 1, 1}, []uint64{1, 1, 1, 1, 1, 1, 1, 1}},
		{[]byte{0xFF}, []uint{2, 2, 2, 2}, []uint64{0x3, 0x3, 0x3, 0x3}},
		{[]byte{0xFF}, []uint{3, 3, 2}, []uint64{0x7, 0x7, 0x3}},
		{[]byte{0xFF}, []uint{4, 4}, []uint64{0xF, 0xF}},
		{[]byte{0xFF}, []uint{5, 3}, []uint64{0x1F, 0x7}},
		{[]byte{0xFF}, []uint{6, 2}, []uint64{0x3F, 0x3}},
		{[]byte{0xFF}, []uint{7, 1}, []uint64{0x7F, 0x1}},
		{[]byte{0xFF}, []uint{8}, []uint64{0xFF}},

		{[]byte{0xAA}, []uint{1, 1, 1, 1, 1, 1, 1, 1}, []uint64{1, 0, 1, 0, 1, 0, 1, 0}},
		{[]byte{0xAA}, []uint{2, 2, 2, 2}, []uint64{0x2, 0x2, 0x2, 0x2}},
		{[]byte{0xAA}, []uint{3, 3, 2}, []uint64{0x5, 0x2, 0x2}},
		{[]byte{0xAA}, []uint{4, 4}, []uint64{0xA, 0xA}},
		{[]byte{0xAA}, []uint{5, 3}, []uint64{0x15, 0x2}},
		{[]byte{0xAA}, []uint{6, 2}, []uint64{0x2A, 0x2}},
		{[]byte{0xAA}, []uint{7, 1}, []uint64{0x55, 0x0}},
		{[]byte{0xAA}, []uint{8}, []uint64{0xAA}},

		{
			[]byte{0xAA, 0x55},
			[]uint{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[]uint64{1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1},
		},

		{
			[]byte{0xAA, 0x55},
			[]uint{7, 8, 1},
			[]uint64{0x55, 0x2A, 0x1},
		},

		{
			[]byte{0xAA, 0x55},
			[]uint{3, 3, 3, 3, 3, 1},
			[]uint64{0x5, 0x2, 0x4, 0x5, 0x2, 0x1},
		},

		{
			[]byte{0xAA, 0x55},
			[]uint{16},
			[]uint64{0xAA55},
		},

		{
			[]byte{0xAA, 0x55, 0xAA, 0x55, 0xAA, 0x55, 0xAA, 0x55},
			[]uint{32, 32},
			[]uint64{0xAA55AA55, 0xAA55AA55},
		},

		{
			[]byte{0xAA, 0x55, 0xAA, 0x55, 0xAA, 0x55, 0xAA, 0x55},
			[]uint{33, 31},
			[]uint64{0x154AB54AB, 0x2A55AA55},
		},
	}

	for _, test := range tests {
		r := NewReader(bytes.NewReader(test.data))
		if len(test.ns) != len(test.vals) {
			panic("Number of reads does not match number of results")
		}
		for i, n := range test.ns {
			m, err := r.Uint64(n)
			if err != nil {
				panic("Unexpected error: " + err.Error())
			}
			if m != test.vals[i] {
				t.Errorf("%v with reads %v: read %d gave %x, expected %x\n", test.data, test.ns, i, m, test.vals[i])
				break
			}
		}
	}
}

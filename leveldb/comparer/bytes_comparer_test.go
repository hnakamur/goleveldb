package comparer

import (
	"bytes"
	"testing"
)

func TestBytesComparerSeparator(t *testing.T) {
	cases := []struct {
		a    []byte
		b    []byte
		want []byte
	}{
		{a: nil, b: nil, want: nil},
		{a: []byte{0x01}, b: []byte{0x01}, want: nil},
		{a: []byte{0x01, 0x02}, b: []byte{0x01}, want: nil},
		{a: []byte{0x01}, b: []byte{0x02}, want: nil},
		{a: []byte{0x01}, b: []byte{0x03}, want: []byte{0x02}},
		{a: []byte{0xfd}, b: []byte{0xff}, want: []byte{0xfe}},
		{a: []byte{0xfe}, b: []byte{0xff}, want: nil},
		{a: []byte{0xff}, b: []byte{0xff}, want: nil},
		{a: []byte{0x01, 0x02}, b: []byte{0x01, 0x01}, want: nil},
		{a: []byte{0x01, 0x02}, b: []byte{0x01, 0x02}, want: nil},
		{a: []byte{0x01, 0x02}, b: []byte{0x01, 0x03}, want: nil},
		{a: []byte{0x01, 0x02}, b: []byte{0x01, 0x04}, want: []byte{0x01, 0x03}},
		{a: []byte{0x01, 0xfd}, b: []byte{0x01, 0xff}, want: []byte{0x01, 0xfe}},
		{a: []byte{0x01, 0xfe}, b: []byte{0x01, 0xff}, want: nil},
		{a: []byte{0x01, 0xff}, b: []byte{0x01, 0xff}, want: nil},
	}
	comp := bytesComparer{}
	for _, c := range cases {
		got := comp.Separator(nil, c.a, c.b)
		if !bytes.Equal(got, c.want) {
			t.Errorf("unexpected result for bytesComparer.Separator. a=%v, b=%v, got=%v, want=%v", c.a, c.b, got, c.want)
		}
	}
}

func TestBytesComparerSuccessor(t *testing.T) {
	cases := []struct {
		b    []byte
		want []byte
	}{
		{b: nil, want: nil},
		{b: []byte{0xff}, want: nil},
		{b: []byte{0xfe}, want: []byte{0xff}},
		{b: []byte{0xff, 0xff}, want: nil},
		{b: []byte{0xff, 0xff, 0xfe}, want: []byte{0xff, 0xff, 0xff}},
	}
	comp := bytesComparer{}
	for _, c := range cases {
		got := comp.Successor(nil, c.b)
		if !bytes.Equal(got, c.want) {
			t.Errorf("unexpected result for bytesComparer.Successor. b=%v, got=%v, want=%v", c.b, got, c.want)
		}
	}
}

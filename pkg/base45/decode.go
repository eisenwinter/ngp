package base45

import (
	"bytes"
	"errors"
)

//Base45Decoder allows the decoing of a base45 encoded string
type Base45Decoder interface {
	Decode(input string) ([]byte, error)
	NewReader(input string) Base45Reader
}

type Base45Reader interface {
	Read(b []byte) (n int, err error)
}

type base45Decoder struct{}
type base45Reader struct {
	decoded []byte
	len     int
	err     error
}

func (r *base45Reader) Read(b []byte) (n int, err error) {
	b = r.decoded
	return r.len, r.err
}

func (*base45Decoder) Decode(input string) ([]byte, error) {
	return b45decode(input)
}

func (*base45Decoder) NewReader(input string) Base45Reader {
	b, err := b45decode(input)
	return &base45Reader{
		decoded: b,
		len:     len(b),
		err:     err,
	}
}

func NewDecoder() Base45Decoder {
	return &base45Decoder{}
}

var set = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:")

func b45decode(s string) ([]byte, error) {
	in := []byte(s)
	len := len(in)
	if len%3 == 1 {
		return nil, errors.New("Invalid base45")
	}
	out := make([]byte, 0)
	for i := 0; i < len; i += 3 {
		if len-i >= 3 {
			var x int
			x = int(bytes.IndexByte(set, in[i]) + bytes.IndexByte(set, in[i+1])*45 + bytes.IndexByte(set, in[i+2])*45*45)
			if x > 0xFFFF {
				return nil, errors.New("Invalid base45 - overflow")
			}
			out = append(out, byte(x/256))
			out = append(out, byte(x%256))
		} else {
			x := bytes.IndexByte(set, in[i]) + bytes.IndexByte(set, in[i+1])*45
			if x > 0xff {
				return nil, errors.New("Invalid base45 - overflow")
			}
			out = append(out, byte(x))
		}
	}
	return out, nil
}

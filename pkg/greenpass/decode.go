package greenpass

import (
	"bytes"
	"compress/zlib"
	"errors"
	"strings"

	"github.com/eisenwinter/ngp/pkg/base45"
	"github.com/eisenwinter/ngp/pkg/euspec"
	"github.com/fxamacker/cbor/v2"
)

//GreenpassDeocer decodes greenpass information (only considering V's - vaccinations)
type GreenpassDecoder interface {
	Decode(input []byte) (*euspec.EuJsonPayload, error)
}

type greenpassDecoder struct{}

func New() GreenpassDecoder {
	return &greenpassDecoder{}
}

type cose struct {
	_           struct{} `cbor:",toarray"`
	Protected   []byte
	Unprotected coseHeader
	Payload     []byte
	Signature   []byte
}

type healtJsonPayload struct {
	Json euspec.EuJsonPayload `cbor:"1,keyasint,omitempty"`
}

type healthJson struct {
	Content healtJsonPayload `cbor:"-260,keyasint,omitempty"`
}

type coseHeader struct {
	Alg int `cbor:"1,keyasint,omitempty"`
}

func (*greenpassDecoder) Decode(input []byte) (*euspec.EuJsonPayload, error) {
	decoder := base45.NewDecoder()
	if !strings.HasPrefix(string(input), "HC1:") {
		return nil, errors.New("Invalid Header (not HC1)")
	}
	decoded, err := decoder.Decode(string(input)[4:])
	if err != nil {
		return nil, errors.New("b45 decoding failed")
	}
	if decoded[0] == 0x78 {
		b := bytes.NewReader(decoded)
		z, err := zlib.NewReader(b)
		if err != nil {
			return nil, err
		}
		cborDecoder := cbor.NewDecoder(z)
		var v cose
		err = cborDecoder.Decode(&v)
		if err != nil {
			return nil, err
		}
		var j healthJson
		err = cbor.Unmarshal(v.Payload, &j)
		if err != nil {
			return nil, err
		}
		return &j.Content.Json, nil
	}
	return nil, errors.New("no zlib header found")
}

package messagepack

import (
	"github.com/andrei-dascalu/shortener/src/shortener"
	"github.com/vmihailenco/msgpack/v5"
)

type Serializer struct{}

func (r *Serializer) Decode(input []byte) (*shortener.LinkRedirect, error) {
	redirect := &shortener.LinkRedirect{}

	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, err
	}

	return redirect, nil
}

func (r *Serializer) Encode(input *shortener.LinkRedirect) ([]byte, error) {
	data, err := msgpack.Marshal(input)
	if err != nil {
		return nil, err
	}

	return data, nil
}

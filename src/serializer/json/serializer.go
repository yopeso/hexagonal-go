package json

import (
	"github.com/andrei-dascalu/shortener/src/shortener"
	jsoniter "github.com/json-iterator/go"
)

type Serializer struct{}

func (r *Serializer) Decode(input []byte) (*shortener.LinkRedirect, error) {
	redirect := &shortener.LinkRedirect{}

	var jsonDeserializer = jsoniter.ConfigCompatibleWithStandardLibrary

	if err := jsonDeserializer.Unmarshal(input, redirect); err != nil {
		return nil, err
	}

	return redirect, nil
}

func (r *Serializer) Encode(input *shortener.LinkRedirect) ([]byte, error) {
	var jsonDeserializer = jsoniter.ConfigCompatibleWithStandardLibrary

	data, err := jsonDeserializer.Marshal(input)

	if err != nil {
		return nil, err
	}

	return data, nil
}

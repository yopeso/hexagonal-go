package shortener

type LinkSerializer interface {
	Decode(input []byte) (*LinkRedirect, error)
	Encode(input *LinkRedirect) ([]byte, error)
}

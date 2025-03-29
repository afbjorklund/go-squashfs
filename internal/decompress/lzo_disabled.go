package decompress

import "errors"

type Lzo struct{}

func NewLzo() (Lzo, error) {
	return Lzo{}, errors.New("lzo compression is disable")
}

func (l Lzo) Decompress(data []byte) ([]byte, error) {
	return nil, errors.New("lzo compression is disable")
}

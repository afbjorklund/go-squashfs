package decompress

import (
	"io"

	"github.com/pierrec/lz4/v4"
)

type Lz4 struct{}

func (l Lz4) Reader(r io.Reader) (io.ReadCloser, error) {
	return io.NopCloser(lz4.NewReader(r)), nil
}

func (l Lz4) Reset(old, src io.Reader) error {
	old.(*lz4.Reader).Reset(src)
	return nil
}

func (l Lz4) Decode(in []byte, outSize int) (out []byte, err error) {
	out = make([]byte, outSize)
	outLen, err := lz4.UncompressBlock(in, out)
	if outLen < outSize {
		out = out[:outLen]
	}
	return
}

package hashlib

import (
	"encoding/hex"
	"io"

	"github.com/codahale/blake2"
)

// Sum hash sum.
func Sum(f io.Reader) []byte {
	hash := blake2.New(nil)
	io.Copy(hash, f)
	return hash.Sum(nil)
}

// String to string
func String(src []byte) string {
	return hex.EncodeToString(src)
}

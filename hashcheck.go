package art

import (
	"crypto/sha256"
	"io/ioutil"
)

func HashCheck(a string) bool {
	hasher := sha256.New()
	s, err := ioutil.ReadFile(a)
	hasher.Write(s)
	CheckErr(err)
	l := hasher.Sum(nil)

	hashfile := []byte{225, 148, 241, 3, 52, 66, 97, 122, 184, 167, 142, 28, 166, 58, 32, 97, 245, 204, 7, 163, 240, 90, 194, 38, 237, 50, 235, 157, 253, 34, 166, 191}

	return string(l) == string(hashfile)
}

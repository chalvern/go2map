package go2map

import (
	"io/ioutil"
)

// mustReadFile panic if reading failed.
func mustReadFile(fileName string) []byte {
	contentsByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return contentsByte
}

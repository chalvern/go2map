package go2map_test

import (
	"fmt"
	"testing"

	"github.com/chalvern/go2map"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
  1: test
  2:
  - 21
  - 22
`

func TestYaml2Map(t *testing.T) {
	m := go2map.Yaml2Map([]byte(data))
	fmt.Printf("m is: %#v", m)
}

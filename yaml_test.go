package go2map_test

import (
	"testing"

	"github.com/chalvern/go2map"
	"github.com/stretchr/testify/assert"
)

func TestYaml2Map(t *testing.T) {

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
	m := go2map.Yaml2Map([]byte(data))
	keys := []string{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	assert.Contains(t, keys, "a")
	assert.Contains(t, keys, "b.c")
	assert.Contains(t, keys, "b.d")
	assert.Contains(t, keys, "b.1")
	assert.Contains(t, keys, "b.2")
}

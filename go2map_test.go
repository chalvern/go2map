package go2map_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chalvern/go2map"
)

func TestFlatMapGetStringOf(t *testing.T) {
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
	b1, err := m.GetStringOf("b.1")
	assert.Nil(t, err)
	assert.Equal(t, "test", b1)

}

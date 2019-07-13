package main

import (
	"fmt"

	"github.com/chalvern/go2map"
)

func main() {
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
	b1, _ := m.GetStringOf("b.1")
	fmt.Println(b1)
}

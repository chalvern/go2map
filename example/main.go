package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/chalvern/go2map"
)

func main() {
	// convert from bytes
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
	fmt.Println(m.MustGetStringOf("b.100"))

	// convert from filePath
	_, file, _, _ := runtime.Caller(0)
	yamlFile := path.Join(path.Dir(file), "example.yaml")
	m2 := go2map.Yaml2MapFromFile(yamlFile)
	jingweiLink, _ := m2.GetStringOf("zh.blog.jingwei.link")
	fmt.Println(jingweiLink)
	fmt.Println(m2.MustGetStringOf("zh.jingwei.blog.jingwei.link"))
}

# go2map

[English doc](./README.md)

在 Golang 项目的开发过程中，我发现 YAML 的使用很令人揪心。可能考虑到效率的问题吧（？），一般在使用 yaml 的时候往往需要显示地定义结构（struct），然后把 yaml 文件映射成为一个结构对象进行使用。**然而，当使用 yaml 文件的时候往往包含了非常多的条目，显然这个时候的 struct 就太难定义了。**

本项目希望能够简化 YAML 的使用——**避免定义复杂的 struct， 同时能够便利地读取 yaml 中的值**。

```yaml
a: Easy!
b:
  c: 2
  d: [3, 4]
  1: test
  2:
  - 21
  - 22
```

比如上面的例子，最简单的方式是通过一次加载（不必定义 struct 的情况下）然后就可以通过 `x.a`、`x.b.c`、`x.b.1` 类似的方式获取到特定路径的值。


## 待跟进

* FlatMap 的方法目前只根据需要实现了部分，待进一步完善（欢迎提交 merge 贡献力量）。
* 完善英文文档（由于本人以中文表达为主，同时精力有限，英文文档无法及时跟进，如果对本项目感兴趣的话欢迎帮忙翻译）。
* 扩展 JSON 类似的转变功能。


## 使用例子

具体可参考 [源文件](./example/main.go) 的内容。

```go
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

```

errwrap
[![Go](https://github.com/issue9/errwrap/workflows/Go/badge.svg)](https://github.com/issue9/errwrap/actions?query=workflow%3AGo)
[![license](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)
[![codecov](https://codecov.io/gh/issue9/errwrap/branch/master/graph/badge.svg)](https://codecov.io/gh/issue9/errwrap)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/issue9/errwrap)](https://pkg.go.dev/github.com/issue9/errwrap)
![Go version](https://img.shields.io/github/go-mod/go-version/issue9/errwrap)
======

errwrap 提供了对常用对象需要连续处理 error 的简单封闭。

```go
buf := new(bytes.Buffer)
w := errwrap.Writer{
    Writer: buf,
}

w.WString("string").
    WBytes([]byte("bytes")).
    Printf("format %d", 123)
if w.Err != nil { // 由此处统一处理错误
    // TODO
}
```

安装
----

```shell
go get github.com/issue9/errwrap
```

版权
----

本项目采用 [MIT](http://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。

# A Tour of Goについて

## インストール方法
公式サイト日本語版(<https://go-tour-jp.appspot.com/welcome/3>)に記載してある  
*go get github.com/atotto/go-tour-jp/gotour*をじっこうすると
```bash
$ go get github.com/atotto/go-tour-jp/gotour
go: go.mod file not found in current directory or any parent directory.
        'go get' is no longer supported outside a module.
        To build and install a command, use 'go install' with a version,
        like 'go install example.com/cmd@latest'
        For more information, see https://golang.org/doc/go-get-install-deprecation
        or run 'go help get' or 'go help install'.
```
となるので英語版公式サイト(<https://go.dev/tour/welcome/3>)にある
*go install golang.org/x/website/tour@latest*
を実行する必要があります。

## 実行方法

```bash
/go/bin/tour
```
でブラウザが開くのでそこで実行可能。

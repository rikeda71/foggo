# foggo

[![ci](https://github.com/s14t284/foggo/actions/workflows/ci.yml/badge.svg)](https://github.com/s14t284/foggo/actions/workflows/ci.yml)
[![Release](https://github.com/s14t284/foggo/actions/workflows/release.yml/badge.svg)](https://github.com/s14t284/foggo/actions/workflows/release.yml)
[![Coverage Status](https://coveralls.io/repos/github/s14t284/foggo/badge.svg?branch=main)](https://coveralls.io/github/s14t284/foggo?branch=main)


Golang の構造体から `Functional Option Pattern` コードを自動生成する cli 

## Installation

```shell
$ go install golang.org/x/tools/cmd/goimports@latest  # foggo use 'goimports' command
$ go install github.com/s14t284/foggo@latest
```

## Usage

__foggo__ では `fop` と `afop` サブコマンドを提供しています。

```shell
Usage:
  foggo (fop|afop) [flags]

Flags:
  -h, --help   help for fop

Global Flags:
  -p, --package string   Package name having target struct (default ".")
  -s, --struct string    Target struct name (required)
```

### Generate with command line

1. 以下のように構造体を用意します。optionを提供したくないフィールドには構造体タグに `foggo:"-"` を付与してください。

    ```go
    // ./image/image.go
    package image
    
    type Image struct {
        Width  int
        Height int
        Src    string `foggo:"-"`
        Alt    string
    }
    ```

2. `foggo fop` コマンドを実行。

    ```shell
    # struct パラメータには構造体名を指定します
    # package パラメータには構造体が配置されている相対パスを指定します
    $ foggo fop --struct Image --package image
    ```

3. `foggo` コマンドにより、以下のような Functional Option Pattern のコードが `./image/image_gen.go` に自動生成されます。

    ```go
    // Code generated by foggo; DO NOT EDIT.

    package image

    type ImageOption func(*Image)

    func NewImage(options ...ImageOption) *Image {
        s := &Image{}
    
        for _, option := range options {
            option(s)
        }
    
        return s
    }
    
    func WithWidth(Width int) ImageOption {
        return func(args *Image) {
            args.Width = Width
        }
    }
    
    func WithHeight(Height int) ImageOption {
        return func(args *Image) {
            args.Height = Height
        }
    }

    func WithAlt(Alt string) ImageOption {
        return func(args *Image) {
            args.Alt = Alt
        }
    }
    ```
   
4. あとは `Functional Option Pattern` を使って実装するだけです。

    ```go
    package main
   
    import "github.com/user/project/image"
    
    func main() {
	    image := NewImage(
	    	WithWidth(1280),
	    	WithHeight(720),
	    	WithAlt("alt title"),
	    )
	    image.Src = "./image.png"
        ...
    }
    ```

### Generate with `go:generate`

1. `go:generate` コメントを付与して構造体を定義します。packageパラメータには何も指定しなくて大丈夫です。

    ```go
    // ./image/image.go
    package image
    
    //go:generate foggo fop --struct Image 
    type Image struct {
        Width  int
        Height int
        // don't want to create option, specify `foggo:"-"` as the structure tag 
        Src    string `foggo:"-"`
        Alt    string
    }
    ```

2. `go generate ./...` コマンドを実行してください。

    ```shell
    $ go generate ./...
    ```

3. `go:generate foggo` コメントが付与されている全ての構造体に対して `Functional Option Pattern` のコードが自動生成されます。

### Generate with `afop` command

`afop` は `Applicable Functional Option Pattern` のコードを自動生成するコマンドです。

1. `go:generate` コメントを付与して構造体を定義します。サブコマンドとして `afop` コマンドを指定します

    ```go
    // ./image/image.go
    package image
    
    //go:generate foggo afop --struct Image 
    type Image struct {
        Width  int
        Height int
        // don't want to create option, specify `foggo:"-"` as the structure tag 
        Src    string `foggo:"-"`
        Alt    string
    }
    ```

2. `go generate ./...` コマンドを実行してください。

    ```shell
    $ go generate ./...
    ```

3. `go:generate foggo` コメントが付与されている全ての構造体に対して `Applicable Functional Option Pattern` のコードが自動生成されます。

    ```go
    // Code generated by foggo; DO NOT EDIT.

    package image

    type ImageOption interface {
        apply(*Image)
    }

    type WidthOption struct {
        Width int
    }

    func (o WidthOption) apply(s *Image) {
        s.Width = o.Width
    } 

    type HeightOption struct {
        Height int
    }

    func (o HeightOption) apply(s *Image) {
        s.Height = o.Height
    }

    type AltOption struct {
        Alt string
    }

    func (o AltOption) apply(s *Image) {
        s.Alt = o.Alt
    }

    func NewImage(options ...ImageOption) *Image {
        s := &Image{}

        for _, option := range options {
            option.apply(s)
        }

        return s
    }
    ```

4. あとは `Applicable Functional Option Pattern` を使って実装するだけです。

    ```go
    package main
   
    import "github.com/user/project/image"
    
    func main() {
        image := NewImage(
            WidthOption(1280),
            HeightOption(720),
            AltOption("alt title"),
        )
        image.Src = "./image.png"
        ...
    }
    ```


## Functional Option Pattern ?
`Functional Option Pattern`(`FOP`) は Golang でよく使われるデザインパターンの一種です。

Golang では python や ruby で利用できるキーワード引数のようなオプション引数を提供していません。
`FOP` を使うことで、オプション引数を再現します。

以下の記事が詳しいです。

- [Goでオプショナルパラメータをどう扱うか](https://raahii.github.io/posts/optional-parameters-in-go/)
- [Functional Options Pattern に次ぐ、オプション引数を実現する方法](https://ww24.jp/2019/07/go-option-pattern)
- https://commandcenter.blogspot.jp/2014/01/self-referential-functions-and-design.html
- https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

### Applicable Functional Option Pattern ?

`Applicable Functional Option Pattern`(`AFOP`) は生成されるオプションがテスト可能な `FOP` です。
`FOP` ではオプションを関数として定義します。
そのため、同一引数を持つオプション関数同士を比較しても等しくないと判定されてしまいます。

`AFOP` ではオプションを 単一のパラメータを持ち、`apply` メソッドを実装した構造体として定義します。
Go言語では構造体は比較可能であるため、同一引数を持つオプション同士を比較することができます。（すなわちテスト可能です）

`AFOP` については以下の記事が詳しいです。

- [Functional Options Pattern に次ぐ、オプション引数を実現する方法](https://ww24.jp/2019/07/go-option-pattern) 
  - `Applicable Functional Option Pattern` はこの記事で命名されているため名称を拝借しました
- https://github.com/uber-go/guide/blob/master/style.md#functional-options

## References

- https://github.com/moznion/gonstructor
# foggo

__foggo__ is generator of `Functional Option Pattern` from struct field in Golang code.

## Installation

```shell
$ go install golang.org/x/tools/cmd/goimports@latest  # foggo use 'goimports' command
$ go install github.com/s14t284/foggo@latest
```

## Usage

__foggo__ provides `foc` subcommand. (`afoc` subcommand will be provided in the future)

```shell
Usage:
  foggo foc [flags]

Flags:
  -h, --help   help for foc

Global Flags:
  -p, --package string   Package name having target struct (default ".")
  -s, --struct string    Target struct name (required)
```

### Generate with command line

1. prepare a struct type.

    ```go
    // ./image/image.go
    package image
    
    type Image struct {
        Width  int
        Height int
        // don't want to create option, specify `foggo:"-"` as the structure tag 
        Src    string `foggo:"-"`
        Alt    string
    }
    ```

2. execute `foggo foc` command.

    ```shell
    # struct must be set struct type name 
    # package must be package path
    $ foggo foc --struct Image --package image
    ~~~ success to write functional option pattern code to /path/to/image_gen.go
    ```

3. then `foggo` generates functional option pattern code to `./image/image_gen.go`.

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
   
4. write Golang code using `functional option parameter`

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

1. prepare a struct type with `go:generate`.

    ```go
    // ./image/image.go
    package image
    
    //go:generate foggo foc --struct Image 
    type Image struct {
        Width  int
        Height int
        // don't want to create option, specify `foggo:"-"` as the structure tag 
        Src    string `foggo:"-"`
        Alt    string
    }
    ```

2. execute `go generate ./...` command.

    ```shell
    $ go generate ./...
    ```

3. the `foggo` generate functional option pattern code to all files written `go:generate`. 

### go:generate mode

## Functional Option Pattern ?
`Functional Option Pattern` is one of the most common design patterns used in Golang code.

Golang cannot provide optional arguments such as keyword arguments (available in python, ruby, ...).
`Functional Option Pattern` is the technique for achieving optional arguments.

For more information, please refer to the following articles.

- https://commandcenter.blogspot.jp/2014/01/self-referential-functions-and-design.html
- https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

## References

- https://github.com/moznion/gonstructor

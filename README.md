# Go gdub
Based on the original [gdub](https://github.com/dougborg/gdub), this project is a golang gradle wrapper wrapper :smile:. 
As the name suggests it is a wrapper for the gradle wrapper which allows you to execute gradle tasks from anywhere in your project

## Pros and Cons
I recommend heading over to [gdub](https://github.com/dougborg/gdub) as it explained really well over there and anything I add will probably have come from there.

## Why? gdub already exists
As much as I love gdub, it really isn't a viable option for someone who wants a pure windows shell (no babun, cygwin or minGW) and so the main reason I wrote this was
to allow me to have the gdub functionality on windows specifically powershell.

## Why Go?
Originally the plan was to write this in powershell or batch as the main goal was to run in a windows shell but as I am learning Go I thought it would be a good exercise for me to practice my golang. Go can
also build binaries for practically all environments and architectures and so it was the perfect language for what I was trying to do.

## Installation

### macOS / Linux

```shell
curl -sL https://github.com/aaron-vaz/gw/raw/master/install.sh | bash
```

### Alternative methods
#### Pre-built Binaries
There are pre-built [binaries](https://github.com/aaron-vaz/gw/releases/latest) available,
simply download the binary for your system rename and place in a location on the `$PATH`

#### Go get
Another way to install go-gdub is to use golang's inbuilt `go get`
> go get github.com/aaron-vaz/gw/cmd/gw

Though this does require you to have `go` installed and a valid `$GOPATH` environment variable and `$GOPATH/bin` added to your `$PATH`. please refer to the [go docs](https://golang.org/doc/install) for help

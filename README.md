[![build status](https://gitlab.com/aaronvaz/gw/badges/master/build.svg)](https://gitlab.com/aaronvaz/gw/commits/master) [![coverage report](https://gitlab.com/aaronvaz/gw/badges/master/coverage.svg)](https://gitlab.com/aaronvaz/gw/commits/master)

# Go gdub
Based on the orginal [gdub](https://github.com/dougborg/gdub), this project is a golang gradle wrapper wrapper :smile:. 
As the name suggests it is a wrapper for the gradle wrapper which allows you to execute gradle tasks from anywhere in your project

## Pros and Cons
I recommed heading over to [gdub](https://github.com/dougborg/gdub) as it explained really well over there and anything I add will probably have come from there.

## Why? gdub already exists
As much as I love gdub, it really isnt a viable option for someone who wants a pure windows shell (no babun, cygwin or minGW) and so the main reason I wrote this was
to allow me to have the gdub functionality on windows specifically powershell.

## Why Go?
Originally the plan was to write this in powershell or batch as the main goal was to run in a windows shell but as I am learning Go I thought it would be a good exercise for me to pratice my golang. Golang can
also build binaries for pratically all environments and architectures and so it was the perfect language for what I was trying to do.

## Installation
### Pre-built binaries
Currently the easiest way to install Go gdub is using the [built binaries](https://gitlab.com/aaronvaz/gw/builds). Simply download the binary for your system, rename it and then add it to your `PATH`

### Go get
Another way is use golang's inbuilt `go get`
>go get gitlab.com/aaronvaz/gw

Though this does require you to have `go` installed and a valid `GOPATH` environment variable and `GOPATH/bin` added to your `PATH`. please refer to the [go docs](https://golang.org/doc/install) for help

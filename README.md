# goenv

## What?

Goenv is a go package that provides virtualenv-like functionality for Go.

## Why?

Because global dependencies are evil.

## Features

- Identical basic functionality as virtualenv
- Written (mostly) in Go
- Installable as a Go package
- No need to add any source / eval code to your shell init script

## Quick start

First, ensure your `PATH` includes the /bin directory in your global `GOPATH`, with something like:

```script
PATH=PATH:$GOPATH/bin
```

Install this package

```
$ go get github.com/crsmithdev/goenv
```

Create (or enter) a directory and set up a goenv:

```
$ mkdir myproject
$ cd myproject
$ goenv local
```

Activate the goenv:

```
$ . local/bin/activate
```

Packages installed with `go get` will now be installed in the `local` directory.

Deactivate the goenv:

```
deactivate
```

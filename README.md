# goenv

[![Travis](https://travis-ci.org/crsmithdev/goenv.svg?branch=master)](https://travis-ci.org/crsmithdev/goenv)
[![GoDoc](https://godoc.org/github.com/crsmithdev/goenv?status.svg)](https://godoc.org/github.com/crsmithdev/goenv)

## What?

Goenv is a Go package that provides virtualenv-like functionality for Go projects, isolating dependencies into workspaces for safer and simpler management.

## Why?

Go's package management expects that all go source files and packages will exist under a single system directory, GOPATH.  This makes it easy to install and find packages, but means that any  projects being worked on will share the same GOPATH, which can cause issues if different versions of packages are required and make it difficult to create isolated, reproducible builds.

## Features

- Similar functionality to Python's virtualenv and virtualenvwrapper.
- Separates development directory from import path - e.g., develop in `~/myproject`, but import from `github.com/me/myproject`
- Isolates dependencies from other projects.
- Does not interfere with any `go` command functionality.
- Written in Go, installable with `go get`.

## Quick start

First, ensure your `PATH` includes the /bin directory in your global `GOPATH`, with something like:

```bash
PATH=PATH:$GOPATH/bin
```

Install this package:

```bash
$ go get github.com/crsmithdev/goenv
```

Create or enter your project directory:

```bash
$ mkdir -p ~/myproject && cd ~/myproject
```

Create a goenv:

```bash
$ goenv myproject github.com/me/myproject
```

Activate the goenv:

```bash
$ . goenv/activate
```

Install packages with `go get` or other dependency managment tools.

```bash
(myproject) $ go get github.com/hoisie/redis
```

When finished, deactivate the goenv:

```bash
(myproject) $ deactivate
```

## Todo

- `destroy`command
- Testing

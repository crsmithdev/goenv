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
go get github.com/crsmithdev/goenv
```

Within your project directory, reate a goenv:

```bash
goenv github.com/me/myproject
```

Activate the goenv:

```bash
. goenv/activate
```

Install packages with `go get` or other dependency managment tools.

```bash
go get github.com/hoisie/redis
```

When finished, deactivate the goenv:

```bash
deactivate
```

## Commands

### Init

```bash
Usage: goenv init [-g][-s][-p][-n] [import path]

Init initializes a goenv and creates an initialization script that
activates it.  This script creates, if needed, a GOPATH directory
structure, symlinks the project into that structure at the specified
input path, and then alters the current session's GOPATH environment
variable to point to it.

The goenv can be deactivated with 'deactivate'.

Init supports the following options:

    -n
        the name of the environment, defaulting to the name
        of the current working directory.

    -g
        the GOPATH to create, defaulting to ~/.goenv/<name>

    -p
        the project path, defaulting to the current working
        directory.

    -s
        the full path to the initialization script, defaulting
        to ./goenv/activate.
```



## Todo

- other commands?
- tests

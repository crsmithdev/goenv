# goenv

## What?

Goenv is a Go package that provides virtualenv-like functionality for Go projects, isolating dependencies into workspaces for safer and simpler management.

## Why?

Go's package management expects that all go source files and packages will exist under a single system directory, GOPATH.  This makes it easy to install and find packages, but means that any  projects being worked on will share the same GOPATH, which can cause issues if different versions of packages are required and make it difficult to create isolated, reproducible builds.

## Features

- Similar functionality to Python's virtualenv and virtualenvwrapper.
- Separates development directory from import path - e.g., develop in `~/myproject`, but import from `github.com/me/myproject`
- Isolates dependencies from other projects.
- Written in Go, installable with `go get`.

## Quick start

First, ensure your `PATH` includes the /bin directory in your global `GOPATH`, with something like:

```shell
PATH=PATH:$GOPATH/bin
```

Install this package:

```shell
$ go get github.com/crsmithdev/goenv
```

Create or enter your project directory:

```shell
$ mkdir -p ~/myproject && cd ~/myproject
```

Create a goenv:

```
$ goenv myproject github.com/me/myproject
```

Activate the goenv:

```
$ . goenv/activate
```

Install other packages with `go get` or other dependency managment tools.

```
(myproject) $ go get github.com/hoisie/redis
```

Finally, when finished, deactivate the goenv:

```
(myproject) $ deactivate
```

Your GOPATH is now back to what it was before.

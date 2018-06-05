Red Hat Container Manager
=========================

This program allows you to retrieve information about the container project form Red Hat Connect and trigger manual builds without needing to login in the website.

## Requirements

- [Golang](https://golang.org/)

## Build and Install

```sh
go get github.com/tembleking/rhc-manager
go install github.com/tembleking/rhc-manager/cmd/rhc-manager
```

## Usage

You can find your `$GOPATH` with `go env | grep GOPATH`

**Note**: it's recommended to add `$GOPATH/bin` to your `$PATH` variable.

Get project info: 

```sh
$GOPATH/bin/rhc-manager -id <project_id>
```

Trigger manual build:

```sh
$GOPATH/bin/rhc-manager -id <project_id> -b <build_tag>
```


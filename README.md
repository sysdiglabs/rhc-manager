Red Hat Container Manager
=========================

This program allows you to retrieve information about the container project form Red Hat Connect and trigger manual builds without needing to log in to the website.

## Requirements

- [Golang](https://golang.org/)

## Build and Install

```sh
go get github.com/sysdiglabs/rhc-manager
go install github.com/sysdiglabs/rhc-manager/cmd/rhc-manager
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

## Testing

You can execute the automated tests with:

```sh
go test -v
```

And also checkout the code coverage with:

```sh
go test -coverprofile cover.out
go tool cover -html cover.out
```

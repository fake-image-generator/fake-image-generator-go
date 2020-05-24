# fake-image-generator-go
[![GoDoc](https://godoc.org/github.com/fake-image-generator/fake-image-generator-go?status.svg)](https://godoc.org/github.com/fake-image-generator/fake-image-generator-go)![Go version](https://img.shields.io/github/go-mod/go-version/fake-image-generator/fake-image-generator-go) [![Github Actions](https://img.shields.io/github/workflow/status/fake-image-generator/fake-image-generator-go/Go)](https://github.com/fake-image-generator/fake-image-generator-go/actions?query=workflow%3AGo)



<img align="left" width="100" height="100" src="fake-image-generator.png">

Generate a fake JPG or PNG image in any size between 1 KB and 2 GB.

This was made for generating big files that pose as images for testing purposes (for testing an image upload feature in an app, for example).

### Requirements

This utility was tested in Go 1.13.4.

### Installation

```
go install github.com/fake-image-generator/fake-image-generator-go
```

### Build

```
go get -d ./...
go build
```

### Usage

```
fake-image-generator-go -size=15500 -extension=png -output=C:/
```
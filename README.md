# fake-image-generator-go
Generate a fake JPG or PNG image in any size between 1 KB and 2 GB.

This was made for generating big files that pose as images for testing purposes (for testing an image upload feature in an app, for example).

### Requirements

This utility was tested in Go 1.13.4.

### Installation

TODO

### Build

```
go get -d ./...
go build
```

### Usage

```
fake-image-generator-go -size=15500 -extension=png -output=C:/
```
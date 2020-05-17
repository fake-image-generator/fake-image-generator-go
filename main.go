// Small utility that generates a fake JPG or PNG image in any size between 1 KB and 2 GB.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	size := flag.String("size", "1000", "Size in bytes for the fake image")
	extension := flag.String("extension", "jpg", "Extension to generate the image. Available options are `jpg` and `png`")
	output := flag.String("output", "C:", "Output path for the fake image")

	flag.Parse()

	i, err := strconv.ParseInt(*size, 10, 32)
	check(err)

	if i < 1000 {
		err := errors.New("Size must be greater than 1000 bytes (1 KB)")
		check(err)
	}

	bytes := make([]byte, i)

	f, err := os.Open(fmt.Sprintf("Untitled.%s", *extension))
	check(err)

	buffer := make([]byte, 600)
	_, err = f.Read(buffer)
	check(err)

	for i, _ := range buffer {
		bytes[i] = buffer[i]
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/image.%s", *output, *extension), bytes, 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

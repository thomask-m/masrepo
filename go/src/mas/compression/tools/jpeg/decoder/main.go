package main

import (
	"fmt"
	_ "image"
	_ "image/jpeg"
	"os"

	jpg "github.com/thomask-m/masrepo/go/src/mas/compression/jpeg"
)

func main() {
	// example picture we'll decompress for now
	// TODO: the CLI tool should take in a path to a valid jpeg file
	r, err := os.Open("/Users/thomaskim/Downloads/SampleJPGImage_50kbmb.jpg")
	if err != nil {
		fmt.Println("something went wrong opening the file")
		return
	}
	defer r.Close()

	_, err = jpg.Decode(r)
	if err != nil {
		fmt.Printf("something went wrong decoding the image: %s\n", err)
		return
	}
}

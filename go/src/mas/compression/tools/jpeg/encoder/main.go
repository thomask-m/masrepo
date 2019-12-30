package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"

	jpg "github.com/thomask-m/masrepo/go/src/mas/compression/jpeg"
)

func main() {
	fmt.Println("lets get this to compile and run")

	// plan of attack:
	// we'll decode a regular jpeg image using image.Decode, get the image struct from it
	// and then we'll pass it back into the jpg.Encode call and see if we get the same jpeg image back.

	// There will probably need to be an actual validation layer that sees if the images actually hold
	// the same data but I'm honestly not sure how to do that yet....
	r, err := os.Open("/Users/thomaskim/Downloads/SampleJPGImage_50kbmb.jpg")
	if err != nil {
		fmt.Println("something went wrong opening the file")
		return
	}
	defer r.Close()

	m, s, err := image.Decode(r)
	if err != nil {
		fmt.Printf("something went wrong decoding the image: %s\n", err)
		return
	}
	fmt.Printf("image format: %s\n", s)
	jpg.Encode(os.Stdout, m, nil)
}

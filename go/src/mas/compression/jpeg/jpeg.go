package jpeg

// This library implements the same API as Go standard library "image/jpeg"
import (
	"fmt"
	"image"
	"image/color"
	"io"
)

const DefaultQuality = 75

type FormatError string

type UnsupportedError string

type Options struct {
	Quality int
}

type img struct {
	// fill in this struct
}

type colormodel struct {
	// fill in this struct
}

type kolor struct {
	// fill in this struct
}

func (k kolor) RGBA() (r, g, b, a uint32) {
	// fill in this struct
	return 0, 0, 0, 1
}

func (cm colormodel) Convert(c color.Color) color.Color {
	// fill in this func
	return c
}

func (i img) ColorModel() color.Model {
	// fill in this func
	return colormodel{}
}

func (i img) Bounds() image.Rectangle {
	// fill in this func
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: 2,
			Y: 2,
		},
	}
}

func (i img) At(x, y int) color.Color {
	// fill in this func
	return kolor{}
}

func Decode(r io.Reader) (image.Image, error) {
	mdata := make([]byte, 1000)
	_, err := r.Read(mdata)
	count := 0
	for err == nil {
		_, err = r.Read(mdata)
		for i := 0; i < 1000; i++ {
			fmt.Printf("%d ", mdata[i])
		}
		count++
	}
	if err != io.EOF {
		return img{}, err
	}
	return img{}, nil
}

func DecodeConfig(r io.Reader) (image.Config, error) {
	// fill in this func
	return image.Config{
		ColorModel: colormodel{},
		Width:      0,
		Height:     0,
	}, nil
}

func rgbToYCbCr(r, g, b uint8) (uint8, uint8, uint8) {
	r1 := float64(r)
	g1 := float64(g)
	b1 := float64(b)

	// "good enough" coefficients for converting from rgb to y'cbcr
	yy := 0.299*r1 + 0.587*g1 + 0.114*b1
	cb := 128 - (0.168736 * r1) - (0.331264 * g1) + (0.5 * b1)
	cr := 128 + (0.5 * r1) - (0.418688 * g1) - (0.081312 * b1)
	return uint8(yy), uint8(cb), uint8(cr)
}

func Encode(w io.Writer, m image.Image, o *Options) error {
	mrect := m.Bounds()
	mdata := [][]uint8{}
	// step 1: convert rgb values to ycbcr
	for y := mrect.Min.Y; y < mrect.Max.Y; y++ {
		for x := mrect.Min.X; x < mrect.Max.X; x++ {
			r, g, b, _ := m.At(x, y).RGBA()
			y, cb, cr := rgbToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			mdata = append(mdata, []uint8{y, cb, cr})
		}
	}
	// step 2: downsample using 4:2:0 scheme
	return nil
}

func (e FormatError) Error() string {
	// fill in this func
	return ""
}

func (e UnsupportedError) Error() string {
	// fill in this func
	return ""
}

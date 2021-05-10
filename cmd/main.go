package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/uberswe/centertext"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "./assets/OpenSans-Regular.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", 125, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
	text     = flag.String("text", "text", "The text you would like to add to the image")
)

func main() {
	flag.Parse()
	fmt.Printf("Loading fontfile %q\n", *fontfile)
	b, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := truetype.Parse(b)
	if err != nil {
		log.Println(err)
		return
	}

	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{},
		Max: image.Point{
			X: 640,
			Y: 640,
		},
	})

	out, err := centertext.OnImage(img, *f, *text)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Save that RGBA image to disk.
	outFile, err := os.Create("out.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	bf := bufio.NewWriter(outFile)
	err = png.Encode(bf, out)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = bf.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Wrote out.png OK.")

}

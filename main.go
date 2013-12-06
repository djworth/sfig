package main

import (
	"flag"
	"image"
	"image/gif"
	"log"
	"os"
)

var path string
var output string

func init() {
	flag.StringVar(&path, "path", "", "Path to the gif file to reverse")
	flag.StringVar(&output, "output", "reversed.gif", "Reversed version of the gif")
}

//Return a the reversed gif.GIF
func reverse(g gif.GIF) *gif.GIF {

	reversed := &gif.GIF{Image: []*image.Paletted{}, Delay: g.Delay, LoopCount: g.LoopCount}
	length := len(g.Image)

	for i := length; i > 0; i-- {
		g.Image[length-i] = g.Image[i-1]
	}

	return reversed
}

func main() {

	flag.Parse()

	file, err := os.Open(path) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	if image, err := gif.DecodeAll(file); err == nil {
		reverse(*image)

		if f, err := os.Create(output); err == nil {
			defer f.Close()
			gif.EncodeAll(f, image)
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

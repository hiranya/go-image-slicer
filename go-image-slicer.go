package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"strings"
)

var flag_x_slices, flag_y_slices int
var flag_filename string

func Init() {
	log.Println("go-image-slicer: Slices an image in to individual image files")
	flag.IntVar(&flag_x_slices, "hslices", 0, "[Required] The number of horizontal slices/blocks you require")
	flag.IntVar(&flag_y_slices, "vslices", 0, "[Required] The number of vertical slices/blocks you require")
	flag.StringVar(&flag_filename, "file", "", "[Required] Filename of the JPEG image file to slice")
	flag.Parse()

	if flag_x_slices == 0 || flag_y_slices == 0 || flag_filename == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	Init()

	reader, err := os.Open(flag_filename)
	outputFileNameFormat := "%03d.jpg"
	outputFileNameIndex := 1

	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	log.Print("Slicing: ", flag_filename)

	x_width, y_height := bounds.Dx()/flag_x_slices, bounds.Dy()/flag_y_slices

	imgNew := image.NewRGBA(image.Rect(0, 0, x_width, y_height))
	draw.Draw(imgNew, imgNew.Bounds(), img, image.Point{0, 480}, draw.Src)

	dirName := strings.Split(reader.Name(), ".")[0]

	err = os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	for y := 0; y < bounds.Dy(); y += y_height {
		for x := 0; x < bounds.Dx(); x += x_width {

			draw.Draw(imgNew, imgNew.Bounds(), img, image.Point{x, y}, draw.Src)

			outFilename := fmt.Sprintf(dirName+"/"+outputFileNameFormat, outputFileNameIndex)
			outFile, err := os.Create(outFilename)
			outputFileNameIndex++

			if err != nil {
				log.Fatal(err)
			}
			defer outFile.Close()

			jpeg.Encode(outFile, imgNew, &jpeg.Options{100})
		}
	}

}

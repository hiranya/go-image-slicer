package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"strings"
)

func main() {
	x_slices := 9
	y_slices := 5

	reader, err := os.Open("ContactSheet-001.jpg")
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
	log.Print(bounds)

	x_width, y_height := bounds.Dx()/x_slices, bounds.Dy()/y_slices

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

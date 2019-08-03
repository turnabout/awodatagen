// Generates sprite sheets & data used by AWO (outputs visuals.json & spritesheet.png)
package main

import (
	"image"
	"image/draw"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

const ssOutputDefault string = "./spritesheet.png"
const awoSSEnvVariable string = "AWO_SPRITESHEET"

func main() {
	// Grab some existing images
	img1 := getImage("./raw_inputs/units/AntiAir/0/1.png")
	img2 := getImage("./raw_inputs/units/AntiAir/0/2.png")

	// Create output image
	outputImg := image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: 255, Y: 255},
	})

	// Grab rectangle for both images
	img1Rect := img1.Bounds()
	img2Rect := img2.Bounds()

	// Move image 2 to the right by 16 pixels
	img2Rect.Min.X += 16
	img2Rect.Max.X += 16

	// Draw the two images into the output image
	draw.Draw(outputImg, img1Rect, img1, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(outputImg, img2Rect, img2, image.Point{X: 0, Y: 0}, draw.Src)

	// Export the output image
	// Use either the AWO spritesheet env variable path or the default output path
	var output string
	var exists bool

	if output, exists = os.LookupEnv(awoSSEnvVariable); !exists {
		output = ssOutputDefault
	}

	writeImage(output, outputImg)
}

// Gets the image stored at the given path
func getImage(path string) image.Image {
	// Load img file
	imgFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	// Decode img
	img, _, err := image.Decode(imgFile)

	if err != nil {
		log.Fatal(err)
	}

	return img
}

// Write a given image to the given path
func writeImage(path string, outputImg image.Image) {
	out, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	if png.Encode(out, outputImg) != nil {
		log.Fatal(err)
	}
}

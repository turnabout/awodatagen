// Generates sprite sheets & data used by AWO (outputs visuals.json & spritesheet.png)
package main

import (
	"image"
	"image/draw"
	"image/png"
	_ "image/png"
	"log"
	"os"
	"path"
	"runtime"
)

// Default name for the resulting spritesheet output
const ssOutputDefaultName string = "spritesheet.png"

// Environment variable holding the path where the sprite sheet should be output
const ssOutputEnvVar string = "AWO_SPRITESHEET"

// The base path of awssgen
var dirPath string = getDirPath()

// Grab this directory's full path
func getDirPath() string {
	// Grab awssgen's directory path
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		panic("No caller information")
	}

	return path.Dir(filename)
}

func main() {
	// Grab some existing images
	img1 := getImage(dirPath + "/raw_inputs/units/AntiAir/0/1.png")
	img2 := getImage(dirPath + "/raw_inputs/units/AntiAir/0/2.png")

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
	// Use either the AWO spritesheet environment variable path or this directory as a default
	var outputPath string
	var envExists bool

	if outputPath, envExists = os.LookupEnv(ssOutputEnvVar); !envExists {
		// Environment variable for spritesheet output doesn't exist, output it in this directory directly
		outputPath = dirPath + "/" + ssOutputDefaultName
	}

	writeImage(outputPath, outputImg)
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

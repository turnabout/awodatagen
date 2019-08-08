// Generates sprite sheets & data used by AWO (outputs visuals.json & spritesheet.png)
package main

import (
	"image"
	"image/png"
	_ "image/png"
	"log"
	"os"
	"path"
	"runtime"
	"sort"
)

// Object used to store all of the game's visual data
type visualsData struct {
	units [][]unitFrame
}

// Default name for the resulting spritesheet output
const ssOutputDefaultName string = "spritesheet.png"

// Environment variable holding the path where the sprite sheet should be output
const ssOutputEnvVar string = "AWO_SPRITESHEET"

// Directory containing spritesheet images
const ssImagesDirName string = "/raw_inputs"

// The base path of awssgen
var baseDirPath string = getDirPath()

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
	generateUnits()


	/*
	// Grab some existing images
	img1 := getImage(baseDirPath + "/raw_inputs/units/AntiAir/0/1.png")
	img2 := getImage(baseDirPath + "/raw_inputs/units/AntiAir/0/2.png")

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
		outputPath = baseDirPath + "/" + ssOutputDefaultName
	}

	writeImage(outputPath, outputImg)
	*/
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

// Get a slice of sorted keys from the given map
func getMapSortedKeys(m map[int]string) []int {
	sortedKeys := make([]int, 0)

	// Add all keys from the map
	for k := range m {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Ints(sortedKeys)

	return sortedKeys
}

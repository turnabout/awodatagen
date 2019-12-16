// Functions used to process inputs
package genio

import (
    "image"
    "log"
    "os"
)

// Gets the image stored at the given path
func getImage(path string) image.Image {
    // Load Image file
    imgFile, err := os.Open(path)

    if err != nil {
        log.Fatal(err)
    }

    // Decode Image
    img, _, err := image.Decode(imgFile)

    if err != nil {
        log.Fatal(err)
    }

    return img
}

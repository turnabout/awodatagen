// Functions used to process inputs
package genio

import (
    "image"
    "log"
    "os"
    "path/filepath"
)

// Gets image stored at the given path
func GetImage(path string) image.Image {

    if extension := filepath.Ext(path); extension != ".png" {
        log.Fatalf
    }

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

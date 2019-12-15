package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "path"
    "strconv"
    "strings"
)

// Generate Src visual data JSON & sprite sheet
func getUiData(packedFrameImgs *[]FrameImage) *UiData {
    vData := UiData{
        // Src: *getTilesSrcVData(packedFrameImgs),
    }

    return &vData
}

// Gathers data on every single UI image
func getUISrcFrameImgs(frameImgs *[]FrameImage) {
    uiDir := baseDirPath + inputsDirName + uiDirName + "/"

    // Loop all UI folders
    folders, err := ioutil.ReadDir(uiDir)
    if err != nil { log.Fatal(err) }

    for _, uiSubDir := range folders {

        // Ensure all looped values are directories
        if !uiSubDir.IsDir() {
            log.Fatal(
                fmt.Sprintf(
                    "Found file '%s' in UI directory (should only contain subdirectories)\n",
                    uiSubDir.Name(),
                ),
            )
        }

        // Get frame images from the sub directory
        gatherUiSubDirFrameImgs(frameImgs, uiSubDir.Name(), uiDir + uiSubDir.Name() + "/")
    }
}

func gatherUiSubDirFrameImgs(frameImgs *[]FrameImage, dirName string, dirPath string) {

    // Get the UI Element corresponding to this directory
    var uiElement UiElement
    var ok bool

    if uiElement, ok = uiElementsReverseStrings[dirName]; !ok {
        log.Fatal(fmt.Sprintf("UI Element directory '%s' not part of the UiElement enum\n", dirName))
    }

    // Loop all frames for this UI element
    uiSubDirFiles, err := ioutil.ReadDir(dirPath)
    if err != nil {log.Fatal(err)}

    for _, file := range uiSubDirFiles {
        appendUiFrameImg(dirPath, file.Name(), -1, uiElement, frameImgs)
    }
}

func appendUiFrameImg(dirPath string, fileName string, frameIndex int, uiElement UiElement, frameImgs* []FrameImage) {

    // Create the frame image for this UI element
    imageObj := getImage(dirPath + fileName)

    // If frame index not given, the frame index should be the file's name itself
    if frameIndex == -1 {
        var err error
        frameIndex, err = strconv.Atoi(strings.TrimSuffix(fileName, path.Ext(fileName)))
        if err != nil {log.Fatal(err)}
    }

    *frameImgs = append(*frameImgs, FrameImage{
        Image: imageObj,
        Width: imageObj.Bounds().Max.X,
        Height: imageObj.Bounds().Max.Y,
        MetaData: FrameImageMetaData{
            Type: uint8(uiElement),
            Index: frameIndex,
            FrameImageType: UiElementFrameImage,
        },
    })
}

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

    // Gather frame images from the elements found in the UI directory
    folders, err := ioutil.ReadDir(uiDir)
    if err != nil { log.Fatal(err) }

    for _, uiDirElement := range folders {

        if uiDirElement.IsDir() {
            gatherUiSubDirFrameImgs(frameImgs, uiDirElement.Name(), uiDir + uiDirElement.Name() + "/")
        } else {
            appendUiFrameImg(uiDir, uiDirElement.Name(), 0, UiElementNone, frameImgs)
        }
    }
}

func gatherUiSubDirFrameImgs(frameImgs *[]FrameImage, dirName string, dirPath string) {

    // Get the UI Element corresponding to this directory
    uiElement := getUiElementByString(dirName)

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

    // If ui element not given, the ui element should be the file's name itself
    if int(uiElement) == UiElementNone {
        uiElement = getUiElementByString(fileName)
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

func getUiElementByString(str string) UiElement {
    var ok bool
    var uiElement UiElement

    if uiElement, ok = uiElementsReverseStrings[str]; !ok {
        log.Fatal(fmt.Sprintf("UI Element string '%s' not part of the UiElement enum\n", str))
    }

    return uiElement
}

package uigen

import (
    "io/ioutil"
    "log"
    "path"
    "strconv"
    "strings"
)

// Generate Src visual data JSON & sprite sheet
func getUiData(packedFrameImgs *[]FrameImage) *UiData {
    var uiData *UiData = getUiBaseData(packedFrameImgs)

    return uiData
}

// Gathers data on every single UI image
func getUISrcFrameImgs(frameImgs *[]FrameImage) {

    // Gather frame images from the elements found in the UI directory
    folders, err := ioutil.ReadDir(getFullProjectPath(uiDir))
    if err != nil { log.Fatal(err) }

    for _, uiDirElement := range folders {
        if uiDirElement.IsDir() {
            gatherUiSubDirFrameImgs(
                frameImgs,
                uiDirElement.Name(),
                getFullProjectPath(uiDir, uiDirElement.Name()),
            )
        } else {
            appendUiFrameImg(
                getFullProjectPath(uiDir),
                uiDirElement.Name(),
                0,
                UiElementNone,
                frameImgs,
            )
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
    imageObj := getImage(path.Join(dirPath, fileName))

    // If frame index not given, the frame index should be the file's name itself
    if frameIndex == -1 {
        var err error
        frameIndex, err = strconv.Atoi(strings.TrimSuffix(fileName, path.Ext(fileName)))
        if err != nil {log.Fatal(err)}
    }

    // If ui element not given, the ui element should be the file's name itself
    if int(uiElement) == UiElementNone {
        uiElement = getUiElementByString(strings.TrimSuffix(fileName, path.Ext(fileName)))
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

func getUiBaseData(packedFrameImgs *[]FrameImage) *UiData {

    // UI Element Type -> UI Element Frames
    uiData := make(UiData, UiElementCount)

    // Process frame images
    for _, frameImg := range *packedFrameImgs {

        // Ignore non-UI element frame images
        if frameImg.MetaData.FrameImageType != UiElementFrameImage {
            continue
        }

        uiElement := UiElement(frameImg.MetaData.Type)
        uiElFrame := frameImg.MetaData.Index

        // Add any frames missing up until the one we're adding
        if missingFrames := (uiElFrame + 1) - len(uiData[uiElement]); missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                uiData[uiElement] = append(uiData[uiElement], Frame{})
            }
        }

        // Add the Frame data to the animation slice, and record it to the visual data
        frame := Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Width: frameImg.Width,
            Height: frameImg.Height,
        }

        uiData[uiElement][uiElFrame] = frame
    }

    return &uiData
}

func getUiElementByString(str string) UiElement {
    var ok bool
    var uiElement UiElement

    if uiElement, ok = uiElementsReverseStrings[str]; !ok {
        log.Fatalf("UI Element string '%s' not part of the UiElement enum\n", str)
    }

    return uiElement
}

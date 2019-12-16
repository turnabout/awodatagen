package uigen

import (
    "github.com/turnabout/awossgen"
    "github.com/turnabout/awossgen/pkg/packer"
)

// Generate Src visual data JSON & sprite sheet
func getUiData(packedFrameImgs *[]packer.FrameImage) *awossgen.UIData {
    var uiData *awossgen.UIData = getUiBaseData(packedFrameImgs)

    return uiData
}

func getUiBaseData(packedFrameImgs *[]packer.FrameImage) *awossgen.UIData {

    // UI Element Type -> UI Element Frames
    uiData := make(awossgen.UIData, awossgen.UIElementCount)

    // Process frame images
    for _, frameImg := range *packedFrameImgs {

        // Ignore non-UI element frame images
        if frameImg.MetaData.FrameImageDataType != uint8(awossgen.UIDataType) {
            continue
        }

        uiElement := awossgen.UIElement(frameImg.MetaData.Type)
        uiElFrame := frameImg.MetaData.Index

        // Add any frames missing up until the one we're adding
        if missingFrames := (uiElFrame + 1) - len(uiData[uiElement]); missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                uiData[uiElement] = append(uiData[uiElement], awossgen.Frame{})
            }
        }

        // Add the Frame data to the animation slice, and record it to the visual data
        frame := awossgen.Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Width: frameImg.Width,
            Height: frameImg.Height,
        }

        uiData[uiElement][uiElFrame] = frame
    }

    return &uiData
}


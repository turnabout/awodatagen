package uigen

import (
	"github.com/turnabout/awodatagen/internal/packer"
	"github.com/turnabout/awodatagen/pkg/framedata"
)

// Generates the UI-related game data
func GetUIData(packedFrameImgs *[]packer.FrameImage) *UIData {
	var UIData *UIData = getUIBaseData(packedFrameImgs)

	return UIData
}

func getUIBaseData(packedFrameImgs *[]packer.FrameImage) *UIData {

	// UI Element Type -> UI Element Frames
	UIData := make(UIData, UIElementCount)

	// Process frame images
	for _, frameImg := range *packedFrameImgs {

		// Ignore non-UI element frame images
		if frameImg.MetaData.FrameImageDataType != uint8(framedata.UIDataType) {
			continue
		}

		// Get metadata on the UI element this frame image represents
		uiElement := UIElement(frameImg.MetaData.Type)
		uiElFrame := frameImg.MetaData.Index

		// Add any frames missing up until the one we're adding
		if missingFrames := (uiElFrame + 1) - len(UIData[uiElement]); missingFrames > 0 {
			for i := 0; i < missingFrames; i++ {
				UIData[uiElement] = append(UIData[uiElement], framedata.Frame{})
			}
		}

		// Add the Frame data to the animation slice, and record it to the visual data
		frame := framedata.Frame{
			X:      frameImg.X,
			Y:      frameImg.Y,
			Width:  frameImg.Width,
			Height: frameImg.Height,
		}

		UIData[uiElement][uiElFrame] = frame
	}

	return &UIData
}

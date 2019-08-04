package main

// Object used to store all of the game's visual data
type visualsData struct {
	units [][]unitFrameData
}

// Used to store a unit's frame's visual data within the game's sprite sheet
type unitFrameData struct{
	x, y, w, h int
}

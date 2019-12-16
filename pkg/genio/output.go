// Functions used to process outputs
package genio

import (
    "encoding/json"
    "fmt"
    "github.com/turnabout/awossgen"
    "image"
    "image/png"
    "io/ioutil"
    "log"
    "os"
    "path"
    "regexp"
)

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

// Attach the JSON data at the given file path and stores the result in the value pointed to by v
func AttachJSONData(jsonPath string , v interface{}) {
    data, err := ioutil.ReadFile(jsonPath)
    awossgen.LogFatalIfErr(err)

    // Make Regexp used to remove comments
    re := regexp.MustCompile(`//.*`)

    // Unmarshal and store the result
    err = json.Unmarshal(re.ReplaceAll(data, []byte("")), v)
    awossgen.LogFatalIfErr(err, jsonPath)
}

// Output the visuals data JSON
func OutputJSON(visualData *awossgen.GameData) {
    // data, err := json.Marshal(visualData)
    data, err := json.MarshalIndent(visualData, "", "\t")

    awossgen.LogFatalIfErr(err)

    // Use either the awossgen JSON environment variable path or this directory as a default
    var jsonOutputPath string
    var envExists bool

    // If environment variable doesn't exist, output in this directory directly
    if jsonOutputPath, envExists = os.LookupEnv(awossgen.JSONOutputEnvVar); !envExists {
        jsonOutputPath = path.Join(".", awossgen.JSONOutputDefaultName)
    }

    err = ioutil.WriteFile(jsonOutputPath, data, 0644)
    awossgen.LogFatalIfErr(err)

    fmt.Printf("Output %s\n", jsonOutputPath)
}

// Output the game sprite sheet
func OutputSpriteSheet(ss *image.RGBA) {
    // Use either the awossgen sprite sheet environment variable path or this directory as a default
    var ssOutputPath string
    var envExists bool

    // If environment variable doesn't exist, output in this directory directly
    if ssOutputPath, envExists = os.LookupEnv(awossgen.SSOutputEnvVar); !envExists {
        ssOutputPath = path.Join(".", awossgen.SSOutputDefaultName)
    }

    writeImage(ssOutputPath, ss)
    fmt.Printf("Output %s\n", ssOutputPath)
}

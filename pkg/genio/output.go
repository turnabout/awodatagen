// Functions used to process outputs
package genio

import (
    "encoding/json"
    "fmt"
    "github.com/turnabout/awodatagen"
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
    awodatagen.LogFatalIfErr(err)

    // Make Regexp used to remove comments
    re := regexp.MustCompile(`//.*`)

    // Unmarshal and store the result
    err = json.Unmarshal(re.ReplaceAll(data, []byte("")), v)
    awodatagen.LogFatalIfErr(err, jsonPath)
}

// Output the visuals data JSON
func OutputJSON(visualData *awodatagen.GameData) {
    // data, err := json.Marshal(visualData)
    data, err := json.MarshalIndent(visualData, "", "\t")

    awodatagen.LogFatalIfErr(err)

    // Use either the awodatagen JSON environment variable path or this directory as a default
    var jsonOutputPath string
    var envExists bool

    // If environment variable doesn't exist, output in this directory directly
    if jsonOutputPath, envExists = os.LookupEnv(awodatagen.JSONOutputEnvVar); !envExists {
        jsonOutputPath = path.Join(".", awodatagen.JSONOutputDefaultName)
    }

    err = ioutil.WriteFile(jsonOutputPath, data, 0644)
    awodatagen.LogFatalIfErr(err)

    fmt.Printf("Output %s\n", jsonOutputPath)
}

// Output the game sprite sheet
func OutputSpriteSheet(ss *image.RGBA) {
    // Use either the awodatagen sprite sheet environment variable path or this directory as a default
    var ssOutputPath string
    var envExists bool

    // If environment variable doesn't exist, output in this directory directly
    if ssOutputPath, envExists = os.LookupEnv(awodatagen.SSOutputEnvVar); !envExists {
        ssOutputPath = path.Join(".", awodatagen.SSOutputDefaultName)
    }

    writeImage(ssOutputPath, ss)
    fmt.Printf("Output %s\n", ssOutputPath)
}

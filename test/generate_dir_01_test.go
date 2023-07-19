package test

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var stRootDir string
var stSeparator string
var iJsonData map[string]any

const stJsonFilename = "dir.json"

func loadJson() {
	stSeparator = string(filepath.Separator)
	stWorkDir, _ := os.Getwd()
	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)]

	gnJsonBytes, _ := os.ReadFile(stWorkDir + stSeparator + stJsonFilename)
	err := json.Unmarshal(gnJsonBytes, &iJsonData)
	if err != nil {
		panic("Load json file error: " + err.Error())
	}
}

func parseMap(m map[string]any, parentPath string) {
	for _, v := range m {
		switch v.(type) {
		case string:
			path, _ := v.(string)
			if path == "" {
				continue
			}
			parentPath = parentPath + stSeparator + path
			//create the dir
			createDir(parentPath)
			continue

		case []any:
			for _, val := range v.([]any) {
				parseMap(val.(map[string]any), parentPath)
			}
		}
	}
}

func createDir(path string) {
	if path == "" {
		return
	}

	err := os.Mkdir(path, 0755)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
func TestGenerateDir01(t *testing.T) {
	//load json file to the map(iJsonData)
	loadJson()
	//parse the map
	parseMap(iJsonData, stRootDir)
}

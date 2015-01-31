package brewci

import (
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

// builds the map of tea brewing times
// data is based on the file at '~/.shinshutsu.yml'
func BuildTeasMap() map[string][]int {
	filePath := buildFilePath()
	yamlFile := buildYamlFile(filePath)
	return buildTeas(yamlFile)
}

func buildTeas(data []byte) (teas map[string][]int) {
	err := yaml.Unmarshal([]byte(data), &teas)
	handleError(err)
	return
}

func buildYamlFile(filePath string) (yamlFile []byte) {
	yamlFile, err := ioutil.ReadFile(filePath)
	handleError(err)
	return
}

func buildFilePath() (filePath string) {
	filePath, err := filepath.Abs(getBrewingDataFilePath())
	handleError(err)
	return
}

var getBrewingDataFilePath = func() (filePath string) {
	filePath, err := homedir.Expand("~/.shinshutsu.yml")
	handleError(err)
	return
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gitlab.com/SiegfriedEhret/gitzytout/pkg/gitconfig"

	"github.com/ghodss/yaml"
)

// Config represents the yaml content
type Config struct {
	Main    string   `json:"main"`
	Mirrors []string `json:"mirrors"`
}

func (c Config) String() string {
	return c.Main + "; " + strings.Join(c.Mirrors, ", ")
}

func inArray(array []string, value string) bool {
	for _, v := range array {
		if strings.Compare(value, v) == 0 {
			return true
		}
	}
	return false
}

func maybeAddOrigin(main string) {
	remoteURL := gitconfig.GetRemoteOrigin()

	if strings.Compare(main, remoteURL) != 0 {
		errMain := gitconfig.AddOrigin(main)
		if errMain != nil {
			fmt.Println("Error while adding origin: "+main, errMain.Error())
		}
	}
}

func maybeAddPushUrls(main string, mirrors []string) {
	pushUrls := gitconfig.GetPushURL()

	things := []string{main}
	things = append(things, mirrors...)

	for _, mirror := range things {
		if !inArray(pushUrls, mirror) {
			err := gitconfig.AddPushURL(mirror)
			if err != nil {
				fmt.Println("Error while adding push url: `"+mirror, err.Error())
			}
		}
	}
}

func main() {
	fmt.Println("gitzytout\n=========")

	const configPath string = "gitzytout.yaml"

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Can't read config %s", configPath)
	}

	var conf Config
	unmarshalErr := yaml.Unmarshal(data, &conf)

	if unmarshalErr != nil {
		fmt.Println("Can't decode yaml !")
	}

	maybeAddOrigin(conf.Main)
	maybeAddPushUrls(conf.Main, conf.Mirrors)

	fmt.Println("Done!")
}

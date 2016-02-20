package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gitlab.com/SiegfriedEhret/gitzytout/pkg/gitconfig"

	"github.com/ghodss/yaml"
)

const (
	configPath string = "gitzytout.yaml"
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

func maybeAddOrigin(main string) (err error) {
	remoteURL := gitconfig.GetRemoteOrigin()

	if strings.Compare(main, remoteURL) != 0 {
		if err := gitconfig.AddOrigin(main); err != nil {
			fmt.Println("Error while adding origin: "+main, err.Error())
		}
	}

	return
}

func maybeAddPushUrls(main string, mirrors []string) (errors []error) {
	pushUrls := gitconfig.GetPushURL()
	things := []string{main}
	things = append(things, mirrors...)

	for _, mirror := range things {
		if !inArray(pushUrls, mirror) {
			if err := gitconfig.AddPushURL(mirror); err != nil {
				errors = append(errors, err)
			}
		}
	}

	return
}

func main() {
	fmt.Println("gitzytout\n=========")

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("Can't read config " + configPath)
	}

	var conf Config
	unmarshalErr := yaml.Unmarshal(data, &conf)

	if unmarshalErr != nil {
		log.Fatal("Can't decode yaml !")
	}

	errMain := maybeAddOrigin(conf.Main)
	errMirrors := maybeAddPushUrls(conf.Main, conf.Mirrors)

	if errMain != nil {
		log.Println("Failed to set up the main repository", conf.Main)
	}

	if len(errMirrors) > 0 {
		for i := 0; i < len(errMirrors); i++ {
			log.Println("Failed to write push url: ", errMirrors[i])
		}
	}

	fmt.Println("Done!")
}

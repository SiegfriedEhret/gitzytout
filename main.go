package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"os/user"
	"strings"

	"github.com/ghodss/yaml"
)

type Config struct {
	Main    string   `json:"main"`
	Mirrors []string `json:"mirrors"`
}

func (c Config) String() string {
	return strings.Join(c.Mirrors, ", ")
}

func GetUserDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("!! Can't get current user", err)
		return ""
	}

	return usr.HomeDir
}

func main() {
	fmt.Println("  gitzytout\n  =========")

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

	fmt.Println(conf)

	// git remote set-url origin --push --add git@gitlab.com:SiegfriedEhret/test.git
	// git remote set-url origin --push --add git@github.com:SiegfriedEhret/test.git

	for _, mirror := range conf.Mirrors {
		cmd := exec.Command("git", "remote", "set-url", "origin", "--push", "--add", mirror)
		fmt.Println(cmd)
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error while running `" + strings.Join(cmd.Args, " ") + "`")
		}
	}
}

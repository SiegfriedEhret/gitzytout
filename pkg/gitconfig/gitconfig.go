// Package gitconfig expose a few methods to interact with a git configuration
package gitconfig

import (
	"fmt"
	"os/exec"
	"strings"
)

// AddOrigin sets the remote origin
func AddOrigin(main string) error {
	output, err := exec.Command("git", "remote", "add", "origin", main).Output()

	if err != nil && len(output) != 0 {
		fmt.Println("Something weird happend while adding origin", err.Error(), "\nOutput is:", output)
	}

	return err
}

// AddPushURL adds a url to push your code
func AddPushURL(mirror string) error {
	output, err := exec.Command("git", "remote", "set-url", "origin", "--push", "--add", mirror).Output()

	if err != nil && len(output) != 0 {
		fmt.Println("Something weird happend while adding push url: ", err.Error(), "\nOutput is:", output)
	}

	return err
}

// GetPushURL returns an array of available push urls
func GetPushURL() []string {
	output, err := exec.Command("git", "config", "--get-all", "remote.origin.pushurl").Output()

	if err != nil && len(output) != 0 {
		fmt.Println("Something weird happend while reading available git push urls: ", err.Error(), "\nOutput is: ", output)
	}

	stringOutput := string(output)
	stringArray := strings.Split(stringOutput, "\n")

	return stringArray
}

// GetRemoteOrigin returns the configured remote origin url
func GetRemoteOrigin() string {
	output, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()

	if err != nil && len(output) != 0 {
		fmt.Println("Something weird happend while reading available remote origin", err.Error(), "\nOutput is: ", output)
	}

	return strings.Trim(string(output), "\n")
}

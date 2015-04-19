package timesync

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	URL    string
	User   string
	Editor string
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		return false
	}
}

func GetConfig() (Config, error) {
	config := Config{}

	// Check for the config file at known locations. If it does not exist there
	// then assume this is a development environment and act accordingly.
	path := ""
	homeDir := os.Getenv("HOME")
	if pathExists(homeDir + "/.timesync-cli.json") {
		path = homeDir + "/.timesync-cli.json"
	} else if pathExists("./.timesync-cli.json") {
		path = "./.timesync-cli.json"
	} else { // We assume this is a development environment
		config.User = ""
		config.Editor = ""
		config.URL = "http://localhost:8000"
	}

	if path != "" {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return config, err
		}
		err = json.Unmarshal(data, &config)
		if err != nil {
			return config, err
		}
	}

	// If the user or the editor are not set, get them from the
	// environment.
	if config.User == "" {
		config.User = os.Getenv("USER")
	}

	if config.Editor == "" {
		config.Editor = os.Getenv("EDITOR")
		if config.Editor == "" {
			config.Editor = "vi"
		}
	}

	return config, nil
}

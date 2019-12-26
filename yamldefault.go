package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type NotesConfig struct {
	DefaultLocation string            `yaml:"DEFAULT"`
	Locations       map[string]string `yaml:"LOCATIONS"`
}

func logErr(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func expandPath(path string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if path == "~" {
		return dir
	} else if strings.HasPrefix(path, "~/") {
		return filepath.Join(dir, path[2:])
	}
	return path
}

func overwriteDefaultLocation(t NotesConfig, osArgs []string) NotesConfig {
	if len(osArgs) > 2 {
		t.DefaultLocation = osArgs[2]
	}
	return t
}

func main() {
	path := expandPath(os.Args[1])
	data, err := ioutil.ReadFile(path)
	logErr(err)

	config := NotesConfig{}
	err = yaml.Unmarshal([]byte(data), &config)
	logErr(err)

	config = overwriteDefaultLocation(config, os.Args)

	yaml_string, err := yaml.Marshal(&config)
	logErr(err)
	err = ioutil.WriteFile(path, yaml_string, 0644)
	logErr(err)

	directory := expandPath(config.Locations[config.DefaultLocation])
	fmt.Printf(directory)
}

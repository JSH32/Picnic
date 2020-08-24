package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Riku32/Picnic/logger"
	"gopkg.in/yaml.v2"
)

// Prop : command property object
type Prop struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Category    string   `yaml:"category"`
	Aliases     []string `yaml:"aliases"`
}

// Command : command object
type Command struct {
	Prop    Prop
	Command string
}

// Loader : create a command map from loaded commands
func Loader() Handler {
	files, err := ioutil.ReadDir("./commands")
	if err != nil {
		logger.Panic(err.Error())
	}

	commands := make(map[string]Command)

	for _, f := range files {
		if !f.IsDir() {
			logger.Warn(fmt.Sprintf("%s is not a directory", f.Name()))
			continue
		}

		propf, err := os.Open(fmt.Sprintf("./commands/%s/prop.yaml", f.Name()))
		if err != nil {
			logger.Warn(fmt.Sprintf("No prop file found in %s command", f.Name()))
			continue
		}
		defer propf.Close()

		commandf, err := ioutil.ReadFile(fmt.Sprintf("./commands/%s/command.js", f.Name()))
		if err != nil {
			logger.Warn(fmt.Sprintf("No command script found in %s command", f.Name()))
			continue
		}

		var prop Prop

		err = yaml.NewDecoder(propf).Decode(&prop)
		if err != nil {
			logger.Warn(fmt.Sprintf("Invalid prop for %s command", f.Name()))
			continue
		}

		command := Command{
			Command: string(commandf),
			Prop:    prop,
		}

		if _, ok := commands[command.Prop.Name]; ok {
			logger.Warn(fmt.Sprintf("%s has already been registered", f.Name()))
			continue
		}

		commands[command.Prop.Name] = command
	}

	return Handler{
		Commands: commands,
	}
}

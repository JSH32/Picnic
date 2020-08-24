package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Riku32/Picnic/stdlib/logger"
	babel "github.com/jvatic/goja-babel"
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
func Loader(transpile bool) Handler {
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

		var script = string(commandf)

		// ES6/2015 support
		if transpile {
			script, err = babel.TransformString(string(commandf), map[string]interface{}{
				"plugins": []string{
					"transform-es2015-arrow-functions",
					"transform-es2015-block-scoped-functions",
					"transform-es2015-block-scoping",
					"transform-es2015-classes",
					"transform-es2015-computed-properties",
					"transform-es2015-destructuring",
					"transform-es2015-duplicate-keys",
					"transform-es2015-for-of",
					"transform-es2015-function-name",
					"transform-es2015-instanceof",
					"transform-es2015-literals",
					"transform-es2015-object-super",
					"transform-es2015-parameters",
					"transform-es2015-shorthand-properties",
					"transform-es2015-spread",
					"transform-es2015-sticky-regex",
					"transform-es2015-template-literals",
					"transform-es2015-typeof-symbol",
					"transform-es2015-unicode-regex",
				},
			})
			if err != nil {
				logger.Error(fmt.Sprintf("Unable to transpile module %s : %s", f.Name(), err.Error()))
				continue
			}
		}

		command := Command{
			Command: script,
			Prop:    prop,
		}

		if _, ok := commands[command.Prop.Name]; ok {
			logger.Warn(fmt.Sprintf("%s has already been registered", f.Name()))
			continue
		}

		commands[command.Prop.Name] = command
	}

	logger.Info(fmt.Sprintf("%d commands have been registered!", len(commands)))

	return Handler{
		Commands: commands,
	}
}

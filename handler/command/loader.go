package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Riku32/Picnic/stdlib/logger"
	babel "github.com/Riku32/goja-babel"
	"github.com/dop251/goja_nodejs/require"
	"gopkg.in/yaml.v2"
)

var (
	es6plugin = []string{
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
	}
	minifyplugin = []string{
		"transform-inline-environment-variables",
		"transform-member-expression-literals",
		"transform-merge-sibling-variables",
		"transform-minify-booleans",
		"minify-constant-folding",
		"minify-dead-code-elimination",
		"minify-flip-comparisons",
		"minify-guarded-expressions",
		"minify-infinity",
		"minify-numeric-literals",
		"minify-replace",
		"minify-simplify",
		"minify-type-constructors",
		"transform-property-literals",
	}
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

// ModuleStore : store for native javascript modules
type ModuleStore map[string]string

// SourceLoader : loader for loading native javascript modules
func (m ModuleStore) SourceLoader(filename string) ([]byte, error) {
	if data, ok := m[filename]; ok {
		return []byte(data), nil
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, require.ModuleFileDoesNotExistError
	}
	script, err := babel.TransformString(string(data), map[string]interface{}{
		"plugins": append(es6plugin, minifyplugin...),
	})
	if err != nil {
		logger.Panic(fmt.Sprintf("Unable to compile stdlib %s : %s", filename, err.Error()))
	}
	m[filename] = script
	return []byte(script), nil
}

// Loader : create a command map from loaded commands
func Loader(transpile bool) Handler {
	/*
		JS standard library initializer
		This gets minified
		Included in all files automatically
	*/
	files, err := ioutil.ReadDir("./jlib")
	if err != nil {
		logger.Panic(err.Error())
	}

	// var libprepack string

	// for _, f := range files {
	// 	if f.IsDir() {
	// 		logger.Warn(fmt.Sprintf("%s is a directory", f.Name()))
	// 		continue
	// 	}

	// 	scriptr, err := ioutil.ReadFile(fmt.Sprintf("./jlib/%s", f.Name()))
	// 	if err != nil {
	// 		logger.Warn(fmt.Sprintf("Unable to read script %s", f.Name()))
	// 		continue
	// 	}

	// 	libprepack += string(scriptr)
	// }

	// jlib, err := babel.TransformString(string(libprepack), map[string]interface{}{
	// 	"plugins": append(es6plugin, minifyplugin...),
	// })
	// if err != nil {
	// 	logger.Panic("Unable to compile standard library : " + err.Error())
	// }

	/*
		Read command scripts
		This gets optionally babel packed
	*/
	files, err = ioutil.ReadDir("./commands")
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
				"plugins": es6plugin,
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

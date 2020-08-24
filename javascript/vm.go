package javascript

import (
	"fmt"

	"github.com/Riku32/Picnic/handler/command"
	"github.com/Riku32/Picnic/stdlib/logger"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type Vm struct {
	runtime *goja.Runtime
}

func NewVM() Vm {
	registry := new(require.Registry)
	vm := goja.New()

	registry.Enable(vm)

	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	runtime := Vm{
		runtime: vm,
	}

	runtime.setglobals()

	return runtime
}

func (vm Vm) setglobals() {
	vm.runtime.Set("logger", logger.JSLogger{})
}

// Execute : execute a js script
func (vm Vm) Execute(command command.Command) {
	_, err := vm.runtime.RunString(command.Command)
	if err != nil {
		logger.Error(fmt.Sprintf("[%s] %s", command.Prop.Name, err.Error()))
	}
}

// SetGlobal : set a variable in
func (vm Vm) SetGlobal(key string, value interface{}) {
	vm.runtime.Set(key, value)
}

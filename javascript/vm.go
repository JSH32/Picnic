package javascript

import (
	"github.com/Riku32/Picnic/logger"
	"github.com/dop251/goja"
)

type Vm struct {
	runtime *goja.Runtime
}

func NewVM() Vm {
	vm := goja.New()
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	runtime := Vm{
		runtime: vm,
	}

	runtime.setglobals()

	return runtime
}

func (vm Vm) setglobals() {
	vm.runtime.Set("logger", logger.NewJSLogger())
}

// Execute : execute a js script
func (vm Vm) Execute(script string) {
	_, err := vm.runtime.RunString(script)
	if err != nil {
		logger.Error(err.Error())
	}
}

// SetGlobal : set a variable in
func (vm Vm) SetGlobal(key string, value interface{}) {
	vm.runtime.Set(key, value)
}

package javascript

import (
	"fmt"

	"github.com/Riku32/Picnic/handler/command"
	"github.com/Riku32/Picnic/stdlib/http"
	"github.com/Riku32/Picnic/stdlib/logger"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
)

type Vm struct {
	runtime *eventloop.EventLoop
}

func NewVM() Vm {
	vm := eventloop.NewEventLoop()

	// require.RegisterNativeModule("xf", func(f *goja.Runtime, x *goja.Object) {
	// 	x.Set("lmao", "sdf")
	// })
	nvstore := make(command.ModuleStore)
	rm := require.NewRegistry(require.WithGlobalFolders("jlib"), require.WithLoader(nvstore.SourceLoader))

	vm.Run(func(vm *goja.Runtime) {
		vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
		rm.Enable(vm)
	})

	runtime := Vm{
		runtime: vm,
	}

	return runtime
}

func (vm Vm) setglobals() {
	vm.SetGlobal("logger", logger.JSLogger{})
	vm.SetGlobal("http", http.Http{})
}

// Execute : execute a js script
func (vm Vm) Execute(command command.Command) {
	vm.runtime.Run(func(vm *goja.Runtime) {
		_, err := vm.RunString(command.Command)
		if err != nil {
			logger.Error(fmt.Sprintf("[%s] %s", command.Prop.Name, err.Error()))
		}
	})
}

// GetCore : for those times when you really need it
func (vm Vm) GetCore() *eventloop.EventLoop {
	return vm.runtime
}

// SetGlobal : set a variable in
func (vm Vm) SetGlobal(key string, value interface{}) {
	vm.runtime.Run(func(vm *goja.Runtime) {
		vm.Set(key, value)
	})
}

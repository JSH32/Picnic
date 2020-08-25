package test

import (
	"fmt"

	"github.com/dop251/goja"
)

func Tx(call goja.ConstructorCall) *goja.Object {
	call.This.Set("send", func(str string) {
		fmt.Println(str)
	})

	return nil
}

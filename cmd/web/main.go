// +build js,wasm

package main

import (
	"github.com/flostadler/name-generator/pkg/generators/animals"
	"github.com/flostadler/name-generator/pkg/generators/dc"
	"github.com/flostadler/name-generator/pkg/generators/docker"
	"github.com/flostadler/name-generator/pkg/generators/got"
	"github.com/flostadler/name-generator/pkg/generators/marvel"
	"github.com/flostadler/name-generator/pkg/generators/hp"
	"syscall/js"
)

func main() {
	c := make(chan bool)

	gen := js.Global().Get("Object").New()
	gen.Set("dockerName", js.FuncOf(func (this js.Value, inputs []js.Value) interface{} {
		return docker.Generate()
	}))
	gen.Set("marvelCharacter", js.FuncOf(func (this js.Value, inputs []js.Value) interface{} {
		return marvel.Generate()
	}))
	gen.Set("dcCharacter", js.FuncOf(func (this js.Value, inputs []js.Value) interface{} {
		return dc.Generate()
	}))
	gen.Set("animal", js.FuncOf(func (this js.Value, inputs []js.Value) interface{} {
		return animals.Generate()
	}))
	gen.Set("got", js.FuncOf(func (this js.Value, inputs []js.Value) interface{} {
		return got.Generate()
	}))
	gen.Set("hp", js.FuncOf(func (this js.Value, inputs []js.Value) interface{} {
		return hp.Generate()
	}))

	js.Global().Set("NameGenerator", gen)
	<-c
}
